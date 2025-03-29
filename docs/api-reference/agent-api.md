---
version: v0.1.0
last_updated: 2023-03-29
applies_to: crules v0.1.0+
---

# Agent API Reference

> 📘 This document provides detailed API reference for the agent system components of crules.

## Overview

The Agent API provides functionality for working with the agent system in crules, including:

- **Registry**: Manages the collection of available agents
- **Loader**: Handles loading and initializing agents
- **Context**: Provides shared context between agents
- **Agent Definitions**: Defines agent metadata and content

## Agent Types

### AgentDefinition

```go
// AgentDefinition represents a single agent's definition and metadata
type AgentDefinition struct {
    ID             string   `json:"id"`
    Name           string   `json:"name"`
    Description    string   `json:"description"`
    Capabilities   []string `json:"capabilities"`
    Version        string   `json:"version"`
    DefinitionPath string   `json:"-"` // Path to the .mdc file
    Content        string   `json:"-"` // The actual content of the agent definition
}
```

The `AgentDefinition` struct represents a single agent's metadata and content.

### Agent

```go
// Agent represents a loaded and initialized agent
type Agent struct {
    Definition *AgentDefinition
    Context    *AgentContext
}
```

The `Agent` struct represents a loaded and initialized agent.

### AgentContext

```go
// AgentContext represents shared context between agents
type AgentContext struct {
    Data        map[string]interface{}
    LastUpdated time.Time
}
```

The `AgentContext` struct provides a shared context for agents to store and retrieve data.

## Registry

The `Registry` manages the collection of available agents.

### Types

```go
// Registry manages the collection of available agents
type Registry struct {
    agents   map[string]*AgentDefinition
    rulesDir string
    config   *utils.Config
}
```

### Functions

#### NewRegistry

```go
func NewRegistry(config *utils.Config, rulesDir string) (*Registry, error)
```

Creates a new agent registry.

**Parameters:**
- `config`: Configuration options
- `rulesDir`: The directory containing agent definition files

**Returns:**
- `*Registry`: A pointer to the new registry
- `error`: An error if initialization fails

**Example:**
```go
config := utils.LoadConfig()
registry, err := agent.NewRegistry(config, "/path/to/rules")
if err != nil {
    log.Fatalf("Failed to create registry: %v", err)
}
```

#### GetAgent

```go
func (r *Registry) GetAgent(id string) (*AgentDefinition, bool)
```

Returns an agent by ID.

**Parameters:**
- `id`: The ID of the agent to retrieve

**Returns:**
- `*AgentDefinition`: A pointer to the agent definition
- `bool`: Whether the agent exists

**Example:**
```go
agentDef, exists := registry.GetAgent("wizard")
if !exists {
    log.Fatalf("Agent 'wizard' not found")
}
fmt.Printf("Agent: %s - %s\n", agentDef.Name, agentDef.Description)
```

#### ListAgents

```go
func (r *Registry) ListAgents() []*AgentDefinition
```

Returns all available agents.

**Returns:**
- `[]*AgentDefinition`: A slice of agent definitions

**Example:**
```go
agents := registry.ListAgents()
for _, agent := range agents {
    fmt.Printf("Agent: %s - %s\n", agent.Name, agent.Description)
}
```

#### AgentExists

```go
func (r *Registry) AgentExists(id string) bool
```

Checks if an agent with the given ID exists.

**Parameters:**
- `id`: The ID of the agent to check

**Returns:**
- `bool`: Whether the agent exists

**Example:**
```go
if registry.AgentExists("wizard") {
    fmt.Println("Wizard agent is available")
} else {
    fmt.Println("Wizard agent is not available")
}
```

## Loader

The `Loader` handles loading and initializing agents.

### Types

```go
// Loader handles loading and initializing agents
type Loader struct {
    registry *Registry
    config   *utils.Config
}
```

### Functions

#### NewLoader

```go
func NewLoader(registry *Registry, config *utils.Config) *Loader
```

Creates a new agent loader.

**Parameters:**
- `registry`: The agent registry
- `config`: Configuration options

**Returns:**
- `*Loader`: A pointer to the new loader

**Example:**
```go
config := utils.LoadConfig()
registry, _ := agent.NewRegistry(config, "/path/to/rules")
loader := agent.NewLoader(registry, config)
```

#### LoadAgent

```go
func (l *Loader) LoadAgent(id string) (*Agent, error)
```

Loads and initializes an agent by ID.

**Parameters:**
- `id`: The ID of the agent to load

**Returns:**
- `*Agent`: A pointer to the loaded agent
- `error`: An error if loading fails

**Example:**
```go
agent, err := loader.LoadAgent("wizard")
if err != nil {
    log.Fatalf("Failed to load agent: %v", err)
}
fmt.Printf("Loaded agent: %s\n", agent.Definition.Name)
```

#### LoadAgentWithContext

```go
func (l *Loader) LoadAgentWithContext(id string, context *AgentContext) (*Agent, error)
```

Loads an agent and initializes it with the provided context.

**Parameters:**
- `id`: The ID of the agent to load
- `context`: The context to initialize the agent with

