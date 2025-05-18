package ui

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
)


var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}


const (
	SuccessSymbol = "✓"
	ErrorSymbol   = "✗"
	WarningSymbol = "⚠"
	PendingSymbol = "⋯"
)


type TerminalAnimator struct {
	spinnerIndex int
	spinnerSpeed time.Duration
	isActive     bool
	messages     map[string]string
	statuses     map[string]string
}


func NewAnimator() *TerminalAnimator {
	return &TerminalAnimator{
		spinnerIndex: 0,
		spinnerSpeed: 100 * time.Millisecond,
		isActive:     false,
		messages:     make(map[string]string),
		statuses:     make(map[string]string),
	}
}


func (ta *TerminalAnimator) StartAnimation(title string) {
	ta.isActive = true
	HeaderStyle.Printf("\n%s\n\n", title)
	go ta.animate()
}


func (ta *TerminalAnimator) AddItem(id string, message string) {
	ta.messages[id] = message
	ta.statuses[id] = "pending"
}



func (ta *TerminalAnimator) UpdateStatus(id string, status string) {
	ta.statuses[id] = status
}


func (ta *TerminalAnimator) StopAnimation(summary string) {
	ta.isActive = false
	
	time.Sleep(ta.spinnerSpeed * 2)

	
	ta.clearAnimationArea()
	ta.renderFinalState()

	
	fmt.Println()
	SuccessStyle.Println(summary)
	fmt.Println()
}


func (ta *TerminalAnimator) animate() {
	for ta.isActive {
		ta.clearAnimationArea()
		ta.renderAnimationFrame()
		ta.spinnerIndex = (ta.spinnerIndex + 1) % len(spinnerFrames)
		time.Sleep(ta.spinnerSpeed)
	}
}


func (ta *TerminalAnimator) clearAnimationArea() {
	
	linesUp := len(ta.messages)
	if linesUp > 0 {
		
		fmt.Printf("\033[%dA\033[J", linesUp+1)
	}
}


func (ta *TerminalAnimator) renderAnimationFrame() {
	
	ids := make([]string, 0, len(ta.messages))
	for id := range ta.messages {
		ids = append(ids, id)
	}

	
	sort.Strings(ids)

	
	for _, id := range ids {
		message := ta.messages[id]
		status := ta.statuses[id]

		switch status {
		case "success":
			SuccessStyle.Printf("%s ", SuccessSymbol)
			fmt.Printf("%s\n", message)
		case "error":
			ErrorStyle.Printf("%s ", ErrorSymbol)
			fmt.Printf("%s\n", message)
		case "warning":
			WarnStyle.Printf("%s ", WarningSymbol)
			fmt.Printf("%s\n", message)
		default: 
			color.New(color.FgBlue).Printf("%s ", spinnerFrames[ta.spinnerIndex])
			fmt.Printf("%s\n", message)
		}
	}
	
	fmt.Println()
}


func (ta *TerminalAnimator) renderFinalState() {
	
	ids := make([]string, 0, len(ta.messages))
	for id := range ta.messages {
		ids = append(ids, id)
	}

	
	sort.Strings(ids)

	
	successCount := 0
	errorCount := 0
	warningCount := 0

	
	for _, id := range ids {
		message := ta.messages[id]
		status := ta.statuses[id]

		switch status {
		case "success":
			SuccessStyle.Printf("%s ", SuccessSymbol)
			fmt.Printf("%s\n", message)
			successCount++
		case "error":
			ErrorStyle.Printf("%s ", ErrorSymbol)
			fmt.Printf("%s\n", message)
			errorCount++
		case "warning":
			WarnStyle.Printf("%s ", WarningSymbol)
			fmt.Printf("%s\n", message)
			warningCount++
		default:
			
			WarnStyle.Printf("%s ", PendingSymbol)
			fmt.Printf("%s\n", message)
			warningCount++
		}
	}

	
	fmt.Println()
	fmt.Printf("Summary: ")
	if successCount > 0 {
		SuccessStyle.Printf("%d successful ", successCount)
	}
	if warningCount > 0 {
		WarnStyle.Printf("%d warnings ", warningCount)
	}
	if errorCount > 0 {
		ErrorStyle.Printf("%d errors ", errorCount)
	}
	fmt.Println()
}
