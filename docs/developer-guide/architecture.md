---
version: v0.1.0
last_updated: 2023-03-29
applies_to: crules v0.1.0+
---

# Architecture

> 🏗️ This document describes the high-level architecture of the crules tool and how its components interact.

## Overview

The crules tool is designed with a modular architecture that separates concerns between different components. The architecture follows these key principles:

1. **Command-driven interface**: Core functionality is exposed through a hierarchical command structure
2. **Separation of concerns**: Each component has a specific responsibility
3. **Extensibility**: The system is designed to be extended with new capabilities
4. **Configuration management**: Central configuration that persists user settings

## System Architecture

![Architecture Overview](../assets/images/diagrams/architecture-overview.png)
*Figure 1: High-level architecture of the crules system*

The system is organized into several interconnected components that work together to provide the functionality of crules. The interactive diagram below represents the high-level architecture with color-coded components:

```mermaid
graph TD
    User((👤 User)):::user --> CLI[/Command Line Interface\]:::cli
    
    subgraph Core["🔧 Core System"]
        CLI --> Init[Initialize]:::command
        CLI --> Sync[Synchronize]:::command
        CLI --> Merge[Merge]:::command
        CLI --> Config[Configure]:::command
        CLI --> AgentCmd[Agent Commands]:::command
        
        AgentCmd --> AgentList[List Agents]:::agentOp
        AgentCmd --> AgentSelect[Select Agent]:::agentOp
        AgentCmd --> AgentInfo[Agent Info]:::agentOp
        
        Init --> Registry[(Project Registry)]:::registry
        Sync --> Registry
        Merge --> Registry
        
        Registry --> FileOps[File Operations]:::fileOp
    end
    
    subgraph Storage["💾 Storage"]
        MainLoc[("📁 Main Rules\nLocation")]:::storage <--> FileOps
        Projects[("📁 Project\nRules")]:::storage <--> FileOps
        ConfigFile[("⚙️ Configuration\nFile")]:::storage <--> Config
    end
    
    subgraph UI["🖥️ User Interface"]
        AgentSelect --> TermUI{Terminal UI}:::ui
        AgentList --> TermUI
        AgentInfo --> TermUI
    end
    
    %% Enhanced style definitions
    classDef user fill:#ff9966,stroke:#ff6600,stroke-width:2px,color:#333,font-weight:bold
    classDef cli fill:#66b3ff,stroke:#0066cc,stroke-width:2px,color:#333,font-weight:bold
    classDef command fill:#ffcc99,stroke:#ff9933,stroke-width:2px,color:#333
    classDef agentOp fill:#cc99ff,stroke:#9933ff,stroke-width:2px,color:#333
    classDef registry fill:#ff99cc,stroke:#ff3399,stroke-width:2px,color:#333
    classDef fileOp fill:#99ccff,stroke:#3399ff,stroke-width:2px,color:#333
    classDef storage fill:#99ff99,stroke:#33cc33,stroke-width:2px,color:#333
    classDef ui fill:#ffff99,stroke:#cccc00,stroke-width:2px,color:#333,font-weight:bold
    
    %% Style assignment for subgraphs
    style Core fill:#f0f0ff,stroke:#9999cc,stroke-width:2px,color:#333,font-weight:bold
    style Storage fill:#f0fff0,stroke:#99cc99,stroke-width:2px,color:#333,font-weight:bold
    style UI fill:#fffff0,stroke:#cccc99,stroke-width:2px,color:#333,font-weight:bold
```

*Figure 2: Interactive diagram of system components and their relationships*

## Key Components

### Command Line Interface (CLI)

![CLI Component](../assets/images/diagrams/cli-component.png)
*Figure 3: Command Line Interface component*

The CLI component provides the user interface for the tool. It:
- Processes command line arguments
- Manages the command hierarchy
- Handles input/output formatting
- Provides help and documentation

The CLI is implemented using the Cobra library, which provides a structured approach to defining commands, subcommands, and flags.

<details>
  <summary>📺 View CLI Interaction Example</summary>
  <img src="../assets/gifs/usage/command-usage.gif" alt="CLI interaction example" width="600">
</details>

### Core System

The core system implements the main functionality of the tool:

| Component | Responsibility | Key Functions |
|-----------|----------------|---------------|
| **Project Registry** | Manages project information | Register, list, and validate projects |
| **Rule Management** | Handles rule files | Parse, validate, and synchronize rules |
| **Configuration** | Manages user settings | Load, save, and validate configuration |
| **File Operations** | Handles file system interaction | Copy, compare, and synchronize files |

### Agent System

![Agent System](../assets/images/diagrams/agent-system.png)
*Figure 4: Agent System architecture*

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

The detailed flow of the agent selection process is illustrated in the diagram below:

