# Getting Started with crules

> 🚀 This guide will help you get up and running with crules quickly, showing you how to set up your environment, initialize projects, and work with agents.

## Overview

Crules is a command-line tool that simplifies the management of Cursor rules across multiple projects. It solves several key challenges:

- **Centralized Rule Management**: Maintain all your AI agent rules in one main location
- **Automatic Synchronization**: Ensure all projects have the latest versions of your rules
- **Consistent AI Experience**: Create a uniform AI agent experience across all your projects
- **Enhanced Collaboration**: Share rules with team members easily

## Installation

If you haven't installed crules yet, follow the [Installation Guide](./installation.md) for step-by-step instructions.

Quick installation options:

```bash
# Using Homebrew (macOS)
brew install crules

# Manual installation
curl -sL https://github.com/org/crules/releases/latest/download/install.sh | bash
```

Verify your installation:

```bash
crules --version
```

## Basic Usage

### Initial Setup

When you first use crules, you'll need to set up your main rules location. This is where all your rules will be centrally managed.

```bash
# Initialize rules in the current project
crules init
```

The first time you run this command, you'll be guided through setting up your main rules location. Options include:

1. Creating an empty directory structure
2. Fetching rules from a git repository
3. Importing rules from a URL

```mermaid
sequenceDiagram
    participant User as 👤 User
    participant CLI as 🖥️ CLI
    participant System as 🔧 Core System
    participant Files as 📁 Files
    
    User->>CLI: crules init
    CLI->>System: Check main location
    Note over System: First-time setup
    System->>User: Display options
    User->>System: Select option
    
    alt Create Empty Directory
        System->>Files: Create directory structure
    else Fetch from Git Repository
        User->>System: Enter git URL
        System->>Files: Clone repository
    else Import from URL
        User->>System: Enter URL
        System->>Files: Import rules
    end
    
    System->>Files: Create project files
    System->>User: Setup complete
    
    style User fill:#ff9966,stroke:#ff6600
    style CLI fill:#66b3ff,stroke:#0066cc
    style System fill:#9966ff,stroke:#6600cc
    style Files fill:#66cc99,stroke:#009966
```

<details>
  <summary>📺 View First-Time Setup Process</summary>
  <img src="../assets/gifs/installation/first-time-setup.gif" alt="First-time setup process" width="600">
</details>

### Project Management

Once your main rules location is set up, you can initialize rules in any project:

```bash
cd /path/to/your/project
crules init
```

This will:
1. Copy rules from your main location to the current project
2. Register the project in the crules registry
3. Update your `.gitignore` to exclude the `.cursor/` directory

View all registered projects:

```bash
crules list
```

Clean up non-existent projects:

```bash
crules clean
```

### Rule Synchronization

To ensure your project has the latest rules from your main location:

```bash
crules sync
```

When you've made changes to rules in your project that you want to share:

```bash
crules merge
```

This will:
1. Copy your local rules to the main location
2. Synchronize the changes to all registered projects

<details>
  <summary>📺 View Synchronization Process</summary>
  <img src="../assets/gifs/usage/sync-process.gif" alt="Synchronization process" width="600">
</details>

## Working with Agents

### Listing Available Agents

To see all available agents:

```bash
crules agent
```

This will display a formatted table of all agents with their information.

### Agent Selection

To interactively select and load an agent:

```bash
crules agent select
```

This will present a menu of available agents. Use the arrow keys to navigate, then press Enter to select.

<details>
  <summary>📺 View Agent Selection Process</summary>
  <img src="../assets/gifs/usage/agent-selection-process.gif" alt="Agent selection process" width="600">
</details>

### Getting Agent Information

To view detailed information about a specific agent:

```bash
crules agent info <id>
```

Replace `<id>` with either:
- The agent's string ID (e.g., `wizard`)
- The agent's numeric index from the list (e.g., `1`)

## Importing Agents

You can import agent rules from a URL:

```bash
crules import https://cursor.directory/example-agent
```

This will:
1. Fetch the content from the URL
2. Parse the content to extract rules
3. Store the rules in your main location
4. Optionally store them in your current project if it has a rules directory

## Common Workflows

### Setting Up a New Project

```bash
# Create a new project directory
mkdir my-new-project
cd my-new-project

# Initialize git repository
git init

# Initialize rules
crules init

# Start using agents in your project
crules agent select
```

### Creating Custom Agents

1. Create a new `.mdc` file in your main rules location
2. Define your agent with the appropriate metadata and instructions
3. Use `crules merge` to share it with all your projects

## Troubleshooting

### Common Issues

- **Rules not syncing**: Ensure your main location is properly set up and accessible
- **Agent not found**: Check that the agent ID is correct and the file exists in your rules directory
- **Permission errors**: Verify you have write access to the relevant directories

If you encounter persistent issues, check the [Troubleshooting Guide](./troubleshooting.md) for more detailed solutions.

## Next Steps

Now that you've got crules up and running, you might want to explore:

- [Configuration](./configuration.md) to customize your crules setup
- [Commands](./commands.md) to learn all available commands
- [Agent System](./agents.md) to understand how to work with AI agents

## Navigation

- Previous: [Installation](./installation.md)
- Next: [Configuration](./configuration.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md) 