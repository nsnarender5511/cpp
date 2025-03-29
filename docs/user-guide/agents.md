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

- `crules agent list` - Display all available agents in a concise list
- `crules agent info <id>` - Show detailed information about a specific agent
- `crules agent select` - Interactively select and load an agent

### Agent List

The `list` command displays all available agents in a concise format:

```
$ crules agent list

Available agents (6):
   1. 🧙‍♂️ Technical Wizard Agent          2. ✨ Feature Planner Agent
   3. 🔍 Fix Planner Agent               4. 🛠️ Implementer Agent
   5. 🏃 Runner Agent                    6. 📚 Documentation Agent
```

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