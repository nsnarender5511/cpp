package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"crules/internal/agent"
	"crules/internal/core"
	"crules/internal/ui"
	"crules/internal/utils"
	"crules/internal/version"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Exit codes
const (
	ExitSuccess    = 0
	ExitUsageError = 1
	ExitInitError  = 10
	ExitMergeError = 11
	ExitSyncError  = 12
	ExitListError  = 13
	ExitCleanError = 14
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
		}
		// Keep non-flag arguments
		args = append(args, arg)
	}

	// Replace os.Args with filtered args
	os.Args = args

	// Check for version flag
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("crules version %s\n", version.GetVersion())
		os.Exit(ExitSuccess)
	}

	// Get app paths and initialize logger
	appName := os.Getenv("APP_NAME")
	appPaths := utils.GetAppPaths(appName)
	utils.InitLogger(appPaths)

	// Configure console output
	utils.SetDebugConsole(debugFlag)
	utils.SetVerbose(verboseFlag)

	utils.Info("Starting crules")

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

	// Handle commands
	command := os.Args[1]
	utils.Info("Executing command | command=" + command)

	switch command {
	case "init":
		handleInit(manager)
	case "merge":
		handleMerge(manager)
	case "sync":
		handleSync(manager)
	case "list":
		handleList(manager)
	case "clean":
		handleClean(manager)
	case "agent":
		handleAgent(manager, appPaths)
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
	utils.Error(commandName + " failed: " + err.Error())
	ui.Error("%s failed: %v", commandName, err)
	os.Exit(exitCode)
}

func printUsage() {
	ui.Header("Usage: crules [OPTIONS] <command>")

	ui.Plain("\nOptions:")
	ui.Plain("  --verbose    Show informational messages on console")
	ui.Plain("  --debug      Show debug messages on console")
	ui.Plain("  --version    Show version information")
	ui.Plain("  -v           Show version information")

	ui.Plain("\nCommands:")
	ui.Plain("  init         Initialize current directory with rules from main location")
	ui.Plain("  merge        Merge current rules to main location and sync to all locations")
	ui.Plain("  sync         Force sync from main location to current directory")
	ui.Plain("  list         Display all registered projects")
	ui.Plain("  clean        Remove non-existent projects from registry")
	ui.Plain("  agent        Interactively select and use agents")
}

func handleInit(manager *core.SyncManager) {
	utils.Debug("Handling init command")
	if err := manager.Init(); err != nil {
		handleCommandError("Init", err, ExitInitError)
	}
	utils.Info("Init command completed successfully")
	ui.Success("Successfully initialized rules from main location")
}

func handleMerge(manager *core.SyncManager) {
	utils.Debug("Handling merge command")
	if err := manager.Merge(); err != nil {
		handleCommandError("Merge", err, ExitMergeError)
	}
	utils.Info("Merge command completed successfully")
	ui.Success("Successfully merged rules to main location")
}

func handleSync(manager *core.SyncManager) {
	utils.Debug("Handling sync command")
	if err := manager.Sync(); err != nil {
		handleCommandError("Sync", err, ExitSyncError)
	}
	utils.Info("Sync command completed successfully")
	ui.Success("Successfully synced rules from main location")
}

func handleList(manager *core.SyncManager) {
	utils.Debug("Handling list command")

	projects := manager.GetRegistry().GetProjects()
	count := len(projects)

	if count == 0 {
		ui.Info("No projects registered.")
		utils.Info("List command completed - no projects found")
		return
	}

	ui.Header("Registered projects (%d):", count)
	validCount := 0

	for i, project := range projects {
		exists := utils.DirExists(project)
		if exists {
			validCount++
			ui.Plain("  %d. %s", i+1, project)
		} else {
			ui.Plain("  %d. %s %s", i+1, project, ui.ErrorStyle.Sprint("(not found)"))
		}
	}

	if validCount < count {
		ui.Warning("\n%d project(s) could not be found. Run 'crules clean' to remove them.", count-validCount)
	}

	utils.Info("List command completed successfully | project_count=" + fmt.Sprintf("%d", count) + ", valid_count=" + fmt.Sprintf("%d", validCount))
}

func handleClean(manager *core.SyncManager) {
	utils.Debug("Handling clean command")

	removedCount, err := manager.Clean()
	if err != nil {
		handleCommandError("Clean", err, ExitCleanError)
	}

	if removedCount == 0 {
		ui.Info("No projects needed to be removed.")
	} else {
		ui.Success("Successfully removed %d non-existent project(s) from registry.", removedCount)
	}

	utils.Info("Clean command completed successfully | removed_count=" + fmt.Sprintf("%d", removedCount))
}

