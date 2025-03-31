---
title: Utils API
description: Utility functions and helpers
applies_to: cursor++ v0.1.0+
---

# Utils API

The Utils API provides various utility functions for working with the cursor++ system.

## File System Utilities

### `EnsureDirectoryExists`

```go
func EnsureDirectoryExists(path string) error
```

Creates a directory if it doesn't exist.

**Parameters:**
- `path` - The directory path to create

**Returns:**
- `error` - An error if the directory couldn't be created

### `WriteFile`

```go
func WriteFile(filePath string, content []byte) error
```

Writes content to a file, creating the file if it doesn't exist.

**Parameters:**
- `filePath` - The path to the file
- `content` - The content to write

**Returns:**
- `error` - An error if the file couldn't be written

## Configuration

The Configuration utilities handle loading, storing, and accessing application configuration.

### Types

```go
// Config represents the application configuration
type Config struct {
    MainLocation     string `json:"mainLocation"`
    RulesDirName     string `json:"rulesDirName"`
    RegistryFileName string `json:"registryFileName"`
    LastSelectedAgent string `json:"lastSelectedAgent"`
    DirPermission    int    `json:"dirPermission"`
    FilePermission   int    `json:"filePermission"`
    Theme            Theme  `json:"theme"`
}

// Theme represents UI theme configuration
type Theme struct {
    UseColoredOutput bool   `json:"useColoredOutput"`
    AccentColor      string `json:"accentColor"`
}
```

### Functions

#### LoadConfig

```go
func LoadConfig() *Config
```

Loads the configuration from the default location or creates a default configuration.

**Returns:**
- `*Config`: A pointer to the loaded configuration

**Example:**
```go
config := utils.LoadConfig()
fmt.Printf("Main location: %s\n", config.MainLocation)
```

#### SaveConfig

```go
func SaveConfig(config *Config) error
```

Saves the configuration to the default location.

**Parameters:**
- `config`: The configuration to save

**Returns:**
- `error`: An error if saving fails

**Example:**
```go
config := utils.LoadConfig()
config.Theme.AccentColor = "blue"
err := utils.SaveConfig(config)
if err != nil {
    log.Fatalf("Failed to save config: %v", err)
}
```

#### GetConfigPath

```go
func GetConfigPath(appName string) string
```

Returns the path to the configuration file for the specified application.

**Parameters:**
- `appName`: The name of the application

**Returns:**
- `string`: The path to the configuration file

**Example:**
```go
configPath := utils.GetConfigPath("cursor++")
fmt.Println(configPath)
// Output: /home/user/.config/cursor++/config.json (on Linux)
```

## App Paths

The App Paths utilities handle platform-specific path resolution.

### Types

```go
// AppPaths represents application-specific paths
type AppPaths struct {
    configDir  string
    dataDir    string
    cacheDir   string
    appName    string
}
```

### Functions

#### GetAppPaths

```go
func GetAppPaths(appName string) AppPaths
```

Returns application-specific paths for the specified application.

**Parameters:**
- `appName`: The name of the application

**Returns:**
- `AppPaths`: The application-specific paths

**Example:**
```go
paths := utils.GetAppPaths("cursor++")
fmt.Println(paths.ConfigDir)
// Output: /home/user/.config/cursor++
```

#### GetConfigDir

```go
func (ap AppPaths) GetConfigDir() string
```

Returns the configuration directory for the application.

**Returns:**
- `string`: The configuration directory path

**Example:**
```go
paths := utils.GetAppPaths("cursor++")
fmt.Println(paths.ConfigDir)
// Output: /home/user/.config/cursor++
```

#### GetDataDir

```go
func (ap AppPaths) GetDataDir() string
```

Returns the data directory for the application.

**Returns:**
- `string`: The data directory path

**Example:**
```go
paths := utils.GetAppPaths("cursor++")
fmt.Println(paths.DataDir)
// Output: /home/user/.local/share/cursor++
```

#### GetCacheDir

```go
func (ap AppPaths) GetCacheDir() string
```

