# Commands Reference

> 🔍 This document provides detailed information about all available commands in crules.

## Basic Usage

```bash
crules [OPTIONS] <command>
```

## Global Options

| Option | Description |
|--------|-------------|
| `--verbose` | Show informational messages on console |
| `--debug` | Show debug messages on console (implies verbose) |
| `--version`, `-v` | Show version information |

## Core Commands

### `init`

Initializes the current directory with rules from the main location.

```bash
crules init
```

This command will:
1. Copy rules from the main location to the current directory
2. Create a `.cursor/rules` directory if it doesn't exist
3. Register the current project in the registry

If the rules directory already exists, you'll be prompted to confirm overwriting it.

### `merge`

Merges current rules to the main location and syncs them to all locations.

```bash
crules merge
```

This command will:
1. Copy rules from the current directory to the main location
2. Sync the updated rules to all registered projects

The current directory must have a `.cursor/rules` directory with rules.

### `sync`

Forces synchronization from the main location to the current directory.

```bash
crules sync
```

This command will:
1. Copy rules from the main location to the current directory
2. Overwrite any existing rules in the current directory

### `list`

Displays all registered projects.

```bash
crules list
```

This command will show:
1. A list of all registered project paths
2. An indication of which projects exist and which are missing
3. A count of valid and invalid projects

### `clean`

Removes non-existent projects from the registry.

```bash
crules clean
```

This command will:
1. Check all registered projects to verify they exist
2. Remove any projects that no longer exist
3. Report how many projects were removed

## Agent Commands

The agent commands provide interactive access to the Agent System, allowing you to discover and use specialized AI agents for different tasks.

The `agent` command displays information about available agents and allows you to interact with them.

```
crules agent [subcommand]
```

#### Subcommands:

- **`agent`** - Lists all available agents in a clean tabular format. The display adapts based on your terminal width to show an optimal amount of information.
- **`agent select`** - Interactively select an agent from the list.
- **`agent info <agent-id>`** - Show detailed information about a specific agent.

#### Examples:

**List all available agents:**
```
crules agent
```

This will display a formatted table of all available agents with their reference IDs. The table will adapt to your terminal width:

- In narrow terminals: A simplified table with just agent numbers and IDs
- In medium terminals: A table with agent numbers, names, and reference syntax
- In wide terminals: A comprehensive table including version information

**Get detailed information about a specific agent:**
```
crules agent info wizard
```

**Interactively select an agent:**
```
crules agent select
```

## Using Agents in Chat

You can directly reference agents in the chatbox using the `@` symbol followed by the agent ID. For example:

- `@wizard.mdc` - Invokes the Technical Wizard agent
- `@refactoring-guru.mdc` - Invokes the Refactoring Guru agent
- `@quick-answer-agent.mdc` - Invokes the Quick Answer Agent

This approach allows you to quickly switch between specialized agents without running additional commands.

## Command Output

All commands provide:
- Success or error messages
- Details about the operation performed
- Warnings about potential issues

Adding the `--verbose` flag will show additional information about the operation, while `--debug` shows even more detailed diagnostic information.

## Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Usage error |
| 10 | Init error |
| 11 | Merge error |
| 12 | Sync error |
| 13 | List error |
| 14 | Clean error |
| 15 | Agent error |
| 20 | Setup error |

## See Also

- [Agent System Overview](./agents.md) - Detailed information about the Agent System
- [Configuration](./configuration.md) - How to configure crules
- [Troubleshooting](./troubleshooting.md) - Common issues and solutions
