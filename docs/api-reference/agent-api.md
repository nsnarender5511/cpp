---
version: v0.1.0
last_updated: 2023-03-29
applies_to: crules v0.1.0+
---

# Agent API

> 🤖 The Agent API provides methods for discovering, loading, and managing AI agent definitions in cursor++.

## Overview

The Agent API is located in the `cursor++/internal/agent` package and provides functionality for:

- Discovering available agents from `.mdc` files in the rules directory
- Loading agent definitions with their metadata and capabilities
- Managing agent selection and preferences
- Handling agent-related operations in the cursor++ system

## Key Types

### Agent

The `Agent` struct represents a single AI agent with its metadata and definition:

```go
type Agent struct {
    // Unique identifier for the agent
    ID string
    
    // Human-readable name of the agent
    Name string
    
    // Short description of the agent
    Description string
    
    // Version of the agent
    Version string
    
    // Categories this agent belongs to
    Categories []string
    
    // Capabilities this agent has
    Capabilities []string
    
    // Path to the .mdc file containing the agent definition
    FilePath string
    
    // Raw content of the agent definition
    Content string
}
```

### Registry

The `Registry` interface provides methods for managing agents:

```go
type Registry interface {
    // LoadAgents loads all agents from the rules directory
    LoadAgents(ctx context.Context) ([]Agent, error)
    
    // GetAgentByID returns an agent by its ID
    GetAgentByID(id string) (Agent, bool)
    
    // GetAgentByName returns an agent by its name (case-insensitive)
    GetAgentByName(name string) (Agent, bool)
    
    // GetAgentByIndex returns an agent by its position in the list
    GetAgentByIndex(index int) (Agent, bool)
    
    // GetAllAgents returns all loaded agents
    GetAllAgents() []Agent
    
    // GetLastSelectedAgent returns the last selected agent
    GetLastSelectedAgent() (Agent, bool)
    
    // SetLastSelectedAgent sets the last selected agent
    SetLastSelectedAgent(agent Agent) error
}
```

## Key Functions

### Creating and Loading

```go
// NewRegistry creates a new agent registry
func NewRegistry() (Registry, error)

// NewRegistryWithOptions creates a new agent registry with custom options
func NewRegistryWithOptions(options RegistryOptions) (Registry, error)

// ParseAgentFile parses an agent from a .mdc file
func ParseAgentFile(filePath string) (Agent, error)

// LoadAgentContent loads the full content of an agent definition
func LoadAgentContent(agent *Agent) error
```

### Agent Management

```go
// FilterAgentsByCategory returns agents that belong to the specified category
func FilterAgentsByCategory(agents []Agent, category string) []Agent

// FilterAgentsByCapability returns agents that have the specified capability
func FilterAgentsByCapability(agents []Agent, capability string) []Agent

// SortAgentsByName sorts agents alphabetically by name
func SortAgentsByName(agents []Agent) []Agent

// SortAgentsByCategory sorts agents by category, then by name
func SortAgentsByCategory(agents []Agent) []Agent
```

### Display and Selection

```go
// DisplayAgentList shows a list of agents in the terminal
func DisplayAgentList(agents []Agent, options AgentDisplayOptions) error

// SelectAgent presents an interactive UI for selecting an agent
func SelectAgent(registry Registry) (Agent, error)

// DisplayAgentInfo shows detailed information about an agent
func DisplayAgentInfo(agent Agent) error
```

## Examples

### Loading All Agents

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "cursor++/internal/agent"
)

func main() {
    // Create a context
    ctx := context.Background()
    
    // Create a registry
    registry, err := agent.NewRegistry()
    if err != nil {
        log.Fatalf("Failed to create agent registry: %v", err)
    }
    
    // Load all agents
    agents, err := registry.LoadAgents(ctx)
    if err != nil {
        log.Fatalf("Failed to load agents: %v", err)
    }
    
    // Display the loaded agents
    fmt.Printf("Loaded %d agents:\n", len(agents))
    for i, a := range agents {
        fmt.Printf("%d. %s (ID: %s)\n", i+1, a.Name, a.ID)
    }
}
```

### Selecting an Agent

```go
package main

