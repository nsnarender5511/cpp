package cli

import "flag"

var (
	DebugFlag          *bool
	MultiAgentModeFlag *bool
	VersionShortFlag   *bool
	AgentCmd           *flag.FlagSet
)


func SetupFlags() {
	DebugFlag = flag.Bool("debug", false, "Show debug messages on console")
	MultiAgentModeFlag = flag.Bool("multi-agent", false, "Enable multi-agent mode for collaborative tasks.")
	VersionShortFlag = flag.Bool("v", false, "Display the application version and exit (shorthand).")

	AgentCmd = flag.NewFlagSet("agent", flag.ExitOnError)

	
	flag.Usage = func() {
		PrintUsage() 
	}
}


func ParseFlags() {
	flag.Parse()
}
