package ui

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)


var Version = "1.0.0" 


func PrintBanner() {
	
	fig := figure.NewFigure("vibe", "big", true)

	
	bannerText := fig.String()

	
	blueColor := color.New(color.FgBlue, color.Bold).SprintFunc()
	cyanColor := color.New(color.FgCyan, color.Bold).SprintFunc()

	
	lines := strings.Split(bannerText, "\n")
	gradientBanner := ""

	for i, line := range lines {
		if line == "" {
			continue
		}

		
		if i < len(lines)/2 {
			gradientBanner += blueColor(line) + "\n"
		} else {
			gradientBanner += cyanColor(line) + "\n"
		}
	}

	
	fmt.Println()
	fmt.Print(gradientBanner)

	
	subtitle := fmt.Sprintf("Multi-Agent System v%s", Version)

	
	bannerWidth := 60 

	
	boxWidth := len(subtitle) + 8
	boxPadding := strings.Repeat(" ", (bannerWidth-boxWidth)/2)
	boxTop := boxPadding + "╭" + strings.Repeat("─", boxWidth) + "╮"
	boxMiddle := boxPadding + "│  " + subtitle + "  │"
	boxBottom := boxPadding + "╰" + strings.Repeat("─", boxWidth) + "╯"

	
	color.New(color.FgCyan).Println(boxTop)
	color.New(color.FgCyan, color.Bold).Println(boxMiddle)
	color.New(color.FgCyan).Println(boxBottom)

	
	fmt.Println()
	borderWidth := bannerWidth
	border := strings.Repeat("━", borderWidth)
	color.New(color.FgBlue).Println(border)
	fmt.Println()
}
