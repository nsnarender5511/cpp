package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"vibe/internal/agent"
	"vibe/internal/core"
	"vibe/internal/ui"
	"vibe/internal/utils"
)


func HandleAgentCmd(initializer *core.AgentInitializer, appPaths utils.AppPaths, verbose bool, args []string) {
	utils.Debug("Handling agent command")
	configManager := utils.NewConfigManager()
	if err := configManager.Load(); err != nil {
		HandleCommandError("Agent", fmt.Errorf("failed to load configuration: %w", err), ExitConfigError)
		return
	}
	config := configManager.GetConfig()

	
	chosenDir, err := FindValidRulesDir(config.RulesDirName, config.AgentsDirName)
	if err != nil {
		ui.Warning("Error finding rules directory: %v", err)
		ui.Plain("Run %s to initialize the agent system in this directory or ensure system agents are installed.", ui.SuccessStyle.Sprint("vibe init"))
		return
	}
	if chosenDir == "" {
		ui.Warning("No local or system agent definitions found.")
		ui.Plain("Run %s to initialize the agent system in this directory.", ui.SuccessStyle.Sprint("vibe init"))
		return
	}

	utils.Debug("Using agent rules directory: " + chosenDir)
	agentRegistry, err := agent.NewRegistry(config, chosenDir)
	if err != nil {
		HandleCommandError("Agent", err, ExitAgentError)
		return
	}

	var filteredArgs []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if len(filteredArgs) < 1 {
		displayAgentListCmd(agentRegistry)
		return
	}

	subCommand := filteredArgs[0]
	utils.Info("Executing agent sub-command | sub_command=" + subCommand)

	switch subCommand {
	case "list":
		displayAgentListCmd(agentRegistry)
	case "select":
		handleAgentSelectCmd(agentRegistry, config, appPaths)
	case "info":
		if len(filteredArgs) < 2 {
			ui.Error("Missing agent ID/index. Usage: vibe agent info <agent-id | agent-index>")
			PrintAgentUsage(agentRegistry)
			os.Exit(ExitUsageError)
		}
		handleAgentInfoCmd(agentRegistry, filteredArgs[1], verbose)
	case "run": 
		if len(filteredArgs) < 2 {
			ui.Error("Missing agent ID/name. Usage: vibe agent run <agent-id | agent-name>")
			PrintAgentUsage(agentRegistry)
			os.Exit(ExitUsageError)
		}
		
		
		ui.Info("Executing agent: %s (run functionality is conceptual in this refactor)", filteredArgs[1])
		handleAgentRunCmd(agentRegistry, config, appPaths, filteredArgs[1])
	case "help", "--help", "-h":
		PrintAgentUsage(agentRegistry)
	default:
		ui.Warning("Unknown agent sub-command: %s", subCommand)
		PrintAgentUsage(agentRegistry)
		os.Exit(ExitUsageError)
	}
}


func displayAgentListCmd(registry *agent.Registry) {
	utils.Debug("Displaying agent list | registry path: " + registry.GetRulesDir())
	agents := registry.ListAgents()
	count := len(agents)

	if count == 0 {
		ui.Header("Agent List")
		ui.Warning("No agents found in the current directory or system path.")
		ui.Plain("\nRun %s to initialize this project with agents, or check system installation.", ui.SuccessStyle.Sprint("vibe init"))
		return
	}

	termWidth := ui.GetTerminalWidth()
	configManager := utils.NewConfigManager()
	if err := configManager.Load(); err != nil {
		utils.Warn("Failed to load configuration for agent list: " + err.Error())
	}
	config := configManager.GetConfig()
	lastSelectedAgent := config.LastSelectedAgent

	options := ui.DefaultAgentDisplayOptions()
	options.TermWidth = termWidth
	options.GroupByCategory = termWidth > 100
	options.SelectedAgentID = lastSelectedAgent

	if err := ui.DisplayAgentListEnhanced(agents, options); err != nil {
		HandleCommandError("AgentListDisplay", err, ExitAgentError)
	}
}


