package ui

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
)

// Spinner animation frames
var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// Animation status symbols
const (
	SuccessSymbol = "✓"
	ErrorSymbol   = "✗"
	WarningSymbol = "⚠"
	PendingSymbol = "⋯"
)

// TerminalAnimator manages terminal animations for improved user experience
type TerminalAnimator struct {
	spinnerIndex int
	spinnerSpeed time.Duration
	isActive     bool
	messages     map[string]string
	statuses     map[string]string
}

// NewAnimator creates a new terminal animator
func NewAnimator() *TerminalAnimator {
	return &TerminalAnimator{
		spinnerIndex: 0,
		spinnerSpeed: 100 * time.Millisecond,
		isActive:     false,
		messages:     make(map[string]string),
		statuses:     make(map[string]string),
	}
}

// StartAnimation begins a loading animation for a sequence of items
func (ta *TerminalAnimator) StartAnimation(title string) {
	ta.isActive = true
	HeaderStyle.Printf("\n%s\n\n", title)
	go ta.animate()
}

// AddItem adds a new item to the animation with "pending" status
func (ta *TerminalAnimator) AddItem(id string, message string) {
	ta.messages[id] = message
	ta.statuses[id] = "pending"
}

// UpdateStatus updates the status of an item
// status can be "success", "error", "warning", or "pending"
func (ta *TerminalAnimator) UpdateStatus(id string, status string) {
	ta.statuses[id] = status
}

// StopAnimation ends the animation and displays final statuses
func (ta *TerminalAnimator) StopAnimation(summary string) {
	ta.isActive = false
	// Wait a moment to ensure the last animation frame completes
	time.Sleep(ta.spinnerSpeed * 2)

	// Clear the animation area and display final state
	ta.clearAnimationArea()
	ta.renderFinalState()

	// Display summary
	fmt.Println()
	SuccessStyle.Println(summary)
	fmt.Println()
}

// animate handles the animation loop in a goroutine
func (ta *TerminalAnimator) animate() {
	for ta.isActive {
		ta.clearAnimationArea()
		ta.renderAnimationFrame()
		ta.spinnerIndex = (ta.spinnerIndex + 1) % len(spinnerFrames)
		time.Sleep(ta.spinnerSpeed)
	}
}

// clearAnimationArea clears the animation area for redrawing
func (ta *TerminalAnimator) clearAnimationArea() {
	// Move cursor up by the number of items plus a blank line
	linesUp := len(ta.messages)
	if linesUp > 0 {
		// +1 for the blank line after the last item
		fmt.Printf("\033[%dA\033[J", linesUp+1)
	}
}

// renderAnimationFrame renders the current state of the animation
func (ta *TerminalAnimator) renderAnimationFrame() {
	// Sort the ids to ensure consistent order
	ids := make([]string, 0, len(ta.messages))
	for id := range ta.messages {
		ids = append(ids, id)
	}

	// Sort ids alphabetically for consistent display
	sort.Strings(ids)

	// Display each item with appropriate animation
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
		default: // pending
			color.New(color.FgBlue).Printf("%s ", spinnerFrames[ta.spinnerIndex])
			fmt.Printf("%s\n", message)
		}
	}
	// Add a blank line after the items
	fmt.Println()
}

// renderFinalState renders the final state without animations
func (ta *TerminalAnimator) renderFinalState() {
	// Sort the ids to ensure consistent order
	ids := make([]string, 0, len(ta.messages))
	for id := range ta.messages {
		ids = append(ids, id)
	}

	// Sort ids alphabetically for consistent display
	sort.Strings(ids)

	// Get counts for summary
	successCount := 0
	errorCount := 0
	warningCount := 0

	// Display each item with final status
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
			// Any remaining pending items shown as warnings
			WarnStyle.Printf("%s ", PendingSymbol)
			fmt.Printf("%s\n", message)
			warningCount++
		}
	}

	// Add a summary line
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
