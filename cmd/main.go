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

	"cursor++/internal/agent"
	"cursor++/internal/constants"
	"cursor++/internal/core"
	"cursor++/internal/ui"
	"cursor++/internal/utils"
	"cursor++/internal/version"
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
		fmt.Printf("cursor++ version %s\n", version.GetVersion())
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

	utils.Info("Starting cursor++")

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
	ui.Header("Usage: cursor++ [OPTIONS] <command>")

	ui.Plain("\nOptions:")
	ui.Plain("  --verbose        Show informational messages on console")
	ui.Plain("  --debug          Show debug messages on console")
	ui.Plain("  --multi-agent    Enable multi-agent mode for this session")
	ui.Plain("  --verbose-errors Display detailed error messages on failure")
	ui.Plain("  --version        Show version information")
	ui.Plain("  -v               Show version information")

	ui.Plain("\nCommands:")
	ui.Plain("  init         Initialize current directory with cursor++ agents")
	ui.Plain("  agent        Interactively select and use agents for cursor++ IDE")
}

func handleInit(manager *core.AgentInitializer) {
	utils.Debug("Handling init command")

	// Print a blank line before starting for better spacing
	fmt.Println()

	if err := manager.Init(); err != nil {
		handleCommandError("Init", err, ExitInitError)
	}

	// Space after the animation
	fmt.Println()

	// Get current directory for agent listing
	currentDir, err := os.Getwd()
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
					ui.Plain("Use %s to see all available agents", ui.SuccessStyle.Sprint("cursor++ agent list"))
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
	ui.Plain("1. Use %s to see available agents", ui.SuccessStyle.Sprint("cursor++ agent list"))
	ui.Plain("2. Start using agents by referencing them with %s in your chat", ui.SuccessStyle.Sprint("@agent-name.mdc"))
	ui.Plain("3. For more information about an agent, use %s", ui.SuccessStyle.Sprint("cursor++ agent info <agent-id>"))

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

	// Use local rules directory instead of global
	localRulesDir := filepath.Join(currentDir, config.RulesDirName, config.AgentsDirName)
	utils.Debug("Using local rules directory | path=" + localRulesDir)

	// Get agent registry from local directory
	registry, err := agent.NewRegistry(config, localRulesDir)
	if err != nil {
		handleCommandError("Agent", err, ExitAgentError)
	}

	// Check if there are no arguments after "agent"
	if len(args) < 1 {
		displayAgentList(registry)
		return
	}

	// Handle sub-commands
	subCommand := args[0]
	utils.Info("Executing agent sub-command | sub_command=" + subCommand)

	switch subCommand {
	case "list":
		displayAgentList(registry)
	case "select":
		handleAgentSelect(registry, config, appPaths)
	case "info":
		if len(args) < 2 {
			ui.Error("Missing agent ID. Usage: cursor++ agent info <agent-id>")
			os.Exit(ExitUsageError)
		}
		handleAgentInfo(registry, args[1], verbose)
	default:
		ui.Warning("Unknown agent sub-command: %s", subCommand)
		printAgentUsage()
		os.Exit(ExitUsageError)
	}
}

// displayAgentList shows available agents
func displayAgentList(registry *agent.Registry) {
	utils.Debug("Displaying agent list")

	// Get agents from registry
	agents := registry.ListAgents()
	count := len(agents)

	if count == 0 {
		ui.Warning("No agents found.")
		utils.Info("Agent list completed - no agents found")
		return
	}

	// Get terminal width for display formatting
	termWidth := getTerminalWidth()

	// Get last selected agent from config
	configManager := utils.NewConfigManager()
	if err := configManager.Load(); err != nil {
		utils.Warn("Failed to load configuration: " + err.Error())
	}
	config := configManager.GetConfig()
	lastSelectedAgent := config.LastSelectedAgent

	// Create display options
	options := ui.DefaultAgentDisplayOptions()
	options.TermWidth = termWidth
	options.GroupByCategory = termWidth > 100 // Use categories for wider terminals
	options.CompactMode = termWidth < 80
	options.SelectedAgentID = lastSelectedAgent

	// Display with enhanced UI
	ui.DisplayAgentListEnhanced(agents, options)

	utils.Info("Agent list completed successfully | agent_count=" + fmt.Sprintf("%d", count))
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

// displayAgentContent displays agent content with proper formatting
// If paginated is true, the content will be displayed with pagination
// Returns error if content loading fails
func displayAgentContent(agentDef *agent.AgentDefinition, paginated bool) error {
	// Get content from file if not loaded yet
	content := agentDef.Content
	if content == "" {
		// Validate path for security
		if !validatePath(agentDef.DefinitionPath) {
			return fmt.Errorf(constants.ErrInvalidAgentDefPath)
		}

		contentBytes, err := os.ReadFile(agentDef.DefinitionPath)
		if err != nil {
			utils.Errorf(constants.ErrFailedLoadAgentContent, err)
			return fmt.Errorf(constants.ErrFailedLoadAgentContent, err)
		}
		content = string(contentBytes)
	}

	// Add multi-agent information
	config := utils.NewConfigManager().GetConfig()
	if config.MultiAgentEnabled {
		content = strings.Replace(content,
			"# Agent System Integration:",
			"# Multi-Agent System Integration:\n\n> Note: This agent is configured to work in multi-agent mode with other cursor++ agents.",
			1)
	}

	// Split into lines
	lines := strings.Split(content, "\n")

	// Display with formatting
	if paginated {
		// Define page size as a constant or derive from terminal size
		pageSize := calculatePageSize()
		currentLine := 0
		pageNum := 1

		for currentLine < len(lines) {
			endLine := currentLine + pageSize
			if endLine > len(lines) {
				endLine = len(lines)
			}

			for i := currentLine; i < endLine; i++ {
				formatAndDisplayLine(lines[i])
			}

			currentLine = endLine
			pageNum++

			// If there are more pages, prompt to continue
			if currentLine < len(lines) {
				ui.Plain("")
				ui.Prompt("Press Enter to continue, or q to quit: ")
				var key string
				_, err := fmt.Scanln(&key)
				if err == nil && (key == "q" || key == "Q") {
					break
				}
			}
		}
	} else {
		// Display all content without pagination
		ui.Plain("")
		for _, line := range lines {
			formatAndDisplayLine(line)
		}
	}

	return nil
}

// calculatePageSize determines an appropriate page size based on terminal dimensions
// Default is 20 lines, but will adjust if terminal size is available
func calculatePageSize() int {
	// Default size
	const defaultPageSize = 20

	// Try to get terminal size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return defaultPageSize
	}

	// Parse output (rows columns)
	parts := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(parts) != 2 {
		return defaultPageSize
	}

	// Get rows and subtract buffer for prompt and status
	rows, err := strconv.Atoi(parts[0])
	if err != nil || rows <= 5 {
		return defaultPageSize
	}

	// Leave space for prompts and headers
	return rows - 5
}

