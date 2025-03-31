# Command Reference

> 📋 This guide provides a comprehensive reference for all commands available in the cursor++ tool.

## Command Overview

cursor++ provides the following commands to manage your agents:

| Command | Description |
|---------|-------------|
| `init` | Initialize current directory with cursor++ agents |
| `agent` | Interactively select and use agents for cursor++ IDE |

## Global Options

These options can be used with any command:

| Option | Description |
|--------|-------------|
| `--verbose` | Show informational messages on console |
| `--debug` | Show debug messages on console (implies verbose) |
| `--multi-agent` | Enable multi-agent mode for this session |
| `--verbose-errors` | Display detailed error messages on failure |
| `--version`, `-v` | Show version information |

```bash
# Show version information
cursor++ --version

# Enable verbose output
cursor++ --verbose agent 

# Enable debug output
cursor++ --debug init

# Enable multi-agent mode
cursor++ --multi-agent init
```

## Command Details

### `init` Command

Initializes the current directory with cursor++ agents.

```bash
cursor++ init
```

**Behavior:**
- Creates the agent rules directory in the current project
- Guides you through setting up the main agent rules location if it doesn't exist
- Creates the `.cursor/rules` directory if it doesn't exist
- Updates `.gitignore` to exclude the `.cursor/` directory if needed
- Displays available agents after initialization
- Performs verification steps to ensure successful initialization
- Provides detailed feedback if issues are detected

**Verification Steps:**
- Checks that all required directories exist
- Verifies permissions on directories
- Counts files in the rules directory to confirm proper setup
- Provides diagnostic information if verification fails

**First-time Setup Options:**
1. Create an empty directory structure
2. Fetch agents from a git repository
3. Cancel operation

**Example Output:**
```
✨ cursor++ - Cursor IDE Enhancement Tool v1.0.0 ✨

⚠️ Warning: Main agents location does not exist: /Users/username/.cursor/rules

Choose an option:
1. Create empty directory structure
2. Fetch from git repository
3. Cancel operation

> 2

Enter git repository URL: [git@github.com:cursor-ai/cursor-plus-plus.git]

✓ Successfully initialized agents in /path/to/project/.cursor/rules
✓ Verification complete: 12 rule files found and configured correctly

Next Steps:
1. Use cursor++ agent to see available agents
2. Start using agents by referencing them with @agent-name.mdc in your chat
3. For more information about an agent, use cursor++ agent info <agent-id>
```

### `agent` Command

The `agent` command provides access to the Agent System, allowing you to view, select, and interact with agents.

#### Basic Usage

```bash
cursor++ agent
```

Without additional arguments, this lists all available agents.

#### Subcommands

| Subcommand | Description |
|------------|-------------|
| (no subcommand) | Display all available agents (default behavior) |
| `info <id>` | Show detailed information about a specific agent |
| `select` | Interactively select and load an agent |

#### Listing All Agents

```bash
cursor++ agent
```

**Example Output:**
```
✨ cursor++ - Cursor IDE Enhancement Tool v1.0.0 ✨

+-----+---------------------+--------------------+----------+
| No. | Agent Name          | Reference ID       | Version  |
+-----+---------------------+--------------------+----------+
| 1   | Document Syncer     | @doc-syncer.mdc    | 1.0      |
+-----+---------------------+--------------------+----------+
```

> **Note**: The `agent` command will automatically search for rules in multiple locations, checking first in the project-specific location (`.cursor/rules`), then in the user's home directory (`~/.cursor/rules`), and finally in the default system-wide location (`/usr/local/share/cursor-rules`). This ensures that agents are found regardless of where they are stored.

#### `agent info` Subcommand

Shows detailed information about a specific agent.

```bash
cursor++ agent info <id>
```

The `<id>` can be:
- A string ID (e.g., `doc-syncer`)
- A numeric index from the list (e.g., `1`)

**Example Output:**
```
✨ cursor++ - Cursor IDE Enhancement Tool v1.0.0 ✨

Agent details:
  ID:          doc-syncer
  Name:        🔄 Document Syncer Agent
  Version:     1.0

Description:
  The Document Syncer Agent specializes in maintaining synchronization 
  between documentation in a codebase and its associated documentation system. 
  It ensures documentation stays updated and consistent.

Capabilities:
  - Context Detection and Analysis
  - Codebase-to-Docs Synchronization
  - Website Support File Synchronization
  - Conflict Resolution

File: /Users/username/.cursor/rules/doc-syncer.mdc
```

#### `agent select` Subcommand

Interactively selects and loads an agent.

```bash
cursor++ agent select
```

**Behavior:**
- Displays a terminal UI for selecting an agent
- Shows agent details after selection
- Optionally displays the full agent definition

**Example Output:**
```
✨ cursor++ - Cursor IDE Enhancement Tool v1.0.0 ✨

Select an agent:
> 1. 🔄 Document Syncer Agent

[Use arrow keys to navigate, Enter to select]
```

## Exit Codes

The cursor++ tool uses the following exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Usage error |
| 10 | Init error |
| 15 | Agent error |
| 20 | Setup error |
| 25 | Config error |

## Command Workflow Examples

### Basic Project Setup

```bash
# Create a new project
mkdir my-project
cd my-project

# Initialize git repository
git init

# Initialize agents
cursor++ init

# Select an agent to work with
cursor++ agent select
```

### Viewing and Using Agents

```bash
# List available agents
cursor++ agent

# Get detailed information about an agent
cursor++ agent info doc-syncer

# Select a specific agent
cursor++ agent select
```

## Using Agents in Cursor Chat

The most common way to use agents is directly in the Cursor chat interface:

```
@doc-syncer.mdc I need help synchronizing documentation between my codebase and website
```

This approach allows you to immediately invoke the agent's capabilities without using command-line tools.

## See Also

- [Configuration](./configuration.md) for details on configuring cursor++
- [Agent System](./agents.md) for information on working with agents
- [Examples](../examples/) for usage examples and workflows

## Navigation

- Previous: [Configuration](./configuration.md)
- Next: [Agent System](./agents.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)