# Configuration Guide

> ⚙️ This guide explains how to configure the crules tool to suit your workflow and preferences.

## Configuration Overview

crules uses a configuration file to store:

- The main rules location
- Registered project locations
- User preferences

## Configuration File Location

The configuration file is stored in a platform-specific location:

- **Windows**: `%APPDATA%\crules\config.json`
- **macOS**: `~/Library/Application Support/crules/config.json`
- **Linux**: `~/.config/crules/config.json`

## Configuration Structure

The configuration file uses JSON format and has the following structure:

```json
{
  "mainLocation": "/path/to/main/rules/directory",
  "locations": {
    "/path/to/project1": {
      "path": "/path/to/project1",
      "createdAt": "2023-03-29T10:00:00Z"
    },
    "/path/to/project2": {
      "path": "/path/to/project2",
      "createdAt": "2023-03-29T11:00:00Z"
    }
  },
  "selectedAgent": "wizard"
}
```

### Key Configuration Fields

- **`mainLocation`**: The path to the main rules directory where rules are synchronized from.
- **`locations`**: A map of registered project paths to location objects that include:
  - **`path`**: The path to the project.
  - **`createdAt`**: Timestamp of when the project was registered.
- **`selectedAgent`**: The ID of the currently selected agent.

## Configuring with Environment Variables

You can override certain configuration values using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `CRULES_MAIN_LOCATION` | Main rules location | Platform-specific user config directory |
| `CRULES_RULES_DIR` | Rules directory name | `.cursor/rules` |
| `CRULES_LOG_LEVEL` | Logging level | `info` |

## Initial Configuration

The first time you run `crules init` in a project directory, crules will:

1. Create a configuration file if it doesn't exist
2. Set the main location to the default location or the one specified by environment variables
3. Register the current project in the configuration

```bash
# Initialize crules in the current directory
crules init
```

## Changing the Main Location

To change the main rules location:

```bash
# Set a different main location using an environment variable
export CRULES_MAIN_LOCATION=/path/to/new/main/location
crules init
```

## Managing Project Locations

Projects are automatically registered when you run `crules init` in their directories. You can view all registered projects with:

```bash
crules list
```

To remove non-existent projects from the registry:

```bash
crules clean
```

## Agent Configuration

### Selecting an Agent

You can select an agent with the interactive selector:

```bash
crules agent select
```

This updates the `selectedAgent` field in the configuration file.

### Viewing the Selected Agent

To see which agent is currently selected, use:

```bash
crules agent info
```

Without specifying an agent ID, this shows information about the currently selected agent.

## Advanced Configuration

### Rules Directory

By default, rules are stored in the `.cursor/rules` directory in each project. You can change this with the `CRULES_RULES_DIR` environment variable:

```bash
export CRULES_RULES_DIR=.custom-rules-dir
```

### Logging Level

You can change the logging level to control the verbosity of output:

```bash
export CRULES_LOG_LEVEL=debug
```

Available log levels (from most to least verbose):
- `trace`
- `debug`
- `info` (default)
- `warn`
- `error`
- `fatal`
- `panic`

## Configuration Backup

It's a good practice to back up your configuration file, especially if you have many registered projects. You can copy the configuration file to a safe location:

```bash
# On macOS
cp ~/Library/Application\ Support/crules/config.json ~/crules-config-backup.json

# On Linux
cp ~/.config/crules/config.json ~/crules-config-backup.json

# On Windows (PowerShell)
Copy-Item "$env:APPDATA\crules\config.json" -Destination "$env:USERPROFILE\crules-config-backup.json"
```

## Troubleshooting Configuration Issues

If you encounter configuration-related issues, you can:

1. Check the configuration file contents:
   ```bash
   cat ~/.config/crules/config.json
   ```

2. Reset the configuration by removing the file:
   ```bash
   rm ~/.config/crules/config.json
   ```
   Then reinitialize crules with `crules init`.

3. See the [Troubleshooting Guide](./troubleshooting.md) for more detailed help.

## Next Steps

Now that you understand how to configure crules, you can:

- Learn about the [Agent System](./agents.md)
- Review the [Command Reference](./commands.md)
- Try out some [Examples](../examples/agent-workflows.md)
