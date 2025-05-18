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


const (
	ExitSuccess     = 0
	ExitUsageError  = 1
	ExitInitError   = 10
	ExitAgentError  = 15
	ExitSetupError  = 20
	ExitConfigError = 25
)


func HandleCommandError(commandName string, err error, exitCode int) {
	errMsg := err.Error()
	
	verboseErrors := utils.IsVerboseErrors()

	
	utils.Error(commandName + " failed: " + errMsg)

	
	if verboseErrors {
		
		detailedErr := errMsg
		if stackTracer, ok := err.(interface{ StackTrace() string }); ok {
			detailedErr += "\n" + stackTracer.StackTrace()
		}
		ui.Error("%s failed with details:\n%s", commandName, detailedErr)
	} else {
		
		ui.Error("%s failed: %v", commandName, err)
	}
	os.Exit(exitCode)
}



func FindValidRulesDir(rulesDirName string, agentsDirName string) (string, error) {
	
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current directory: %w", err)
	}

	projectRulesBaseDir := filepath.Join(currentDir, rulesDirName)           
	projectAgentsSubDir := filepath.Join(projectRulesBaseDir, agentsDirName) 

	
	if utils.DirExists(projectAgentsSubDir) {
		hasMDC, _ := utils.HasMDCFiles(projectAgentsSubDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in project subdirectory: " + projectAgentsSubDir)
			return projectAgentsSubDir, nil
		}
	}

	
	if utils.DirExists(projectRulesBaseDir) {
		hasMDC, _ := utils.HasMDCFiles(projectRulesBaseDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in project base directory: " + projectRulesBaseDir)
			return projectRulesBaseDir, nil
		}
	}

	
	
	
	appPaths := utils.GetAppPaths(utils.DefaultAgentsDirName)              
	systemRulesBaseDir := appPaths.GetRulesDir(rulesDirName)               
	systemAgentsSubDir := filepath.Join(systemRulesBaseDir, agentsDirName) 

	
	if utils.DirExists(systemAgentsSubDir) {
		hasMDC, _ := utils.HasMDCFiles(systemAgentsSubDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in system subdirectory: " + systemAgentsSubDir)
			return systemAgentsSubDir, nil
		}
	}

	
	if utils.DirExists(systemRulesBaseDir) {
		hasMDC, _ := utils.HasMDCFiles(systemRulesBaseDir)
		if hasMDC {
			utils.Debug("Found valid agent rules in system base directory: " + systemRulesBaseDir)
			return systemRulesBaseDir, nil
		}
	}

	utils.Debug("No valid agent rules directory found in local project or system paths.")
	return "", nil 
}


func cleanAgentNameCmd(name string) string {
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	caser := cases.Title(language.Und, cases.NoLower)
	return caser.String(name)
}
