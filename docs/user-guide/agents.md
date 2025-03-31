---
version: v1.0.0
last_updated: 2023-03-31
applies_to: cursor++ v1.0.0+
---

# Agent System

> 🧠 The Agent System in cursor++ allows you to work with specialized AI agents for different tasks, from planning and implementing features to fixing issues and documenting code.

## Overview

The cursor++ Agent System provides an interactive way to discover, select, and use various AI agents defined in your `.cursor/rules` directory. Each agent is specialized for a particular task, such as:

- Planning new features 
- Fixing bugs and issues
- Implementing code based on plans
- Running and testing implementations
- Creating and synchronizing documentation
- Reviewing code

Agents are defined in Markdown (`.mdc`) files that contain both the agent's metadata (name, description, capabilities) and its full definition.

## Agent Ecosystem

The agents in cursor++ work together as an ecosystem, with different agents specializing in different phases of the development lifecycle. The diagram below illustrates how these agents interact:

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
    DocSyncer[🔄 Document\nSyncer]:::support
    QuickAnswer[⚡ Quick\nAnswer]:::support
    
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
        DocSyncer
        QuickAnswer
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
    DocAgent ==> DocSyncer
    GitCommitter ==> User
    
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
    linkStyle 4,5,6,7,8,9,10,11,12,13,14,15 stroke:#0066cc,stroke-width:2px
    linkStyle 16,17 stroke:#ff6600,stroke-width:2px,stroke-dasharray: 5 5
    
    %% Title
    subgraph "🤖 Agent Relationships and Workflow"
    end
    style "🤖 Agent Relationships and Workflow" fill:none,stroke:none
```

*Figure 1: The cursor++ agent ecosystem and interactions*

In this ecosystem:
- The **Technical Wizard** acts as a coordinator, helping users decide which specialized agents to use
- **Planning agents** (purple) work on designing and planning tasks
- **Implementation agents** (blue) execute the plans and verify the results
- **Support agents** (green) handle ancillary tasks like documentation, code review, and version control
- The solid lines show typical workflows, while dotted lines represent feedback loops

This modular approach allows you to use the right agent for each specific task while maintaining a coherent development workflow.

## Using Agents

You can interact with agents using the following commands:

- `cursor++ agent` - Display all available agents
- `cursor++ agent info <id>` - Show detailed information about a specific agent
- `cursor++ agent select` - Interactively select and load an agent

You can also reference agents directly in the chatbox using the `@` symbol (e.g., `@wizard.mdc`).

## Listing Available Agents

To see all available agents, use:

```bash
cursor++ agent
```

This will display a formatted table of all agents with their information. The table adapts to your terminal width to provide the optimal display format.

## Available Agents

cursor++ includes a rich ecosystem of specialized agents:

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

## Referencing Agents

There are multiple ways to reference agents in cursor++, providing flexibility based on your preference and needs.

### Referencing by String ID

You can reference agents by their unique string ID, which is derived from the filename without the `.mdc` extension:

```bash
cursor++ agent info wizard
```

This method is stable across sessions and reorderings of the agent list.

### Referencing by Numeric Index

You can also reference agents by their position number shown in the agent list:

```bash
cursor++ agent info 1  # References the first agent in the list
```

This method is convenient for quick access when you can see the number in the list but might not remember the exact ID.

### Agent @ References

The quickest way to invoke a specific agent in the chatbox is by using the `@` reference. Simply type `@` followed by the agent ID to invoke that agent:

```
@wizard.mdc I need help designing a new API endpoint
```

```
@quick-answer-agent.mdc What does the HTTP 418 status code mean?
```

This method allows you to quickly switch between different agent specializations without running additional commands.

## Creating Your Own Agents

You can create your own agents by adding new `.mdc` files to your `.cursor/rules` directory. Each agent definition should include:

1. **Metadata**: Agent name, purpose, and version
2. **Role Description**: Clear explanation of the agent's responsibilities
3. **Core Capabilities**: What tasks the agent can perform
4. **Boundaries**: What actions the agent should avoid

For example, to create a new documentation agent:

```markdown
# 📝 Documentation Agent

## 🎯 Role:
You are a **Documentation Agent**, specialized in creating and improving documentation.

## 🛠️ Core Responsibilities:
- Create clear, concise documentation
- Improve existing documentation
- Format documentation according to project standards
- Ensure consistency across documentation files

## 🚫 Boundaries:
- Do not modify code files
- Do not delete existing documentation without confirmation
```

## Agent Collaboration

Agents are designed to collaborate. When implementing complex workflows, you can chain agents together:

1. Use the **Feature Planner** to design a new feature
2. Use the **Implementer** to build the feature
3. Use the **Runner** to test the implementation
4. Use the **Code Reviewer** to verify code quality
5. Use the **Documentation Agent** to document the feature
6. Use the **Git Committer** to create commit messages

## Recommended Workflows

For the best results with cursor++, consider these common agent workflows:

### New Feature Development
1. Start with the **Technical Wizard** to discuss high-level approach
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

### Documentation Management
1. Use the **Document Syncer** to identify inconsistencies
2. Use the **Documentation Agent** to update documentation
3. Use the **Document Reviewer** to verify quality

---

## Navigation

- Previous: [Commands](./commands.md)
- Next: [Troubleshooting](./troubleshooting.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md) 