```mermaid
flowchart TD
    %% Define the main flow
    Start([🚀 Start]):::start --> Command["💻 crules agent select"]:::command
    Command --> LoadConfig["⚙️ Load Configuration"]:::process
    LoadConfig --> ScanRules["🔍 Scan for Rule Files"]:::process
    ScanRules --> ExtractMeta["📋 Extract Agent Metadata"]:::process
    ExtractMeta --> BuildMenu["🖥️ Build Selection Menu"]:::process
    BuildMenu --> DisplayUI["👁️ Display Terminal UI"]:::ui
    
    %% User interaction
    DisplayUI --> UserSelect{"🤔 User Selection"}:::decision
    UserSelect -->|"Choose Agent"| ValidAgent["✅ Valid Agent Selected"]:::valid
    UserSelect -->|"Cancel"| NoSelection["❌ No Selection Made"]:::invalid
    
    %% Process selection
    ValidAgent --> SaveConfig["💾 Save Selected Agent to Config"]:::process
    SaveConfig --> LoadAgent["📂 Load Agent Definition"]:::process
    LoadAgent --> Complete([🎉 Complete]):::end
    
    %% Handle no selection
    NoSelection --> ExitNoChange["🚫 Exit Without Changes"]:::process
    ExitNoChange --> Exit([🚪 Exit]):::end
    
    %% Error handling branch
    ScanRules -->|"No Rules Found"| NoRules["⚠️ No Rules Found"]:::error
    NoRules --> CreatePrompt["❓ Create Rules?"]:::decision
    CreatePrompt -->|"Yes"| InitRules["📝 Initialize Rules"]:::process
    CreatePrompt -->|"No"| Exit
    InitRules --> Command
    
    %% Styles
    classDef start fill:#4CAF50,stroke:#45a049,stroke-width:2px,color:white,font-weight:bold
    classDef end fill:#607D8B,stroke:#546E7A,stroke-width:2px,color:white,font-weight:bold
    classDef command fill:#2196F3,stroke:#1E88E5,stroke-width:2px,color:white
    classDef process fill:#9C27B0,stroke:#8E24AA,stroke-width:2px,color:white
    classDef decision fill:#FF9800,stroke:#F57C00,stroke-width:2px,color:white
    classDef valid fill:#4CAF50,stroke:#45a049,stroke-width:2px,color:white
    classDef invalid fill:#F44336,stroke:#E53935,stroke-width:2px,color:white
    classDef error fill:#F44336,stroke:#E53935,stroke-width:2px,color:white
    classDef ui fill:#00BCD4,stroke:#00ACC1,stroke-width:2px,color:white
    
    %% Add tooltips for better accessibility
    linkStyle 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16 stroke-width:2px,fill:none,stroke:#9E9E9E
    
    %% Add a title
    subgraph "🤖 Agent Selection Process"
    Start
    end
    style "🤖 Agent Selection Process" fill:#ECEFF1,stroke:#CFD8DC,stroke-width:2px,color:#263238,font-weight:bold
```

*Figure 5: Detailed flowchart of the agent selection process*

<details>
  <summary>📺 View Agent Selection Process</summary>
  <img src="../assets/gifs/usage/agent-selection-process.gif" alt="Agent selection process" width="600">
</details>

### Storage Layer

The storage layer manages the persistence of rules and configuration:

| Storage Location | Purpose | Content |
|------------------|---------|---------|
| **Main Rules Location** | Central repository of rules | All rule files (.mdc) |
| **Project Rules** | Project-specific rules | Local copies of rules |
| **Configuration File** | User settings | JSON configuration data |

## Data Flow

### Rule Synchronization Flow

![Synchronization Flow](../assets/images/diagrams/rule-sync-flow.png)
*Figure 6: Rule synchronization flow between main location and projects*

The synchronization process follows a sequential flow as illustrated in this diagram:

```mermaid
sequenceDiagram
    participant User as 👤 User
    participant CLI as 🖥️ CLI
    participant Core as 🔧 Core System
    participant MainLoc as 📁 Main Location
    participant ProjLoc as 📁 Project Location

    %% Define styles using note annotations
    note over User: fill:#ff9966,stroke:#ff6600
    note over CLI: fill:#66b3ff,stroke:#0066cc
    note over Core: fill:#cc99ff,stroke:#9933ff
    note over MainLoc: fill:#99ff99,stroke:#33cc33
    note over ProjLoc: fill:#99ff99,stroke:#33cc33
    
    %% Sequence starts
    User->>+CLI: 🚀 Execute Sync Command
    CLI->>+Core: 📥 Process Command
    
    %% Loading phase
    Core->>Core: 🔍 Load Configuration
    
    %% Validation phase
    alt Main Location Not Set
        Core-->>CLI: ❌ Error: Main Location Not Set
        CLI-->>User: ❌ Display Error
    else Main Location Set
        Core->>+MainLoc: 📂 Scan Files
        MainLoc-->>-Core: 📋 File List & Metadata
        Core->>+ProjLoc: 📂 Scan Files
        ProjLoc-->>-Core: 📋 File List & Metadata
        
        %% Comparison phase
        Core->>Core: 🔍 Compare File Lists
        
        %% Synchronization phase
        loop For Each Updated File
            Core->>+MainLoc: 📤 Read File
            MainLoc-->>-Core: 📄 File Content
            Core->>+ProjLoc: 📥 Write File
            ProjLoc-->>-Core: ✅ Write Confirmation
        end
        
        %% Completion phase
        Core-->>-CLI: 📊 Sync Results
        CLI-->>-User: ✅ Display Success
    end
    
    %% Final confirmation
    Note over User,ProjLoc: 🎉 Synchronization Complete!
```

