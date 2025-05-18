package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	
	SuccessStyle = color.New(color.FgGreen, color.Bold)
	ErrorStyle   = color.New(color.FgRed, color.Bold)
	WarnStyle    = color.New(color.FgYellow)
	InfoStyle    = color.New(color.FgCyan)
	HeaderStyle  = color.New(color.FgBlue, color.Bold)
	PromptStyle  = color.New(color.FgMagenta, color.Bold)

	
	successStyle = SuccessStyle
	errorStyle   = ErrorStyle
	warnStyle    = WarnStyle
	infoStyle    = InfoStyle
	headerStyle  = HeaderStyle
	promptStyle  = PromptStyle
)


func Success(format string, a ...interface{}) {
	successStyle.Printf("✓ "+format+"\n", a...)
}


func Error(format string, a ...interface{}) {
	errorStyle.Printf("✗ "+format+"\n", a...)
}


func Warning(format string, a ...interface{}) {
	warnStyle.Printf("⚠ "+format+"\n", a...)
}


func Info(format string, a ...interface{}) {
	infoStyle.Printf("ℹ "+format+"\n", a...)
}


func Header(format string, a ...interface{}) {
	headerStyle.Printf(format+"\n", a...)
}


func Plain(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}


func Prompt(format string, a ...interface{}) {
	promptStyle.Printf("> "+format, a...)
}
