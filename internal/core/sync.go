package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cursor++/internal/git"
	"cursor++/internal/ui"
	"cursor++/internal/utils"
)

// AgentInitializer handles agent system initialization
type AgentInitializer struct {
	agentPath string
	registry  *Registry
	config    *utils.Config
	appPaths  utils.AppPaths
	gitMgr    *git.GitManager
}

// NewAgentInitializer creates a new agent initializer
func NewAgentInitializer() (*AgentInitializer, error) {
	// Load config
	cm := utils.NewConfigManager()
	if err := cm.Load(); err != nil {
		return nil, wrapOpError("NewAgentInitializer", "config", err, "failed to load configuration")
	}
	config := cm.GetConfig()
	utils.Debug("Loaded configuration")

	// Get app name from environment
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = utils.DefaultAgentsDirName
	}

	// Get system paths
	appPaths := utils.GetAppPaths(appName)
	utils.Debug("Loaded system paths for current platform")

	// Use OS-specific paths
	agentPath := appPaths.GetRulesDir(config.RulesDirName)
	registryPath := appPaths.GetRegistryFile(config.RegistryFileName)

	// Ensure required directories exist
	if err := utils.EnsureDirExists(appPaths.ConfigDir, config.DirPermission); err != nil {
		return nil, wrapOpError("NewAgentInitializer", appPaths.ConfigDir, err, "failed to create config directory")
	}

	if err := utils.EnsureDirExists(appPaths.DataDir, config.DirPermission); err != nil {
		return nil, wrapOpError("NewAgentInitializer", appPaths.DataDir, err, "failed to create data directory")
	}

	if err := utils.EnsureDirExists(appPaths.LogDir, config.DirPermission); err != nil {
		return nil, wrapOpError("NewAgentInitializer", appPaths.LogDir, err, "failed to create log directory")
	}

	utils.Debug("Using agent path | path=" + agentPath)

	// Ensure agent directory exists
	if err := os.MkdirAll(agentPath, config.DirPermission); err != nil {
		return nil, wrapOpError("NewAgentInitializer", agentPath, err, "failed to create agent directory")
	}

	// Load or create registry
	registry, err := LoadRegistry(registryPath, config)
	if err != nil {
		return nil, wrapOpError("NewAgentInitializer", registryPath, err, "failed to load registry")
	}
	utils.Debug("Registry loaded successfully")

	return &AgentInitializer{
		agentPath: agentPath,
		registry:  registry,
		config:    config,
		appPaths:  appPaths,
		gitMgr:    git.NewGitManager(nil),
	}, nil
}

