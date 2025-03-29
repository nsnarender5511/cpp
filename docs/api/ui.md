# UI API Reference

> 📌 **Summary**: Terminal user interface API documentation for the crules system.

## Overview

The UI package provides terminal-based user interface components and interactions for the crules system, enabling interactive agent selection and displaying information in a user-friendly format.

## UI Components

### Terminal Interface

Functions for setting up and managing the terminal interface.

```go
// NewTerminal creates a new terminal interface
func NewTerminal() (*Terminal, error)

// Close closes the terminal interface
func (t *Terminal) Close()

// Clear clears the terminal screen
func (t *Terminal) Clear()
```

### Agent Selection

Components for displaying and selecting agents interactively.

```go
// SelectAgent displays the agent selection interface and returns the selected agent
func (t *Terminal) SelectAgent(agents []Agent) (Agent, error)

// DisplayAgentInfo displays detailed information about an agent
func (t *Terminal) DisplayAgentInfo(agent Agent)
```

### List Display

Components for displaying lists of items such as projects or agents.

```go
// DisplayList displays a list of items with optional headers
func (t *Terminal) DisplayList(items []string, header string)

// DisplayTable displays tabular data with headers
func (t *Terminal) DisplayTable(headers []string, rows [][]string)
```

### Progress Display

Components for displaying progress information for long-running operations.

```go
// StartProgress starts a progress indicator with the given message
func (t *Terminal) StartProgress(message string) *Progress

// UpdateProgress updates the progress indicator with a new message
func (p *Progress) Update(message string)

// StopProgress stops the progress indicator
func (p *Progress) Stop()
```

## Style Constants

The UI package defines several style constants for consistent visual styling:

- `StyleHeader`: Style for headers
- `StyleSuccess`: Style for success messages
- `StyleError`: Style for error messages
- `StyleWarning`: Style for warning messages
- `StyleInfo`: Style for informational messages
- `StyleHighlight`: Style for highlighted items

## Example Usage

```go
// Create a new terminal interface
terminal, err := ui.NewTerminal()
if err != nil {
    log.Fatalf("Failed to create terminal: %v", err)
}
defer terminal.Close()

// Display a list of projects
projects := []string{"/path/to/project1", "/path/to/project2"}
terminal.DisplayList(projects, "Registered Projects")

// Display progress for a long-running operation
progress := terminal.StartProgress("Synchronizing rules...")
// ... perform synchronization ...
progress.Update("Finalizing synchronization...")
// ... complete synchronization ...
progress.Stop()
```
