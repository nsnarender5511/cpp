package core

import (
	"fmt"
	"os"
	"path/filepath"

	"cursor++/internal/agent"
	"cursor++/internal/git"
	"cursor++/internal/ui"
	"cursor++/internal/utils"
)

// SyncManager handles all sync operations
type SyncManager struct {
	mainPath string
	registry *Registry
	config   *utils.Config
	appPaths utils.AppPaths
}

// NewSyncManager creates a new sync manager
func NewSyncManager() (*SyncManager, error) {
	// Load config
	config := utils.LoadConfig()
	utils.Debug("Loaded configuration")

	// Get app name from environment
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = utils.DefaultAppName
	}

	// Get system paths
	appPaths := utils.GetAppPaths(appName)
	utils.Debug("Loaded system paths for current platform")

	// Use OS-specific paths
	mainPath := appPaths.GetRulesDir(config.RulesDirName)
	registryPath := appPaths.GetRegistryFile(config.RegistryFileName)

	// Ensure required directories exist
	if err := utils.EnsureDirExists(appPaths.ConfigDir, config.DirPermission); err != nil {
		utils.Error("Cannot create config directory | path=" + appPaths.ConfigDir + ", error=" + err.Error())
		return nil, fmt.Errorf("cannot create config directory: %v", err)
	}

	if err := utils.EnsureDirExists(appPaths.DataDir, config.DirPermission); err != nil {
		utils.Error("Cannot create data directory | path=" + appPaths.DataDir + ", error=" + err.Error())
		return nil, fmt.Errorf("cannot create data directory: %v", err)
	}

	if err := utils.EnsureDirExists(appPaths.LogDir, config.DirPermission); err != nil {
		utils.Error("Cannot create log directory | path=" + appPaths.LogDir + ", error=" + err.Error())
		return nil, fmt.Errorf("cannot create log directory: %v", err)
	}

	utils.Debug("Using main rules path | path=" + mainPath)

	// Ensure main directory exists
	if err := os.MkdirAll(mainPath, config.DirPermission); err != nil {
		utils.Error("Cannot create main directory | path=" + mainPath + ", error=" + err.Error())
		return nil, fmt.Errorf("cannot create main directory: %v", err)
	}

	// Load or create registry
	registry, err := LoadRegistry(registryPath, config)
	if err != nil {
		utils.Error("Cannot load registry | error=" + err.Error())
		return nil, fmt.Errorf("cannot load registry: %v", err)
	}
	utils.Debug("Registry loaded successfully")

	return &SyncManager{
		mainPath: mainPath,
		registry: registry,
		config:   config,
		appPaths: appPaths,
	}, nil
}

// Init copies rules from main location to current directory
func (sm *SyncManager) Init() error {
	currentDir, err := os.Getwd()
	if err != nil {
		utils.Error("Cannot get current directory | error=" + err.Error())
		return fmt.Errorf("cannot get current directory: %v", err)
	}

	targetPath := filepath.Join(currentDir, sm.config.RulesDirName)
	utils.Debug("Init target path | path=" + targetPath)

	// Check if the main rules location exists
	mainLocationNeedsSetup := false

	if !utils.DirExists(sm.mainPath) {
		utils.Debug("Main rules location does not exist | path=" + sm.mainPath)
		ui.Warning("Main rules location does not exist: %s", sm.mainPath)
		mainLocationNeedsSetup = true
	} else {
		// Main location exists, but check if it has any .mdc files
		hasMDCFiles, err := utils.HasMDCFiles(sm.mainPath)
		if err != nil {
			utils.Error("Failed to check for .mdc files | path=" + sm.mainPath + ", error=" + err.Error())
			return fmt.Errorf("failed to check for .mdc files: %v", err)
		}

		if !hasMDCFiles {
			utils.Debug("Main rules location exists but contains no .mdc files | path=" + sm.mainPath)
			ui.Warning("Main rules location exists but contains no rules: %s", sm.mainPath)
			mainLocationNeedsSetup = true
		}
	}

	// If main location doesn't exist or is empty, offer options
	if mainLocationNeedsSetup {
		if !sm.offerMainLocationOptions() {
			ui.Info("Operation cancelled by user")
			return fmt.Errorf("operation cancelled by user")
		}
	}

	// Check if target already exists and list its contents
	if utils.DirExists(targetPath) {
		utils.Debug("Rules directory already exists | path=" + targetPath)

		// List files that will be overwritten
		files, err := utils.ListDirectoryContents(targetPath)
		if err != nil {
			utils.Error("Failed to list directory contents | path=" + targetPath + ", error=" + err.Error())
			return fmt.Errorf("failed to list directory contents: %v", err)
		}

		if len(files) > 0 {
			ui.Header("The following files will be deleted:")
			ui.DisplayFileTable(files)

			ui.Plain("")
			if !ui.PromptYesNo("Do you want to continue and overwrite these files?") {
				ui.Info("Operation cancelled by user")
				return fmt.Errorf("operation cancelled by user")
			}
		} else {
			ui.Info("Destination directory exists but is empty")
		}
	}

	// Copy from main to current
	utils.Debug("Copying rules to current directory | source=" + sm.mainPath + ", target=" + targetPath)
	if err := utils.CopyDir(sm.mainPath, targetPath); err != nil {
		utils.Error("Failed to copy rules | source=" + sm.mainPath + ", target=" + targetPath + ", error=" + err.Error())
		return fmt.Errorf("failed to copy rules: %v", err)
	}

	// After copying, use the animation to show agent loading process
	animator := ui.NewAnimator()

	// Create the registry to identify agents
	agentRegistry, err := sm.createAgentRegistryWithAnimation(animator, targetPath)
	if err != nil {
		utils.Error("Failed to create agent registry | error=" + err.Error())
		return fmt.Errorf("failed to create agent registry: %v", err)
	}

	// Register this project
	utils.Debug("Registering project | project=" + currentDir)
	if err := sm.registry.AddProject(currentDir); err != nil {
		utils.Error("Failed to register project | project=" + currentDir + ", error=" + err.Error())
		return err
	}

	// Ensure .cursor/ is in gitignore
	utils.Debug("Ensuring .cursor/ is in gitignore | project=" + currentDir)
	if err := utils.EnsureGitIgnoreEntry(currentDir, ".cursor/"); err != nil {
		utils.Warn("Failed to update gitignore | project=" + currentDir + ", error=" + err.Error())
		// Don't fail the operation just for gitignore issues
		// Just log a warning
	} else {
		utils.Info("Updated gitignore to exclude .cursor/ directory")
	}

	// Get agent count for final summary
	agentCount := len(agentRegistry.ListAgents())

	utils.Info("Rules initialized successfully | project=" + currentDir)
	ui.Success("Successfully initialized rules in %s with %d agents", targetPath, agentCount)
	return nil
}

