# Configuration Guide

> ⚙️ This guide explains how to configure cursor++ to match your workflow and preferences.

## Configuration Overview

cursor++ uses a configuration system that allows you to customize its behavior. Configuration settings define:

- Where your agent definitions are stored
- How the agent system behaves
- Visual and formatting preferences

## Configuration File

cursor++ automatically creates a configuration file at:

- **macOS/Linux**: `~/.config/cursor++/config.env`
- **Windows**: `%APPDATA%\cursor++\config.env`

This file is created when you first run cursor++. You can edit it directly or use the command-line interface to modify settings.

### Example Configuration File

```env
AGENTS_DIR_NAME=agents
REGISTRY_FILE_NAME=registry.json
RULES_DIR_NAME=.cursor/rules
SOURCE_FOLDER=default
DIR_PERMISSION=755
FILE_PERMISSION=644
MULTI_AGENT_ENABLED=true
LAST_SELECTED_AGENT=wizard
```

## Core Settings

### Rules Directory Name

The `RULES_DIR_NAME` setting defines the directory name used to store rules within each project.

```env
RULES_DIR_NAME=.cursor/rules
```

By default, this is set to `.cursor/rules`, which creates a hidden directory in your project root. This directory is typically added to `.gitignore`.

### Agents Directory Name

The `AGENTS_DIR_NAME` setting defines the directory name within the rules directory where agent definitions are stored.

```env
AGENTS_DIR_NAME=cursor-rules
```

By default, this is set to `cursor-rules`, which is a dedicated directory for Cursor-specific rules. This directory is where agent definition files (`.mdc`) are stored and should be included in version control if you want to share your agents with your team.

> **Note**: In versions prior to v1.1, this was set to "test" by default. The change to "cursor-rules" provides better clarity about its purpose and improves compatibility with standard practices.

### Source Folder

The `SOURCE_FOLDER` setting defines the subfolder within a cloned repository that contains the agent definitions to be used.

```env
SOURCE_FOLDER=default
```

By default, this is set to `default`. When initializing from a git repository that contains multiple folders, the system will clone the entire repository but only copy files from this specified subfolder to your rules directory. This allows repositories to maintain multiple sets of agent definitions while your local setup uses only the ones you need.

> **Note**: If the source folder doesn't exist in the cloned repository, the operation will fail with an error message.

### Registry File Name

The `REGISTRY_FILE_NAME` setting defines the name of the file used to store the agent registry.

```env
REGISTRY_FILE_NAME=registry.json
```

The registry file keeps track of available agents and their metadata.

### Multi-Agent Mode

The `MULTI_AGENT_ENABLED` setting controls whether multiple agents can be used simultaneously.

```env
MULTI_AGENT_ENABLED=true
```

When enabled, you can reference multiple agents in your conversations.

### Last Selected Agent

The `LAST_SELECTED_AGENT` setting stores the ID of the last agent you selected.

```env
LAST_SELECTED_AGENT=wizard
```

This is used to remember your last selection when using the agent system.

## Permissions

### Directory Permission

The `DIR_PERMISSION` setting defines the permission mode used when creating directories.

```env
DIR_PERMISSION=755
```

This value is equivalent to `0755` in octal notation (read/write/execute for owner, read/execute for group and others).

### File Permission

The `FILE_PERMISSION` setting defines the permission mode used when creating files.

```env
FILE_PERMISSION=644
```

This value is equivalent to `0644` in octal notation (read/write for owner, read for group and others).

## Environment Variables

cursor++ respects the following environment variables:

| Variable | Description |
|----------|-------------|
| `APP_NAME` | Overrides the application name (`cursor++` by default) |
| `COLUMNS` | Defines the terminal width (useful for testing) |
| `NO_COLOR` | Disables colored output when set to any value |
| `AGENTS_DIR_NAME` | Overrides the agents directory name |
| `REGISTRY_FILE_NAME` | Overrides the registry file name |
| `RULES_DIR_NAME` | Overrides the rules directory name |
| `SOURCE_FOLDER` | Specifies which subfolder to use from cloned repositories |
| `DIR_PERMISSION` | Overrides the directory permission mode (in octal) |
| `FILE_PERMISSION` | Overrides the file permission mode (in octal) |
| `MULTI_AGENT_ENABLED` | Controls multi-agent mode (`true`, `1`, or `yes` to enable) |
| `LAST_SELECTED_AGENT` | Sets the last selected agent ID |

Example:

```bash
# Override the application name
export APP_NAME=custom-cursor++

# Set terminal width
export COLUMNS=120

# Disable colored output
export NO_COLOR=1

# Enable multi-agent mode
export MULTI_AGENT_ENABLED=true
```

## Configuration Strategies

### Per-User Configuration

The default configuration is stored per user, ensuring each user can have their own settings.

### Command-Line Flags

You can override certain configuration options using command-line flags:

```bash
# Enable multi-agent mode for the current session
cursor++ --multi-agent init
```

## Advanced Configuration

### Agent Definitions

Agent definition files (`.mdc`) are stored in the agents directory within your rules directory. These files define the capabilities and behavior of each agent.

Example agent definition structure:

```
# 🧙‍♂️ Technical Wizard Agent Prompt

## 🎯 Role:
You are a wise **Technical Wizard Agent**, a versatile technical expert...

## 🛠️ Core Responsibilities:
...
```

## See Also

- [Commands](./commands.md) for details on cursor++ commands
- [Agent System](./agents.md) for information on working with agents
- [Building from Source](../developer-guide/building.md) for developer configuration

## Navigation

- Previous: [Installation](./installation.md)
- Next: [Commands](./commands.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