func handleAgentSelectCmd(registry *agent.Registry, config *utils.Config, appPaths utils.AppPaths) {
	utils.Debug("Handling agent select command")
	agents := registry.ListAgents()
	if len(agents) == 0 {
		ui.Warning("No agents available to select.")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	go func() {
		<-interruptCh
		utils.Info("User interrupted agent selection")
		cancel()
	}()

	selector := ui.NewAgentSelector(agents)
	selectedAgent, selectErr := selector.RunWithContext(ctx)

	if selectErr != nil {
		HandleCommandError("AgentSelect", selectErr, ExitAgentError)
		return
	}

	loader := agent.NewLoader(registry, config)
	agentCtx := agent.NewAgentContext() 
	loadedAgent, err := loader.LoadAgentWithContextCancellation(ctx, selectedAgent.ID, agentCtx)
	if err != nil {
		HandleCommandError("AgentLoad", err, ExitAgentError)
		return
	}

	utils.Debug("Agent loaded with context | last_updated=" + loadedAgent.Context.GetLastUpdated().String())

	config.LastSelectedAgent = selectedAgent.ID
	configManager := utils.NewConfigManager()
	configManager.SetConfig(config)
	if err := configManager.Save(); err != nil {
		utils.Warn("Failed to save last selected agent: " + err.Error())
	} else {
		utils.Info("Last selected agent saved | agent_id=" + selectedAgent.ID)
	}

	fmt.Println()
	ui.Success("Agent '%s' selected and loaded successfully!", selectedAgent.Name)

	ui.Prompt("View agent details? (y/n): ")
	var input string
	_, err = fmt.Scanln(&input)

	if err == nil && (input == "y" || input == "Y") {
		if errDisplay := ui.DisplayAgentInfoEnhanced(selectedAgent, true); errDisplay != nil {
			HandleCommandError("AgentInfoDisplay", errDisplay, ExitAgentError)
		}
	} else {
		if errDisplay := ui.DisplayAgentInfoEnhanced(selectedAgent, false); errDisplay != nil {
			HandleCommandError("AgentInfoDisplay", errDisplay, ExitAgentError)
		}
	}
	utils.Info("Agent select completed successfully | selected_agent=" + selectedAgent.ID)
}


func handleAgentInfoCmd(registry *agent.Registry, agentParam string, verbose bool) {
	utils.Debugf("Handling agent info subcommand | agent_param=%s verbose=%t", agentParam, verbose)
	var agentDef *agent.AgentDefinition
	var err error

	agents := registry.ListAgents()
	if len(agents) == 0 {
		ui.Error("No agents available in the current directory or system path.")
		ui.Plain("\nTry running %s first to initialize the agent system.", ui.SuccessStyle.Sprint("vibe init"))
		return
	}

	if index, convErr := strconv.Atoi(agentParam); convErr == nil {
		if index <= 0 || index > len(agents) {
			ui.Error("Invalid agent index: %d. Use a number between 1 and %d.", index, len(agents))
			os.Exit(ExitAgentError)
		}
		agentDef = agents[index-1]
	} else {
		agentDef, err = registry.GetAgent(agentParam) 
		if err != nil {
			
			foundByName := false
			cleanedParam := cleanAgentNameCmd(agentParam) 
			for _, ag := range agents {
				
				if strings.EqualFold(cleanAgentNameCmd(ag.Name), cleanedParam) || strings.EqualFold(ag.ID, agentParam) {
					agentDef = ag
					foundByName = true
					break
				}
			}
			if !foundByName {
				ui.Error("Agent '%s' not found by ID, name, or index.", agentParam)
				PrintAgentUsage(registry)
				os.Exit(ExitAgentError)
				return
			}
		}
	}

	ui.Header("Agent: %s", agentDef.Name)
	if err := ui.DisplayAgentInfoEnhanced(agentDef, verbose); err != nil {
		HandleCommandError("AgentInfoDisplay", err, ExitAgentError)
	}
}




func handleAgentRunCmd(registry *agent.Registry, config *utils.Config, appPaths utils.AppPaths, agentNameOrID string) {
	utils.Debugf("Attempting to run agent: %s", agentNameOrID)

	var agentDef *agent.AgentDefinition
	var err error
	agents := registry.ListAgents()
	if len(agents) == 0 {
		ui.Error("No agents available to run.")
		return
	}

	
	agentDef, err = registry.GetAgent(agentNameOrID)
	if err != nil {
		foundByNameOrID := false
		cleanedParam := cleanAgentNameCmd(agentNameOrID) 
		for _, ag := range agents {
			
			if strings.EqualFold(cleanAgentNameCmd(ag.Name), cleanedParam) || strings.EqualFold(ag.ID, agentNameOrID) {
				agentDef = ag
				foundByNameOrID = true
				break
			}
		}
		if !foundByNameOrID {
			ui.Error("Agent '%s' not found to run.", agentNameOrID)
			PrintAgentUsage(registry)
			return
		}
	}

	ui.Info("Found agent: %s (ID: %s)", agentDef.Name, agentDef.ID)
	ui.Info("Simulating agent run... (Actual run logic would be implemented here)")

	
	
	
	
	
	
	
	
	
	
	
	
	

	
	ui.Success("Conceptual run of agent '%s' completed.", agentDef.Name)

	
	config.LastSelectedAgent = agentDef.ID
	configManager := utils.NewConfigManager()
	configManager.SetConfig(config)
	if errSave := configManager.Save(); errSave != nil {
		utils.Warn("Failed to save last selected agent after run: " + errSave.Error())
	}
}
