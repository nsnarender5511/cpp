package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"

	"vibe/internal/agent"
	"vibe/internal/core"
	"vibe/internal/ui"
	"vibe/internal/utils"
	"vibe/internal/version"
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

// getTerminalWidth returns the width of the terminal in characters
func getTerminalWidth() int {
	// First check if COLUMNS environment variable is set (useful for testing)
	if colEnv := os.Getenv("COLUMNS"); colEnv != "" {
		if width, err := strconv.Atoi(colEnv); err == nil && width > 0 {
			return width
		}
	}

	// Try using the stty command to get terminal size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		parts := strings.Split(strings.TrimSpace(string(out)), " ")
		if len(parts) == 2 {
			width, err := strconv.Atoi(parts[1])
			if err == nil && width > 0 {
				return width
			}
		}
	}

	// If everything fails, return a default width
	return 80
}

func main() {
	// Setup command-line flags using the standard flag package
	debugFlag := flag.Bool("debug", false, "Show debug messages on console")
	verboseFlag := flag.Bool("verbose", false, "Show informational messages on console")
	multiAgentFlag := flag.Bool("multi-agent", false, "Enable multi-agent mode for this session")
	verboseErrorsFlag := flag.Bool("verbose-errors", false, "Display detailed error messages on failure")
	versionFlag := flag.Bool("version", false, "Show version information")

	// Add a short version flag
	versionShortFlag := flag.Bool("v", false, "Show version information (short flag)")

	// Custom usage function to match our existing help format
	flag.Usage = func() {
		printUsage()
	}

	// Parse the flags
	flag.Parse()

	// Debug implies verbose
	if *debugFlag {
		*verboseFlag = true
	}

	// Check for version flag
	if *versionFlag || *versionShortFlag {
		fmt.Printf("vibe version %s\n", version.GetVersion())
		os.Exit(ExitSuccess)
	}

	// Get remaining arguments after flag parsing
	args := flag.Args()

	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths)

	// Configure console output
	utils.SetDebugConsole(*debugFlag)
	utils.SetVerbose(*verboseFlag)
	utils.SetVerboseErrors(*verboseErrorsFlag)

	utils.Info("Starting vibe")

	// Print usage if no arguments
	if len(args) < 1 {
		utils.Debug("No command provided, showing usage")
		printUsage()
		os.Exit(ExitUsageError)
	}

	// Display banner for all commands
	ui.PrintBanner()

	// Create new sync manager
	utils.Debug("Initializing sync manager")
	initializer, err := core.NewAgentInitializer()
	if err != nil {
		utils.Error("Error initializing: " + err.Error())
		ui.Error("Error initializing: %v", err)
		os.Exit(ExitSetupError)
	}

	// Set multi-agent mode if flag is provided
	if *multiAgentFlag {
		utils.Info("Multi-agent mode explicitly enabled via flag")
		// Update the config to enable multi-agent mode
		cm := utils.NewConfigManager()
		if err := cm.Load(); err != nil {
			utils.Error("Failed to load configuration: " + err.Error())
			ui.Error("Failed to load configuration: %v", err)
			os.Exit(ExitConfigError)
		}
		config := cm.GetConfig()
		config.MultiAgentEnabled = true

		// Save configuration to make the change permanent
		cm.SetConfig(config)
		if err := cm.Save(); err != nil {
			utils.Warn("Failed to save multi-agent configuration: " + err.Error())
		} else {
			utils.Info("Multi-agent mode permanently enabled")
		}
	}

	// Handle commands
	command := args[0]
	utils.Info("Executing command | command=" + command)

	switch command {
	case "init":
		handleInit(initializer)
	case "agent":
		handleAgent(initializer, appPaths, *verboseFlag, args[1:])
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
		ui.Plain("\nUse --verbose-errors flag for detailed error information")
	}
	os.Exit(exitCode)
}

func printUsage() {
	ui.Header("Usage: vibe [OPTIONS] <command>")

	ui.Plain("\nOptions:")
	ui.Plain("  --verbose        Show informational messages on console")
	ui.Plain("  --debug          Show debug messages on console")
	ui.Plain("  --multi-agent    Enable multi-agent mode for this session")
	ui.Plain("  --verbose-errors Display detailed error messages on failure")
	ui.Plain("  --version        Show version information")
	ui.Plain("  -v               Show version information")

	ui.Plain("\nCommands:")
	ui.Plain("  init         Initialize current directory with vibe agents")
	ui.Plain("  agent        Interactively select and use agents for vibe IDE")
}

