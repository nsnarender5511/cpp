package main

import (
	"fmt"
	"os"

	"crules/internal/core"
	"crules/internal/ui"
	"crules/internal/utils"
)

// Exit codes
const (
	ExitSuccess    = 0
	ExitUsageError = 1
	ExitInitError  = 10
	ExitMergeError = 11
	ExitSyncError  = 12
	ExitListError  = 13
	ExitCleanError = 14
	ExitSetupError = 20
)

func main() {
	// Parse debug/verbose flags
	debugFlag := false
	verboseFlag := false

	// Process each argument to find flags
	args := make([]string, 0, len(os.Args))
	for _, arg := range os.Args {
		if arg == "--debug" {
			debugFlag = true
			verboseFlag = true // Debug implies verbose
			continue
		} else if arg == "--verbose" {
			verboseFlag = true
			continue
		}
		// Keep non-flag arguments
		args = append(args, arg)
	}

	// Replace os.Args with filtered args
	os.Args = args

	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths)

	// Configure console output
	utils.SetDebugConsole(debugFlag)
	utils.SetVerbose(verboseFlag)

	utils.Info("Starting crules")

	// Print usage if no arguments
	if len(os.Args) < 2 {
		utils.Debug("No command provided, showing usage")
		printUsage()
		os.Exit(ExitUsageError)
	}

	// Display banner for all commands
	ui.PrintBanner()

	// Create new sync manager
	utils.Debug("Initializing sync manager")
	manager, err := core.NewSyncManager()
	if err != nil {
		utils.Error("Error initializing: " + err.Error())
		ui.Error("Error initializing: %v", err)
		os.Exit(ExitSetupError)
	}

	// Handle commands
	command := os.Args[1]
	utils.Info("Executing command | command=" + command)

	switch command {
	case "init":
		handleInit(manager)
	case "merge":
		handleMerge(manager)
	case "sync":
		handleSync(manager)
	case "list":
		handleList(manager)
	case "clean":
		handleClean(manager)
	default:
		utils.Warn("Unknown command received | command=" + command)
		ui.Warning("Unknown command: %s", command)
		printUsage()
		os.Exit(ExitUsageError)
	}

	utils.Info("Command completed successfully | command=" + command)
}

// handleCommandError handles command errors consistently
func handleCommandError(commandName string, err error, exitCode int) {
	utils.Error(commandName + " failed: " + err.Error())
	ui.Error("%s failed: %v", commandName, err)
	os.Exit(exitCode)
}

func printUsage() {
	ui.Header("Usage: crules [OPTIONS] <command>")

	ui.Plain("\nOptions:")
	ui.Plain("  --verbose    Show informational messages on console")
	ui.Plain("  --debug      Show debug messages on console")

	ui.Plain("\nCommands:")
	ui.Plain("  init         Initialize current directory with rules from main location")
	ui.Plain("  merge        Merge current rules to main location and sync to all locations")
	ui.Plain("  sync         Force sync from main location to current directory")
	ui.Plain("  list         Display all registered projects")
	ui.Plain("  clean        Remove non-existent projects from registry")
}

func handleInit(manager *core.SyncManager) {
	utils.Debug("Handling init command")
	if err := manager.Init(); err != nil {
		handleCommandError("Init", err, ExitInitError)
	}
	utils.Info("Init command completed successfully")
	ui.Success("Successfully initialized rules from main location")
}

func handleMerge(manager *core.SyncManager) {
	utils.Debug("Handling merge command")
	if err := manager.Merge(); err != nil {
		handleCommandError("Merge", err, ExitMergeError)
	}
	utils.Info("Merge command completed successfully")
	ui.Success("Successfully merged rules to main location")
}

func handleSync(manager *core.SyncManager) {
	utils.Debug("Handling sync command")
	if err := manager.Sync(); err != nil {
		handleCommandError("Sync", err, ExitSyncError)
	}
	utils.Info("Sync command completed successfully")
	ui.Success("Successfully synced rules from main location")
}

func handleList(manager *core.SyncManager) {
	utils.Debug("Handling list command")

	projects := manager.GetRegistry().GetProjects()
	count := len(projects)

	if count == 0 {
		ui.Info("No projects registered.")
		utils.Info("List command completed - no projects found")
		return
	}

	ui.Header("Registered projects (%d):", count)
	validCount := 0

	for i, project := range projects {
		exists := utils.DirExists(project)
		if exists {
			validCount++
			ui.Plain("  %d. %s", i+1, project)
		} else {
			ui.Plain("  %d. %s %s", i+1, project, ui.ErrorStyle.Sprint("(not found)"))
		}
	}

	if validCount < count {
		ui.Warning("\n%d project(s) could not be found. Run 'crules clean' to remove them.", count-validCount)
	}

	utils.Info("List command completed successfully | project_count=" + fmt.Sprintf("%d", count) + ", valid_count=" + fmt.Sprintf("%d", validCount))
}

func handleClean(manager *core.SyncManager) {
	utils.Debug("Handling clean command")

	removedCount, err := manager.Clean()
	if err != nil {
		handleCommandError("Clean", err, ExitCleanError)
	}

	if removedCount == 0 {
		ui.Info("No projects needed to be removed.")
	} else {
		ui.Success("Successfully removed %d non-existent project(s) from registry.", removedCount)
	}

	utils.Info("Clean command completed successfully | removed_count=" + fmt.Sprintf("%d", removedCount))
}
