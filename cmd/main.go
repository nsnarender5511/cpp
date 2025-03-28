package main

import (
	"fmt"
	"os"

	"crules/internal/core"
	"crules/internal/utils"
)

func main() {
	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths)

	utils.Info("Starting crules")

	// Print usage if no arguments
	if len(os.Args) < 2 {
		utils.Debug("No command provided, showing usage")
		printUsage()
		os.Exit(1)
	}

	// Create new sync manager
	utils.Debug("Initializing sync manager")
	manager, err := core.NewSyncManager()
	if err != nil {
		utils.Error("Error initializing: " + err.Error())
		fmt.Printf("Error initializing: %v\n", err)
		os.Exit(1)
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
	default:
		utils.Warn("Unknown command received | command=" + command)
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}

	utils.Info("Command completed successfully | command=" + command)
}

func printUsage() {
	fmt.Println("Usage: crules <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  init     Initialize current directory with rules from main location")
	fmt.Println("  merge    Merge current rules to main location and sync to all locations")
	fmt.Println("  sync     Force sync from main location to current directory")
}

func handleInit(manager *core.SyncManager) {
	utils.Debug("Handling init command")
	if err := manager.Init(); err != nil {
		utils.Error("Init failed: " + err.Error())
		fmt.Printf("Init failed: %v\n", err)
		os.Exit(1)
	}
	utils.Info("Init command completed successfully")
	fmt.Println("Successfully initialized rules from main location")
}

func handleMerge(manager *core.SyncManager) {
	utils.Debug("Handling merge command")
	if err := manager.Merge(); err != nil {
		utils.Error("Merge failed: " + err.Error())
		fmt.Printf("Merge failed: %v\n", err)
		os.Exit(1)
	}
	utils.Info("Merge command completed successfully")
	fmt.Println("Successfully merged rules to main location")
}

func handleSync(manager *core.SyncManager) {
	utils.Debug("Handling sync command")
	if err := manager.Sync(); err != nil {
		utils.Error("Sync failed: " + err.Error())
		fmt.Printf("Sync failed: %v\n", err)
		os.Exit(1)
	}
	utils.Info("Sync command completed successfully")
	fmt.Println("Successfully synced rules from main location")
}
