package core

import (
	"fmt"
	"os"
	"path/filepath"

	"crules/internal/utils"
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
	}

	if err := utils.EnsureDirExists(appPaths.DataDir, config.DirPermission); err != nil {
		utils.Error("Cannot create data directory | path=" + appPaths.DataDir + ", error=" + err.Error())
		return nil, fmt.Errorf("cannot create data directory: %v", err)
	}

	if err := utils.EnsureDirExists(appPaths.LogDir, config.DirPermission); err != nil {
		utils.Error("Cannot create log directory | path=" + appPaths.LogDir + ", error=" + err.Error())
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

	// Check if target already exists
	if utils.DirExists(targetPath) {
		utils.Debug("Rules directory already exists | path=" + targetPath)
		if !utils.ConfirmOverwrite(sm.config.RulesDirName) {
			utils.Info("Operation cancelled by user")
			return fmt.Errorf("operation cancelled by user")
		}
	}

	// Copy from main to current
	utils.Debug("Copying rules to current directory | source=" + sm.mainPath + ", target=" + targetPath)
	if err := utils.CopyDir(sm.mainPath, targetPath); err != nil {
		utils.Error("Failed to copy rules | source=" + sm.mainPath + ", target=" + targetPath + ", error=" + err.Error())
		return fmt.Errorf("failed to copy rules: %v", err)
	}

	// Register this project
	utils.Debug("Registering project | project=" + currentDir)
	if err := sm.registry.AddProject(currentDir); err != nil {
		utils.Error("Failed to register project | project=" + currentDir + ", error=" + err.Error())
		return err
	}

	utils.Info("Rules initialized successfully | project=" + currentDir)
	return nil
}

// Merge copies current rules to main and syncs to all locations
func (sm *SyncManager) Merge() error {
	currentDir, err := os.Getwd()
	if err != nil {
		utils.Error("Cannot get current directory | error=" + err.Error())
		return fmt.Errorf("cannot get current directory: %v", err)
	}

	sourcePath := filepath.Join(currentDir, sm.config.RulesDirName)
	utils.Debug("Checking for rules in current directory | path=" + sourcePath)
	if !utils.DirExists(sourcePath) {
		utils.Error("Rules not found in current directory | path=" + sourcePath)
		return fmt.Errorf("%s not found in current directory", sm.config.RulesDirName)
	}

	// Copy to main
	utils.Debug("Copying rules to main location | source=" + sourcePath + ", target=" + sm.mainPath)
	if err := utils.CopyDir(sourcePath, sm.mainPath); err != nil {
		utils.Error("Failed to copy to main | source=" + sourcePath + ", target=" + sm.mainPath + ", error=" + err.Error())
		return fmt.Errorf("failed to copy to main: %v", err)
	}
	utils.Info("Rules merged to main location | source=" + sourcePath)

	// Sync to all registered projects
	utils.Debug("Starting sync to all registered projects")
	return sm.syncToAll()
}

// Sync forces sync from main to current
func (sm *SyncManager) Sync() error {
	currentDir, err := os.Getwd()
	if err != nil {
		utils.Error("Cannot get current directory | error=" + err.Error())
		return fmt.Errorf("cannot get current directory: %v", err)
	}

	targetPath := filepath.Join(currentDir, sm.config.RulesDirName)
	utils.Debug("Syncing rules from main location | source=" + sm.mainPath + ", target=" + targetPath)

	if err := utils.CopyDir(sm.mainPath, targetPath); err != nil {
		utils.Error("Failed to sync rules | source=" + sm.mainPath + ", target=" + targetPath + ", error=" + err.Error())
		return err
	}

	utils.Info("Rules synced successfully | target=" + targetPath)
	return nil
}

// syncToAll syncs main rules to all registered projects
func (sm *SyncManager) syncToAll() error {
	projects := sm.registry.GetProjects()
	utils.Debug("Syncing to all projects | count=" + fmt.Sprintf("%d", len(projects)))

	succeeded := 0
	failed := 0

	for _, project := range projects {
		targetPath := filepath.Join(project, sm.config.RulesDirName)
		utils.Debug("Syncing to project | project=" + project + ", target=" + targetPath)

		if err := utils.CopyDir(sm.mainPath, targetPath); err != nil {
			utils.Warn("Failed to sync to project | project=" + project + ", error=" + err.Error())
			fmt.Printf("Warning: failed to sync to %s: %v\n", project, err)
			failed++
		} else {
			succeeded++
		}
	}

	utils.Info("Sync to all projects completed | successful=" + fmt.Sprintf("%d", succeeded) + ", failed=" + fmt.Sprintf("%d", failed))
	return nil
}

// GetRegistry returns the registry instance
func (sm *SyncManager) GetRegistry() *Registry {
	return sm.registry
}