// Init initializes the agent system in the current directory
func (ai *AgentInitializer) Init() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return wrapOpError("Init", "cwd", err, "failed to get current directory")
	}

	// Add detailed debug info about configuration
	if utils.IsDebug() {
		utils.Debugf("Init configuration details | agentPath=%s | rulesDirName=%s | dataDir=%s | sourceFolder=%s",
			ai.agentPath, ai.config.RulesDirName, ai.appPaths.DataDir, ai.config.SourceFolder)
		utils.Debugf("Registry location | path=%s | projects=%d",
			ai.appPaths.GetRegistryFile(ai.config.RegistryFileName), ai.registry.GetProjectCount())
	}

	// Add verbose info
	if utils.IsVerbose() {
		utils.Info("Initializing agent system in current directory")
		utils.Infof("Using system agent path: %s", ai.agentPath)
	}

	// The target path is the rules directory in the current project
	targetPath := filepath.Join(currentDir, ai.config.RulesDirName)
	utils.Debug("Init target path | path=" + targetPath)

	// Check if the agent location exists and has valid definitions
	needsSetup := false

	if !utils.DirExists(ai.agentPath) {
		utils.Debug("Agent location does not exist | path=" + ai.agentPath)
		ui.Warning("Agent location does not exist: %s", ai.agentPath)
		needsSetup = true
	} else {
		// Check source folder if it's specified
		if ai.config.SourceFolder != "" {
			sourceFolderPath := filepath.Join(ai.agentPath, ai.config.SourceFolder)
			if !utils.DirExists(sourceFolderPath) {
				utils.Debug("Source folder does not exist | path=" + sourceFolderPath)
				ui.Warning("Source folder does not exist: %s", sourceFolderPath)
				needsSetup = true
			} else {
				hasMDCFiles, err := utils.HasMDCFiles(sourceFolderPath)
				if err != nil {
					return wrapOpError("Init", sourceFolderPath, err, "failed to check for agent definitions in source folder")
				}

				if !hasMDCFiles {
					utils.Debug("Source folder exists but contains no definitions | path=" + sourceFolderPath)
					ui.Warning("Source folder exists but contains no definitions: %s", sourceFolderPath)
					needsSetup = true
				} else if utils.IsVerbose() {
					utils.Info("Found existing agent definitions in source folder")
				}
			}
		} else {
			hasMDCFiles, err := utils.HasMDCFiles(ai.agentPath)
			if err != nil {
				return wrapOpError("Init", ai.agentPath, err, "failed to check for agent definitions")
			}

			if !hasMDCFiles {
				utils.Debug("Agent location exists but contains no definitions | path=" + ai.agentPath)
				ui.Warning("Agent location exists but contains no definitions: %s", ai.agentPath)
				needsSetup = true
			} else if utils.IsVerbose() {
				utils.Info("Found existing agent definitions")
			}
		}
	}

	if needsSetup {
		if !ai.handleInitialSetup() {
			return wrapValidationError("setup", "agent system initialization cancelled")
		}
	}

	// Log copy operation details
	if utils.IsVerbose() {
		if ai.config.SourceFolder != "" {
			utils.Infof("Copying agent definitions from '%s' subfolder to project directory: %s",
				ai.config.SourceFolder, targetPath)
		} else {
			utils.Infof("Copying all agent definitions to project directory: %s", targetPath)
		}
	}

	// Debug with more details about the copy operation
	if utils.IsDebug() {
		utils.Debugf("Copy operation details | source=%s | sourceFolder=%s | target=%s | permission=%o",
			ai.agentPath, ai.config.SourceFolder, targetPath, ai.config.DirPermission)
	}

	// Ensure the target directory exists
	if err := os.MkdirAll(targetPath, ai.config.DirPermission); err != nil {
		return wrapOpError("Init", targetPath, err, "failed to create target directory")
	}

	// Copy agent definitions selectively to the rules directory
	if err := utils.CopyDirSelective(ai.agentPath, targetPath, ai.config.SourceFolder); err != nil {
		return wrapOpError("Init", targetPath, err, "failed to copy agent definitions")
	}

	if utils.IsDebug() {
		// Verify the copy operation
		mdcFiles := 0
		err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mdc") {
				mdcFiles++
				utils.Debugf("Copied MDC file: %s", path)
			}
			return nil
		})
		if err != nil {
			utils.Debugf("Error walking target directory: %v", err)
		} else {
			utils.Debugf("Total MDC files copied: %d", mdcFiles)
		}
	}

	// Add project to registry
	if err := ai.registry.AddProject(currentDir); err != nil {
		return wrapOpError("Init", currentDir, err, "failed to register project")
	}

	if utils.IsVerbose() {
		utils.Info("Project registered in agent registry")
	}

	ui.Success("Successfully initialized agent system in %s", currentDir)
	return nil
}

