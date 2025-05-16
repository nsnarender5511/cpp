package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"vibe/internal/ui"
	"vibe/internal/utils"
)

// Exit codes
const (
	ExitSuccess     = 0
	ExitUsageError  = 1
	ExitInitError   = 10
	ExitAgentError  = 15
	ExitSetupError  = 20
	ExitConfigError = 25
)

// HandleCommandError handles command errors consistently
func HandleCommandError(commandName string, err error, exitCode int) {
	errMsg := err.Error()
	// Check if we should display detailed error information
	verboseErrors := utils.IsVerboseErrors()

	// Log the full error regardless of console output settings
	utils.Error(commandName + " failed: " + errMsg)

	// For UI display, use different formatting based on verbosity
	if verboseErrors {
		// Show detailed error with stack trace or additional context if available
		detailedErr := errMsg
		if stackTracer, ok := err.(interface{ StackTrace() string }); ok {
			detailedErr += "\n" + stackTracer.StackTrace()
		}
		ui.Error("%s failed with details:\n%s", commandName, detailedErr)
	} else {
		// Show simplified error message
		ui.Error("%s failed: %v", commandName, err)
	}
	os.Exit(exitCode)
}

// FindValidRulesDir tries to find a valid rules directory, checking local project paths first, then system paths.
// It returns the path to the first valid directory found, or an empty string if none are found.
func FindValidRulesDir(rulesDirName string, agentsDirName string) (string, error) {
	// 1. Check local project directory: .cursor/rules/agents (or configured names)
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current directory: %w", err)
	}

	projectRulesBaseDir := filepath.Join(currentDir, rulesDirName)           // e.g., ./ .cursor/rules
	projectAgentsSubDir := filepath.Join(projectRulesBaseDir, agentsDirName) // e.g., ./ .cursor/rules/vibe

	// Check project agents subdirectory first (e.g., ./.cursor/rules/vibe)
	if utils.DirExists(projectAgentsSubDir) {
		hasMDC, _ := utils.HasMDCFiles(projectAgentsSubDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in project subdirectory: " + projectAgentsSubDir)
			return projectAgentsSubDir, nil
		}
	}

	// Check project base rules directory (e.g., ./.cursor/rules - where init places them)
	if utils.DirExists(projectRulesBaseDir) {
		hasMDC, _ := utils.HasMDCFiles(projectRulesBaseDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in project base directory: " + projectRulesBaseDir)
			return projectRulesBaseDir, nil
		}
	}

	// 2. Check system-wide directory (conceptual - adapt to actual system path logic if it exists)
	// This part needs to align with how system-wide agents are actually stored if vibe supports it.
	// For now, let's assume AppPaths provides a way to get a system rules dir.
	appPaths := utils.GetAppPaths(utils.DefaultAgentsDirName)              // Use a default app name for system paths
	systemRulesBaseDir := appPaths.GetRulesDir(rulesDirName)               // e.g., ~/Library/Application Support/vibe/.cursor/rules
	systemAgentsSubDir := filepath.Join(systemRulesBaseDir, agentsDirName) // e.g., ~/Library/Application Support/vibe/.cursor/rules/vibe

	// Check system agents subdirectory
	if utils.DirExists(systemAgentsSubDir) {
		hasMDC, _ := utils.HasMDCFiles(systemAgentsSubDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in system subdirectory: " + systemAgentsSubDir)
			return systemAgentsSubDir, nil
		}
	}

	// Check system base rules directory
	if utils.DirExists(systemRulesBaseDir) {
		hasMDC, _ := utils.HasMDCFiles(systemRulesBaseDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in system base directory: " + systemRulesBaseDir)
			return systemRulesBaseDir, nil
		}
	}

	utils.Debug("No valid agent rules directory found in local project or system paths.")
	return "", nil // No valid directory found
}

// cleanAgentNameCmd cleans up the agent name for display.
func cleanAgentNameCmd(name string) string {
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	caser := cases.Title(language.Und, cases.NoLower)
	return caser.String(name)
}