// validatePath checks if a file path is safe to access
// Returns false for paths containing directory traversal patterns or absolute paths
func validatePath(path string) bool {
	// Clean the path first to normalize it
	cleanPath := filepath.Clean(path)

	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		utils.Errorf(constants.ErrFailedGetCurrentDir, err)
		return false
	}

	// Verify the path is within expected directory
	config := utils.NewConfigManager().GetConfig()

	// Get absolute path of local agents directory
	localAgentsDir := filepath.Join(currentDir, config.RulesDirName, config.AgentsDirName)
	absAgentsDir, err := filepath.Abs(localAgentsDir)
	if err != nil {
		utils.Errorf(constants.ErrFailedGetAgentsDirPath, err)
		return false
	}

	// Get absolute path of file
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		utils.Errorf(constants.ErrFailedGetAbsPath, err)
		return false
	}

	// Use filepath.Rel to check if path is inside the agents directory
	relPath, err := filepath.Rel(absAgentsDir, absPath)
	if err != nil || strings.HasPrefix(relPath, "..") || filepath.IsAbs(relPath) {
		utils.Errorf(constants.ErrPathSecurityViolation, path)
		utils.Debug("Expected prefix: " + absAgentsDir)
		utils.Debug("Actual path: " + absPath)
		return false
	}

	return true
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
	utils.Debug("Handling agent info subcommand | agent_param=" + agentParam)

	var agentDef *agent.AgentDefinition
	var err error

	// Check if the input is a numeric index
	if index, convErr := strconv.Atoi(agentParam); convErr == nil {
		// Validate the index is positive
		if index <= 0 {
			ui.Error("Invalid agent index: %d. Agent indexes start at 1.", index)
			os.Exit(ExitAgentError)
		}

		// Get agent by numeric index
		agents := registry.ListAgents()
		if index <= len(agents) {
			agentDef = agents[index-1] // Convert to 0-based index
		} else {
			ui.Error("Agent index %d is out of range. Use a number between 1 and %d.", index, len(agents))
			os.Exit(ExitAgentError)
		}
	} else {
		// Get agent by string ID
		agentDef, err = registry.GetAgent(agentParam)
		if err != nil {
			ui.Error("Agent '%s' not found: %v", agentParam, err)
			os.Exit(ExitAgentError)
		}
	}

	// Use enhanced agent info display with verbose parameter
	err = ui.DisplayAgentInfoEnhanced(agentDef, verbose)
	if err != nil {
		handleCommandError("Agent display", err, ExitAgentError)
	}

	utils.Info("Agent info completed successfully | agent_param=" + agentParam + ", agent_id=" + agentDef.ID)
}

func printAgentUsage() {
	ui.Header("Usage: cursor++ agent [subcommand]")

	ui.Plain("\nSubcommands:")
	ui.Plain("  <none>       List all available agents (default)")
	ui.Plain("  select       Interactively select an agent")
	ui.Plain("  info <id>    Display detailed information about a specific agent")

	ui.Plain("\nExample usage:")
	ui.Plain("  cursor++ agent               # List all available agents")
	ui.Plain("  cursor++ agent select        # Interactively select an agent")
	ui.Plain("  cursor++ agent info wizard   # Show detailed info about the wizard agent")
	ui.Plain("  cursor++ agent info 1        # Show detailed info about the first agent in the list")
	ui.Plain("\nYou can also reference agents in the chatbox using @ (example: @wizard.mdc)")
}
