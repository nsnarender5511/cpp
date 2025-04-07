package ui

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

// Version information
var Version = "1.0.0" // This would be injected at build time

// PrintBanner displays the application banner
func PrintBanner() {
	// Create a new figure with a clearer font
	fig := figure.NewFigure("vibe", "big", true)

	// Convert to string and get the banner text
	bannerText := fig.String()

	// Create color functions for gradient effect
	blueColor := color.New(color.FgBlue, color.Bold).SprintFunc()
	cyanColor := color.New(color.FgCyan, color.Bold).SprintFunc()

	// Apply gradient effect to the banner
	lines := strings.Split(bannerText, "\n")
	gradientBanner := ""

	for i, line := range lines {
		if line == "" {
			continue
		}

		// Apply gradient based on line position
		if i < len(lines)/2 {
			gradientBanner += blueColor(line) + "\n"
		} else {
			gradientBanner += cyanColor(line) + "\n"
		}
	}

	// Print the banner with gradient
	fmt.Println()
	fmt.Print(gradientBanner)

	// Add a styled subtitle with version
	subtitle := fmt.Sprintf("Multi-Agent System v%s", Version)

	// Center the subtitle under the banner
	bannerWidth := 60 // Approximate width of the banner

	// Create a styled box for the subtitle
	boxWidth := len(subtitle) + 8
	boxPadding := strings.Repeat(" ", (bannerWidth-boxWidth)/2)
	boxTop := boxPadding + "╭" + strings.Repeat("─", boxWidth) + "╮"
	boxMiddle := boxPadding + "│  " + subtitle + "  │"
	boxBottom := boxPadding + "╰" + strings.Repeat("─", boxWidth) + "╯"

	// Print the styled box
	color.New(color.FgCyan).Println(boxTop)
	color.New(color.FgCyan, color.Bold).Println(boxMiddle)
	color.New(color.FgCyan).Println(boxBottom)

	// Add a decorative border
	fmt.Println()
	borderWidth := bannerWidth
	border := strings.Repeat("━", borderWidth)
	color.New(color.FgBlue).Println(border)
	fmt.Println()
}
