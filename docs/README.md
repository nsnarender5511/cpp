# crules Documentation

> 📚 Welcome to the crules documentation! This guide will help you understand how to use and extend the crules tool for managing and synchronizing Cursor rules across multiple projects.

## Overview

Crules is a command-line tool that simplifies the management of Cursor rules across multiple projects. It provides a centralized location for your rules and ensures they stay in sync across all your projects.

## Documentation Sections

- [**User Guide**](./user-guide/): Instructions for end users on how to install, configure, and use crules
  - [Installation](./user-guide/installation.md): How to install crules on different platforms
  - [Configuration](./user-guide/configuration.md): Configuring crules for your environment
  - [Commands](./user-guide/commands.md): Detailed documentation of all available commands
  - [Troubleshooting](./user-guide/troubleshooting.md): Common issues and their solutions

- [**Developer Guide**](./developer-guide/): Documentation for developers who want to extend or modify crules
  - [Architecture](./developer-guide/architecture.md): Overview of the system architecture
  - [Code Structure](./developer-guide/code-structure.md): Understanding the codebase organization
  - [Contributing](./developer-guide/contributing.md): Guidelines for contributors
  - [Testing](./developer-guide/testing.md): How to test your changes

- [**API Reference**](./api/): Technical reference for the internal APIs

- [**Examples**](./examples/): Practical examples and use cases

## Agent System

One of the key features of crules is its Agent System, which provides an interactive way to work with specialized AI agents. These agents are defined in `.mdc` files and can be used for various tasks like planning features, fixing issues, implementing code, and more.

For detailed information about the Agent System, see:
- [Agent System Overview](./user-guide/agents.md)
- [Using the Agent Commands](./user-guide/commands.md#agent-commands)
- [Extending the Agent System](./developer-guide/extending-agents.md)
