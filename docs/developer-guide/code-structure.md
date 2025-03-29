# Code Structure

> 📁 This guide explains the organization of the codebase and the purpose of each component.

## Directory Structure

The crules codebase is organized into the following main directories:

```
crules/
├── cmd/            # Command-line interface entry points
├── internal/       # Internal packages
│   ├── agent/      # Agent system implementation
│   ├── cli/        # CLI command implementations
│   ├── config/     # Configuration management
│   ├── projects/   # Project management
│   └── utils/      # Utility functions and helpers
├── docs/           # Documentation
└── assets/         # Static assets
```

## Key Components

### Command Line Interface (cmd/)

The `cmd/` directory contains the entry points for the command-line interface:

- `cmd/crules/main.go`: The main entry point for the crules CLI tool

### Internal Packages (internal/)

#### Agent System (internal/agent/)

The agent system is implemented in the `internal/agent/` package:

- `registry.go`: Manages the collection of available agents
- `loader.go`: Handles loading agents from their definition files
- `types.go`: Defines the data structures for agent metadata
- `parser.go`: Parses agent definition files (`.mdc`) to extract metadata
- `selection.go`: Implements the interactive agent selection UI

#### CLI Commands (internal/cli/)

The `internal/cli/` package implements the command handlers:

- `root.go`: Defines the root command and global options
- `init.go`: Implements the `init` command
- `merge.go`: Implements the `merge` command
- `sync.go`: Implements the `sync` command
- `list.go`: Implements the `list` command
- `clean.go`: Implements the `clean` command
- `agent.go`: Implements the agent-related commands

#### Configuration (internal/config/)

The `internal/config/` package handles configuration management:

- `config.go`: Defines the configuration data structures
- `loader.go`: Loads configuration from files

#### Project Management (internal/projects/)

The `internal/projects/` package manages project registration:

- `registry.go`: Handles project registration and tracking
- `sync.go`: Synchronizes rules between projects

#### Utilities (internal/utils/)

The `internal/utils/` package provides common utility functions:

- `fs.go`: File system operations
- `ui.go`: User interface helpers
- `terminal.go`: Terminal interaction utilities

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

1. **New Commands**: Add new subcommands by creating a new file in `internal/cli/`
2. **Agent Capabilities**: Extend the agent system by modifying the parser or adding new agent types
3. **UI Components**: Add new UI components in `internal/utils/ui.go`

## Development Workflow

When working on crules, the typical workflow is:

1. Make code changes
2. Build with `go build ./cmd/crules`
3. Run tests with `go test ./...`
4. Test manually with the built binary

## See Also

- [Architecture](./architecture.md)
- [Extending Agents](./extending-agents.md)
- [Contributing Guidelines](./contributing.md)
