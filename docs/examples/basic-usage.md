# Basic Usage Examples

> 🧪 This guide provides practical examples of common cursor++ workflows.

## Overview

The examples in this guide demonstrate real-world scenarios and common tasks you'll encounter when using cursor++. Each example includes:

- A brief description of the task
- Step-by-step instructions
- Expected output
- Tips and variations

## Example 1: Setting Up a New Project

This example shows how to initialize a new project with agents.

### Scenario

You're starting a new project and want to set it up with cursor++ agents.

### Steps

1. Create a new project directory:
   ```bash
   mkdir ~/projects/new-project
   cd ~/projects/new-project
   ```

2. Initialize a git repository (optional but recommended):
   ```bash
   git init
   ```

3. Initialize cursor++:
   ```bash
   cursor++ init
   ```

4. Verify initialization by checking the agents list:
   ```bash
   cursor++ agent
   ```

### Expected Output

```
✨ cursor++ - Cursor Agent Manager ✨

Initializing directory with cursor++ agents...

✓ Successfully initialized with cursor++ agents

Found 1 available agent in the current directory.
Run 'cursor++ agent' to see available agents.
```

### Tips

- If you're using version control, make sure to add your preferred agent configuration to the repository
- You can use `cursor++ agent select` to set a default agent for your project

## Example 2: Using the Agent Selection Interface

This example demonstrates how to interactively select and use agents.

### Scenario

You want to choose an agent for a specific task from the available agents.

### Steps

1. Navigate to your project:
   ```bash
   cd ~/projects/my-project
   ```

2. List available agents:
   ```bash
   cursor++ agent
   ```

3. Select an agent:
   ```bash
   cursor++ agent select
   ```

4. Use arrow keys to navigate to your desired agent and press Enter to select it.

### Expected Output

For `cursor++ agent`:
```
✨ cursor++ - Cursor Agent Manager ✨

Available agents:
1. 🔄 Document Syncer Agent
```

For `cursor++ agent select`:
```
✨ cursor++ - Cursor Agent Manager ✨

Select an agent:
> 1. 🔄 Document Syncer Agent

[Use arrow keys to navigate, Enter to select]
```

After selection:

```
Selected agent: 🔄 Document Syncer Agent

Agent details:
  ID:          doc-syncer
  Name:        🔄 Document Syncer Agent
  Version:     1.0

Description:
  This agent specializes in synchronizing documentation between codebases
  and documentation systems, ensuring consistency and up-to-date information.

Agent saved as default.
```

### Tips

- Use `cursor++ agent info <id>` to see agent details without selecting it
- The last selected agent is remembered for future use
- Agent selection works with both string IDs and numeric indices

## Example 3: Getting Agent Information

This example shows how to view detailed information about a specific agent.

### Scenario

You want to learn more about an agent's capabilities before selecting it for your task.

### Steps

1. First, list available agents to find the ID or index:
   ```bash
   cursor++ agent
   ```

2. Get detailed information about a specific agent:
   ```bash
   cursor++ agent info doc-syncer
   ```
   
   Or by index:
   ```bash
   cursor++ agent info 1
   ```

### Expected Output

```
✨ cursor++ - Cursor Agent Manager ✨

Agent details:
  ID:          doc-syncer
  Name:        🔄 Document Syncer Agent
  Version:     1.0

Description:
  This agent specializes in synchronizing documentation between codebases
  and documentation systems, ensuring consistency and up-to-date information.

Capabilities:
  ✅ Context Detection and Analysis
  ✅ Codebase-to-Docs Synchronization  
  ✅ Website Support File Synchronization
  ✅ Conflict Resolution
  
Instructions:
  When synchronizing documentation, I will:
  1. Determine execution context (codebase or website)
  2. Map documentation files to website references
  3. Detect changes between contexts
  4. Update documentation while maintaining cross-references
```

### Tips

- Use this command to explore the capabilities of the agent
- Review the agent's instructions to understand how it processes requests

## Example 4: Using the Document Syncer Agent in Cursor

This example shows how to use the Document Syncer agent directly in the Cursor chat interface.

### Scenario

You need to synchronize documentation between your codebase and documentation system.

### Steps

1. In the Cursor IDE, open your project with documentation files.

2. In the chat interface, reference the agent using the `@` syntax:
   ```
   @doc-syncer.mdc I need to check for inconsistencies between our code and documentation
   ```

3. The agent will analyze your codebase and documentation, then provide a synchronization plan.

4. Follow up with specific instructions:
   ```
   @doc-syncer.mdc Please update the README.md file to match our current implementation
   ```

### Expected Output

The agent will respond in the chat interface with its analysis and proposed changes. You can review and approve these changes before they're applied to your documentation.

### Tips

- Be specific about which files need synchronization
- Provide context about the codebase and documentation structure
- Ask the agent to create a synchronization plan before making changes

## Navigation

- Previous: [Troubleshooting](../user-guide/troubleshooting.md)
- Next: [Advanced Usage](./advanced-usage.md)
- Up: [Examples](./README.md)
- Home: [Documentation Home](../README.md)
