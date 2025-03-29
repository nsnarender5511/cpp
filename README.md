# crules

> 🧩 A powerful tool for managing AI agent rules across multiple projects

[![Go Report Card](https://goreportcard.com/badge/github.com/nsnarender5511/crules)](https://goreportcard.com/report/github.com/nsnarender5511/crules)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

crules (Cursor Rules) is a command-line tool that helps you manage AI agent rules across multiple projects. It provides commands for initializing, syncing, and merging rule files, as well as an interactive agent selection system.

**Features Overview**: Command-line tool for managing AI agent rules across projects

## 🌟 Features

- **Rule Synchronization**: Keep rules in sync across multiple projects
- **Agent System**: Work with specialized AI agents for different tasks
- **Interactive Selection**: Choose agents through an intuitive terminal UI
- **Project Management**: Register and track projects with rule directories

**Try it out!** The tool provides an intuitive interface for working with AI agent rules.

## 📦 Installation

### Using Homebrew (macOS and Linux)

```bash
# Add the tap
brew tap nsnarender5511/tap

# Install crules
brew install crules
```

**Note**: Follow these commands to install via Homebrew.

### Manual Installation

1. Download the appropriate binary for your operating system from the [Releases page](https://github.com/nsnarender5511/crules/releases).
2. Extract the archive (if applicable).
3. Move the `crules` binary to a location in your PATH.

**Note**: These steps will install crules manually on your system.

## Homebrew Tap Setup (for maintainers)

To set up the Homebrew tap repository:

1. Create a new GitHub repository named `homebrew-tap` under your GitHub account.
2. Initialize it with a README.md file.
3. Once the repository is created, you can release crules using GoReleaser:

```bash
# Tag the release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# Run GoReleaser (requires a GITHUB_TOKEN)
export GITHUB_TOKEN=your_github_token
goreleaser release --clean
```

This will automatically build the binaries, create the release on GitHub, and update the Homebrew tap.

## 🚀 Quick Start

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

**Quick Start Guide**: These commands will help you get started with crules quickly.

## 🤖 Agent System

The Agent System allows you to work with specialized AI agents for different tasks in software development:

| Agent | Description |
|-------|-------------|
| **Technical Wizard** | Provides high-level technical guidance and coordinates other agents |
| **Feature Planner** | Breaks down requirements into implementation tasks |
| **Fix Planner** | Analyzes issues and develops fix strategies |
| **Implementer** | Translates plans into code |
| **Runner** | Executes and tests code |
| **Documentation** | Creates and maintains documentation |
| **Code Reviewer** | Performs thorough code reviews |
| **Git Committer** | Creates conventional format commit messages |

**Note**: Use the agent selection interface to choose the right agent for your task.

![Agent Selection](./docs/assets/gifs/usage/agent-selection.gif)

Learn more about the Agent System in the [documentation](./docs/user-guide/agents.md).

## 📋 Commands

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

**Note**: See the command documentation for detailed usage information.

![Command Usage](./docs/assets/gifs/usage/command-usage.gif)

For detailed information about commands, see the [Command Reference](./docs/user-guide/commands.md).

## ⚙️ Configuration

crules stores its configuration in `~/.config/crules/config.json`. This file contains:

- The main rules location
- Registered project locations

You can view your current configuration with:

```bash
crules config show
```

**Note**: Use configuration commands to manage your crules setup.

![Configuration](./docs/assets/gifs/usage/configuration.gif)

## 📋 Documentation

Our comprehensive documentation is available in the [docs](./docs) directory:

- **[Documentation Map](./docs/documentation-map.md)**: Complete overview of all documentation
- **[User Guide](./docs/user-guide/)**: Instructions for using crules
  - [Installation](./docs/user-guide/installation.md)
  - [Configuration](./docs/user-guide/configuration.md)
  - [Commands](./docs/user-guide/commands.md)
  - [Agent System](./docs/user-guide/agents.md)
  - [Troubleshooting](./docs/user-guide/troubleshooting.md)

- **[Developer Guide](./docs/developer-guide/)**: Information for developers
  - [Architecture](./docs/developer-guide/architecture.md)
  - [Code Structure](./docs/developer-guide/code-structure.md)
  - [Contributing](./docs/developer-guide/contributing.md)
  - [Testing](./docs/developer-guide/testing.md)
  - [Extending Agents](./docs/developer-guide/extending-agents.md)

- **[Examples](./docs/examples/)**: Usage examples and workflows
  - [Basic Usage](./docs/examples/basic-usage.md)
  - [Advanced Usage](./docs/examples/advanced-usage.md)
  - [Agent Workflows](./docs/examples/agent-workflows.md)

- **[API Reference](./docs/api-reference/)**: Technical reference for the internal APIs
  - [Core API](./docs/api-reference/core-api.md)
  - [Agent API](./docs/api-reference/agent-api.md)
  - [UI API](./docs/api-reference/ui-api.md)
  - [Utils API](./docs/api-reference/utils-api.md)
  - [Git API](./docs/api-reference/git-api.md)
  - [Version API](./docs/api-reference/version-api.md)

The documentation includes detailed explanations, examples, screenshots, and animated GIFs to help you understand how to use crules effectively.

## 👥 Contributing

Contributions are welcome! Please see our [Contributing Guidelines](./docs/developer-guide/contributing.md) for more information on how to get started.

## 📋 Release Process

There are two ways to create a new release:

### 1. Local Release (using GoReleaser locally)

```bash
# Create a .env file with your GitHub token
make release TAG=v0.0.1
```

### 2. GitHub Actions Release (recommended)

```bash
# Trigger the GitHub Actions release workflow
make release-github TAG=v0.0.1
```

For detailed release instructions, see [RELEASING.md](RELEASING.md).

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 