*Figure 7: Sequence diagram showing the synchronization workflow*

The synchronization process follows these steps:

1. **Initialization**:
   - User invokes `crules sync` command
   - System loads configuration
   - System identifies main and project locations

2. **Comparison**:
   - System scans both locations
   - Files are compared based on existence and modification time
   - System generates a list of files to update

3. **Synchronization**:
   - Files are copied from source to destination
   - System updates file attributes as needed
   - System logs the synchronization actions

4. **Reporting**:
   - System summarizes the changes made
   - User receives confirmation of synchronization

<details>
  <summary>📺 View Synchronization Process</summary>
  <img src="../assets/gifs/workflows/sync-process.gif" alt="Synchronization process" width="600">
</details>

### Agent Selection Flow

The agent selection process follows this flow:

1. **Initialization**:
   - User invokes `crules agent select` command
   - System loads agent registry
   - System identifies available agents

2. **User Interface**:
   - Terminal UI displays agent options
   - User navigates and selects an agent
   - Selection is confirmed by user

3. **Processing**:
   - Selected agent ID is saved to configuration
   - Agent definition is loaded

4. **Reporting**:
   - System confirms agent selection
   - Agent details are displayed to user

## Technical Architecture

### Packaging Structure

```
crules/
├── cmd/
│   └── main.go        # Entry point
├── internal/
│   ├── agent/         # Agent system
│   ├── core/          # Core functionality
│   ├── git/           # Git integration
│   ├── ui/            # User interface
│   ├── utils/         # Utilities
│   └── version/       # Version information
└── .goreleaser.yml    # Release configuration
```

### Dependencies

The project relies on the following key external dependencies:

| Dependency | Purpose | Usage |
|------------|---------|-------|
| **Cobra** | Command line interface | Defining and handling commands |
| **Viper** | Configuration | Managing configuration files |
| **Bubble Tea** | Terminal UI | Interactive terminal interfaces |
| **Lip Gloss** | Terminal styling | Styling terminal output |
| **Afero** | File system abstraction | File operations |

## Error Handling

The error handling strategy follows these principles:

1. **Errors are propagated up**: Lower-level components return errors to be handled at higher levels
2. **Descriptive error messages**: Errors include context to help diagnose issues
3. **Graceful degradation**: The system attempts to continue operation when possible
4. **User-friendly messaging**: Error messages are translated into user-friendly terms in the CLI

![Error Handling](../assets/images/diagrams/error-handling.png)
*Figure 8: Error handling flow in the system*

## Extension Points

The architecture includes several extension points:

| Extension Point | Purpose | How to Extend |
|-----------------|---------|---------------|
| **Commands** | Add new functionality | Create new command handlers in cmd/ |
| **Agents** | Add new agent types | Create new .mdc files in the rules directory |
| **UI Components** | Enhance user interface | Extend the UI package in internal/ui/ |
| **Storage Backends** | Support different storage options | Implement storage interfaces in internal/core/ |

## Future Directions

![Future Roadmap](../assets/images/diagrams/future-roadmap.png)
*Figure 9: Future architecture roadmap*

The architecture is designed to support future enhancements:

1. **Plugin System**: Support for external plugins
2. **Remote Agents**: Support for remotely-hosted agents
3. **Collaboration Features**: Sharing and collaboration on rules
4. **Web Interface**: A browser-based interface for managing rules
5. **API Server**: REST API for programmatic access

## See Also

- [Code Structure](./code-structure.md)
- [Extending Agents](./extending-agents.md)
- [Contributing Guidelines](./contributing.md)
- [API Reference](../api-reference/core-api.md)

## 🔄 Next Steps:
After completing the document review and presenting findings:

"The **Documentation Agent** would be ideal for implementing these recommendations. They can apply the suggested improvements to enhance documentation clarity, accuracy, and organization based on this analysis.

use @documentation-agent.mdc to invoke"

## Navigation

- Next: [Extending Agents](./extending-agents.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md)
