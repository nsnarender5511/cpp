---
version: v0.1.0
last_updated: 2023-03-29
applies_to: crules v0.1.0+
---

# Agent System

> 🧠 The Agent System in crules allows you to work with specialized AI agents for different tasks, from planning and implementing features to fixing issues and documenting code.

## Overview

The crules Agent System provides an interactive way to discover, select, and use various AI agents defined in your `.cursor/rules` directory. Each agent is specialized for a particular task, such as:

- Planning new features 
- Fixing bugs and issues
- Implementing code based on plans
- Running and testing implementations
- Creating documentation
- Reviewing code

Agents are defined in Markdown (`.mdc`) files that contain both the agent's metadata (name, description, capabilities) and its full definition.

## Agent Ecosystem

The agents in crules work together as an ecosystem, with different agents specializing in different phases of the development lifecycle. The diagram below illustrates how these agents interact:

```mermaid
graph TD
    %% Main agent nodes
    User((👤 User)):::user
    Wizard([🧙 Technical\nWizard]):::wizard
    
    %% Planning agents
    FeaturePlanner[💡 Feature\nPlanner]:::planner
    FixPlanner[🔧 Fix\nPlanner]:::planner
    ArchPlanner[🏗️ Architecture\nPlanner]:::planner
    RefactorGuru[♻️ Refactoring\nGuru]:::planner
    ScraperPlanner[🕸️ Scraper\nPlanner]:::planner
    
    %% Implementation agents
    Implementer[⚙️ Implementer]:::implementer
    Runner[🏃 Runner]:::implementer
    
    %% Support agents
    CodeReviewer[🔍 Code\nReviewer]:::support
    GitCommitter[📝 Git\nCommitter]:::support
    DocAgent[📚 Documentation\nAgent]:::support
    GitAgent[📊 Git\nAgent]:::support
    
    %% Coordination flow
    User --> Wizard
    Wizard --> Planning
    Wizard --> Implementation
    Wizard --> Support
    
    %% Planning subgraph
    subgraph Planning[" Planning Phase "]
        direction TB
        FeaturePlanner
        FixPlanner
        ArchPlanner
        RefactorGuru
        ScraperPlanner
    end
    
    %% Implementation subgraph
    subgraph Implementation[" Implementation Phase "]
        direction TB
        Implementer
        Runner
    end
    
    %% Support subgraph
    subgraph Support[" Support Phase "]
        direction TB
        CodeReviewer
        GitCommitter
        DocAgent
        GitAgent
    end
    
    %% Inter-agent workflows
    FeaturePlanner ==> Implementer
    FixPlanner ==> Implementer
    ArchPlanner ==> FeaturePlanner
    ArchPlanner ==> Implementer
    RefactorGuru ==> Implementer
    ScraperPlanner ==> Implementer
    
    Implementer ==> Runner
    Runner ==> CodeReviewer
    CodeReviewer ==> GitCommitter
    Implementer ==> DocAgent
    GitCommitter ==> GitAgent
    
    %% Feedback loops
    CodeReviewer -.-> Implementer
    Runner -.-> FixPlanner
    
    %% Styles for nodes
    classDef user fill:#ff9966,stroke:#ff6600,stroke-width:2px,color:#333,font-weight:bold
    classDef wizard fill:#ff66b2,stroke:#cc0066,stroke-width:3px,color:white,font-weight:bold
    classDef planner fill:#9966ff,stroke:#6600cc,stroke-width:2px,color:white
    classDef implementer fill:#66b3ff,stroke:#0066cc,stroke-width:2px,color:white
    classDef support fill:#66cc99,stroke:#009966,stroke-width:2px,color:white
    
    %% Styles for subgraphs
    style Planning fill:#f5f0ff,stroke:#d5c0ff,stroke-width:2px,color:#333,font-weight:bold
    style Implementation fill:#f0f5ff,stroke:#c0d5ff,stroke-width:2px,color:#333,font-weight:bold
    style Support fill:#f0fff5,stroke:#c0ffd5,stroke-width:2px,color:#333,font-weight:bold
    
    %% Styles for linkages
    linkStyle 0,1,2,3 stroke:#ff66b2,stroke-width:3px
    linkStyle 4,5,6,7,8,9,10,11,12,13,14 stroke:#0066cc,stroke-width:2px
    linkStyle 15,16 stroke:#ff6600,stroke-width:2px,stroke-dasharray: 5 5
    
    %% Title
    subgraph "🤖 Agent Relationships and Workflow"
    end
    style "🤖 Agent Relationships and Workflow" fill:none,stroke:none
```