**Returns:**
- `*Agent`: A pointer to the loaded agent
- `error`: An error if loading fails

**Example:**
```go
context := agent.NewAgentContext()
context.Set("key", "value")
agent, err := loader.LoadAgentWithContext("wizard", context)
if err != nil {
    log.Fatalf("Failed to load agent: %v", err)
}
```

## Agent Context

The `AgentContext` provides a shared context for agents to store and retrieve data.

### Functions

#### NewAgentContext

```go
func NewAgentContext() *AgentContext
```

Creates a new agent context.

**Returns:**
- `*AgentContext`: A pointer to the new context

**Example:**
```go
context := agent.NewAgentContext()
```

#### Set

```go
func (c *AgentContext) Set(key string, value interface{})
```

Stores a value in the context.

**Parameters:**
- `key`: The key to store the value under
- `value`: The value to store

**Example:**
```go
context := agent.NewAgentContext()
context.Set("username", "john")
context.Set("preferences", map[string]string{"theme": "dark"})
```

#### Get

```go
func (c *AgentContext) Get(key string) (interface{}, bool)
```

Retrieves a value from the context.

**Parameters:**
- `key`: The key to retrieve the value for

**Returns:**
- `interface{}`: The retrieved value
- `bool`: Whether the value exists

**Example:**
```go
context := agent.NewAgentContext()
context.Set("username", "john")

value, exists := context.Get("username")
if exists {
    username := value.(string)
    fmt.Printf("Username: %s\n", username)
}
```

## Agent Manager Interface

The `AgentManager` interface defines methods for agent management.

### Interface

```go
// AgentManager interface defines methods for agent management
type AgentManager interface {
    ListAgents() ([]*AgentDefinition, bool)
    GetAgent(id string) (*AgentDefinition, bool)
    LoadAgent(id string) (*Agent, error)
}
```

This interface can be implemented by custom agent management systems.

## Agent File Format

Agent definition files (`.mdc`) follow a specific format:

```markdown
# Agent Name

## 🎯 Role:
One-line description of the agent's primary role

## 📝 Description:
A more detailed description of what the agent does, its purpose, and when to use it.

## ✅ Capabilities:

### ✅ Capability 1
Description of the first capability

### ✅ Capability 2
Description of the second capability
```

The agent definition files are parsed to extract:
- **Name**: The name of the agent (level-1 heading)
- **Role**: A description of the agent's role (heading with 🎯 emoji)
- **Description**: A detailed description of the agent
- **Capabilities**: A list of the agent's capabilities (headings with ✅ emoji)

## Integration with Core System

The Agent system integrates with the Core system in several ways:

1. **Agent Discovery**: The Registry scans the rules directory for `.mdc` files
2. **Agent Loading**: The Loader retrieves agent definitions from the Registry
3. **Agent Selection**: The UI displays agent information from the Registry
4. **Agent Persistence**: The selected agent is stored in the configuration

**Example (Integration Flow):**
```go
// Initialize the system
manager, _ := core.NewSyncManager()
config := utils.LoadConfig()
appPaths := utils.GetAppPaths("crules")
rulesDir := appPaths.GetRulesDir(config.RulesDirName)

// Create registry and loader
registry, _ := agent.NewRegistry(config, rulesDir)
loader := agent.NewLoader(registry, config)

// List available agents
agents := registry.ListAgents()
for i, agent := range agents {
    fmt.Printf("%d. %s - %s\n", i+1, agent.Name, agent.Description)
}

// Load selected agent
selectedAgent, _ := loader.LoadAgent("wizard")
fmt.Printf("Selected agent: %s\n", selectedAgent.Definition.Name)
```

## Extending the Agent System

You can extend the Agent system in several ways:

1. **Custom Agent Definitions**: Create custom `.mdc` files with agent definitions
2. **Custom Agent Parsers**: Implement custom parsing logic for agent definition files
3. **Custom Agent Managers**: Implement the `AgentManager` interface for custom agent management

### Creating Custom Agent Parsers

To create a custom agent parser:

```go
func customAgentParser(content []byte, path string) (*AgentDefinition, error) {
    // Parse the content and extract agent metadata
    // ...
    
    return &AgentDefinition{
        ID:             id,
        Name:           name,
        Description:    description,
        Capabilities:   capabilities,
        Version:        version,
        DefinitionPath: path,
        Content:        string(content),
    }, nil
}
```

### Implementing Custom Agent Managers

To implement a custom agent manager:

```go
type CustomAgentManager struct {
    // Custom fields
}

func (m *CustomAgentManager) ListAgents() ([]*AgentDefinition, bool) {
    // Custom implementation
    return agents, true
}

func (m *CustomAgentManager) GetAgent(id string) (*AgentDefinition, bool) {
    // Custom implementation
    return agent, true
}

func (m *CustomAgentManager) LoadAgent(id string) (*Agent, error) {
    // Custom implementation
    return agent, nil
}
```

## Best Practices

1. **Error Handling**: Always check and handle errors returned by Agent API functions
2. **Context Sharing**: Use the context system to share data between agents
3. **Agent Definition**: Follow the format guidelines for agent definition files
4. **Registry Management**: Scan for new agents when rules are updated

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