Returns the cache directory for the application.

**Returns:**
- `string`: The cache directory path

**Example:**
```go
paths := utils.GetAppPaths("cursor++")
fmt.Println(paths.CacheDir)
// Output: /home/user/.cache/cursor++
```

#### GetRulesDir

```go
func (ap AppPaths) GetRulesDir(rulesDirName string) string
```

Returns the rules directory path based on the configured rules directory name.

**Parameters:**
- `rulesDirName`: The name of the rules directory

**Returns:**
- `string`: The rules directory path

**Example:**
```go
paths := utils.GetAppPaths("cursor++")
fmt.Println(paths.LogDir)
// Output: /home/user/.local/state/cursor++/logs
```

## File Operations

The File Operations utilities provide functions for working with files and directories.

### Functions

#### CreateDirIfNotExists

```go
func CreateDirIfNotExists(path string, perm os.FileMode) error
```

Creates a directory if it doesn't exist.

**Parameters:**
- `path`: The path to the directory
- `perm`: The permission mode to use when creating the directory

**Returns:**
- `error`: An error if creation fails

**Example:**
```go
err := utils.CreateDirIfNotExists("/path/to/dir", 0755)
if err != nil {
    log.Fatalf("Failed to create directory: %v", err)
}
```

#### FileExists

```go
func FileExists(path string) bool
```

Checks if a file exists.

**Parameters:**
- `path`: The path to the file

**Returns:**
- `bool`: Whether the file exists

**Example:**
```go
if utils.FileExists("/path/to/file") {
    fmt.Println("File exists")
} else {
    fmt.Println("File does not exist")
}
```

#### DirectoryExists

```go
func DirectoryExists(path string) bool
```

Checks if a directory exists.

**Parameters:**
- `path`: The path to the directory

**Returns:**
- `bool`: Whether the directory exists

**Example:**
```go
if utils.DirectoryExists("/path/to/dir") {
    fmt.Println("Directory exists")
} else {
    fmt.Println("Directory does not exist")
}
```

#### CopyFile

```go
func CopyFile(src, dst string) error
```

Copies a file from the source path to the destination path.

**Parameters:**
- `src`: The source file path
- `dst`: The destination file path

**Returns:**
- `error`: An error if copying fails

**Example:**
```go
err := utils.CopyFile("/path/to/source", "/path/to/destination")
if err != nil {
    log.Fatalf("Failed to copy file: %v", err)
}
```

#### CopyDir

```go
func CopyDir(src, dst string) error
```

Copies a directory recursively from the source path to the destination path.

**Parameters:**
- `src`: The source directory path
- `dst`: The destination directory path

**Returns:**
- `error`: An error if copying fails

**Example:**
```go
err := utils.CopyDir("/path/to/source/dir", "/path/to/destination/dir")
if err != nil {
    log.Fatalf("Failed to copy directory: %v", err)
}
```

## Terminal UI Helpers

The Terminal UI Helpers provide utilities for displaying formatted text in the terminal.

### Functions

#### GetTerminalWidth

```go
func GetTerminalWidth() int
```

Returns the width of the terminal in characters.

**Returns:**
- `int`: The terminal width

**Example:**
```go
width := utils.GetTerminalWidth()
fmt.Printf("Terminal width: %d\n", width)
```

#### PrintBanner

```go
func PrintBanner(appName string, version string)
```

Prints a formatted application banner with the application name and version.

**Parameters:**
- `appName`: The name of the application
- `version`: The version of the application

**Example:**
```go
utils.PrintBanner("cursor++", "v0.1.0")
// Output:
// ✨ cursor++ v0.1.0 ✨
```

#### PrintSuccess

```go
func PrintSuccess(message string)
```

Prints a success message with a checkmark.

**Parameters:**
- `message`: The message to print

**Example:**
```go
utils.PrintSuccess("Operation completed successfully")
```

#### PrintError

```go
func PrintError(message string)
```

Prints an error message with a cross mark.