func handleInit(manager *core.AgentInitializer) {
	utils.Debug("Handling init command")

	// Print a blank line before starting for better spacing
	fmt.Println()

	if err := manager.Init(); err != nil {
		handleCommandError("Init", err, ExitInitError)
	}

	// Add .cursor to .gitignore and .cursorignore files
	pattern := ".cursor"
	cursorignorePattern := ".cursorignore"
	currentDir, err := os.Getwd()
	if err == nil {
		gitignorePath := filepath.Join(currentDir, ".gitignore")
		cursorignorePath := filepath.Join(currentDir, ".cursorignore")

		// Ensure .cursor is in .gitignore
		if err := utils.EnsurePathInFile(gitignorePath, pattern); err != nil {
			utils.Warn("Failed to update .gitignore with .cursor: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursor is in .gitignore")
		}

		// Ensure .cursorignore is in .gitignore
		if err := utils.EnsurePathInFile(gitignorePath, cursorignorePattern); err != nil {
			utils.Warn("Failed to update .gitignore with .cursorignore: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursorignore is in .gitignore")
		}

		// Ensure .cursor is in .cursorignore
		if err := utils.EnsurePathInFile(cursorignorePath, pattern); err != nil {
			utils.Warn("Failed to update .cursorignore: " + err.Error())
		} else {
			utils.Debug("Successfully ensured .cursor is in .cursorignore")
		}
	} else {
		utils.Warn("Could not get current directory to update ignore files: " + err.Error())
	}

	// Space after the animation
	fmt.Println()

	// Get current directory for agent listing
	currentDir, err = os.Getwd()
	if err != nil {
		utils.Warn("Could not get current directory for agent listing: " + err.Error())
	} else {
		// Display agent list from local directory
		configManager := utils.NewConfigManager()
		if err := configManager.Load(); err != nil {
			utils.Warn("Failed to load configuration: " + err.Error())
		}
		config := configManager.GetConfig()
		localRulesDir := filepath.Join(currentDir, config.RulesDirName, config.AgentsDirName)
		registry, regErr := agent.NewRegistry(config, localRulesDir)
		if regErr != nil {
			utils.Warn("Could not list agents: " + regErr.Error())
		} else {
			// Get actual agent count
			agents := registry.ListAgents()
			agentCount := len(agents)

			// Only show agents if there are any
			if agentCount > 0 {
				ui.Header("Available Agents (%d):", agentCount)

				// Display the agents in a compatible format
				if agentCount <= 5 {
					// Show compact list for small number of agents
					options := ui.DefaultAgentDisplayOptions()
					options.CompactMode = true
					ui.DisplayAgentListEnhanced(agents, options)
				} else {
					// Show summary for large number of agents
					ui.Plain("Use %s to see all available agents", ui.SuccessStyle.Sprint("vibe agent list"))
					// Show first few agents as examples
					for i := 0; i < 3 && i < len(agents); i++ {
						ui.Plain("• %s (%s)",
							ui.InfoStyle.Sprint(cleanAgentName(agents[i].Name)),
							fmt.Sprintf("@%s.mdc", agents[i].ID))
					}
					if len(agents) > 3 {
						ui.Plain("• %s", ui.WarnStyle.Sprint("...and more"))
					}
				}
				fmt.Println()
			}
		}
	}

	// Display next steps instructions with formatting
	ui.Header("Next Steps:")
	ui.Plain("1. Use %s to see available agents", ui.SuccessStyle.Sprint("vibe agent list"))
	ui.Plain("2. Start using agents by referencing them with %s in your chat", ui.SuccessStyle.Sprint("@agent-name.mdc"))
	ui.Plain("3. For more information about an agent, use %s", ui.SuccessStyle.Sprint("vibe agent info <agent-id>"))

	fmt.Println()
	utils.Info("Init command completed successfully")
}

