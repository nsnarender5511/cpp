# crules

[![Go](https://github.com/cursor-ai/crules/workflows/Go/badge.svg)](https://github.com/cursor-ai/crules/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-0.1.0-blue)]()
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

## 📚 Overview

crules is a command-line tool that enhances the Cursor IDE experience by providing a synchronized multi-agent system. It enables multiple specialized AI agents to work together seamlessly, providing:

- 🤖 **Specialized Agents**: Each agent has unique capabilities and expertise
- 🔄 **Multi-Agent Collaboration**: Agents can share context and work together
- 📂 **Project Synchronization**: Keep agents in sync across your workspace
- 📝 **Integrated Workflow**: Seamless integration with your development process

---

## 🚀 Getting Started

### Installation Options

#### Download Release

1. Visit the [Releases](https://github.com/cursor-ai/crules/releases) page
2. Download the appropriate version for your operating system
3. Extract the archive and follow the installation instructions

#### Manual Installation

1. Clone this repository
2. Build the binary using `go build`
3. Move the `crules` binary to a location in your PATH.

**Note**: These steps will install crules manually on your system.

### Quick Start Commands

#### Initialize in Current Directory
```
crules init
```

#### View Available Agents
```
crules agent list
```

#### Select an Agent Interactively
```
crules agent select
```

#### View Detailed Agent Information
```
crules agent info wizard
```

**Quick Start Guide**: These commands will help you get started with crules quickly.

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
| `init` | Initializes the current directory with crules agents |
| `agent` | Manages agents (list, select, info) |

### Configuration File

crules stores its configuration in `~/.config/crules/config.json`. This file contains:

- Agent preferences
- System settings
- User customizations

---

## 📖 Documentation

- **[Installation](./docs/installation/)**: Detailed installation instructions
- **[User Guide](./docs/user-guide/)**: Instructions for using crules
- **[Agent Reference](./docs/agents/)**: Detailed information about available agents
- **[API Documentation](./docs/api/)**: Reference for programmatic integration
- **[FAQ](./docs/faq/)**: Frequently Asked Questions

---

## 🤝 Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) before submitting a PR.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 