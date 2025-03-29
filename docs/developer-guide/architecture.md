# Architecture

> 🏗️ This document describes the high-level architecture of the crules tool and how its components interact.

## Overview

The crules tool is designed with a modular architecture that separates concerns between different components. The architecture follows these key principles:

1. **Command-driven interface**: Core functionality is exposed through a hierarchical command structure
2. **Separation of concerns**: Each component has a specific responsibility
3. **Extensibility**: The system is designed to be extended with new capabilities
4. **Configuration management**: Central configuration that persists user settings

## System Components

<!-- If you have created a PNG diagram, uncomment the following line -->
<!-- ![Architecture Diagram](../assets/architecture-diagram.png) -->

<!-- If a PNG diagram is not available, the ASCII art version is displayed below -->
```
+---------------------------------------+
|            Command Line UI             |
+-----------------+---------------------+
                  |
                  v
+---------------------------------------+
|          Command Handlers             |<---------+
+-----------------+---------------------+          |
                  |                                |
                  v                                |
+---------------------------------------+          |
|           Core Services               |    +-----+-----+
+-----------------+---------------------+    |           |
                  |                          | Config    |
                  v                          |           |
+---------------------------------------+    +-----------+
|           Agent System                |
|                                       |    +-----------+
| +-------------+   +--------------+    |    |           |
| |  Registry   |<->|   Parser     |    |<-->| File      |
| +-------------+   +--------------+    |    | System    |
| +-------------+   +--------------+    |    |           |
| |  Selector   |<->|   Loader     |    |    +-----------+
| +-------------+   +--------------+    |
+---------------------------------------+
```

### Component Diagram

The crules system is composed of the following high-level components:

```
┌─────────────────────────┐
│     Command Line UI     │
└───────────┬─────────────┘
            │
┌───────────▼─────────────┐    ┌─────────────────────────┐
│     Command Handlers    │◄───┤      Configuration      │
└───────────┬─────────────┘    └─────────────────────────┘
            │
┌───────────▼─────────────┐    ┌─────────────────────────┐
│     Core Services       │◄───┤      File System        │
└───────────┬─────────────┘    └─────────────────────────┘
            │
┌───────────▼─────────────┐
│     Agent System        │
└─────────────────────────┘
```

## Key Components

### Command Line UI (CLI)

The CLI component provides the user interface for the tool. It:
- Processes command line arguments
- Manages the command hierarchy
- Handles input/output formatting
- Provides help and documentation

The CLI is implemented using the Cobra library, which provides a structured approach to defining commands, subcommands, and flags.

### Command Handlers

Each command in the system is handled by a dedicated handler that implements its specific logic. Command handlers:
- Validate input parameters
- Coordinate between services to perform operations
- Format and display results

### Core Services

The core services implement the main functionality of the tool:
- **Project Management**: Handles project registration and tracking
- **Sync Service**: Manages synchronization between different locations
- **Rule Management**: Handles parsing and processing of rules

### Agent System

The Agent System is a specialized component that:
- Discovers available agents from rule files
- Manages agent metadata and capabilities
- Provides selection and loading of agents
- Handles agent persistence

The Agent System has the following subcomponents:
- **Registry**: Manages the collection of available agents
- **Parser**: Extracts agent information from rule files
- **Selector**: Provides an interactive UI for agent selection
- **Loader**: Loads agent definitions for use

### Configuration

The Configuration component:
- Manages persistent configuration
- Handles reading and writing configuration files
- Provides access to configuration parameters throughout the system

### File System

The File System component:
- Abstracts file system operations
- Handles file reading, writing, and copying
- Manages directory creation and traversal

## Data Flow

### Agent Selection Flow

The agent selection process follows this flow:

1. User invokes the `crules agent select` command
2. The CLI routes to the agent select command handler
3. The handler initializes the Agent Selector
4. The Selector loads available agents from the Registry
5. The Registry scans the rules directory for agent definitions
6. The Parser extracts metadata from rule files
7. The Selector displays the UI for user selection
8. User selects an agent
9. The selected agent ID is saved to the configuration
10. The CLI displays confirmation of the selection

### Project Synchronization Flow

The project synchronization process follows this flow:

1. User invokes the `crules sync` command
2. The CLI routes to the sync command handler
3. The handler loads the configuration
4. The configuration provides the main location and current location
5. The sync service compares files between locations
6. The sync service copies modified files from main to current
7. The CLI displays the sync results

## Communication Between Components

Components communicate through well-defined interfaces:

- **Command to Service**: Commands invoke services through their public methods
- **Service to Service**: Services interact through interfaces to maintain loose coupling
- **Service to Configuration**: Services access configuration through a central provider
- **Service to File System**: File operations are abstracted through a file system interface

## Error Handling

The error handling strategy follows these principles:

1. **Errors are propagated up**: Lower-level components return errors to be handled at higher levels
2. **Descriptive error messages**: Errors include context to help diagnose issues
3. **Graceful degradation**: The system attempts to continue operation when possible
4. **User-friendly messaging**: Error messages are translated into user-friendly terms in the CLI

## Configuration Management

Configuration is managed at several levels:

1. **Global Configuration**: Stored in a central location
2. **Project Configuration**: Specific to each project
3. **Runtime Configuration**: Passed via command-line arguments
4. **Environment Variables**: Used for system-specific settings

## Extension Points

The architecture includes several extension points:

1. **New Commands**: The command system can be extended with new commands
2. **Agent Types**: The agent system can support new types of agents
3. **Output Formats**: Results can be formatted in different ways
4. **Rule Processors**: New rule processing capabilities can be added

## Future Directions

The architecture is designed to support future enhancements:

1. **Plugin System**: Support for external plugins
2. **Remote Agents**: Support for remotely-hosted agents
3. **Collaboration Features**: Sharing and collaboration on rules
4. **Web Interface**: A browser-based interface for managing rules

## See Also

- [Code Structure](./code-structure.md)
- [Extending Agents](./extending-agents.md)
- [Contributing Guidelines](./contributing.md)
