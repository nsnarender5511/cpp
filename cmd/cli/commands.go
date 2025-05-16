package cli

import (
	"flag" // Required for flag.Args()
	"os"

	"vibe/internal/core"
	"vibe/internal/ui" // For ui.Warning
	"vibe/internal/utils"
)

// DispatchCommand routes the command to the appropriate handler.
func DispatchCommand(initializer *core.AgentInitializer, appPaths utils.AppPaths) {
	args := flag.Args() // Get non-flag arguments

	if len(args) < 1 {
		utils.Debug("No command provided in DispatchCommand, showing usage") // Should have been caught in main
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
		// The 'verbose' argument to HandleAgentCmd was used to control detail in 'agent info'.
		// Since debug now implies full verbosity, pass the state of DebugFlag.
		debugMode := DebugFlag != nil && *DebugFlag
		HandleAgentCmd(initializer, appPaths, debugMode, commandArgs)
	default:
		utils.Warn("Unknown command received | command=" + command)
		ui.Warning("Unknown command: %s", command)
		PrintUsage()
		os.Exit(ExitUsageError)
	}
	utils.Info("Command completed successfully | command=" + command)
}