// handleInitialSetup manages the initial setup of the agent system
func (ai *AgentInitializer) handleInitialSetup() bool {
	// Default repository URL instead of prompting the user
	defaultRepoURL := "https://github.com/nsnarender5511/AgenticSystem"

	ui.Info("\nNo agent definitions found. Automatically cloning from default repository...")
	ui.Info("Repository URL: %s", defaultRepoURL)

	// Add more verbose information
	if utils.IsVerbose() {
		utils.Infof("Target agent path: %s", ai.agentPath)
		if ai.config.SourceFolder != "" {
			utils.Infof("Will use '%s' subfolder for rules", ai.config.SourceFolder)
		}
	}

	// Add detailed debug information
	if utils.IsDebug() {
		utils.Debugf("Clone operation details | repo=%s | path=%s | permission=%o | sourceFolder=%s",
			defaultRepoURL, ai.agentPath, ai.config.DirPermission, ai.config.SourceFolder)
		utils.Debugf("Agent registry details | projects=%d | registryPath=%s",
			ai.registry.GetProjectCount(), ai.appPaths.GetRegistryFile(ai.config.RegistryFileName))
	}

	if err := ai.cloneRepository(defaultRepoURL); err != nil {
		ui.Error(err.Error())
		return false
	}

	// Verify and log the clone results
	if utils.DirExists(ai.agentPath) {
		// Check if source folder exists
		if ai.config.SourceFolder != "" {
			sourceFolderPath := filepath.Join(ai.agentPath, ai.config.SourceFolder)
			if !utils.DirExists(sourceFolderPath) {
				ui.Error("Source folder '%s' does not exist in the cloned repository", ai.config.SourceFolder)
				return false
			}

			// Verify it has MDC files
			hasMDCFiles, _ := utils.HasMDCFiles(sourceFolderPath)
			if !hasMDCFiles {
				ui.Warning("Source folder '%s' exists but contains no agent definitions", ai.config.SourceFolder)
			}
		}

		if utils.IsVerbose() {
			fileCount, _ := utils.CountFiles(ai.agentPath)
			utils.Infof("Repository cloned successfully with %d files", fileCount)

			if ai.config.SourceFolder != "" {
				sourceFolderPath := filepath.Join(ai.agentPath, ai.config.SourceFolder)
				sourceFileCount, _ := utils.CountFiles(sourceFolderPath)
				utils.Infof("Source folder '%s' contains %d files", ai.config.SourceFolder, sourceFileCount)
			}
		}

		if utils.IsDebug() {
			if ai.config.SourceFolder != "" {
				sourceFolderPath := filepath.Join(ai.agentPath, ai.config.SourceFolder)
				mdcFiles, _ := utils.CountFilesByExt(sourceFolderPath, ".mdc")
				utils.Debugf("Source folder definition details | path=%s | total files=%d | mdc files=%d",
					sourceFolderPath, utils.CountFilesRecursive(sourceFolderPath), mdcFiles)
			} else {
				mdcFiles, _ := utils.CountFilesByExt(ai.agentPath, ".mdc")
				utils.Debugf("Agent definition details | total files=%d | mdc files=%d",
					utils.CountFilesRecursive(ai.agentPath), mdcFiles)
			}
		}
	}

	ui.Success("Successfully cloned repository to %s", ai.agentPath)
	return true
}

func (ai *AgentInitializer) readUserInput(input *string) error {
	if _, err := fmt.Scanln(input); err != nil {
		return wrapOpError("readUserInput", "stdin", err, "failed to read user input")
	}
	return nil
}

func (ai *AgentInitializer) createEmptyDirectory() error {
	if err := os.MkdirAll(ai.agentPath, ai.config.DirPermission); err != nil {
		return wrapOpError("createEmptyDirectory", ai.agentPath, err, "failed to create agent directory")
	}
	return nil
}

func (ai *AgentInitializer) cloneRepository(repoURL string) error {
	ui.Info("Cloning repository %s to %s...", repoURL, ai.agentPath)

	if err := os.MkdirAll(filepath.Dir(ai.agentPath), ai.config.DirPermission); err != nil {
		return wrapOpError("cloneRepository", ai.agentPath, err, "failed to create parent directory")
	}

	if err := ai.gitMgr.CloneOrPull(context.Background(), repoURL, ai.agentPath); err != nil {
		return wrapOpError("cloneRepository", repoURL, err, "failed to clone repository")
	}

	return nil
}

// GetRegistry returns the registry
func (ai *AgentInitializer) GetRegistry() *Registry {
	return ai.registry
}
