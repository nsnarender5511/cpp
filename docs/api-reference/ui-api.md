---
title: UI API
description: User interface utilities
applies_to: cursor++ v0.1.0+
---

# UI API

The UI API provides utilities for building user interfaces in cursor++.

## Terminal UI Components

### `PrintSuccess`

```go
func PrintSuccess(message string)
```

Prints a success message to the terminal.

**Parameters:**
- `message` - The success message to display

### `PrintError`

```go
func PrintError(message string)
```

Prints an error message to the terminal.

**Parameters:**
- `message` - The error message to display

### `PrintWarning`

```go
func PrintWarning(message string)
```

Prints a warning message to the terminal.

**Parameters:**
- `message` - The warning message to display

## Overview

The UI API provides components and utilities for building the terminal user interface in cursor++. It includes functions for displaying banners, messages, and interactive components like selectors.

This reference covers the most important types and functions in the UI package located at `/internal/ui`.

## Color Utilities

### `Colors` Type

```go
type Colors struct {
    Reset      string
    Bold       string
    Underline  string
    Reversed   string
    Black      string
    Red        string
    Green      string
    Yellow     string
    Blue       string
    Purple     string
    Cyan       string
    White      string
    BlackBg    string
    RedBg      string
    GreenBg    string
    YellowBg   string
    BlueBg     string
    PurpleBg   string
    CyanBg     string
    WhiteBg    string
}
```

The `Colors` struct provides ANSI color codes for terminal output styling.

### `GetColors` Function

```go
func GetColors() Colors
```

Returns a struct containing ANSI color codes for terminal output. These colors can be used to format text in terminal output.

#### Example

```go
colors := ui.GetColors()
fmt.Printf("%sError:%s Something went wrong\n", colors.Red, colors.Reset)
```

### `DisableColors` Function

```go
func DisableColors()
```

Disables colors in terminal output by setting all color codes to empty strings.

#### Example

```go
// Disable colors for terminals that don't support them
if !supportsColors {
    ui.DisableColors()
}
```

## Banner Functions

### `PrintBanner` Function

```go
func PrintBanner(version string)
```

Prints the application banner with version information.

#### Parameters

- `version string`: The application version to display

#### Example

```go
ui.PrintBanner("v0.1.0")
```

Output:
```
✨ cursor++ - Cursor Rules Manager v0.1.0 ✨
```

### `PrintAppName` Function

```go
func PrintAppName()
```

Prints just the application name in a stylized format.

#### Example

```go
ui.PrintAppName()
```

Output:
```
✨ cursor++ ✨
```

## Message Functions

### `PrintSuccess` Function

```go
func PrintSuccess(message string)
```

Prints a success message with a green checkmark.

#### Parameters

- `message string`: The success message to display

#### Example

```go
ui.PrintSuccess("Successfully initialized rules")
```

Output:
```
✓ Successfully initialized rules
```

### `PrintError` Function

```go
func PrintError(message string)
```

Prints an error message in red.

#### Parameters

- `message string`: The error message to display

#### Example

```go
ui.PrintError("Failed to initialize rules")
```

Output:
```
Error: Failed to initialize rules
```

### `PrintWarning` Function

```go
func PrintWarning(message string)
```

Prints a warning message in yellow.

#### Parameters

- `message string`: The warning message to display

#### Example

```go
ui.PrintWarning("Some projects could not be found")
```

Output:
```
⚠️ Some projects could not be found
```

### `PrintInfo` Function

```go
func PrintInfo(message string)
```

Prints an informational message in blue.

#### Parameters

- `message string`: The informational message to display

#### Example

```go
ui.PrintInfo("Syncing rules from main location...")
```

Output:
```
ℹ️ Syncing rules from main location...
```

## Interactive Components

### `Selector` Type

```go
type Selector struct {
    Items       []string
    Prompt      string
    Selected    int
    TermWidth   int
    MaxItems    int
    ShowNumbers bool
}
```

The `Selector` struct provides an interactive list selector for terminal UIs.

### `NewSelector` Function

```go
func NewSelector(items []string, prompt string) *Selector
```

Creates a new selector with the provided items and prompt.

#### Parameters

- `items []string`: List of items to select from
- `prompt string`: Prompt to display above the selection list

#### Returns

- `*Selector`: A new selector instance

#### Example

```go
agents := []string{"Technical Wizard", "Feature Planner", "Implementer"}
selector := ui.NewSelector(agents, "Select an agent:")
```