import (
    "fmt"
    "log"
    
    "cursor++/internal/agent"
)

func main() {
    // Create a registry
    registry, err := agent.NewRegistry()
    if err != nil {
        log.Fatalf("Failed to create agent registry: %v", err)
    }
    
    // Present UI for selecting an agent
    selectedAgent, err := agent.SelectAgent(registry)
    if err != nil {
        log.Fatalf("Failed to select agent: %v", err)
    }
    
    // Display the selected agent
    fmt.Printf("Selected agent: %s\n", selectedAgent.Name)
    fmt.Printf("Description: %s\n", selectedAgent.Description)
    
    // Save the selection as the last selected agent
    if err := registry.SetLastSelectedAgent(selectedAgent); err != nil {
        log.Fatalf("Failed to save agent selection: %v", err)
    }
}
```

### Displaying Agent Information

```go
package main

import (
    "context"
    "log"
    
    "cursor++/internal/agent"
)

func main() {
    // Create a context
    ctx := context.Background()
    
    // Create a registry
    registry, err := agent.NewRegistry()
    if err != nil {
        log.Fatalf("Failed to create agent registry: %v", err)
    }
    
    // Load agents
    if _, err := registry.LoadAgents(ctx); err != nil {
        log.Fatalf("Failed to load agents: %v", err)
    }
    
    // Get an agent by ID
    wizard, exists := registry.GetAgentByID("wizard")
    if !exists {
        log.Fatalf("Wizard agent not found")
    }
    
    // Display detailed agent information
    if err := agent.DisplayAgentInfo(wizard); err != nil {
        log.Fatalf("Failed to display agent info: %v", err)
    }
}
```

## Reference Tables

### Agent Categories

Common agent categories used in cursor++:

| Category | Description |
|----------|-------------|
| `planning` | Agents focused on architectural and feature planning |
| `implementation` | Agents that write or modify code |
| `review` | Agents that review code or documentation |
| `support` | Agents that provide auxiliary functionality |
| `coordination` | Agents that coordinate between other agents |

### Agent Display Options

Options for customizing agent display:

| Option | Type | Description |
|--------|------|-------------|
| `TermWidth` | `int` | Width of the terminal in characters |
| `ShowCategories` | `bool` | Whether to display agent categories |
| `ShowCapabilities` | `bool` | Whether to display agent capabilities |
| `CompactMode` | `bool` | Use compact display format |
| `GroupByCategory` | `bool` | Group agents by their categories |

## Error Handling

The Agent API uses standard Go error handling patterns:

```go
agents, err := registry.LoadAgents(ctx)
if err != nil {
    if errors.Is(err, agent.ErrNoAgentsFound) {
        fmt.Println("No agents found in the rules directory")
    } else if errors.Is(err, agent.ErrInvalidAgentFile) {
        fmt.Println("Found invalid agent definitions")
    } else {
        fmt.Printf("Failed to load agents: %v\n", err)
    }
}
```

Common errors returned by the Agent API include:

- `ErrNoAgentsFound`: No agent definitions found in the rules directory
- `ErrInvalidAgentFile`: Invalid agent definition file
- `ErrAgentNotFound`: Agent with specified ID, name, or index not found
- `ErrAgentParseFailed`: Failed to parse agent definition

## Design Patterns

The Agent API follows these design patterns:

1. **Repository Pattern**: The Registry interface provides a uniform way to access agents
2. **Factory Pattern**: NewRegistry functions create instances of the Registry interface
3. **Options Pattern**: Functions accept option structs for customization
4. **Iterator Pattern**: The Registry allows iteration over all agents

## Thread Safety

The Registry implementation is not thread-safe. When using the Agent API in a concurrent context, synchronize access to the registry using mutexes or other synchronization mechanisms.

## See Also

- [Core API Reference](./core-api.md)
- [Utility API Reference](./utils-api.md)
- [Extending the Agent System](../developer-guide/extending-agents.md)
- [Agent System Overview](../user-guide/agents.md)

---

## Navigation

- Previous: [Core API](./core-api.md)
- Next: [UI API](./ui-api.md)
- Up: [API Reference](./README.md)
- Home: [Documentation Home](../README.md) 