# cursor++

[![Go](https://github.com/cursor-ai/cursor-plus-plus/workflows/Go/badge.svg)](https://github.com/cursor-ai/cursor-plus-plus/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-1.0.0-blue)]()
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

## 📚 Overview

cursor++ is a command-line tool that enhances the Cursor IDE experience by providing a synchronized multi-agent system. It enables multiple specialized AI agents to work together seamlessly, providing:

- 🤖 **Specialized Agents**: Each agent has unique capabilities and expertise
- 🔄 **Multi-Agent Collaboration**: Agents can share context and work together
- 📂 **Project Synchronization**: Keep agents in sync across your workspace
- 📝 **Integrated Workflow**: Seamless integration with your development process

---

## 🚀 Getting Started

### Installation Options

#### Download Release

1. Visit the [Releases](https://github.com/cursor-ai/cursor-plus-plus/releases) page
2. Download the appropriate version for your operating system
3. Extract the archive and follow the installation instructions

#### Manual Installation

1. Clone this repository
2. Build the binary using `go build`
3. Move the `cursor++` binary to a location in your PATH.

**Note**: These steps will install cursor++ manually on your system.

### Quick Start Commands

#### Initialize in Current Directory
```
cursor++ init
```

#### View Available Agents
```
cursor++ agent
```

#### Select an Agent Interactively
```
cursor++ agent select
```

#### View Detailed Agent Information
```
cursor++ agent info wizard
```

**Quick Start Guide**: These commands will help you get started with cursor++ quickly.

## 🤖 Available Agents

cursor++ includes a rich ecosystem of specialized agents:

| Agent | Icon | Purpose |
|-------|------|---------|
| Technical Wizard | 🧙‍♂️ | High-level technical guidance and coordination |
| Feature Planner | ✨ | Planning feature implementations |
| Fix Planner | 🔍 | Analyzing bugs and planning fixes |
| Architecture Planner | 🏗️ | Designing system architecture | 
| Implementer | 🛠️ | Converting plans into working code |
| Runner | 🏃 | Testing and verifying implementations |
| Code Reviewer | 🔍 | Reviewing code for quality and issues |
| Refactoring Guru | 🔧 | Planning and guiding code refactoring |
| Git Committer | 🔄 | Creating structured commit messages |
| Quick Answer | ⚡ | Providing concise, direct answers |
| Document Syncer | 🔄 | Synchronizing documentation with code |
| Documentation Agent | 📚 | Creating and improving documentation |
| Document Reviewer | 📝 | Reviewing documentation quality |
| Scraper Planner | 🕸️ | Planning data scraping implementations |
| Git Actions Planner | 🚀 | Designing GitHub Actions workflows |
| Agent Selector | 🎯 | Selecting appropriate agents for tasks |

---

## 🛠️ Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `VERBOSE` | Enable verbose output | `false` |
| `DEBUG` | Enable debug logging | `false` |
| `CONFIG_PATH` | Custom config file path | OS-specific |
| `LOG_LEVEL` | Logging level (debug,info,warn,error) | `info` |

### Command Line Arguments

| Argument | Description |
|----------|-------------|
| `--verbose` | Enable verbose output |
| `--debug` | Enable debug logging |
| `--version` | Display version information |
| `--help` | Show help message |

### Commands

| Command | Description |
|---------|-------------|
| `init` | Initializes the current directory with cursor++ agents |
| `agent` | Manages agents (list, select, info) |

### Configuration File

cursor++ stores its configuration in `~/.config/cursor++/config.json`. This file contains:

- Agent preferences
- System settings
- User customizations

---

## 📖 Documentation

- **[Installation](./docs/user-guide/installation.md)**: Detailed installation instructions
- **[User Guide](./docs/user-guide/)**: Instructions for using cursor++
- **[Agent Reference](./docs/user-guide/agents.md)**: Detailed information about available agents
- **[API Documentation](./docs/api-reference/)**: Reference for programmatic integration
- **[FAQ](./docs/user-guide/troubleshooting.md)**: Frequently Asked Questions

---

## 🤝 Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) before submitting a PR.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 