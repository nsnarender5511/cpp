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
	cli.SetupFlags() 
	cli.ParseFlags() 

	
	isDebugSet := cli.DebugFlag != nil && *cli.DebugFlag

	utils.SetDebugConsole(isDebugSet)
	utils.SetVerbose(isDebugSet)       
	utils.SetVerboseErrors(isDebugSet) 

	
	if *cli.VersionShortFlag { 
		fmt.Printf("vibe version %s\n", version.GetVersion())
		os.Exit(0)
	}

	
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths) 

	utils.Info("Starting vibe")

	
	args := flag.Args()
	if len(args) < 1 {
		utils.Debug("No command provided, showing usage")
		cli.PrintUsage()            
		os.Exit(cli.ExitUsageError) 
	}

	
	ui.PrintBanner()

	
	utils.Debug("Initializing agent system")
	initializer, err := core.NewAgentInitializer()
	if err != nil {
		utils.Error("Error initializing agent system: " + err.Error())
		ui.Error("Error initializing agent system: %v", err) 
		os.Exit(cli.ExitSetupError)                          
	}

	
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
		cm.SetConfig(config)              
		if err := cm.Save(); err != nil { 
			utils.Warn("Failed to save multi-agent configuration: " + err.Error())
		} else {
			utils.Info("Multi-agent mode permanently enabled")
		}
	}

	
	
	
	cli.DispatchCommand(initializer, appPaths)
}