func handleAgent(manager *core.SyncManager, appPaths utils.AppPaths) {
	utils.Debug("Handling agent command")

	// Get arguments for subcommands
	subcommand := "" // Empty default means show list
	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	}

	// Get config
	config := utils.LoadConfig()

	// Get rules directory path
	rulesDir := appPaths.GetRulesDir(config.RulesDirName)
	utils.Debug("Using rules directory | path=" + rulesDir)

	// Create agent registry
	registry, err := agent.NewRegistry(config, rulesDir)
	if err != nil {
		handleCommandError("Agent", err, ExitAgentError)
	}

	// Handle subcommands
	switch subcommand {
	case "":
		// Default behavior: show agent list
		displayAgentList(registry)
	case "select":
		handleAgentSelect(registry, config, appPaths)
	case "info":
		if len(os.Args) < 4 {
			ui.Error("Missing agent ID. Usage: crules agent info <agent-id>")
			os.Exit(ExitUsageError)
		}
		handleAgentInfo(registry, os.Args[3])
	default:
		ui.Error("Unknown agent subcommand: %s", subcommand)
		printAgentUsage()
		os.Exit(ExitUsageError)
	}

	utils.Info("Agent command completed successfully | subcommand=" + subcommand)
}

// displayAgentList displays all available agents
func displayAgentList(registry *agent.Registry) {
	utils.Debug("Displaying agent list")

	agents := registry.ListAgents()
	count := len(agents)

	if count == 0 {
		ui.Info("No agents available.")
		utils.Info("Agent list completed - no agents found")
		return
	}

	ui.Header("Available agents (%d):", count)

	// Get terminal width for responsive display
	termWidth := getTerminalWidth()
	utils.Debug(fmt.Sprintf("Terminal width detected: %d", termWidth))

	// For very narrow terminals, use a compact vertical list
	if termWidth < 50 {
		displayCompactAgentList(agents)
	} else {
		// Use the table display for wider terminals
		displayAgentTable(agents, termWidth)
	}

	ui.Plain("\nTip: You can reference agents in the chatbox using the @ syntax shown above")
	ui.Plain("For more details about any agent, use: crules agent info <agent-id>")

	utils.Info("Agent list completed successfully | agent_count=" + fmt.Sprintf("%d", count))
}

