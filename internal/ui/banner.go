package ui

import (
	"strings"
)

// Version information
var Version = "1.0.0" // This would be injected at build time

// PrintBanner displays the application banner
func PrintBanner() {
	banner := `
   _____           _            
  / ____|         | |           
 | |     _ __ _   | | ___  ___  
 | |    | '__| | | |/ _ \/ __| 
 | |____| |  | |_| |  __/\__ \ 
  \_____|_|   \__,_|\___||___/ 
`

	Header(banner)
	Info("Cursor Rules Manager v%s", Version)
	Plain(strings.Repeat("─", 40))
}
