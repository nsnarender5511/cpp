package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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

	targetPath := filepath.Join(currentDir, ai.config.RulesDirName)
	utils.Debug("Init target path | path=" + targetPath)

	// Check if the agent location exists and has valid definitions
	needsSetup := false

	if !utils.DirExists(ai.agentPath) {
		utils.Debug("Agent location does not exist | path=" + ai.agentPath)
		ui.Warning("Agent location does not exist: %s", ai.agentPath)
		needsSetup = true
	} else {
		hasMDCFiles, err := utils.HasMDCFiles(ai.agentPath)
		if err != nil {
			return wrapOpError("Init", ai.agentPath, err, "failed to check for agent definitions")
		}

		if !hasMDCFiles {
			utils.Debug("Agent location exists but contains no definitions | path=" + ai.agentPath)
			ui.Warning("Agent location exists but contains no definitions: %s", ai.agentPath)
			needsSetup = true
		}
	}

	if needsSetup {
		if !ai.handleInitialSetup() {
			return wrapValidationError("setup", "agent system initialization cancelled")
		}
	}

	// Copy agent definitions to project
	if err := utils.CopyDir(ai.agentPath, targetPath); err != nil {
		return wrapOpError("Init", targetPath, err, "failed to copy agent definitions")
	}

	// Add project to registry
	if err := ai.registry.AddProject(currentDir); err != nil {
		return wrapOpError("Init", currentDir, err, "failed to register project")
	}

	ui.Success("Successfully initialized agent system in %s", currentDir)
	return nil
}

// handleInitialSetup manages the initial setup of the agent system
func (ai *AgentInitializer) handleInitialSetup() bool {
	ui.Info("\nNo agent definitions found. Please choose an option:")
	ui.Plain("1. Create empty agent directory")
	ui.Plain("2. Clone from git repository")
	ui.Plain("3. Cancel")
	ui.Prompt("\nEnter choice (1-3): ")

	var choice string
	if err := ai.readUserInput(&choice); err != nil {
		ui.Error(err.Error())
		return false
	}

	switch choice {
	case "1":
		if err := ai.createEmptyDirectory(); err != nil {
			ui.Error(err.Error())
			return false
		}
		ui.Success("Created empty agent directory at %s", ai.agentPath)
		return true

	case "2":
		var repoURL string
		ui.Prompt("Enter git repository URL: ")

		if err := ai.readUserInput(&repoURL); err != nil {
			ui.Error(err.Error())
			return false
		}

		if repoURL == "" {
			ui.Error("Repository URL cannot be empty")
			return false
		}

		if err := ai.cloneRepository(repoURL); err != nil {
			ui.Error(err.Error())
			return false
		}

		ui.Success("Successfully cloned repository to %s", ai.agentPath)
		return true

	case "3":
		ui.Info("Operation cancelled by user")
		return false

	default:
		ui.Error("Invalid choice")
		return false
	}
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
