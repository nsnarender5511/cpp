package cli

import (
	"vibe/internal/agent"
	"vibe/internal/ui"
)

// PrintUsage displays the main help message for the application.
func PrintUsage() {
	ui.Header("Usage: vibe [OPTIONS] <command>")

	ui.Plain("\nGlobal Options:")
	ui.Plain("  --debug                Show debug messages on console.")
	ui.Plain("  --multi-agent          Enable multi-agent mode for collaborative tasks.")
	ui.Plain("  -v                     Display the application version and exit (shorthand).")

	ui.Plain("\nCommands:")
	ui.Plain("  init         Initialize current directory with vibe agents")
	ui.Plain("  agent        Interactively select and use agents for vibe IDE")
}

// PrintAgentUsage displays the help message for the 'agent' command.
func PrintAgentUsage(registry *agent.Registry) {
	ui.Header("Usage: vibe agent <subcommand>")

	ui.Plain("\nSubcommands:")
	ui.Plain("  list         List all available agents")
	ui.Plain("  select       Interactively select an agent to run")
	ui.Plain("  run <name>   Run a specific agent by its name")
	ui.Plain("  info <name>  Show detailed information about a specific agent")

	ui.Plain("\nAvailable Agents (for 'run' and 'info'):")
	if registry != nil {
		agents := registry.ListAgents()
		if len(agents) > 0 {
			for _, ag := range agents {
				ui.Plain("  - %s (%s)", ag.Name, ag.Description)
			}
		} else {
			ui.Plain("  No agents currently registered.")
		}
	} else {
		ui.Plain("  Agent registry not available.")
	}

	ui.Plain("\nExamples:")
	ui.Plain("  vibe agent list")
	ui.Plain("  vibe agent select")
	ui.Plain("  vibe agent run MyAgent")
	ui.Plain("  vibe agent info MyAgent")
}
