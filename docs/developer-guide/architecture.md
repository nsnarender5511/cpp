# Architecture

> 🏗️ **System architecture and design of crules**

## Overview

crules is designed as a command-line tool for managing and synchronizing Cursor rules across multiple projects. The architecture follows a modular approach with clear separation of concerns between the different components.

## System Components

![Architecture Diagram](../assets/architecture-diagram.png)

### High-Level Components

1. **Command Line Interface (CLI)**: Entry point for all user interactions
2. **Core Components**: Core business logic for rule synchronization and registry management
3. **Utility Components**: Helper functions for file operations, logging, etc.
4. **User Interface**: Components for displaying information to the user

## Component Details

### CLI Component

The CLI component is responsible for:
- Parsing command-line arguments and flags
- Validating input
- Routing to the appropriate command handler
- Handling exit codes

**Key Files**: 
- `cmd/main.go`: Entry point of the application

### Core Components

#### Sync Manager

The Sync Manager handles all synchronization operations between the main rule repository and project-specific rules.

**Responsibilities**:
- Initializing rule directories
- Syncing rules between locations
- Merging changes from projects to the main location
- Managing project registration

**Key Files**:
- `internal/core/sync.go`: Implementation of the sync manager

#### Registry

The Registry manages the list of registered projects that have Cursor rules.

**Responsibilities**:
- Maintaining a list of registered projects
- Adding new projects
- Removing non-existent projects
- Persistence of registry data

**Key Files**:
- `internal/core/registry.go`: Implementation of the registry

### Utility Components

#### File Utilities

Provides file system operations abstracted from the core logic.

**Responsibilities**:
- File and directory operations
- Path management
- Permission handling

#### Logging

Handles structured logging for the application.

**Responsibilities**:
- Log formatting and output
- Log level management
- Context enrichment for logs

**Key Files**:
- `internal/utils/logger.go`: Logging implementation

### User Interface Components

Handles user interaction and display of information.

**Responsibilities**:
- Formatted output to the console
- Status messages and errors
- Progress indicators

**Key Files**:
- `internal/ui/console.go`: Console UI implementation

## Data Flow

1. **Command Execution**:
   ```
   User Input -> CLI -> Command Handler -> Core Components -> Result Display
   ```

2. **Rule Synchronization**:
   ```
   Sync Command -> Sync Manager -> File Operations -> Project Updates -> Status Output
   ```

3. **Project Registration**:
   ```
   Init Command -> Sync Manager -> Registry Operations -> File System Changes -> Status Output
   ```

## Design Principles

1. **Separation of Concerns**: Each component has a single responsibility
2. **Platform Independence**: Works across Windows, macOS, and Linux
3. **Configuration Flexibility**: Configurable through environment variables
4. **Robust Error Handling**: Comprehensive error handling and reporting
5. **User-Friendly Output**: Clear, informative console output

## File Organization

```
crules/
├── cmd/                 # Command-line application entry point
│   └── main.go          # Main application file
├── internal/            # Internal packages
│   ├── core/            # Core business logic
│   │   ├── registry.go  # Registry management
│   │   └── sync.go      # Synchronization logic
│   ├── ui/              # User interface components
│   │   └── console.go   # Console UI
│   └── utils/           # Utility functions
│       ├── files.go     # File operations
│       └── logger.go    # Logging utilities
├── go.mod               # Go module file
└── go.sum               # Go module checksum file
```

## Related Documentation

- [Code Structure](code-structure.md) - Detailed code organization
- [Contributing Guide](contributing.md) - How to contribute to the codebase