*Figure 1: The crules agent ecosystem and interactions*

In this ecosystem:
- The **Technical Wizard** acts as a coordinator, helping users decide which specialized agents to use
- **Planning agents** (purple) work on designing and planning tasks
- **Implementation agents** (blue) execute the plans and verify the results
- **Support agents** (green) handle ancillary tasks like documentation, code review, and version control
- The solid lines show typical workflows, while dotted lines represent feedback loops

This modular approach allows you to use the right agent for each specific task while maintaining a coherent development workflow.

![Agent Workflow Example](../assets/images/examples/agent-workflow.png)
*Figure 2: Example of a complete agent-assisted development workflow*

## Using Agents

You can interact with agents using the following commands:

- `crules agent` - Display all available agents (default behavior)
- `crules agent info <id>` - Show detailed information about a specific agent
- `crules agent select` - Interactively select and load an agent

You can also reference agents directly in the chatbox using the `@` symbol (e.g., `@wizard.mdc`).

<details>
  <summary>📺 View Agent Selection Process</summary>
  <img src="../assets/gifs/usage/agent-selection-process.gif" alt="Agent selection process" width="600">
</details>

## Listing Available Agents

To see all available agents, use:

```bash
crules agent
```

This will display a formatted table of all agents with their information. The table adapts to your terminal width to provide the optimal display format:

### Display Examples

**Wide Terminal Display:**
```
+-----+---------------------+--------------------+----------+
| No. | Agent Name          | Reference ID       | Version  |
+-----+---------------------+--------------------+----------+
| 1   | Feature Planner     | @feature-planner.mdc | 1.0    |
| 2   | Fix Planner         | @fix-planner.mdc     | 1.0    |
| 3   | Runner              | @runner.mdc          | 1.0    |
| 4   | Technical Wizard    | @wizard.mdc          | 1.0    |
| ... | ...                 | ...                  | ...    |
+-----+---------------------+--------------------+----------+
```

**Medium Terminal Display:**
```
+-----+----------------+--------------------+
| No. | Name           | Reference          |
+-----+----------------+--------------------+
| 1   | Feature Planner| @feature-planner.mdc |
| 2   | Fix Planner    | @fix-planner.mdc     |
| 3   | Runner         | @runner.mdc          |
| ... | ...            | ...                  |
+-----+----------------+--------------------+
```

**Narrow Terminal Display:**
```
+-----+------------------+
| No. | Agent            |
+-----+------------------+
| 1   | feature-planner  |
| 2   | fix-planner      |
| 3   | runner           |
| ... | ...              |
+-----+------------------+
```

The table format ensures proper alignment and clarity regardless of your terminal size, making agent information easy to read and reference.

## Referencing Agents

There are multiple ways to reference agents in crules, providing flexibility based on your preference and needs.

### Referencing by String ID

You can reference agents by their unique string ID, which is derived from the filename without the `.mdc` extension:

```bash
crules agent info wizard
```

This method is stable across sessions and reorderings of the agent list.

### Referencing by Numeric Index

You can also reference agents by their position number shown in the agent list:

```bash
crules agent info 1  # References the first agent in the list
```

This method is convenient for quick access when you can see the number in the list but might not remember the exact ID.