func handleAgent(manager *core.AgentInitializer, appPaths utils.AppPaths, verbose bool, args []string) {
	utils.Debug("Handling agent command")

	// Get config for agent operations
	configManager := utils.NewConfigManager()
	if err := configManager.Load(); err != nil {
		handleCommandError("Agent", fmt.Errorf("cannot load configuration: %v", err), ExitAgentError)
	}
	config := configManager.GetConfig()

	// Get current directory to use local agents
	currentDir, err := os.Getwd()
	if err != nil {
		handleCommandError("Agent", fmt.Errorf("cannot get current directory: %v", err), ExitAgentError)
	}

	// Try multiple possible locations for agent definitions
	rulesDir := filepath.Join(currentDir, config.RulesDirName)

	// First location: .cursor/rules/cursor-rules (AgentsDirName subfolder)
	localRulesDir := filepath.Join(rulesDir, config.AgentsDirName)

	// Second location: directly in .cursor/rules (where init puts them)
	directRulesDir := rulesDir

	// Choose the appropriate directory based on what exists and has .mdc files
	chosenDir := ""

	// First check the subfolder location
	if utils.DirExists(localRulesDir) {
		// Check if it has .mdc files
		hasMDC, _ := utils.HasMDCFiles(localRulesDir)
		if hasMDC {
			chosenDir = localRulesDir
			if utils.IsDebug() {
				utils.Debugf("Using agent subfolder path: %s", chosenDir)
			}
		}
	}

	// If we haven't found a valid directory yet, check the direct location
	if chosenDir == "" && utils.DirExists(directRulesDir) {
		// Check if it has .mdc files
		hasMDC, _ := utils.HasMDCFiles(directRulesDir)
		if hasMDC {
			chosenDir = directRulesDir
			if utils.IsDebug() {
				utils.Debugf("Using direct rules path: %s", chosenDir)
			}
		}
	}

	if utils.IsVerbose() {
		utils.Infof("Using agent directory: %s", chosenDir)
	}

	if utils.IsDebug() {
		utils.Debugf("Agent configuration: %+v", config)
		utils.Debugf("Directory details: CurrentDir=%s, RulesDirName=%s, AgentsDirName=%s",
			currentDir, config.RulesDirName, config.AgentsDirName)
		utils.Debugf("Checked directories: subfolder=%s, direct=%s, chosen=%s",
			localRulesDir, directRulesDir, chosenDir)

		// List files in the rules directory
		if utils.DirExists(rulesDir) {
			files, err := os.ReadDir(rulesDir)
			if err == nil {
				utils.Debugf("Contents of %s:", rulesDir)
				for _, file := range files {
					utils.Debugf("  - %s (dir: %v)", file.Name(), file.IsDir())
				}
			} else {
				utils.Debugf("Error reading rules directory: %v", err)
			}
		} else {
			utils.Debugf("Rules directory does not exist: %s", rulesDir)
		}
	}

	// Check if we found a valid directory
	if chosenDir == "" {
		ui.Warning("No local agent definitions found in %s or %s", localRulesDir, directRulesDir)
		ui.Plain("Run %s to initialize the agent system in this directory", ui.SuccessStyle.Sprint("vibe init"))
		ui.Plain("Or use %s to get help", ui.SuccessStyle.Sprint("vibe agent --help"))
		return
	}

	// Get agent registry from chosen directory
	registry, err := agent.NewRegistry(config, chosenDir)
	if err != nil {
		handleCommandError("Agent", err, ExitAgentError)
	}

	// Debug the registry scan results
	if utils.IsDebug() {
		agents := registry.ListAgents()
		utils.Debugf("Registry scan complete - found %d agents", len(agents))
		utils.Debugf("Rules directory used: %s", registry.GetRulesDir())

		// Check for .mdc files in the directory
		mdcFiles := 0
		filepath.Walk(chosenDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mdc") {
				mdcFiles++
				utils.Debugf("Found MDC file: %s", path)
			}
			return nil
		})
		utils.Debugf("Manual MDC file count: %d", mdcFiles)
	}

	// Filter out any flag arguments (args starting with -) from the subcommand args
	var filteredArgs []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			filteredArgs = append(filteredArgs, arg)
		} else if utils.IsDebug() {
			utils.Debugf("Skipping flag argument: %s", arg)
		}
	}

	// Check if there are no non-flag arguments after "agent"
	if len(filteredArgs) < 1 {
		if utils.IsVerbose() {
			utils.Info("No subcommand provided, displaying agent list")
		}
		displayAgentList(registry)
		return
	}

	// Handle sub-commands
	subCommand := filteredArgs[0]
	utils.Info("Executing agent sub-command | sub_command=" + subCommand)

	switch subCommand {
	case "list":
		if utils.IsVerbose() {
			utils.Info("Displaying agent list")
		}
		displayAgentList(registry)
	case "select":
		if utils.IsVerbose() {
			utils.Info("Entering agent selection mode")
		}
		handleAgentSelect(registry, config, appPaths)
	case "info":
		if len(filteredArgs) < 2 {
			ui.Error("Missing agent ID. Usage: vibe agent info <agent-id>")
			os.Exit(ExitUsageError)
		}
		if utils.IsVerbose() {
			utils.Infof("Displaying agent info for: %s", filteredArgs[1])
		}
		handleAgentInfo(registry, filteredArgs[1], verbose)
	case "help", "--help", "-h":
		if utils.IsVerbose() {
			utils.Info("Displaying agent usage help")
		}
		printAgentUsage()
	default:
		ui.Warning("Unknown agent sub-command: %s", subCommand)
		printAgentUsage()
		os.Exit(ExitUsageError)
	}
}

