# Core API Reference

> 📘 This document provides detailed API reference for the core components of cursor++.

## Overview

The Core API contains the fundamental components that power the cursor++ system, including:

- **SyncManager**: Handles initialization of agent rules directory
- **Registry**: Manages project registration and tracking
- **Parser**: Processes agent definition files
- **Storage**: Handles file operations and rule storage

## SyncManager

The `SyncManager` is the central component that handles initialization of agent rules, setting up the main location for agent definitions.

### Types

```go
// SyncManager handles initialization operations
type SyncManager struct {
    // internal fields
}
```

### Functions

#### NewSyncManager

```go
func NewSyncManager() (*SyncManager, error)
```

**Returns**:
- `*SyncManager`: A pointer to the new SyncManager instance
- `error`: An error if the operation fails

**Description**:
Creates a new SyncManager instance to handle agent rule initialization.

#### Init

```go
func (sm *SyncManager) Init() error
```

**Returns**:
- `error`: An error if the initialization fails

**Description**:
Initializes the agent rules system by setting up the main location for agent rule storage. This method will guide the user through creating or selecting a main location for agent rules.

#### GetRegistry

```go
func (sm *SyncManager) GetRegistry() (*agent.Registry, error)
```

**Returns**:
- `*agent.Registry`: A pointer to the agent registry
- `error`: An error if the operation fails

**Description**:
Returns the agent registry associated with the SyncManager, which can be used to access and manage agent definitions.

## Registry

The `Registry` keeps track of all projects using cursor++.

### Types

```go
// Registry keeps track of all projects using cursor++
type Registry struct {
    Projects []string `json:"projects"`
    path     string   // path to registry file
    config   *utils.Config
}
```

### Functions

#### LoadRegistry

```go
func LoadRegistry(registryPath string, config *utils.Config) (*Registry, error)
```

Loads or creates the registry.

**Parameters:**
- `registryPath`: The path to the registry file
- `config`: Configuration options

**Returns:**
- `*Registry`: A pointer to the loaded registry
- `error`: An error if loading fails

**Example:**
```go
config := utils.LoadConfig()
registry, err := core.LoadRegistry("/path/to/registry.json", config)
if err != nil {
    log.Fatalf("Failed to load registry: %v", err)
}
```

#### AddProject

```go
func (r *Registry) AddProject(projectPath string) error
```

Adds a project to the registry.

**Parameters:**
- `projectPath`: The path to the project to add

**Returns:**
- `error`: An error if adding the project fails

**Example:**
```go
err := registry.AddProject("/path/to/project")
if err != nil {
    log.Fatalf("Failed to add project: %v", err)
}
```

#### GetProjects

```go
func (r *Registry) GetProjects() []string
```

Returns all registered projects.

**Returns**:
- `[]string`: A slice of project paths

**Description**:
Returns a list of all registered project paths.

**Example:**
```go
projects := registry.GetProjects()
for _, project := range projects {
    fmt.Println(project)
}
```

#### CleanProjects

```go
func (r *Registry) CleanProjects() (int, error)
```

Removes projects that no longer exist.

**Returns:**
- `int`: The number of projects removed
- `error`: An error if cleaning fails

**Example:**
```go
removed, err := registry.CleanProjects()
if err != nil {
    log.Fatalf("Failed to clean projects: %v", err)
}
fmt.Printf("Removed %d non-existent projects\n", removed)
```

#### IsRegistered

```go
func (r *Registry) IsRegistered(path string) bool
```

**Parameters**:
- `path`: The path to check

**Returns**:
- `bool`: True if the path is registered, false otherwise

**Description**:
Checks if a path is registered in the registry.

## Parser

The `Parser` processes content into structured rules.

### Types

```go
// Rule represents a parsed agent rule
type Rule struct {
    Name        string
    Description string
    Content     string
    Format      string
    Source      string
}

// ParserFunc defines the function signature for all domain handlers
type ParserFunc func(content []byte, source string) ([]Rule, error)
```

### Functions

#### ParseRules

```go
func ParseRules(content []byte, source string, isWeb bool) ([]Rule, error)
```

Processes content into structured rules.

**Parameters:**
- `content`: The content to parse
- `source`: The source URL or file
- `isWeb`: Whether the content is from a web source

**Returns:**
- `[]Rule`: A slice of parsed rules
- `error`: An error if parsing fails

**Example:**
```go
content, _ := os.ReadFile("agent.mdc")
rules, err := core.ParseRules(content, "agent.mdc", false)
if err != nil {
    log.Fatalf("Failed to parse rules: %v", err)
}
```

## Storage

The `Storage` component handles file operations and rule storage.

### Functions

#### StoreRules

```go
func StoreRules(rules []Rule, forceOverwrite bool) error
```

Saves parsed rules to the filesystem.

**Parameters:**
- `rules`: The rules to store
- `forceOverwrite`: Whether to overwrite existing files

**Returns:**
- `error`: An error if storing fails

**Example:**
```go
err := core.StoreRules(rules, false)
if err != nil {
    log.Fatalf("Failed to store rules: %v", err)
}
```

#### StoreRulesToPath

```go
func StoreRulesToPath(rules []Rule, baseDir string, forceOverwrite bool) error
```

Saves parsed rules to a specific directory path.

**Parameters:**
- `rules`: The rules to store
- `baseDir`: The directory to store rules in
- `forceOverwrite`: Whether to overwrite existing files

**Returns:**
- `error`: An error if storing fails

**Example:**
```go
err := core.StoreRulesToPath(rules, "/path/to/rules", false)
if err != nil {
    log.Fatalf("Failed to store rules: %v", err)
}
```

## Utility Interfaces

### DomainRegistry

```go
// DomainRegistry maps domains to their specific parser functions
var DomainRegistry = map[string]ParserFunc{
    "cursor.directory": parseCursorDirectory,
    // More domains can be added here later
}
```

The `DomainRegistry` maps domains to their specific parser functions. You can extend this map to support additional domains.

**Example (Adding a new domain):**
```go
// Add support for example.com
DomainRegistry["example.com"] = parseExampleDotCom
```

## Error Handling

The Core API functions return errors that should be handled by the caller. Common error patterns include:

1. **File System Errors**: Errors related to file operations (e.g., file not found, permission denied)
2. **Parse Errors**: Errors related to parsing content (e.g., invalid format, missing required fields)
3. **Network Errors**: Errors related to network operations (e.g., connection timeout, invalid URL)

**Example (Error Handling):**
```go
err := manager.Init()
if err != nil {
    if strings.Contains(err.Error(), "permission denied") {
        // Handle permission error
        log.Fatalf("Permission denied: %v", err)
    } else if strings.Contains(err.Error(), "not found") {
        // Handle not found error
        log.Fatalf("File not found: %v", err)
    } else {
        // Handle other errors
        log.Fatalf("Error initializing: %v", err)
    }
}
```

## Best Practices

1. **Error Handling**: Always check and handle errors returned by Core API functions
2. **Configuration**: Use the configuration system to customize behavior
3. **Registry Management**: Clean the registry periodically to remove non-existent projects
4. **Path Handling**: Use absolute paths when possible to avoid confusion

## See Also

- [Agent API Reference](./agent-api.md)
- [Utility API Reference](./utils-api.md)
- [Architecture](../developer-guide/architecture.md)
- [Building from Source](../developer-guide/building.md)

## Navigation

- Previous: [Building from Source](../developer-guide/building.md)
- Next: [Agent API Reference](./agent-api.md)
- Up: [API Reference](./README.md)
- Home: [Documentation Home](../README.md) 