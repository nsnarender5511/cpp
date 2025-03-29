# Configuration Guide

> ⚙️ This guide explains how to configure crules to match your workflow and preferences.

## Configuration Overview

Crules uses a configuration system that allows you to customize its behavior. Configuration settings define:

- Where your rules are stored
- How rules are synchronized
- How the agent system behaves
- Visual and formatting preferences

## Configuration File

Crules automatically creates a configuration file at:

- **macOS/Linux**: `~/.config/crules/config.json`
- **Windows**: `%APPDATA%\crules\config.json`

This file is created when you first run crules. You can edit it directly or use the command-line interface to modify settings.

### Example Configuration File

```json
{
  "mainLocation": "/Users/username/.cursor/rules",
  "rulesDirName": ".cursor/rules",
  "registryFileName": "registry.json",
  "lastSelectedAgent": "wizard",
  "dirPermission": 493,
  "filePermission": 420,
  "theme": {
    "useColoredOutput": true,
    "accentColor": "cyan"
  }
}
```

## Core Settings

### Main Location

The `mainLocation` setting defines the central location where all your rules are stored. This is the source of truth for rule synchronization.

```json
"mainLocation": "/Users/username/.cursor/rules"
```

The main location is typically set during your first run of `crules init`. You can change it manually in the configuration file.

### Rules Directory Name

The `rulesDirName` setting defines the directory name used to store rules within each project.

```json
"rulesDirName": ".cursor/rules"
```

By default, this is set to `.cursor/rules`, which creates a hidden directory in your project root. This directory is typically added to `.gitignore`.

### Registry File Name

The `registryFileName` setting defines the name of the file used to store the project registry.

```json
"registryFileName": "registry.json"
```

The registry file keeps track of all projects that use crules.

### Last Selected Agent

The `lastSelectedAgent` setting stores the ID of the last agent you selected.

```json
"lastSelectedAgent": "wizard"
```

This is used to remember your last selection when using the agent system.

## Permissions

### Directory Permission

The `dirPermission` setting defines the permission mode used when creating directories.

```json
"dirPermission": 493
```

This value is equivalent to `0755` in octal notation (read/write/execute for owner, read/execute for group and others).

### File Permission

The `filePermission` setting defines the permission mode used when creating files.

```json
"filePermission": 420
```

This value is equivalent to `0644` in octal notation (read/write for owner, read for group and others).

## Theme Settings

### Color Output

The `useColoredOutput` setting enables or disables colored terminal output.

```json
"theme": {
  "useColoredOutput": true
}
```

Set this to `false` if you prefer plain text output without ANSI color codes.

### Accent Color

The `accentColor` setting defines the primary color used for highlighting in the terminal UI.

```json
"theme": {
  "accentColor": "cyan"
}
```

Available options include:
- `"cyan"`
- `"green"`
- `"blue"`
- `"red"`
- `"yellow"`
- `"magenta"`
- `"white"`

## Environment Variables

Crules respects the following environment variables:

| Variable | Description |
|----------|-------------|
| `APP_NAME` | Overrides the application name (`crules` by default) |
| `COLUMNS` | Defines the terminal width (useful for testing) |
| `NO_COLOR` | Disables colored output when set to any value |

Example:

```bash
# Override the application name
export APP_NAME=custom-crules

# Set terminal width
export COLUMNS=120

# Disable colored output
export NO_COLOR=1
```

## Configuration Strategies

### Per-User Configuration

The default configuration is stored per user, ensuring each user can have their own settings.

### Project-Specific Overrides

You can create a project-specific configuration by placing a `crules.config.json` file in your project root directory. This will override the global settings when working in that project.

Example `crules.config.json`:

```json
{
  "rulesDirName": "custom-rules-dir",
  "theme": {
    "accentColor": "green"
  }
}
```

Project-specific configuration only overrides the settings it specifies; all other settings are inherited from the global configuration.

## Advanced Configuration

### Registry Management

The project registry is stored in a JSON file at:

- **macOS/Linux**: `~/.local/share/crules/registry.json`
- **Windows**: `%APPDATA%\crules\registry.json`

This file contains a list of all projects that have been registered with crules. You can manually edit this file to:

- Remove specific projects
- Change project paths
- Reset the registry

Example registry.json:

```json
{
  "projects": [
    "/Users/username/projects/project1",
    "/Users/username/projects/project2",
    "/Users/username/projects/project3"
  ]
}
```

### Log Settings

Logs are stored in:

- **macOS/Linux**: `~/.local/share/crules/logs/`
- **Windows**: `%APPDATA%\crules\logs\`

You can control log verbosity using the `--verbose` and `--debug` flags.

## Configuration Examples

### Minimal Configuration

A minimal configuration focuses only on the essential settings:

```json
{
  "mainLocation": "/Users/username/.cursor/rules",
  "rulesDirName": ".cursor/rules"
}
```

### Development Configuration

A configuration optimized for development:

```json
{
  "mainLocation": "/Users/username/development/rule-repository",
  "rulesDirName": ".cursor/rules",
  "registryFileName": "dev-registry.json",
  "theme": {
    "useColoredOutput": true,
    "accentColor": "blue"
  }
}
```

### Team Configuration

A configuration optimized for team use:

```json
{
  "mainLocation": "/team/shared/rules",
  "rulesDirName": ".cursor/team-rules",
  "registryFileName": "team-registry.json",
  "theme": {
    "useColoredOutput": true,
    "accentColor": "green"
  }
}
```

## Troubleshooting Configuration

### Common Issues

- **Rules not syncing**: Check if your `mainLocation` is correct and accessible
- **Permission errors**: Verify the `dirPermission` and `filePermission` settings
- **Registry not updating**: Ensure the registry file is writable

### Resetting Configuration

To reset your configuration to default values, delete the configuration file and run any crules command. A new configuration file will be created with default values.

```bash
# On macOS/Linux
rm ~/.config/crules/config.json

# On Windows
del %APPDATA%\crules\config.json

# Run any command to regenerate the config
crules --version
```

## See Also

- [Command Reference](./commands.md) for details on all available commands
- [Examples](../examples/) for usage examples and workflows
- [Developer Guide](../developer-guide/) for information on extending crules

## Navigation

- Previous: [Getting Started](./getting-started.md)
- Next: [Commands](./commands.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
