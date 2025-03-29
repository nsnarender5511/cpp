# Code Structure

> 📂 **Detailed overview of the crules codebase organization**

## Directory Structure

```
crules/
├── cmd/                 # Command-line application entry point
│   └── main.go          # Main application file
├── internal/            # Internal packages
│   ├── core/            # Core business logic
│   │   ├── registry.go  # Registry management
│   │   └── sync.go      # Synchronization logic
│   ├── ui/              # User interface components
│   │   └── console.go   # Console UI
│   └── utils/           # Utility functions
│       ├── files.go     # File operations
│       └── logger.go    # Logging utilities
├── .env                 # Environment variables for development
├── .env.example         # Example environment configuration
├── go.mod               # Go module file
└── go.sum               # Go module checksum file
```

## Package Overview

### cmd

The `cmd` package contains the entry point of the application and command-line interface handling.

**Key Files**:
- `main.go`: Main application entry point, command-line argument parsing, and command routing

**Key Functions**:
- `main()`: Application entry point
- `handleInit()`, `handleMerge()`, `handleSync()`, `handleList()`, `handleClean()`: Command handlers
- `printUsage()`: Displays usage information
- `handleCommandError()`: Standardized error handling for commands

### internal/core

The `core` package contains the core business logic of the application.

#### registry.go

**Purpose**: Manages the registry of projects with Cursor rules.

**Key Types**:
- `Registry`: Manages the list of registered projects
- `Project`: Represents a project entry in the registry

**Key Functions**:
- `NewRegistry()`: Creates a new registry instance
- `AddProject()`: Adds a project to the registry
- `RemoveProject()`: Removes a project from the registry
- `GetProjects()`: Returns all registered projects
- `Load()`: Loads the registry from disk
- `Save()`: Saves the registry to disk

#### sync.go

**Purpose**: Handles synchronization of rules between the main location and project locations.

**Key Types**:
- `SyncManager`: Manages rule synchronization operations
- `SyncOptions`: Configuration options for synchronization

**Key Functions**:
- `NewSyncManager()`: Creates a new sync manager
- `Init()`: Initializes a project with rules from the main location
- `Merge()`: Merges project rules to the main location and syncs to other projects
- `Sync()`: Syncs rules from the main location to the current project
- `Clean()`: Removes non-existent projects from the registry

### internal/ui

The `ui` package contains components for user interface and display.

**Key Functions**:
- `PrintBanner()`: Prints the application banner
- `Info()`, `Success()`, `Warning()`, `Error()`: Output functions for different message types
- `Header()`, `Plain()`: Formatting utilities for console output

### internal/utils

The `utils` package contains utility functions used throughout the application.

**Key Functions**:
- File operations: `FileExists()`, `DirExists()`, `CopyFile()`, `CopyDir()`
- Path utilities: `GetAppPaths()`, `GetCurrentDir()`
- Logging: `InitLogger()`, `Info()`, `Debug()`, `Warn()`, `Error()`
- Configuration: `LoadEnv()`, `GetEnv()`

## Flow of Control

### Application Startup

1. Application starts in `cmd/main.go:main()`
2. Command-line flags are parsed (debug, verbose)
3. Environment is loaded and configured
4. Logger is initialized
5. Banner is displayed
6. Sync manager is created
7. Command is determined from arguments
8. Appropriate command handler is called

### Command Execution

1. Command handler (e.g., `handleInit()`) is invoked
2. Command is logged
3. Appropriate sync manager method is called (e.g., `manager.Init()`)
4. Results are displayed to the user
5. Success or error is returned to the OS

## Key Interfaces and Types

### SyncManager

```go
type SyncManager struct {
    registry    *Registry
    mainPath    string
    projectPath string
    // ... other fields
}

func NewSyncManager() (*SyncManager, error)
func (m *SyncManager) Init() error
func (m *SyncManager) Merge() error
func (m *SyncManager) Sync() error
func (m *SyncManager) Clean() (int, error)
func (m *SyncManager) GetRegistry() *Registry
```

### Registry

```go
type Registry struct {
    Projects []string
    path     string
}

func NewRegistry(path string) (*Registry, error)
func (r *Registry) AddProject(path string) bool
func (r *Registry) RemoveProject(path string) bool
func (r *Registry) GetProjects() []string
func (r *Registry) Load() error
func (r *Registry) Save() error
```

## Configuration Management

Configuration is managed through environment variables, loaded from:
1. System environment
2. `.env` file in the project root

Key configuration variables are defined in the [Configuration Guide](../user-guide/configuration.md).

## Error Handling

Error handling follows a consistent pattern:
1. Functions return errors when they encounter problems
2. Command handlers process these errors
3. Errors are logged with appropriate context
4. User-friendly error messages are displayed
5. Application exits with appropriate exit code

## Related Documentation

- [Architecture](architecture.md) - High-level system design
- [Contributing Guide](contributing.md) - How to contribute to the codebase
- [Testing Guide](testing.md) - Testing approach and practices