// displayAgentList shows available agents
func displayAgentList(registry *agent.Registry) {
	if utils.IsDebug() {
		utils.Debug("Displaying agent list | registry path: " + registry.GetRulesDir())
	}

	// Get agents from registry
	agents := registry.ListAgents()
	count := len(agents)

	// Add detailed debug information
	if utils.IsDebug() {
		utils.Debugf("Found %d agents in registry", count)
		if count > 0 {
			for i, agent := range agents {
				utils.Debugf("Agent %d: ID=%s, Name=%s", i+1, agent.ID, agent.Name)
			}
		}
	}

	if count == 0 {
		ui.Header("Agent List")
		ui.Warning("No agents found in the current directory.")

		// Provide helpful information
		ui.Plain("\nPossible reasons:")
		ui.Plain("• This directory has not been initialized with vibe agents")
		ui.Plain("• The agent files (.mdc files) were not properly copied")
		ui.Plain("• The agent directory structure is incorrect")

		ui.Plain("\nSolutions:")
		ui.Plain("1. Run %s to initialize this directory", ui.SuccessStyle.Sprint("vibe init"))
		ui.Plain("2. Check that %s exists", ui.InfoStyle.Sprint(".cursor/rules directory"))
		ui.Plain("3. Verify that this directory contains %s files", ui.InfoStyle.Sprint(".mdc"))

		if utils.IsVerbose() {
			utils.Info("Agent list completed - no agents found")
		}
		return
	}

	// Get terminal width for display formatting
	termWidth := getTerminalWidth()
	if utils.IsDebug() {
		utils.Debugf("Terminal width: %d characters", termWidth)
	}

	// Get last selected agent from config
	configManager := utils.NewConfigManager()
	if err := configManager.Load(); err != nil {
		utils.Warn("Failed to load configuration: " + err.Error())
	}
	config := configManager.GetConfig()
	lastSelectedAgent := config.LastSelectedAgent

	if utils.IsVerbose() && lastSelectedAgent != "" {
		utils.Infof("Last selected agent: %s", lastSelectedAgent)
	}

	// Create display options
	options := ui.DefaultAgentDisplayOptions()
	options.TermWidth = termWidth
	options.GroupByCategory = termWidth > 100 // Use categories for wider terminals
	options.CompactMode = termWidth < 80
	options.SelectedAgentID = lastSelectedAgent

	if utils.IsDebug() {
		utils.Debugf("Agent display options: %+v", options)
	}

	// Display header with count
	ui.Header("Available Agents (%d)", count)

	// Display with enhanced UI
	ui.DisplayAgentListEnhanced(agents, options)

	// Add usage hint after the list
	ui.Plain("\nTip: Use %s to get detailed information about a specific agent",
		ui.SuccessStyle.Sprint("vibe agent info <agent-id>"))
	ui.Plain("Reference agents in your editor using @ (example: @wizard.mdc)")

	if utils.IsVerbose() {
		utils.Info("Agent list completed successfully | agent_count=" + fmt.Sprintf("%d", count))
	}
}

