# Extending the Agent System

> 🧩 This guide explains how to create and extend agents in the crules system, including the format, structure, and best practices for agent definition files.

## Agent Definition Files

Agents in crules are defined using Markdown files with the `.mdc` extension. These files contain both the agent's metadata (name, description, role, capabilities) and the full agent definition (instructions, examples, etc.).

### File Structure

Each agent definition file (`*.mdc`) follows this structure:

```markdown
# Agent Name

## 🎯 Role: 
One-line description of the agent's primary role

## 📝 Description:
A more detailed description of what the agent does, its purpose, and when to use it.

## ✅ Capabilities:
- Capability 1: Description of the first capability
- Capability 2: Description of the second capability
- Capability 3: Description of the third capability

## 🔍 Examples:
Examples of how to use this agent effectively.

## 📋 Instructions:
Detailed instructions for the agent, defining its behavior, constraints, and operation.

## 🧠 Knowledge:
Specific domain knowledge that the agent needs to perform its role effectively.
```

### Required Sections

The following sections are required in every agent definition file:

- **Title** (`# Agent Name`): The name of the agent
- **Role** (`## 🎯 Role:`): A one-line description of the agent's primary role
- **Description** (`## 📝 Description:`): A more detailed description of the agent
- **Capabilities** (`## ✅ Capabilities:`): A list of the agent's key capabilities

### Optional Sections

These sections are optional but recommended:

- **Examples** (`## 🔍 Examples:`): Examples of how to use the agent
- **Instructions** (`## 📋 Instructions:`): Detailed instructions for the agent
- **Knowledge** (`## 🧠 Knowledge:`): Domain-specific knowledge for the agent

## Agent Metadata

### Name and Title

The agent's name is defined by the level-1 heading (`#`) at the beginning of the file. This name is displayed in the agent selection UI and when listing agents.

Examples:
```markdown
# Technical Wizard Agent
```

```markdown
# 🧙‍♂️ Technical Wizard Agent
```

### Role

The role should be a concise one-line description that clearly states the agent's primary function.

Example:
```markdown
## 🎯 Role:
Provides high-level technical guidance and coordinates other specialized agents
```

### Description

The description should provide more detail about what the agent does, when to use it, and its primary value.

Example:
```markdown
## 📝 Description:
The Technical Wizard agent helps with architecture decisions, design patterns, 
and clean code principles. It acts as a coordinator that can suggest when to use 
other specialized agents for specific tasks. Use this agent when you need 
high-level guidance on approaching complex technical problems.
```

### Capabilities

Capabilities should be specific skills or functions that the agent can perform. Each capability should be preceded by a level-3 heading (`###`) with a checkmark emoji.

Example:
```markdown
## ✅ Capabilities:

### ✅ In-Depth Technical Exploration and Analysis
Conducts thorough investigation of technical concepts, frameworks, and approaches, 
providing comprehensive analysis of options and trade-offs.

### ✅ Expert Architectural Guidance
Offers expert guidance on system architecture, component organization, and 
implementation approaches based on industry best practices.

### ✅ Design Patterns Discussion
Explains and recommends appropriate design patterns for specific technical challenges, 
with context-specific implementation guidance.
```

## Creating a Basic Agent

Here's a step-by-step guide to creating a simple agent:

1. Create a new file in the `~/.cursor/rules/` directory (or your configured main location) with the `.mdc` extension:

```bash
touch ~/.cursor/rules/hello-agent.mdc
```

2. Add the required sections to your agent definition:

```markdown
# 👋 Hello Agent

## 🎯 Role:
Provides friendly greetings and introductions

## 📝 Description:
This simple agent demonstrates the basic structure of an agent definition. 
It responds to greetings and introduces itself to the user.

## ✅ Capabilities:

### ✅ Friendly Greetings
Responds to user greetings with appropriate and friendly responses.

### ✅ Self-Introduction
Can introduce itself and explain its capabilities when asked.

## 🔍 Examples:
- "Hello! I'm the Hello Agent. How can I help you today?"
- "Nice to meet you! I'm a simple example agent that shows how agents work."

## 📋 Instructions:
When interacting with the user:
1. Respond to greetings with a friendly greeting of your own
2. If asked about your capabilities, explain that you're a simple example agent
3. Keep responses concise and friendly
4. If asked to do something beyond your capabilities, politely explain your limitations
```

3. Save the file and make it available to crules by running:

```bash
crules merge
```

4. Verify that your agent is recognized:

```bash
crules agent
```

You should see your new agent in the list of available agents.

## Advanced Agent Creation

### Agent Class System

Complex agents can be organized into classes based on their role:

- **Planning Agents**: Focus on planning and design (architecture, features, fixes)
- **Implementation Agents**: Focus on code implementation and execution
- **Support Agents**: Focus on ancillary tasks (documentation, review, git operations)

You can add a class hint to your agent using a comment at the top of the file:

```markdown
<!-- Agent Class: Planning -->
# Feature Planner Agent
```

### Agent Relationships

You can define relationships between agents to create workflows. This is done using the `## 🔄 Related Agents:` section:

```markdown
## 🔄 Related Agents:
- @feature-planner.mdc: For detailed feature planning
- @implementer.mdc: For implementing the plans
- @runner.mdc: For testing the implementation
```

### Agent Versioning

Include a version number to track changes to your agent:

```markdown
## 📌 Version: 1.2.0
```