// displayCompactAgentList renders agents as a simple vertical list for very narrow terminals
func displayCompactAgentList(agents []*agent.AgentDefinition) {
	for i, agent := range agents {
		// Clean up the name if available
		name := cleanAgentName(agent.Name)

		// Display agent number and name
		ui.Plain("%s %s",
			ui.SuccessStyle.Sprintf("%d.", i+1),
			ui.InfoStyle.Sprint(name))

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

	// Configure table based on terminal width
	if termWidth < 80 {
		// For medium width terminals
		headers := table.Row{
			ui.HeaderStyle.Sprint("No."),
			ui.HeaderStyle.Sprint("Name"),
			ui.HeaderStyle.Sprint("Reference"),
		}
		t.AppendHeader(headers)

		for i, agent := range agents {
			// Clean up the name by removing redundant words
			name := cleanAgentName(agent.Name)

			// Truncate name if needed
			if len(name) > 25 {
				name = name[:22] + "..."
			}

			t.AppendRow(table.Row{
				ui.SuccessStyle.Sprintf("%d", i+1),
				ui.InfoStyle.Sprint(name),
				fmt.Sprintf("@%s.mdc", agent.ID),
			})
		}
	} else {
		// For wide terminals, show more details
		headers := table.Row{
			ui.HeaderStyle.Sprint("No."),
			ui.HeaderStyle.Sprint("Name"),
			ui.HeaderStyle.Sprint("Reference ID"),
			ui.HeaderStyle.Sprint("Version"),
		}
		t.AppendHeader(headers)

		for i, agent := range agents {
			// Clean up the name by removing redundant words
			name := cleanAgentName(agent.Name)

			t.AppendRow(table.Row{
				ui.SuccessStyle.Sprintf("%d", i+1),
				ui.InfoStyle.Sprint(name),
				fmt.Sprintf("@%s.mdc", agent.ID),
				ui.WarnStyle.Sprint(agent.Version),
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
func displayAgentContent(agentDef *agent.AgentDefinition, paginated bool) {
	// Get content from file if not loaded yet
	content := agentDef.Content
	if content == "" {
		contentBytes, err := os.ReadFile(agentDef.DefinitionPath)
		if err == nil {
			content = string(contentBytes)
		} else {
			content = "Error loading content: " + err.Error()
		}
	}

	// Split into lines
	lines := strings.Split(content, "\n")

	// Display with formatting
	if paginated {
		// Calculate page size based on terminal height (assuming ~30 lines)
		pageSize := 20
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

	agents := registry.ListAgents()
	if len(agents) == 0 {
		ui.Error("No agents available for selection.")
		os.Exit(ExitAgentError)
	}

	// Create selector and run it
	selector := ui.NewAgentSelector(agents)
	selectedAgent, err := selector.Run()
	if err != nil {
		handleCommandError("Agent select", err, ExitAgentError)
	}

	// Create loader and load the selected agent
	loader := agent.NewLoader(registry, config)
	agent, err := loader.LoadAgent(selectedAgent.ID)
	if err != nil {
		handleCommandError("Agent load", err, ExitAgentError)
	}

	// Log the loaded agent context
	utils.Debug("Agent loaded with context | last_updated=" + agent.Context.LastUpdated.String())

	ui.Success("Agent '%s' loaded successfully!", selectedAgent.Name)

	// Ask user if they want to see the full agent definition
	ui.Prompt("View full agent definition? (y/n): ")
	var input string
	_, err = fmt.Scanln(&input)

	if err == nil && (input == "y" || input == "Y") {
		// Show the full agent definition with pagination
		ui.Header("\nAgent Definition:")
		ui.Plain("")
		displayAgentContent(selectedAgent, true) // Display with pagination
	} else {
		// Just show a brief overview
		ui.Plain("\nAgent overview:")
		ui.Plain("  ID:       %s", selectedAgent.ID)
		ui.Plain("  Name:     %s", selectedAgent.Name)
		ui.Plain("  Version:  %s", selectedAgent.Version)
		ui.Plain("  File:     %s", selectedAgent.DefinitionPath)

		if len(selectedAgent.Capabilities) > 0 {
			ui.Plain("\nCapabilities:")
			for i, capability := range selectedAgent.Capabilities {
				if i < 3 { // Show only first 3 capabilities
					ui.Plain("  - %s", capability)
				} else {
					ui.Plain("  - ...")
					break
				}
			}
		}
	}

	utils.Info("Agent select completed successfully | selected_agent=" + selectedAgent.ID)
}

func handleAgentInfo(registry *agent.Registry, agentID string) {
	utils.Debug("Handling agent info subcommand | agent_id=" + agentID)

	// Get agent by ID
	agentDef, exists := registry.GetAgent(agentID)
	if !exists {
		ui.Error("Agent '%s' not found.", agentID)
		os.Exit(ExitAgentError)
	}

	// Display agent details
	ui.Header("Agent details:")
	ui.Plain("  ID:          %s", agentDef.ID)
	ui.Plain("  Name:        %s", agentDef.Name)
	ui.Plain("  Version:     %s", agentDef.Version)
	ui.Plain("")

	// Display full description with proper line wrapping
	ui.Plain("Description:")
	if agentDef.Description != "" {
		displayAgentContent(agentDef, false) // Display without pagination
	} else {
		ui.Plain("  No description available.")
	}

	if len(agentDef.Capabilities) > 0 {
		ui.Plain("\nCapabilities:")
		for _, capability := range agentDef.Capabilities {
			ui.Plain("  - %s", capability)
		}
	}

	ui.Plain("\nFile: %s", agentDef.DefinitionPath)

	utils.Info("Agent info completed successfully | agent_id=" + agentID)
}

func printAgentUsage() {
	ui.Header("Usage: crules agent [subcommand]")

	ui.Plain("\nSubcommands:")
	ui.Plain("  <none>       List all available agents (default)")
	ui.Plain("  select       Interactively select an agent")
	ui.Plain("  info <id>    Display detailed information about a specific agent")

	ui.Plain("\nExample usage:")
	ui.Plain("  crules agent               # List all available agents")
	ui.Plain("  crules agent select        # Interactively select an agent")
	ui.Plain("  crules agent info wizard   # Show detailed info about the wizard agent")
	ui.Plain("\nYou can also reference agents in the chatbox using @ (example: @wizard.mdc)")
}