// cleanAgentName removes redundant words from agent names
func cleanAgentName(name string) string {
	// Remove common suffixes like "Agent", "Agent Prompt", etc.
	suffixes := []string{
		" Agent Prompt",
		" Agent",
		" Prompt",
	}

	result := name
	for _, suffix := range suffixes {
		if strings.HasSuffix(result, suffix) {
			result = strings.TrimSuffix(result, suffix)
			break
		}
	}

	return result
}

// formatAndDisplayLine formats and displays a single line of agent content
func formatAndDisplayLine(line string) {
	if strings.TrimSpace(line) == "" {
		ui.Plain("") // Empty line
	} else if strings.HasPrefix(strings.TrimSpace(line), "#") {
		// Header line - show without indentation
		ui.Header("%s", strings.TrimSpace(line))
	} else if strings.HasPrefix(strings.TrimSpace(line), "##") {
		// Subheader - show with special format
		ui.Plain("\n%s", ui.InfoStyle.Sprint(strings.TrimSpace(line)))
	} else if strings.HasPrefix(strings.TrimSpace(line), "-") || strings.HasPrefix(strings.TrimSpace(line), "*") {
		// List item
		ui.Plain("  %s", strings.TrimSpace(line))
	} else {
		// Regular line
		ui.Plain("  %s", line)
	}
}

func handleAgentSelect(registry *agent.Registry, config *utils.Config, appPaths utils.AppPaths) {
	utils.Debug("Handling agent select subcommand")

	// Create a cancellation context to handle user interrupts
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	agents := registry.ListAgents()
	if len(agents) == 0 {
		ui.Error("No agents available for selection.")
		os.Exit(ExitAgentError)
	}

	// Add channel to handle user interrupts
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	go func() {
		<-interruptCh
		utils.Info("User interrupted agent selection")
		cancel()
	}()

	// Create selector and run it with context awareness
	selector := ui.NewAgentSelector(agents)
	selectedAgent, selectErr := selector.RunWithContext(ctx)

	// Check for errors or cancellation
	if selectErr != nil {
		handleCommandError("Agent select", selectErr, ExitAgentError)
	}

	// Create loader and load the selected agent with context awareness
	loader := agent.NewLoader(registry, config)

	// Load the agent with context
	agentCtx := agent.NewAgentContext()
	loadedAgent, err := loader.LoadAgentWithContextCancellation(ctx, selectedAgent.ID, agentCtx)
	if err != nil {
		handleCommandError("Agent load", err, ExitAgentError)
	}

	// Log the loaded agent context
	utils.Debug("Agent loaded with context | last_updated=" + loadedAgent.Context.GetLastUpdated().String())

	// Save the selected agent ID to configuration
	config.LastSelectedAgent = selectedAgent.ID
	if err := utils.SaveConfig(config); err != nil {
		utils.Warn("Failed to save last selected agent: " + err.Error())
	} else {
		utils.Info("Last selected agent saved | agent_id=" + selectedAgent.ID)
	}

	// Show success message with enhanced formatting
	fmt.Println()
	ui.Success("Agent '%s' selected and loaded successfully!", selectedAgent.Name)

	// Ask user if they want to see the full agent definition
	ui.Prompt("View agent details? (y/n): ")
	var input string
	_, err = fmt.Scanln(&input)

	if err == nil && (input == "y" || input == "Y") {
		// Use our enhanced agent info display
		err = ui.DisplayAgentInfoEnhanced(selectedAgent, true)
		if err != nil {
			handleCommandError("Agent display", err, ExitAgentError)
		}
	} else {
		// Just show basic agent info with enhanced display
		err = ui.DisplayAgentInfoEnhanced(selectedAgent, false)
		if err != nil {
			handleCommandError("Agent display", err, ExitAgentError)
		}
	}

	utils.Info("Agent select completed successfully | selected_agent=" + selectedAgent.ID)
}

