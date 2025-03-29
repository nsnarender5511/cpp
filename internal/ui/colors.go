package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	// Initialize color outputs - exported for direct use
	SuccessStyle = color.New(color.FgGreen, color.Bold)
	ErrorStyle   = color.New(color.FgRed, color.Bold)
	WarnStyle    = color.New(color.FgYellow)
	InfoStyle    = color.New(color.FgCyan)
	HeaderStyle  = color.New(color.FgBlue, color.Bold)

	// Internal references for consistency
	successStyle = SuccessStyle
	errorStyle   = ErrorStyle
	warnStyle    = WarnStyle
	infoStyle    = InfoStyle
	headerStyle  = HeaderStyle
)

// Success prints a success message
func Success(format string, a ...interface{}) {
	successStyle.Printf("✓ "+format+"\n", a...)
}

// Error prints an error message
func Error(format string, a ...interface{}) {
	errorStyle.Printf("✗ "+format+"\n", a...)
}

// Warning prints a warning message
func Warning(format string, a ...interface{}) {
	warnStyle.Printf("⚠ "+format+"\n", a...)
}

// Info prints an informational message
func Info(format string, a ...interface{}) {
	infoStyle.Printf("ℹ "+format+"\n", a...)
}

// Header prints a header text
func Header(format string, a ...interface{}) {
	headerStyle.Printf(format+"\n", a...)
}

// Plain prints plain text without styling
func Plain(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}