Benefits of numeric index referencing:
- Shorter to type
- Easier for sequential exploration of agents
- Direct visual correspondence with the displayed list

### Agent @ References

The quickest way to invoke a specific agent in the chatbox is by using the `@` reference. Simply type `@` followed by the agent ID to invoke that agent:

```
@wizard.mdc I need help designing a new API endpoint
```

```
@quick-answer-agent.mdc What does the HTTP 418 status code mean?
```

This method allows you to quickly switch between different agent specializations without running additional commands.

### Agent Info

The `info` command provides detailed information about a specific agent:

```
$ crules agent info wizard

Agent details:
  ID:          wizard
  Name:        🧙‍♂️ Technical Wizard Agent
  Version:     1.0

Description:
  [Full agent description with formatted markdown]

Capabilities:
  - In-Depth Technical Exploration and Analysis
  - Expert Architectural Guidance
  - Design Patterns Discussion
  - Clean Code Advisory

File: /Users/username/Library/Application Support/crules/.cursor/rules/wizard.mdc
```

Or using the numeric index:

```
$ crules agent info 6  # Assuming 6 is the position number for the wizard agent

Agent details:
  ID:          wizard
  Name:        🧙‍♂️ Technical Wizard Agent
  Version:     1.0
  
  ...
```

### Agent Selection

The `select` command presents an interactive menu for choosing an agent:

1. Run `crules agent select`
2. Browse the list of available agents
3. Enter the number of the agent you want to select
4. The agent will be loaded, and you'll see its details
5. You can optionally view the full agent definition in a paginated format

![Agent Selection UI](../assets/images/ui/agent-selection-interface.png)
*Figure 3: The agent selection terminal UI*

## Common Agents

The system comes with several pre-defined agents, each specialized in different aspects of software development:

| Agent | Icon | Primary Role | When to Use |
|-------|------|--------------|-------------|
| **Technical Wizard** | 🧙‍♂️ | High-level guidance & coordination | Starting a project, making architectural decisions |
| **Feature Planner** | ✨ | Feature implementation planning | Breaking down new feature requirements |
| **Fix Planner** | 🔍 | Bug analysis and fix planning | Diagnosing and planning fixes for issues |
| **Implementer** | 🛠️ | Code implementation | Translating plans into working code |
| **Runner** | 🏃 | Testing and verification | Running and verifying implementations |
| **Documentation Agent** | 📚 | Documentation creation | Creating and organizing documentation |
| **Code Reviewer** | 👁️ | Code quality assessment | Reviewing code for quality and issues |
| **Git Committer** | 📝 | Version control assistance | Creating commit messages and managing changes |

## Recommended Agent Workflows

For the best results, consider these common workflows:

### New Feature Development

1. Start with the **Technical Wizard** to discuss the high-level approach
2. Use the **Feature Planner** to create a detailed implementation plan
3. Have the **Implementer** translate the plan into code
4. Verify with the **Runner** to test the implementation
5. Let the **Documentation Agent** create documentation for the feature
6. Finish with the **Git Committer** to prepare the changes for commit

### Bug Fixing

1. Begin with the **Fix Planner** to analyze the bug
2. Use the **Implementer** to implement the fix
3. Verify with the **Runner** that the bug is resolved
4. Update documentation with the **Documentation Agent** if needed

![Agent Workflow Animation](../assets/gifs/workflows/agent-workflow.gif)
*Figure 4: Animated demonstration of an agent workflow*

## Adding Custom Agents

You can add your own custom agents by creating new `.mdc` files in your `.cursor/rules` directory. These agents will automatically be discovered by the crules tool and become available in the agent selection menu.

For information on creating custom agents, see the [Extending the Agent System](../developer-guide/extending-agents.md) guide in the Developer Documentation.

---

## Navigation

- Previous: [Commands](./commands.md)
- Next: [Troubleshooting](./troubleshooting.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md) 