Use semantic versioning (MAJOR.MINOR.PATCH) where:
- MAJOR: Breaking changes
- MINOR: New capabilities added in a backward-compatible manner
- PATCH: Backward-compatible bug fixes

### Agent Metadata Table

You can include a metadata table at the top of your file for machine readability:

```markdown
---
name: Technical Wizard Agent
version: 1.0.0
author: Your Name
created: 2023-03-29
updated: 2023-03-30
class: Planning
related: [feature-planner, implementer, code-reviewer]
---

# 🧙‍♂️ Technical Wizard Agent
```

## Best Practices

### Naming Conventions

- Use descriptive names that clearly indicate the agent's function
- Follow a consistent naming pattern for related agents
- Consider adding emojis to names for visual recognition
- Use `-` (hyphen) instead of spaces in filenames

### Content Structure

- Keep the agent's role and description concise
- Organize capabilities from most important to least important
- Use markdown formatting to improve readability:
  - `**bold**` for emphasis
  - `*italic*` for terminology
  - `` `code` `` for code snippets
  - Bullet lists for options or items
  - Numbered lists for steps or sequences

### Performance Optimization

- Only include necessary information in the agent definition
- Break very large agents into multiple specialized agents
- For long knowledge sections, consider using references instead of embedding all content

### Testing Your Agents

After creating an agent, test it thoroughly:

1. Select it using `crules agent select`
2. Verify that it appears correctly in `crules agent info`
3. Test it with various inputs to ensure it behaves as expected
4. Make adjustments as needed and retest

## Example Agents

### Planning Agent Example

```markdown
# 📝 Feature Planner Agent

## 🎯 Role:
Plans the implementation of new features with detailed breakdown of tasks and requirements

## 📝 Description:
The Feature Planner agent helps break down feature requirements into 
implementable components, creates detailed implementation plans, and 
provides guidance on how to approach new feature development. Use this 
agent when you need to plan the implementation of a new feature or enhancement.

## ✅ Capabilities:

### ✅ Requirement Analysis
Analyzes feature requirements and identifies key components and dependencies.

### ✅ Task Breakdown
Breaks down complex features into manageable, implementable tasks.

### ✅ Implementation Planning
Creates detailed step-by-step plans for feature implementation.

## 🔍 Examples:
"Here's a high-level plan for implementing the user authentication feature:
1. Set up the database schema for user accounts
2. Create the authentication API endpoints
3. Implement the frontend login/registration forms
4. Add session management
5. Implement authorization middleware
..."

## 🔄 Related Agents:
- @wizard.mdc: For high-level technical guidance
- @implementer.mdc: For implementing the planned features
- @runner.mdc: For testing the implementation
```

### Implementation Agent Example

```markdown
# ⚙️ Implementer Agent

## 🎯 Role:
Translates detailed plans into working code following established patterns and standards

## 📝 Description:
The Implementer agent is focused on writing clean, efficient code to implement 
plans created by planning agents. It follows established coding standards and 
patterns, and is skilled at translating high-level plans into working implementations.

## ✅ Capabilities:

### ✅ Code Implementation
Translates detailed plans into working code that follows best practices.

### ✅ Pattern Application
Applies appropriate design patterns to solve implementation challenges.

### ✅ Code Integration
Ensures new code integrates well with existing codebase.

## 🔄 Related Agents:
- @feature-planner.mdc: For planning what to implement
- @fix-planner.mdc: For planning bug fixes
- @runner.mdc: For testing implementations
- @code-reviewer.mdc: For reviewing the implementation
```

## Deploying Agents to Teams

### Sharing Agents with Your Team

To share agents with your team:

1. Create your agent definitions in your main rules location
2. Test them thoroughly
3. Share them with your team through:
   - Version control (GitHub, GitLab, etc.)
   - Direct file sharing
   - Importing from a URL

### Setting Up a Team Agent Repository

For larger teams, consider setting up a dedicated repository for agent definitions:

1. Create a git repository for your team's agents
2. Organize agents by type/function in directories
3. Include README files explaining the available agents
4. Use git tags for versioning

Team members can initialize crules with this repository:

```bash
# During first-time setup
crules init

# When prompted, choose to fetch from git repository
# Enter your team's repository URL
```

## Advanced Techniques

### Dynamic Agent Generation

You can generate agent definitions programmatically based on templates:

```bash
#!/bin/bash
# Simple agent generator script

NAME="$1"
FILENAME="$(echo "$NAME" | tr ' ' '-' | tr '[:upper:]' '[:lower:]').mdc"
ROLE="$2"

cat > "$FILENAME" << EOF
# $NAME

## 🎯 Role:
$ROLE

## 📝 Description:
[Add a detailed description here]

## ✅ Capabilities:

### ✅ Capability 1
[Describe this capability]

### ✅ Capability 2
[Describe this capability]
EOF

echo "Agent created: $FILENAME"
```

### Conditional Sections

You can include conditional sections that only apply in certain contexts:

```markdown
## 🔄 Context-Specific Behavior:

### When working with Frontend code:
- Use React best practices
- Follow component-based architecture

### When working with Backend code:
- Ensure proper error handling
- Use appropriate data validation
```

## See Also

- [Agent System](../user-guide/agents.md) for information on using agents
- [Code Structure](./code-structure.md) for details on the codebase organization
- [API Reference](../api-reference/agent-api.md) for the Agent API reference

## Navigation

- Previous: [Architecture](./architecture.md)
- Next: [Contributing](./contributing.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md) 