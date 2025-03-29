# Agent System

> 🧠 The Agent System in crules allows you to work with specialized AI agents for different tasks, from planning and implementing features to fixing issues and documenting code.

## Overview

The crules Agent System provides an interactive way to discover, select, and use various AI agents defined in your `.cursor/rules` directory. Each agent is specialized for a particular task, such as:

- Planning new features 
- Fixing bugs and issues
- Implementing code based on plans
- Running and testing implementations
- Creating documentation
- Reviewing code

Agents are defined in Markdown (`.mdc`) files that contain both the agent's metadata (name, description, capabilities) and its full definition.

## Using Agents

You can interact with agents using the following commands:

- `crules agent` - Display all available agents (default behavior)
- `crules agent info <id>` - Show detailed information about a specific agent
- `crules agent select` - Interactively select and load an agent

You can also reference agents directly in the chatbox using the `@` symbol (e.g., `@wizard.mdc`).

## Listing Available Agents

To see all available agents, use:

```bash
crules agent
```

This will display a formatted table of all agents with their information. The table adapts to your terminal width to provide the optimal display format:

### Display Examples

**Wide Terminal Display:**
```
+-----+---------------------+--------------------+----------+
| No. | Agent Name          | Reference ID       | Version  |
+-----+---------------------+--------------------+----------+
| 1   | Feature Planner     | @feature-planner.mdc | 1.0    |
| 2   | Fix Planner         | @fix-planner.mdc     | 1.0    |
| 3   | Runner              | @runner.mdc          | 1.0    |
| 4   | Technical Wizard    | @wizard.mdc          | 1.0    |
| ... | ...                 | ...                  | ...    |
+-----+---------------------+--------------------+----------+
```

**Medium Terminal Display:**
```
+-----+----------------+--------------------+
| No. | Name           | Reference          |
+-----+----------------+--------------------+
| 1   | Feature Planner| @feature-planner.mdc |
| 2   | Fix Planner    | @fix-planner.mdc     |
| 3   | Runner         | @runner.mdc          |
| ... | ...            | ...                  |
+-----+----------------+--------------------+
```

**Narrow Terminal Display:**
```
+-----+------------------+
| No. | Agent            |
+-----+------------------+
| 1   | feature-planner  |
| 2   | fix-planner      |
| 3   | runner           |
| ... | ...              |
+-----+------------------+
```

The table format ensures proper alignment and clarity regardless of your terminal size, making agent information easy to read and reference.

## Referencing Agents

There are multiple ways to reference agents in crules, providing flexibility based on your preference and needs.

### Referencing by String ID

You can reference agents by their unique string ID, which is derived from the filename without the `.mdc` extension:

```bash
crules agent info wizard
```

This method is stable across sessions and reorderings of the agent list.

### Referencing by Numeric Index

You can also reference agents by their position number shown in the agent list:

```bash
crules agent info 1  # References the first agent in the list
```

This method is convenient for quick access when you can see the number in the list but might not remember the exact ID.

Benefits of numeric index referencing:
- Shorter to type
- Easier for sequential exploration of agents
- Direct visual correspondence with the displayed list

### Agent @ References

The quickest way to invoke a specific agent in the chatbox is by using the `@` reference. Simply type `@` followed by the agent ID to invoke that agent:

```
@wizard.mdc I need help designing a new API endpoint
```

```
@quick-answer-agent.mdc What does the HTTP 418 status code mean?
```

This method allows you to quickly switch between different agent specializations without running additional commands.

### Agent Info

The `info` command provides detailed information about a specific agent:

```
$ crules agent info wizard

Agent details:
  ID:          wizard
  Name:        🧙‍♂️ Technical Wizard Agent
  Version:     1.0

Description:
  [Full agent description with formatted markdown]

Capabilities:
  - In-Depth Technical Exploration and Analysis
  - Expert Architectural Guidance
  - Design Patterns Discussion
  - Clean Code Advisory

File: /Users/username/Library/Application Support/crules/.cursor/rules/wizard.mdc
```

Or using the numeric index:

```
$ crules agent info 6  # Assuming 6 is the position number for the wizard agent

Agent details:
  ID:          wizard
  Name:        🧙‍♂️ Technical Wizard Agent
  Version:     1.0
  
  ...
```

### Agent Selection

The `select` command presents an interactive menu for choosing an agent:

1. Run `crules agent select`
2. Browse the list of available agents
3. Enter the number of the agent you want to select
4. The agent will be loaded, and you'll see its details
5. You can optionally view the full agent definition in a paginated format

## Common Agents

The system comes with several pre-defined agents, each specialized in different aspects of software development:

### 🧙‍♂️ Technical Wizard
- Provides high-level technical guidance and coordinates other agents
- Helps with architecture decisions, design patterns, and clean code principles
- Acts as the primary entry point for complex tasks

### ✨ Feature Planner
- Plans the implementation of new features and enhancements
- Breaks down feature requirements into implementable components
- Creates detailed plans for the Implementer Agent to follow

### 🔍 Fix Planner
- Analyzes bugs and issues to find their root causes
- Develops comprehensive step-by-step plans to fix problems
- Creates detailed guidance for implementing fixes

### 🛠️ Implementer
- Translates detailed plans into working code
- Follows established coding standards and patterns
- Focuses on precise implementation of planning agent instructions

### 🏃 Runner
- Executes and verifies implementations
- Runs tests and validates results
- Provides feedback on the success or failure of implementations

### 📚 Documentation Agent
- Creates and maintains comprehensive documentation
- Ensures documentation is aligned with the current codebase
- Organizes documentation in a logical and accessible structure

## Adding Custom Agents

You can add your own custom agents by creating new `.mdc` files in your `.cursor/rules` directory. These agents will automatically be discovered by the crules tool and become available in the agent selection menu.

For information on creating custom agents, see the [Extending the Agent System](../developer-guide/extending-agents.md) guide in the Developer Documentation. 