// createAgentRegistryWithAnimation creates a registry for the agents with animated progress
func (sm *SyncManager) createAgentRegistryWithAnimation(animator *ui.TerminalAnimator, targetPath string) (*agent.Registry, error) {
	// Get agent directory path - use the targetPath directly since we know it points to .cursor/rules
	agentDir := filepath.Join(targetPath, sm.config.AgentsDirName)
	utils.Debug("Looking for agents in | path=" + agentDir)

	// Start animation
	animator.StartAnimation("Initializing Agents")

	// Add a generic initialization step
	animator.AddItem("init", "Preparing agent system...")
	animator.UpdateStatus("init", "success")

	// Create a new registry
	registry, err := agent.NewRegistry(sm.config, agentDir)
	if err != nil {
		animator.UpdateStatus("init", "error")
		animator.StopAnimation("Failed to initialize agents")
		return nil, err
	}

	// Trigger a rescan to show the animation
	// This is safe because NewRegistry already did the initial scan
	agents := registry.ListAgents()
	count := len(agents)

	// Update animation with results
	animator.StopAnimation(fmt.Sprintf("Successfully initialized %d agents", count))

	return registry, nil
}

// offerMainLocationOptions presents options for an empty or non-existent main location
// Returns true if operation should continue, false if cancelled
func (sm *SyncManager) offerMainLocationOptions() bool {
	ui.Header("Main rules location is empty or does not exist")
	ui.Plain("Options:")
	ui.Plain("  1. Create empty rules directory")
	ui.Plain("  2. Clone from git repository")
	ui.Plain("  3. Cancel operation")

	ui.Prompt("Select an option: ")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		// Create empty directory
		utils.Debug("Creating empty rules directory | path=" + sm.mainPath)
		if err := os.MkdirAll(sm.mainPath, sm.config.DirPermission); err != nil {
			utils.Error("Failed to create rules directory | path=" + sm.mainPath + ", error=" + err.Error())
			ui.Error("Failed to create rules directory: %v", err)
			return false
		}

		// Create agents directory
		agentsDir := filepath.Join(sm.mainPath, sm.config.AgentsDirName)
		if err := os.MkdirAll(agentsDir, sm.config.DirPermission); err != nil {
			utils.Error("Failed to create agents directory | path=" + agentsDir + ", error=" + err.Error())
			ui.Error("Failed to create agents directory: %v", err)
			return false
		}

		ui.Success("Created empty rules directory at %s", sm.mainPath)
		return true

	case "2":
		// Clone from git repository
		ui.Prompt("Enter git repository URL: ")
		var repoURL string
		fmt.Scanln(&repoURL)

		if repoURL == "" {
			ui.Error("Repository URL cannot be empty")
			return false
		}

		ui.Info("Cloning repository %s to %s...", repoURL, sm.mainPath)

		// Ensure the parent directory exists
		if err := os.MkdirAll(filepath.Dir(sm.mainPath), sm.config.DirPermission); err != nil {
			utils.Error("Failed to create parent directory | path=" + filepath.Dir(sm.mainPath) + ", error=" + err.Error())
			ui.Error("Failed to create parent directory: %v", err)
			return false
		}

		if err := git.Clone(repoURL, sm.mainPath); err != nil {
			utils.Error("Failed to clone repository | repo=" + repoURL + ", path=" + sm.mainPath + ", error=" + err.Error())
			ui.Error("Failed to clone repository: %v", err)
			return false
		}

		ui.Success("Successfully cloned repository to %s", sm.mainPath)
		return true

	case "3":
		// Cancel operation
		ui.Info("Operation cancelled by user")
		return false

	default:
		ui.Error("Invalid choice")
		return false
	}
}

// GetRegistry returns the registry
func (sm *SyncManager) GetRegistry() *Registry {
	return sm.registry
}
