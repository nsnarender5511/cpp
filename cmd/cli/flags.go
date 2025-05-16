package cli

import "flag"

var (
	DebugFlag          *bool
	MultiAgentModeFlag *bool
	VersionShortFlag   *bool
	AgentCmd           *flag.FlagSet
)

// SetupFlags defines all the command-line flags for the application.
func SetupFlags() {
	DebugFlag = flag.Bool("debug", false, "Show debug messages on console")
	MultiAgentModeFlag = flag.Bool("multi-agent", false, "Enable multi-agent mode for collaborative tasks.")
	VersionShortFlag = flag.Bool("v", false, "Display the application version and exit (shorthand).")

	AgentCmd = flag.NewFlagSet("agent", flag.ExitOnError)

	// Custom usage function to match our existing help format
	flag.Usage = func() {
		PrintUsage() // This function is in cmd/cli/ui.go
	}
}

// ParseFlags parses the command-line flags.
func ParseFlags() {
	flag.Parse()
}
