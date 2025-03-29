# Extending the Agent System

> 🛠️ This guide explains how to create custom agents and extend the agent system in crules.

## Overview

The Agent System in crules is designed to be extensible. You can create your own specialized agents by defining them in Markdown (`.mdc`) files and placing them in the `.cursor/rules` directory.

## Agent Definition Files

Agents are defined in `.mdc` files (Markdown with custom annotations). Each file contains:

1. Metadata about the agent (name, description, capabilities)
2. The agent's full definition and instructions

### Basic Structure

A minimal agent definition file looks like this:

```markdown
# Agent Name

## 🎯 Role:
Brief description of what this agent does and its purpose.

### ✅ Capability One:
- Detailed point about this capability
- Another point about this capability

### ✅ Capability Two:
- Detailed point about this capability
- Another point about this capability
```

### Required Sections

For proper parsing and display in the crules tool, your agent definition should include:

1. **Title (H1)**: The agent's name, typically with an emoji
2. **Role (H2)**: A section titled "🎯 Role:" that describes the agent's purpose
3. **Capabilities (H3)**: Sections titled "✅ [Capability Name]:" that list the agent's capabilities

### Format Guidelines

- Use emojis to make your agent visually distinctive
- Structure your agent definition with clear headings
- Use bullet points for lists of capabilities or responsibilities
- Keep the first paragraph of the description concise (it will be shown in truncated views)

## How crules Discovers Agents

When you run agent-related commands, crules:

1. Scans the `.cursor/rules` directory for `.mdc` files
2. Parses each file to extract agent metadata
3. Creates an in-memory registry of available agents
4. Uses this registry for listing, selecting, and loading agents

The agent ID is derived from the filename (without the `.mdc` extension).

## Integration with the System

### Agent Registry

The `agent.Registry` type manages the collection of available agents:

```go
// from internal/agent/registry.go
type Registry struct {
    agents   map[string]*AgentDefinition
    rulesDir string
    config   *utils.Config
}
```

### Agent Definition

Agent metadata is stored in the `AgentDefinition` struct:

```go
// from internal/agent/types.go
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

### Agent Loading

When an agent is selected, it's loaded by the `agent.Loader`:

```go
// from internal/agent/loader.go
type Loader struct {
    registry *Registry
    config   *utils.Config
}
```

## Creating Your Own Agent

To create your own custom agent:

1. Create a new `.mdc` file in the `.cursor/rules` directory
2. Name the file to match your agent's ID (e.g., `my-custom-agent.mdc`)
3. Structure the file using the format guidelines above
4. Add detailed instructions for your agent's specific purpose

### Example Custom Agent

Here's an example of a custom agent for SQL query optimization:

```markdown
# 🔍 SQL Optimizer Agent

## 🎯 Role:
You are a specialized **SQL Optimizer Agent**, an expert in database query optimization. Your primary purpose is to analyze SQL queries, identify performance bottlenecks, and suggest optimizations to improve query efficiency.

### ✅ Query Analysis:
- Examine SQL queries to understand their purpose and structure
- Identify inefficient query patterns and potential bottlenecks
- Analyze table structures and join conditions

### ✅ Optimization Recommendations:
- Suggest index improvements for better query performance
- Recommend query restructuring for better execution plans
- Provide alternative query approaches with performance benefits

### ✅ Performance Explanation:
- Explain the performance implications of different SQL constructs
- Clarify how database optimizers handle specific query patterns
- Describe execution plans in understandable terms
```

## Testing Your Custom Agent

After creating your agent definition file:

1. Run `crules agent list` to verify your agent appears in the list
2. Run `crules agent info <your-agent-id>` to check that metadata is parsed correctly
3. Use `crules agent select` to interactively select and test your agent

## Advanced Customization

For more advanced customization of the agent system, you may need to modify the core code:

- `internal/agent/types.go`: Define new agent-related types
- `internal/agent/registry.go`: Customize agent discovery and registration
- `internal/agent/loader.go`: Modify how agents are loaded and initialized

## Best Practices

1. **Keep agent definitions focused**: Create specialized agents with clear purposes
2. **Use descriptive names**: Make agent purposes clear from their names
3. **Structure definitions consistently**: Follow the format guidelines for better user experience
4. **Document capabilities clearly**: Explicitly state what the agent can and cannot do
5. **Use example formatting**: Include example formats for expected inputs/outputs

## See Also

- [Agent System Overview](../user-guide/agents.md)
- [Code Structure](./code-structure.md)
- [Architecture](./architecture.md) 