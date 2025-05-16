package main

import (
	"flag"
	"fmt"
	"os"

	"vibe/cmd/cli"
	"vibe/internal/core"
	"vibe/internal/ui"
	"vibe/internal/utils"
	"vibe/internal/version"
)

func main() {
	cli.SetupFlags() // Defines flags using variables in cli package
	cli.ParseFlags() // Parses flags

	// Configure console output based on DebugFlag
	isDebugSet := cli.DebugFlag != nil && *cli.DebugFlag

	utils.SetDebugConsole(isDebugSet)
	utils.SetVerbose(isDebugSet)       // Debug implies verbose
	utils.SetVerboseErrors(isDebugSet) // Debug implies verbose errors

	// Handle version flag
	if *cli.VersionShortFlag { // Changed from cli.VersionFlag || *cli.VersionShortFlag
		fmt.Printf("vibe version %s\n", version.GetVersion())
		os.Exit(0)
	}

	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths) // Assuming InitLogger doesn't depend on flag values directly before this point

	utils.Info("Starting vibe")

	// Check for command arguments. flag.Args() must be called after flag.Parse().
	args := flag.Args()
	if len(args) < 1 {
		utils.Debug("No command provided, showing usage")
		cli.PrintUsage()            // Use PrintUsage from cli package
		os.Exit(cli.ExitUsageError) // Use exit code from cli package
	}

	// Display banner for all commands
	ui.PrintBanner()

	// Create new agent initializer
	utils.Debug("Initializing agent system")
	initializer, err := core.NewAgentInitializer()
	if err != nil {
		utils.Error("Error initializing agent system: " + err.Error())
		ui.Error("Error initializing agent system: %v", err) // Using internal/ui for direct error
		os.Exit(cli.ExitSetupError)                          // Use exit code from cli package
	}

	// Set multi-agent mode if flag is provided
	if cli.MultiAgentModeFlag != nil && *cli.MultiAgentModeFlag {
		utils.Info("Multi-agent mode explicitly enabled via flag")
		cm := utils.NewConfigManager()
		if err := cm.Load(); err != nil {
			utils.Error("Failed to load configuration for multi-agent mode: " + err.Error())
			ui.Error("Failed to load configuration: %v", err)
			os.Exit(cli.ExitConfigError)
		}
		config := cm.GetConfig()
		config.MultiAgentEnabled = true
		cm.SetConfig(config)              // Set the modified config
		if err := cm.Save(); err != nil { // Then save it
			utils.Warn("Failed to save multi-agent configuration: " + err.Error())
		} else {
			utils.Info("Multi-agent mode permanently enabled")
		}
	}

	// Dispatch the command
	// DispatchCommand will use flag.Args() internally if it needs them again,
	// or they can be passed if preferred.
	cli.DispatchCommand(initializer, appPaths)
}