### `Run` Method

```go
func (s *Selector) Run() (int, error)
```

Runs the selector and returns the index of the selected item.

#### Returns

- `int`: The index of the selected item
- `error`: Any error that occurred

#### Example

```go
index, err := selector.Run()
if err != nil {
    return err
}
fmt.Printf("Selected item: %s\n", agents[index])
```

### `SetShowNumbers` Method

```go
func (s *Selector) SetShowNumbers(show bool)
```

Sets whether to show numbers next to the items.

#### Parameters

- `show bool`: Whether to show numbers

#### Example

```go
selector.SetShowNumbers(true)
```

### `SetMaxItems` Method

```go
func (s *Selector) SetMaxItems(max int)
```

Sets the maximum number of items to display at once.

#### Parameters

- `max int`: Maximum items to display

#### Example

```go
selector.SetMaxItems(10)
```

## Table Formatting

### `FormatTable` Function

```go
func FormatTable(headers []string, rows [][]string, termWidth int) string
```

Formats data as a table with column alignment.

#### Parameters

- `headers []string`: Column headers
- `rows [][]string`: Table rows
- `termWidth int`: Terminal width to fit the table within

#### Returns

- `string`: Formatted table as a string

#### Example

```go
headers := []string{"ID", "Name", "Description"}
rows := [][]string{
    {"1", "Technical Wizard", "Helps with technical planning"},
    {"2", "Implementer", "Implements code based on plans"},
}
table := ui.FormatTable(headers, rows, 80)
fmt.Println(table)
```

Output:
```
ID | Name              | Description
---+-------------------+---------------------------
1  | Technical Wizard  | Helps with technical planning
2  | Implementer       | Implements code based on plans
```

## Confirmation Prompts

### `Confirm` Function

```go
func Confirm(prompt string, defaultValue bool) (bool, error)
```

Displays a confirmation prompt with yes/no options.

#### Parameters

- `prompt string`: The prompt to display
- `defaultValue bool`: The default value (true = yes, false = no)

#### Returns

- `bool`: The user's choice
- `error`: Any error that occurred

#### Example

```go
confirm, err := ui.Confirm("Do you want to continue?", true)
if err != nil {
    return err
}
if confirm {
    // User confirmed
} else {
    // User declined
}
```

## Input Functions

### `ReadLine` Function

```go
func ReadLine(prompt string) (string, error)
```

Reads a line of input from the user with a prompt.

#### Parameters

- `prompt string`: The prompt to display

#### Returns

- `string`: The input provided by the user
- `error`: Any error that occurred

#### Example

```go
name, err := ui.ReadLine("Enter your name: ")
if err != nil {
    return err
}
fmt.Printf("Hello, %s!\n", name)
```

## Progress Indicators

### `StartSpinner` Function

```go
func StartSpinner(message string) *Spinner
```

Starts a spinner with the given message.

#### Parameters

- `message string`: The message to display next to the spinner

#### Returns

- `*Spinner`: A pointer to the spinner instance

#### Example

```go
spinner := ui.StartSpinner("Syncing rules...")
// Do some work
spinner.Stop()
```

### `Spinner` Methods

```go
func (s *Spinner) Stop()
func (s *Spinner) Update(message string)
```

- `Stop()`: Stops the spinner and clears the line
- `Update(message string)`: Updates the spinner message

#### Example

```go
spinner := ui.StartSpinner("Downloading...")
// Do some work
spinner.Update("Processing...")
// Do more work
spinner.Stop()
```

## Best Practices

### Using UI Components

- Use appropriate color functions for different message types (success, error, warning, info)
- Ensure terminal width is considered for table formatting
- Provide clear prompts for interactive components
- Always handle errors from interactive components

### Error Handling

Always check for errors when using interactive components:

```go
selection, err := selector.Run()
if err != nil {
    ui.PrintError(err.Error())
    return err
}
```

### Creating Consistent UIs

For consistent UI appearance:

1. Start with a banner using `PrintBanner()`
2. Use info messages to indicate progress
3. Use success messages to indicate completion
4. Format complex data as tables
5. Use selectors for user choices

---

## Navigation

- Previous: [Agent API Reference](./agent-api.md)
- Next: [Utils API](./utils-api.md)
- Up: [API Reference](./README.md)
- Home: [Documentation Home](../README.md) 