func handleAgentInfo(registry *agent.Registry, agentParam string, verbose bool) {
	if utils.IsDebug() {
		utils.Debugf("Handling agent info subcommand | agent_param=%s verbose=%t", agentParam, verbose)
	}

	var agentDef *agent.AgentDefinition
	var err error

	agents := registry.ListAgents()
	if utils.IsDebug() {
		utils.Debugf("Registry contains %d agents", len(agents))
	}

	// Check if there are no agents available
	if len(agents) == 0 {
		ui.Error("No agents available in the current directory.")
		ui.Plain("\nTry running %s first to initialize the agent system.", ui.SuccessStyle.Sprint("vibe init"))
		return
	}

	// Check if the input is a numeric index
	if index, convErr := strconv.Atoi(agentParam); convErr == nil {
		if utils.IsVerbose() {
			utils.Infof("Looking up agent by numeric index: %d", index)
		}

		// Validate the index is positive
		if index <= 0 {
			ui.Error("Invalid agent index: %d. Agent indexes start at 1.", index)
			if utils.IsDebug() {
				utils.Debug("Agent lookup failed: index <= 0")
			}
			os.Exit(ExitAgentError)
		}

		// Get agent by numeric index
		if index <= len(agents) {
			agentDef = agents[index-1] // Convert to 0-based index
			if utils.IsVerbose() {
				utils.Infof("Found agent at index %d: %s (%s)", index, agentDef.Name, agentDef.ID)
			}
		} else {
			ui.Error("Agent index %d is out of range. Use a number between 1 and %d.", index, len(agents))
			if utils.IsDebug() {
				utils.Debugf("Agent lookup failed: index %d > available agents %d", index, len(agents))
			}
			os.Exit(ExitAgentError)
		}
	} else {
		// Get agent by string ID
		if utils.IsVerbose() {
			utils.Infof("Looking up agent by ID: %s", agentParam)
		}

		agentDef, err = registry.GetAgent(agentParam)
		if err != nil {
			ui.Error("Agent '%s' not found", agentParam)
			if utils.IsVerbose() {
				ui.Plain("\nAvailable agents:")
				for i, agent := range agents {
					ui.Plain("  %d. %s (%s)", i+1, agent.Name, agent.ID)
				}
			}
			if utils.IsDebug() {
				utils.Debugf("Agent lookup error: %v", err)
			}
			os.Exit(ExitAgentError)
		} else if utils.IsVerbose() {
			utils.Infof("Found agent with ID %s: %s", agentDef.ID, agentDef.Name)
		}
	}

	// Display header with agent name before showing details
	ui.Header("Agent: %s", agentDef.Name)

	if utils.IsDebug() {
		utils.Debugf("Displaying agent info | id=%s name=%s verbose=%t",
			agentDef.ID, agentDef.Name, verbose)
	}

	// Use enhanced agent info display with verbose parameter
	err = ui.DisplayAgentInfoEnhanced(agentDef, verbose)
	if err != nil {
		handleCommandError("Agent display", err, ExitAgentError)
	}

	if utils.IsVerbose() {
		utils.Infof("Agent info displayed successfully | agent=%s (%s)", agentDef.Name, agentDef.ID)
	}
}

func printAgentUsage() {
	ui.Header("Usage: vibe agent [OPTIONS] [subcommand]")

	ui.Plain("\nOptions:")
	ui.Plain("  --verbose        Show additional information")
	ui.Plain("  --debug          Show detailed debug information")

	ui.Plain("\nSubcommands:")
	ui.Plain("  <none>       List all available agents (default)")
	ui.Plain("  list         List all available agents")
	ui.Plain("  select       Interactively select an agent")
	ui.Plain("  info <id>    Display detailed information about a specific agent")
	ui.Plain("  help         Show this help message")

	ui.Plain("\nExample usage:")
	ui.Plain("  vibe agent                 # List all available agents")
	ui.Plain("  vibe agent --verbose       # List agents with verbose output")
	ui.Plain("  vibe agent select          # Interactively select an agent")
	ui.Plain("  vibe agent info wizard     # Show info about the wizard agent")
	ui.Plain("  vibe agent info 1 --debug  # Show detailed info with debug output")
	ui.Plain("\nYou can also reference agents in the chatbox using @ (example: @wizard.mdc)")
}
