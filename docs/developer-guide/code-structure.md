---
version: v0.1.0
last_updated: 2023-03-29
applies_to: cursor++ v0.1.0+
---

# Code Structure

> 📚 This document provides an overview of how the codebase is organized, to help developers understand and navigate the project.

## Directory Structure

The cursor++ codebase is organized into the following main directories:

```
cursor++/
├── .github/         # GitHub workflows and templates
├── cmd/             # Command-line interface entry points
├── docs/            # Documentation
├── internal/        # Internal packages
│   ├── agent/       # Agent system implementation
│   ├── constants/   # Common constants
│   ├── core/        # Core logic and functionality
│   ├── git/         # Git integration utilities
│   ├── ui/          # User interface utilities
│   ├── utils/       # Utility functions
│   └── version/     # Version information
└── LICENSE          # License file
```

## Key Components

### Command Line Interface

- `cmd/main.go`: The main entry point for the cursor++ CLI tool, implementing all commands and their functionality

The main file implements the command-line interface using a simple flag-based approach, with commands as the first argument.

### Internal Packages (internal/)

#### Agent System (internal/agent/)

The agent system is implemented in the `internal/agent/` package:

- `registry.go`: Manages the collection of available agents
- `loader.go`: Handles loading agents from their definition files
- `types.go`: Defines the data structures for agent metadata
- `context.go`: Manages agent execution context and environment

#### Core Functionality (internal/core/)

The `internal/core/` package implements the core functionality:

- Handles rule synchronization
- Manages project configuration
- Implements command business logic

#### Git Integration (internal/git/)

The `internal/git/` package provides Git integration:

- Handles Git operations
- Manages version control for rules

#### UI Components (internal/ui/)

The `internal/ui/` package implements the user interface:

- Terminal UI components
- Interactive selection interfaces
- Display formatting

#### Utilities (internal/utils/)

The `internal/utils/` package provides common utility functions:

- File system operations
- Configuration handling
- Logging and error handling

#### Version Information (internal/version/)

The `internal/version/` package manages version information:

- Provides version details
- Handles version checking and comparison

### Core Components

#### Registry (`internal/core/registry.go`)

The Registry manages the list of projects and their associated rules. Key features:

- Project management (add, remove, list)
- Thread-safe operations with mutex protection
- Project count and validation functionality

This component acts as the central database for the cursor++ system.

#### Utils (`internal/utils/`)

Utility functions for file operations, configuration, and other common tasks:

- `config.go`: Configuration management
- `dirutils.go`: Directory utilities, including path validation, file counting, and recursive operations
- `logger.go`: Logging with color support

The `dirutils.go` file includes several useful utilities for directory operations:
- `DirExists`: Checks if a directory exists
- `CountFiles`: Counts files in a directory (non-recursive)
- `CountFilesByExt`: Counts files with a specific extension
- `CountFilesRecursive`: Recursively counts all files in a directory
- `EnsureDirExists`: Creates a directory if it doesn't exist

#### Agent System (`internal/agent/`)

Manages interaction with the agent system:

- `registry.go`: Handles agent discovery and metadata extraction
- `GetRulesDir`: Detects and returns the rules directory, checking multiple possible locations for better reliability

## Key Data Structures

### Agent System

The agent system revolves around these key types:

```go
// Agent Definition
type AgentDefinition struct {
    ID             string   `json:"id"`
    Name           string   `json:"name"`
    Description    string   `json:"description"`
    Capabilities   []string `json:"capabilities"`
    Version        string   `json:"version"`
    DefinitionPath string   `json:"-"`
    Content        string   `json:"-"`
}

// Agent Registry
type Registry struct {
    agents   map[string]*AgentDefinition
    rulesDir string
    config   *utils.Config
}

// Agent Loader
type Loader struct {
    registry *Registry
    config   *utils.Config
}
```

### Configuration

The configuration system uses these types:

```go
// Config represents the global configuration
type Config struct {
    MainLocation string               `json:"mainLocation"`
    Locations    map[string]*Location `json:"locations"`
}

// Location represents a registered project location
type Location struct {
    Path      string    `json:"path"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## Control Flow

### Command Execution

1. `main.go` initializes the root command
2. The root command establishes the command hierarchy
3. When a command is invoked, its `Run` function is called
4. The command implementation performs the requested operation

### Agent Selection Flow

1. `agent.go` defines the `select` subcommand
2. When invoked, it creates a new `selection.Selector`
3. The selector loads agents using the registry
4. It displays the interactive UI for agent selection
5. After selection, it saves the chosen agent's ID for later use

## Extension Points

The codebase is designed with several extension points:

1. **New Commands**: Add new commands by extending the command implementation in `cmd/main.go`
2. **Agent Capabilities**: Extend the agent system by modifying the agent registry in `internal/agent/registry.go`
3. **UI Components**: Add new UI components in `internal/ui/` package

## Development Workflow

When working on cursor++, the typical workflow is:

1. Make code changes
2. Build with `go build ./cmd/cursor++`
3. Run tests with `go test ./...`
4. Test manually with the built binary

## See Also

- [Architecture](./architecture.md)
- [Extending Agents](./extending-agents.md)
- [Contributing Guidelines](./contributing.md)

## Navigation

- Previous: [Building](./building.md)
- Next: [Architecture](./architecture.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md)
