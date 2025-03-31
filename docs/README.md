---
version: v1.0.0
last_updated: 2023-03-31
applies_to: cursor++ v1.0.0+
---

# cursor++ Documentation

> 📚 Welcome to the comprehensive documentation for the cursor++ tool. This documentation aims to help you understand, use, and extend the functionality of cursor++.

![cursor++ Logo](./assets/images/logo.png)

## 📋 Quick Navigation

Looking for something specific? Check out our [Documentation Map](./documentation-map.md) for a complete overview of all documentation resources.

## 📋 Table of Contents

- [User Guide](#user-guide)
- [Developer Guide](#developer-guide)
- [Examples](#examples)
- [API Reference](#api-reference)
- [Contributing](#contributing)

## 🚀 User Guide

The User Guide provides comprehensive information on how to use cursor++ effectively.

### 📥 Installation

Learn how to install cursor++ on different operating systems:

- [Installation Guide](./user-guide/installation.md)
- [System Requirements](./user-guide/installation.md#system-requirements)
- [Environment Setup](./user-guide/installation.md#environment-setup)

### 📝 Getting Started

Quickly get up and running with cursor++:

- [Quick Start](./user-guide/getting-started.md)
- [Basic Usage](./user-guide/getting-started.md#basic-usage)
- [Command Reference](./user-guide/commands.md)

### 🧠 Agent System

Understand how to work with the cursor++ Agent System:

- [Agent Overview](./user-guide/agents.md)
- [Using Agents](./user-guide/agents.md#using-agents)
- [Agent Relationships](./user-guide/agents.md#agent-ecosystem)
- [Creating Your Own Agents](./user-guide/agents.md#creating-your-own-agents)

## 🛠️ Developer Guide

The Developer Guide provides in-depth information for developers who want to extend or modify cursor++.

### 🏗️ Architecture

Understand the internal architecture of cursor++:

- [System Components](./developer-guide/architecture.md)
- [Architecture Diagrams](./developer-guide/architecture.md#architecture)
- [Data Flow](./developer-guide/architecture.md#data-flow)
- [Extension Points](./developer-guide/architecture.md#extension-points)

### 🔧 Building from Source

Learn how to build cursor++ from source:

- [Build Process](./developer-guide/building.md)
- [Dependencies](./developer-guide/building.md#dependencies)
- [Development Environment](./developer-guide/building.md#development-environment)

### 🧩 Extending cursor++

Discover how to extend and customize cursor++:

- [Extending Agents](./developer-guide/extending-agents.md)
- [API Integration](./developer-guide/extending-agents.md#api-integration)
- [Custom Commands](./developer-guide/extending-agents.md#custom-commands)

## 📊 Examples

Explore practical examples of using cursor++:

- [Basic Usage Examples](./examples/basic-usage.md)
- [Advanced Configurations](./examples/advanced-configurations.md)
- [Agent Workflows](./examples/agent-workflows.md)
- [Integration Examples](./examples/integration.md)

## 📘 API Reference

Detailed API documentation for developers:

- [Core API](./api-reference/core-api.md)
- [Agent API](./api-reference/agent-api.md)
- [UI API](./api-reference/ui-api.md)
- [Utils API](./api-reference/utils-api.md)
- [Git API](./api-reference/git-api.md)
- [Version API](./api-reference/version-api.md)

## 🤝 Contributing

Learn how you can contribute to cursor++:

- [Contribution Guidelines](./developer-guide/contributing.md)
- [Code of Conduct](./developer-guide/code-of-conduct.md)
- [Development Workflow](./developer-guide/contributing.md#development-workflow)

## 📊 Available Agents

cursor++ includes a rich ecosystem of specialized agents to assist with various aspects of the development workflow:

### Planning Agents
- **Technical Wizard** (`wizard.mdc`): Provides high-level technical guidance and coordinates other agents
- **Feature Planner** (`feature-planner.mdc`): Designs implementation plans for new features
- **Fix Planner** (`fix-planner.mdc`): Analyzes bugs and plans targeted fixes
- **Architecture Planner** (`architecture-planner.mdc`): Designs high-level system structures and component relationships
- **Scraper Planner** (`scraper-planner.mdc`): Plans implementations for data scraping tasks
- **Git Actions Planner** (`git-actions-planner.mdc`): Designs GitHub Actions workflows
- **Refactoring Guru** (`refactoring-guru.mdc`): Identifies code smells and plans refactoring strategies

### Implementation Agents
- **Implementer** (`implementer.mdc`): Converts plans into working code
- **Runner** (`runner.mdc`): Tests and verifies implementations

### Support Agents
- **Code Reviewer** (`code-reviewer.mdc`): Reviews code for quality and issues
- **Git Committer** (`git-committer.mdc`): Creates structured commit messages
- **Quick Answer** (`quick-answer-agent.mdc`): Provides concise, direct answers
- **Document Syncer** (`document-syncer.mdc`): Synchronizes documentation with code
- **Documentation Agent** (`documentation-agent.mdc`): Creates and improves documentation
- **Document Reviewer** (`document-reviewer-agent.mdc`): Reviews documentation quality
- **Agent Selector** (`agent-selector.mdc`): Selects appropriate agents for specific tasks

## 📈 Version History

- Current Version: 1.0.0
- [Changelog](../CHANGELOG.md)

---

© 2023 cursor++ Contributors | [License](../LICENSE)
