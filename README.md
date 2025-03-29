# crules

> 🧩 A powerful tool for managing AI agent rules across multiple projects

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/crules)](https://goreportcard.com/report/github.com/yourusername/crules)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

crules (Cursor Rules) is a command-line tool that helps you manage AI agent rules across multiple projects. It provides commands for initializing, syncing, and merging rule files, as well as an interactive agent selection system.

## Features

- **Rule Synchronization**: Keep rules in sync across multiple projects
- **Agent System**: Work with specialized AI agents for different tasks
- **Interactive Selection**: Choose agents through an intuitive terminal UI
- **Project Management**: Register and track projects with rule directories

## Installation

### Using Go

```bash
go install github.com/yourusername/crules/cmd/crules@latest
```

### From Binary Releases

Download the appropriate binary for your platform from the [Releases](https://github.com/yourusername/crules/releases) page.

## Quick Start

```bash
# Initialize rules in the current directory
crules init

# List available agents
crules agent list

# Select an agent interactively
crules agent select

# Get detailed information about a specific agent
crules agent info wizard

# Synchronize rules from the main location
crules sync

# Merge rules from the current directory to the main location
crules merge

# List all registered projects
crules list
```

## Agent System

The Agent System allows you to work with specialized AI agents for different tasks in software development:

- **Technical Wizard**: Provides high-level technical guidance
- **Feature Planner**: Breaks down requirements into implementation tasks
- **Fix Planner**: Analyzes issues and develops fix strategies
- **Implementer**: Translates plans into code
- **Runner**: Executes and tests code
- **Documentation**: Creates and maintains documentation

Learn more about the Agent System in the [documentation](./docs/user-guide/agents.md).

## Commands

crules provides several commands to manage your rules:

| Command | Description |
|---------|-------------|
| `init` | Initializes the current directory with rules from the main location |
| `merge` | Merges current rules to the main location and syncs them to all locations |
| `sync` | Forces synchronization from the main location to the current directory |
| `list` | Displays all registered projects |
| `clean` | Removes non-existent projects from the registry |
| `agent list` | Lists all available agents |
| `agent info <id>` | Shows detailed information about a specific agent |
| `agent select` | Interactively selects and loads an agent |

For detailed information about commands, see the [Command Reference](./docs/user-guide/commands.md).

## Configuration

crules stores its configuration in `~/.config/crules/config.json`. This file contains:

- The main rules location
- Registered project locations

## Documentation

For comprehensive documentation, visit the [docs](./docs) directory:

- [User Guide](./docs/user-guide): Instructions for using crules
- [Developer Guide](./docs/developer-guide): Information for developers
- [Examples](./docs/examples): Usage examples and workflows
- [API Reference](./docs/api): API documentation

## Contributing

Contributions are welcome! Please see our [Contributing Guidelines](./docs/developer-guide/contributing.md) for more information.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 