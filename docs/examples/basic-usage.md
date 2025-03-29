# Basic Usage Examples

> 🧪 This guide provides practical examples of common crules workflows.

## Overview

The examples in this guide demonstrate real-world scenarios and common tasks you'll encounter when using crules. Each example includes:

- A brief description of the task
- Step-by-step instructions
- Expected output
- Tips and variations

## Example 1: Setting Up a New Project

This example shows how to initialize a new project with rules from your main location.

### Scenario

You're starting a new project and want to set it up with your existing rules.

### Steps

1. Create a new project directory:
   ```bash
   mkdir ~/projects/new-project
   cd ~/projects/new-project
   ```

2. Initialize a git repository (optional but recommended):
   ```bash
   git init
   ```

3. Initialize crules:
   ```bash
   crules init
   ```

4. Verify initialization:
   ```bash
   ls -la .cursor/rules
   ```

### Expected Output

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Initializing rules in /Users/username/projects/new-project/.cursor/rules...
Registering project in registry...
Updating .gitignore...

✓ Successfully initialized rules in /Users/username/projects/new-project/.cursor/rules
```

### Tips

- If you're using version control, make sure `.cursor/` is in your `.gitignore`
- You can initialize multiple projects and keep them all in sync with your main rules location

## Example 2: Syncing Rules across Projects

This example demonstrates how to keep rules in sync between multiple projects.

### Scenario

You have multiple projects and want to ensure they all have the same latest rules.

### Steps

1. Navigate to your project:
   ```bash
   cd ~/projects/project-one
   ```

2. Sync the latest rules from your main location:
   ```bash
   crules sync
   ```

3. Verify that the rules are up to date:
   ```bash
   ls -la .cursor/rules
   ```

### Expected Output

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Syncing rules from main location...

✓ Successfully synced rules from main location
```

### Tips

- Run `crules sync` regularly to ensure your projects have the latest rules
- You can automate this with a script or git hook

## Example 3: Creating a Custom Agent

This example shows how to create a custom agent for a specific purpose.

### Scenario

You want to create a custom agent that specializes in code review for your team's coding standards.

### Steps

1. Navigate to your main rules location (usually `~/.cursor/rules`):
   ```bash
   cd ~/.cursor/rules
   ```

2. Create a new agent definition file:
   ```bash
   touch code-review-agent.mdc
   ```

3. Edit the file with your preferred editor and add content:
   ```markdown
   # 🔍 Code Review Agent

   ## 🎯 Role:
   Specialized agent for reviewing code according to team standards

   ## 📝 Description:
   This agent specializes in reviewing code according to our team's coding standards,
   focusing on readability, maintainability, and performance.

   ## ✅ Capabilities:

   ### ✅ Style Guide Enforcement
   Ensures code adheres to the team's style guide, including formatting,
   naming conventions, and documentation standards.

   ### ✅ Performance Review
   Identifies potential performance issues such as inefficient algorithms,
   unnecessary computations, and resource leaks.

   ### ✅ Best Practices
   Recommends industry best practices and design patterns appropriate for
   the specific context.

   ## 📋 Instructions:
   When reviewing code, I will:
   1. Check for adherence to the team's style guide
   2. Look for potential performance issues
   3. Suggest improvements based on best practices
   4. Provide clear explanations for all recommendations
   ```

4. Make the agent available to all your projects:
   ```bash
   crules merge
   ```

### Expected Output

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Merging rules to main location...
Syncing changes to 3 registered projects...

✓ Successfully merged rules to main location
```

### Tips

- Test your agent definition before sharing it with your team
- Use emojis in agent names to make them more distinct and memorable
- Include detailed instructions to get consistent results

## Example 4: Using the Agent Selection Interface

This example demonstrates how to interactively select and use agents.

### Scenario

You want to choose an agent for a specific task from the available agents.

### Steps

1. Navigate to your project:
   ```bash
   cd ~/projects/my-project
   ```

2. Open the agent selection interface:
   ```bash
   crules agent select
   ```

3. Use arrow keys to navigate to your desired agent and press Enter to select it.

### Expected Output

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Select an agent:
> 1. 🧙‍♂️ Technical Wizard Agent
  2. 🔍 Code Review Agent
  3. 📝 Feature Planner Agent
  4. 🛠️ Fix Planner Agent

[Use arrow keys to navigate, Enter to select]
```

After selection:

```
Selected agent: 🔍 Code Review Agent

Agent details:
  ID:          code-review-agent
  Name:        🔍 Code Review Agent
  Version:     1.0

Description:
  This agent specializes in reviewing code according to our team's coding standards,
  focusing on readability, maintainability, and performance.

Agent saved as default.
```

### Tips

- Use `crules agent info <id>` to see agent details without selecting it
- The last selected agent is remembered for future use
- Agent selection works with both string IDs and numeric indices

## Example 5: Importing Rules from a URL

This example shows how to import agent rules from a URL.

### Scenario

You've found a useful agent online and want to import it into your rules collection.

### Steps

1. Import the agent from a URL:
   ```bash
   crules import https://cursor.directory/example-agent
   ```

2. Confirm the import when prompted.

3. Verify the agent was imported:
   ```bash
   crules agent
   ```

### Expected Output

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Fetching rules from https://cursor.directory/example-agent...
Parsing rules...

Found 1 rule to import:
  1. Example Agent - An example agent for crules

Do you want to import this rule? [y/N] y

Storing rules to main location...

✓ Successfully imported 1 rule
```

### Tips

- You can import multiple agents from different sources
- Always review imported agents before using them
- Use caution when importing agents from untrusted sources

## Example 6: Managing Projects

This example demonstrates how to manage registered projects.

### Scenario

You want to see which projects are registered with crules and clean up any that no longer exist.

### Steps

1. List all registered projects:
   ```bash
   crules list
   ```

2. Clean up non-existent projects:
   ```bash
   crules clean
   ```

3. Verify the cleanup:
   ```bash
   crules list
   ```

### Expected Output

For `crules list`:

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

Registered projects (3):
  1. /Users/username/projects/project1
  2. /Users/username/projects/project2
  3. /Users/username/old-project (not found)

⚠️ 1 project(s) could not be found. Run 'crules clean' to remove them.
```

For `crules clean`:

```
✨ crules - Cursor Rules Manager v0.1.0 ✨

✓ Successfully removed 1 non-existent project(s) from registry.
```

### Tips

- Run `crules clean` periodically to keep your registry clean
- Use `crules list` to quickly see all your registered projects
- When deleting a project, run `crules clean` to remove it from the registry

## Navigation

- Previous: [Troubleshooting](../user-guide/troubleshooting.md)
- Next: [Advanced Usage](./advanced-usage.md)
- Up: [Examples](./README.md)
- Home: [Documentation Home](../README.md)
