# Command Reference

> 📋 This guide provides a comprehensive reference for all commands available in the crules tool.

## Command Overview

Crules provides several commands to manage your rules and interact with agents:

| Command | Description |
|---------|-------------|
| `init` | Initialize current directory with rules from main location |
| `merge` | Merge current rules to main location and sync to all locations |
| `sync` | Force sync from main location to current directory |
| `list` | Display all registered projects |
| `clean` | Remove non-existent projects from registry |
| `agent` | Interactively select and use agents |
| `import` | Import agent rules from a URL |

## Global Options

These options can be used with any command:

| Option | Description |
|--------|-------------|
| `--verbose` | Show informational messages on console |
| `--debug` | Show debug messages on console (implies verbose) |
| `--version`, `-v` | Show version information |

```bash
# Show version information
crules --version

# Enable verbose output
crules --verbose list

# Enable debug output
crules --debug sync
```

## Command Details

### `init` Command

Initializes the current directory with rules from the main location.

```bash
crules init
```

**Behavior:**
- Copies rules from main location to current directory
- Registers the current project in the crules registry
- Creates the `.cursor/rules` directory if it doesn't exist
- Updates `.gitignore` to exclude the `.cursor/` directory
- If main location doesn't exist, guides you through setup options

**First-time Setup Options:**
1. Create an empty directory structure
2. Fetch rules from a git repository
3. Cancel operation

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

⚠️ Warning: Main rules location does not exist: /Users/username/.cursor/rules

Choose an option:
1. Create empty directory structure
2. Fetch from git repository
3. Cancel operation

> 2

Enter git repository URL: [git@github.com:nsnarender5511/AgenticSystem.git]

✓ Successfully initialized rules in /path/to/project/.cursor/rules
```

### `merge` Command

Merges rules from the current directory to the main location and syncs to all registered projects.

```bash
crules merge
```

**Behavior:**
- Copies rules from current directory to main location
- Synchronizes the changes to all registered projects
- Ensures all projects have the latest versions of your rules

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Merging rules to main location...
Syncing changes to 3 registered projects...

✓ Successfully merged rules to main location
```

### `sync` Command

Forces a sync from the main location to the current directory.

```bash
crules sync
```

**Behavior:**
- Copies rules from main location to current directory
- Overwrites any local changes

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Syncing rules from main location...

✓ Successfully synced rules from main location
```

### `list` Command

Displays all registered projects.

```bash
crules list
```

**Behavior:**
- Lists all projects registered in the crules registry
- Indicates if a project no longer exists

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Registered projects (3):
  1. /Users/username/projects/project1
  2. /Users/username/projects/project2
  3. /Users/username/old-project (not found)

⚠️ 1 project(s) could not be found. Run 'crules clean' to remove them.
```

### `clean` Command

Removes non-existent projects from the registry.

```bash
crules clean
```

**Behavior:**
- Checks all registered projects for existence
- Removes non-existent projects from the registry
- Updates the registry file

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

✓ Successfully removed 1 non-existent project(s) from registry.
```

### `agent` Command

The `agent` command provides access to the Agent System, allowing you to view, select, and interact with agents.

#### Basic Usage

```bash
crules agent
```

Without additional arguments, this lists all available agents.

#### Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | Display all available agents (default behavior) |
| `info <id>` | Show detailed information about a specific agent |
| `select` | Interactively select and load an agent |

#### `agent list` Subcommand

Lists all available agents.

```bash
crules agent list
```

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

+-----+---------------------+--------------------+----------+
| No. | Agent Name          | Reference ID       | Version  |
+-----+---------------------+--------------------+----------+
| 1   | Feature Planner     | @feature-planner.mdc | 1.0    |
| 2   | Fix Planner         | @fix-planner.mdc     | 1.0    |
| 3   | Runner              | @runner.mdc          | 1.0    |
| 4   | Technical Wizard    | @wizard.mdc          | 1.0    |
+-----+---------------------+--------------------+----------+
```

#### `agent info` Subcommand

Shows detailed information about a specific agent.

```bash
crules agent info <id>
```

The `<id>` can be:
- A string ID (e.g., `wizard`)
- A numeric index from the list (e.g., `1`)

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Agent details:
  ID:          wizard
  Name:        🧙‍♂️ Technical Wizard Agent
  Version:     1.0

Description:
  The Technical Wizard Agent provides high-level technical guidance 
  and coordinates other agents. It helps with architecture decisions, 
  design patterns, and clean code principles.

Capabilities:
  - In-Depth Technical Exploration and Analysis
  - Expert Architectural Guidance
  - Design Patterns Discussion
  - Clean Code Advisory

File: /Users/username/.cursor/rules/wizard.mdc
```

#### `agent select` Subcommand

Interactively selects and loads an agent.

```bash
crules agent select
```

**Behavior:**
- Displays a terminal UI for selecting an agent
- Shows agent details after selection
- Optionally displays the full agent definition

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Select an agent:
> 1. 🧙‍♂️ Technical Wizard Agent
  2. ✨ Feature Planner Agent
  3. 🔍 Fix Planner Agent
  4. 🛠️ Implementer Agent

[Use arrow keys to navigate, Enter to select]
```

### `import` Command

Imports agent rules from a URL.

```bash
crules import <url>
```

**Parameters:**
- `<url>`: The URL to import rules from (required)

**Behavior:**
- Fetches content from the specified URL
- Parses the content to extract rules
- Stores the rules in your main location
- Optionally stores them in your current project if it has a rules directory

**Example Output:**
```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Fetching rules from https://cursor.directory/example-agent...
Parsing rules...

Found 1 rule to import:
  1. Example Agent - An example agent for crules

Do you want to import this rule? [y/N] y

Storing rules to main location...

✓ Successfully imported 1 rule
```

## Exit Codes

The crules tool uses the following exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Usage error |
| 10 | Init error |
| 11 | Merge error |
| 12 | Sync error |
| 13 | List error |
| 14 | Clean error |
| 15 | Agent error |
| 16 | Import error |
| 20 | Setup error |

## Command Workflow Examples

### Basic Project Setup

```bash
# Create a new project
mkdir my-project
cd my-project

# Initialize git repository
git init

# Initialize rules
crules init

# Select an agent to work with
crules agent select
```

### Sharing Rules with Team

```bash
# Make changes to rules
# ...

# Merge changes to main location and sync to all projects
crules merge

# In another project, get the latest rules
cd ../another-project
crules sync
```

### Cleaning Up Projects

```bash
# List all registered projects
crules list

# Remove non-existent projects
crules clean

# Verify cleanup
crules list
```

## See Also

- [Configuration](./configuration.md) for details on configuring crules
- [Agent System](./agents.md) for information on working with agents
- [Examples](../examples/) for usage examples and workflows

## Navigation

- Previous: [Configuration](./configuration.md)
- Next: [Agent System](./agents.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
