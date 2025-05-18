package cli

import (
	"flag" 
	"fmt"
	"os"
	"strings"

	"vibe/internal/core"
	"vibe/internal/ui" 
	"vibe/internal/utils"
)


func DispatchCommand(initializer *core.AgentInitializer, appPaths utils.AppPaths) {
	args := flag.Args() 

	if len(args) < 1 {
		utils.Debug("No command provided in DispatchCommand, showing usage") 
		PrintUsage()
		os.Exit(ExitUsageError)
		return
	}

	command := args[0]
	commandArgs := args[1:]

	utils.Info("Executing command | command=" + command)

	switch command {
	case "init":
		HandleInitCmd(initializer)
	case "agent":
		
		
		debugMode := DebugFlag != nil && *DebugFlag
		HandleAgentCmd(initializer, appPaths, debugMode, commandArgs)
	case "clean":
		HandleCleanCmd(commandArgs)
	default:
		utils.Warn("Unknown command received | command=" + command)
		ui.Warning("Unknown command: %s", command)
		PrintUsage()
		os.Exit(ExitUsageError)
	}
	utils.Info("Command completed successfully | command=" + command)
}


func HandleCleanCmd(args []string) {
	if len(args) < 1 {
		ui.Error("Missing path argument for 'clean' command.")
		PrintUsage() 
		os.Exit(ExitUsageError)
		return
	}
	targetPath := args[0]

	info, err := os.Stat(targetPath)
	if err != nil {
		HandleCommandError("Clean", fmt.Errorf("failed to access path %s: %w", targetPath, err), ExitUsageError)
		return
	}

	if info.IsDir() {
		ui.Info(fmt.Sprintf("Cleaning comments from Go files in directory: %s", targetPath))
		if err := cleanDirectory(targetPath); err != nil {
			HandleCommandError("CleanDirectory", err, ExitUsageError)
		}
		return
	}

	if strings.HasSuffix(info.Name(), ".go") {
		ui.Info(fmt.Sprintf("Cleaning comments from file: %s", targetPath))
		if err := cleanFile(targetPath); err != nil {
			HandleCommandError("CleanFile", err, ExitUsageError)
		}
		return
	}

	HandleCommandError("Clean", fmt.Errorf("specified path %s is not a Go file or directory", targetPath), ExitUsageError)
}
