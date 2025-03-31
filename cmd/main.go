package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"

	"cursor++/internal/agent"
	"cursor++/internal/core"
	"cursor++/internal/ui"
	"cursor++/internal/utils"
	"cursor++/internal/version"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Exit codes
const (
	ExitSuccess    = 0
	ExitUsageError = 1
	ExitInitError  = 10
	ExitAgentError = 15
	ExitSetupError = 20
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
	// Parse debug/verbose flags
	debugFlag := false
	verboseFlag := false
	multiAgentFlag := false
	verboseErrorsFlag := false

	// Process each argument to find flags
	args := make([]string, 0, len(os.Args))
	for _, arg := range os.Args {
		if arg == "--debug" {
			debugFlag = true
			verboseFlag = true // Debug implies verbose
			continue
		} else if arg == "--verbose" {
			verboseFlag = true
			continue
		} else if arg == "--multi-agent" {
			multiAgentFlag = true
			continue
		} else if arg == "--verbose-errors" {
			verboseErrorsFlag = true
			continue
		}
		// Keep non-flag arguments
		args = append(args, arg)
	}

	// Replace os.Args with filtered args
	os.Args = args

	// Check for version flag
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("cursor++ version %s\n", version.GetVersion())
		os.Exit(ExitSuccess)
	}

	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths)

	// Configure console output
	utils.SetDebugConsole(debugFlag)
	utils.SetVerbose(verboseFlag)
	utils.SetVerboseErrors(verboseErrorsFlag)

	utils.Info("Starting cursor++")

	// Print usage if no arguments
	if len(os.Args) < 2 {
		utils.Debug("No command provided, showing usage")
		printUsage()
		os.Exit(ExitUsageError)
	}

	// Display banner for all commands
	ui.PrintBanner()

	// Create new sync manager
	utils.Debug("Initializing sync manager")
	manager, err := core.NewSyncManager()
	if err != nil {
		utils.Error("Error initializing: " + err.Error())
		ui.Error("Error initializing: %v", err)
		os.Exit(ExitSetupError)
	}

	// Set multi-agent mode if flag is provided
	if multiAgentFlag {
		utils.Info("Multi-agent mode explicitly enabled via flag")
		// Update the config to enable multi-agent mode
		config := utils.LoadConfig()
		config.MultiAgentEnabled = true

		// Save configuration to make the change permanent
		if err := utils.SaveConfig(config); err != nil {
			utils.Warn("Failed to save multi-agent configuration: " + err.Error())
		} else {
			utils.Info("Multi-agent mode permanently enabled")
		}
	}

	// Handle commands
	command := os.Args[1]
	utils.Info("Executing command | command=" + command)

	switch command {
	case "init":
		handleInit(manager)
	case "agent":
		handleAgent(manager, appPaths, verboseFlag)
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

func handleInit(manager *core.SyncManager) {
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
		config := utils.LoadConfig()
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
					displayCompactAgentList(agents)
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

func handleAgent(manager *core.SyncManager, appPaths utils.AppPaths, verbose bool) {
	utils.Debug("Handling agent command")

	// Get config for agent operations
	config := utils.LoadConfig()

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
	if len(os.Args) < 3 {
		displayAgentList(registry)
		return
	}

	// Handle sub-commands
	subCommand := os.Args[2]
	utils.Info("Executing agent sub-command | sub_command=" + subCommand)

	switch subCommand {
	case "list":
		displayAgentList(registry)
	case "select":
		handleAgentSelect(registry, config, appPaths)
	case "info":
		if len(os.Args) < 4 {
			ui.Error("Missing agent ID. Usage: cursor++ agent info <agent-id>")
			os.Exit(ExitUsageError)
		}
		handleAgentInfo(registry, os.Args[3], verbose)
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
		ui.Info("No agents available.")
		utils.Info("Agent list completed - no agents found")
		return
	}

	// Get terminal width for display formatting
	termWidth := getTerminalWidth()

	// Get last selected agent from config
	config := utils.LoadConfig()
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

// displayCompactAgentList renders agents as a simple vertical list for narrow terminals
func displayCompactAgentList(agents []*agent.AgentDefinition) {
	// Get last selected agent from config
	config := utils.LoadConfig()
	lastSelectedAgent := config.LastSelectedAgent

	for i, agent := range agents {
		// Clean up the name if available
		name := cleanAgentName(agent.Name)

		// Display a star for last selected agent
		selected := ""
		if agent.ID == lastSelectedAgent {
			selected = ui.WarnStyle.Sprint(" ★")
		}

		// Display agent number and name
		ui.Plain("%s %s%s",
			ui.SuccessStyle.Sprintf("%d.", i+1),
			ui.InfoStyle.Sprint(name),
			selected)

		// Display reference on the next line with indentation
		ui.Plain("   %s", fmt.Sprintf("@%s.mdc", agent.ID))

		// Add a separator line between agents (except after the last one)
		if i < len(agents)-1 {
			ui.Plain("   ─────────────────")
		}
	}
}

// displayAgentTable renders agent list as a formatted table
func displayAgentTable(agents []*agent.AgentDefinition, termWidth int) {
	// Get last selected agent from config
	config := utils.LoadConfig()
	lastSelectedAgent := config.LastSelectedAgent

	// Create a new table writer
	t := table.NewWriter()
	// Don't use SetOutputMirror since we'll be manually printing the output

	// Start with a light style
	customStyle := table.StyleLight

	// Enable row separation
	customStyle.Options.SeparateRows = true

	// Setup colors - no alternating row colors
	customStyle.Color.Header = text.Colors{text.Bold}
	customStyle.Color.Row = text.Colors{}
	customStyle.Color.RowAlternate = nil

	// Apply the style
	t.SetStyle(customStyle)

	// Choose appropriate columns based on width
	if termWidth < 80 {
		// For narrow terminals, show minimal info
		headers := table.Row{
			ui.HeaderStyle.Sprint("No."),
			ui.HeaderStyle.Sprint("Name"),
			ui.HeaderStyle.Sprint("Reference"),
			ui.HeaderStyle.Sprint(""),
		}
		t.AppendHeader(headers)

		for i, agent := range agents {
			// Clean up the name by removing redundant words
			name := cleanAgentName(agent.Name)

			// Truncate name if needed
			if len(name) > 25 {
				name = name[:22] + "..."
			}

			// Add a star for last selected agent
			selected := ""
			if agent.ID == lastSelectedAgent {
				selected = ui.WarnStyle.Sprint("★")
			}

			t.AppendRow(table.Row{
				ui.SuccessStyle.Sprintf("%d", i+1),
				ui.InfoStyle.Sprint(name),
				fmt.Sprintf("@%s.mdc", agent.ID),
				selected,
			})
		}
	} else {
		// For wide terminals, show more details
		headers := table.Row{
			ui.HeaderStyle.Sprint("No."),
			ui.HeaderStyle.Sprint("Name"),
			ui.HeaderStyle.Sprint("Reference ID"),
			ui.HeaderStyle.Sprint("Version"),
			ui.HeaderStyle.Sprint(""),
		}
		t.AppendHeader(headers)

		for i, agent := range agents {
			// Clean up the name by removing redundant words
			name := cleanAgentName(agent.Name)

			// Add a star for last selected agent
			selected := ""
			if agent.ID == lastSelectedAgent {
				selected = ui.WarnStyle.Sprint("★")
			}

			t.AppendRow(table.Row{
				ui.SuccessStyle.Sprintf("%d", i+1),
				ui.InfoStyle.Sprint(name),
				fmt.Sprintf("@%s.mdc", agent.ID),
				ui.WarnStyle.Sprint(agent.Version),
				selected,
			})
		}
	}

	// Custom render to stdout
	tableString := t.Render()

	// Replace the first separator line with a double line separator
	lines := strings.Split(tableString, "\n")

	for i, line := range lines {
		// Find the first separator line (should be after the header)
		if i > 0 && strings.Contains(line, "├") {
			// Replace with double line characters
			line = strings.ReplaceAll(line, "├", "╞")
			line = strings.ReplaceAll(line, "┤", "╡")
			line = strings.ReplaceAll(line, "┼", "╪")
			line = strings.ReplaceAll(line, "─", "═")
			lines[i] = line
			break
		}
	}

	// Print the modified table
	fmt.Println(strings.Join(lines, "\n"))
}

// cleanAgentName removes redundant words from agent names
func cleanAgentName(name string) string {
	// Remove common redundant words
	name = strings.ReplaceAll(name, "Agent Prompt", "")
	name = strings.ReplaceAll(name, "Prompt", "")
	name = strings.ReplaceAll(name, "Agent", "")

	// Trim spaces and handle multiple spaces
	name = strings.TrimSpace(name)
	name = strings.Join(strings.Fields(name), " ")

	return name
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
			return fmt.Errorf("invalid agent definition path")
		}

		contentBytes, err := os.ReadFile(agentDef.DefinitionPath)
		if err != nil {
			utils.Error("Failed to load agent content: " + err.Error())
			return fmt.Errorf("failed to load agent content: %w", err)
		}
		content = string(contentBytes)
	}

	// Add multi-agent information
	config := utils.LoadConfig()
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
	// Check for directory traversal attempts
	if strings.Contains(path, "..") {
		utils.Error("Security warning: path contains directory traversal pattern: " + path)
		return false
	}

	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		utils.Error("Failed to get current directory: " + err.Error())
		return false
	}

	// Verify the path is within expected directory
	config := utils.LoadConfig()

	// Get absolute path of local agents directory
	localAgentsDir := filepath.Join(currentDir, config.RulesDirName, config.AgentsDirName)
	absAgentsDir, err := filepath.Abs(localAgentsDir)
	if err != nil {
		utils.Error("Failed to get absolute path for local agents directory: " + err.Error())
		return false
	}

	// Get absolute path of file
	absPath, err := filepath.Abs(path)
	if err != nil {
		utils.Error("Failed to get absolute path: " + err.Error())
		return false
	}

	// Check if file is within agents directory
	if !strings.HasPrefix(absPath, absAgentsDir) {
		utils.Error("Security warning: path is outside agents directory: " + path)
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

	// Create selector and run it in a separate goroutine with context awareness
	selectorDone := make(chan struct{})
	var selectedAgent *agent.AgentDefinition
	var selectErr error

	go func() {
		selector := ui.NewAgentSelector(agents)
		selectedAgent, selectErr = selector.Run()
		close(selectorDone)
	}()

	// Wait for either selector completion or context cancellation
	select {
	case <-selectorDone:
		if selectErr != nil {
			handleCommandError("Agent select", selectErr, ExitAgentError)
		}
	case <-ctx.Done():
		ui.Error("Agent selection was canceled")
		os.Exit(ExitAgentError)
	}

	// Create loader and load the selected agent
	loader := agent.NewLoader(registry, config)

	// Load the agent
	agentCtx := agent.NewAgentContext()
	loadedAgent, err := loader.LoadAgentWithContext(selectedAgent.ID, agentCtx)
	if err != nil {
		handleCommandError("Agent load", err, ExitAgentError)
	}

	// Log the loaded agent context
	utils.Debug("Agent loaded with context | last_updated=" + loadedAgent.Context.LastUpdated.String())

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