**Parameters:**
- `message`: The message to print

**Example:**
```go
utils.PrintError("Operation failed")
```

#### PrintWarning

```go
func PrintWarning(message string)
```

Prints a warning message with a warning symbol.

**Parameters:**
- `message`: The message to print

**Example:**
```go
utils.PrintWarning("Proceed with caution")
```

#### PrintTable

```go
func PrintTable(headers []string, rows [][]string)
```

Prints a formatted table with headers and rows.

**Parameters:**
- `headers`: The table headers
- `rows`: The table rows

**Example:**
```go
headers := []string{"ID", "Name", "Description"}
rows := [][]string{
    {"1", "Item 1", "First item"},
    {"2", "Item 2", "Second item"},
}
utils.PrintTable(headers, rows)
```

## Error Handling

The Error Handling utilities provide functions for common error handling patterns.

### Functions

#### HandleError

```go
func HandleError(err error, exitCode int, message string)
```

Handles an error by printing a message and exiting with the specified code if the error is not nil.

**Parameters:**
- `err`: The error to handle
- `exitCode`: The exit code to use if the error is not nil
- `message`: The message to print if the error is not nil

**Example:**
```go
err := someFunction()
utils.HandleError(err, 1, "Failed to execute function")
```

#### WrapError

```go
func WrapError(err error, message string) error
```

Wraps an error with an additional message.

**Parameters:**
- `err`: The original error
- `message`: The message to add to the error

**Returns:**
- `error`: The wrapped error

**Example:**
```go
originalErr := someFunction()
wrappedErr := utils.WrapError(originalErr, "Additional context")
```

## Logging

The Logging utilities provide functions for logging messages to a file.

### Functions

#### InitLogger

```go
func InitLogger(appName string, level string) error
```

Initializes the logger with the specified application name and log level.

**Parameters:**
- `appName`: The name of the application
- `level`: The log level

**Returns:**
- `error`: An error if initialization fails

**Example:**
```go
err := utils.InitLogger("cursor++", "info")
if err != nil {
    fmt.Println("Failed to initialize logger:", err)
}
// Logger initialized at /home/user/.local/state/cursor++/logs/cursor++.log
```

#### LogDebug

```go
func LogDebug(message string, args ...interface{})
```

Logs a debug message.

**Parameters:**
- `message`: The message to log
- `args`: Additional arguments for formatting the message

**Example:**
```go
utils.LogDebug("Debug message: %s", "details")
```

#### LogInfo

```go
func LogInfo(message string, args ...interface{})
```

Logs an informational message.

**Parameters:**
- `message`: The message to log
- `args`: Additional arguments for formatting the message

**Example:**
```go
utils.LogInfo("Info message: %s", "details")
```

#### LogWarning

```go
func LogWarning(message string, args ...interface{})
```

Logs a warning message.

**Parameters:**
- `message`: The message to log
- `args`: Additional arguments for formatting the message

**Example:**
```go
utils.LogWarning("Warning message: %s", "details")
```

#### LogError

```go
func LogError(message string, args ...interface{})
```

Logs an error message.

**Parameters:**
- `message`: The message to log
- `args`: Additional arguments for formatting the message

**Example:**
```go
utils.LogError("Error message: %s", "details")
```

## Best Practices

1. **Configuration Management**: Use the configuration utilities to maintain a consistent configuration across the application
2. **Path Handling**: Use the app paths utilities to ensure correct path resolution across different platforms
3. **Error Handling**: Use the error handling utilities to provide consistent error messages to users
4. **File Operations**: Check for existence before operating on files and directories
5. **Logging**: Use appropriate log levels for different types of messages

## See Also

- [Core API Reference](./core-api.md)
- [Agent API Reference](./agent-api.md)
- [Configuration Guide](../user-guide/configuration.md)
- [Building from Source](../developer-guide/building.md)

---

## Navigation

- Previous: [Agent API Reference](./agent-api.md)
- Up: [API Reference](./README.md)
- Home: [Documentation Home](../README.md) 