package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"vibe/internal/agent"
	"vibe/internal/core"
	"vibe/internal/ui"
	"vibe/internal/utils"
)


func HandleInitCmd(initializer *core.AgentInitializer) {
	utils.Debug("Handling init command")

	
	fmt.Println()

	if err := initializer.Init(); err != nil {
		HandleCommandError("Init", err, ExitInitError) 
	}

	
	pattern := ".cursor"
	cursorignorePattern := ".cursorignore"
	currentDir, err := os.Getwd()
	if err == nil {
		gitignorePath := filepath.Join(currentDir, ".gitignore")
		cursorignorePath := filepath.Join(currentDir, ".cursorignore")

		
		if err := utils.EnsurePathInFile(gitignorePath, pattern); err != nil {
			utils.Warn("Failed to update .gitignore with .cursor: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursor is in .gitignore")
		}

		
		if err := utils.EnsurePathInFile(gitignorePath, cursorignorePattern); err != nil {
			utils.Warn("Failed to update .gitignore with .cursorignore: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursorignore is in .gitignore")
		}

		
		if err := utils.EnsurePathInFile(cursorignorePath, pattern); err != nil {
			utils.Warn("Failed to update .cursorignore: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursor is in .cursorignore")
		}
	} else {
		utils.Warn("Could not get current directory to update ignore files: " + err.Error())
	}

	
	fmt.Println()

	
	currentDir, err = os.Getwd()
	if err != nil {
		utils.Warn("Could not get current directory for post-init messages: " + err.Error())
	} else {
		
		configManager := utils.NewConfigManager()
		if err := configManager.Load(); err != nil {
			utils.Warn("Failed to load configuration for post-init messages: " + err.Error())
		} else {
			config := configManager.GetConfig()
			localRulesDir := filepath.Join(currentDir, config.RulesDirName, config.AgentsDirName)
			
			agentRegistry, regErr := agent.NewRegistry(config, localRulesDir)
			if regErr != nil {
				utils.Warn("Could not initialize agent registry to list agents after init: " + regErr.Error())
			} else if agentRegistry != nil {
				agents := agentRegistry.ListAgents()
				agentCount := len(agents)

				if agentCount > 0 {
					ui.Header("Available Agents in %s (%d):", currentDir, agentCount)
					
					
					for i := 0; i < 3 && i < len(agents); i++ {
						ui.Plain("• %s (%s)",
							ui.InfoStyle.Sprint(cleanAgentNameCmd(agents[i].Name)), 
							fmt.Sprintf("@%s.mdc", agents[i].ID))
					}
					if len(agents) > 3 {
						ui.Plain("• %s", ui.WarnStyle.Sprint("...and more"))
					}
					fmt.Println()
				}
			}
		}
	}

	ui.Header("Next Steps:")
	ui.Plain("1. Use %s to see available agents", ui.SuccessStyle.Sprint("vibe agent list"))
	ui.Plain("2. Start using agents by referencing them with %s in your chat", ui.SuccessStyle.Sprint("@agent-name.mdc"))
	ui.Plain("3. For more information about an agent, use %s", ui.SuccessStyle.Sprint("vibe agent info <agent-id>"))

	fmt.Println()
}


