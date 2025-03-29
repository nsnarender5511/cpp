# Agent Workflow Examples

> 📋 This guide provides practical examples of how to use the agent system in real-world scenarios.

## Example 1: Implementing a New Feature

This workflow demonstrates how to use multiple agents to plan and implement a new feature.

### Step 1: Plan with the Technical Wizard

First, use the Technical Wizard agent to get high-level guidance:

```bash
$ crules agent select
# Select "Technical Wizard Agent"

User: I need to add a new command to our CLI that will generate a report of all the rules in a project.

Agent: This is a good feature addition. Here's how I'd approach it:

1. Create a new command called "report" in the CLI package
2. Implement a report generator that can scan rules and extract metadata
3. Provide output in multiple formats (text, JSON, markdown)
4. Add appropriate tests

Let's break this down into concrete tasks for implementation.
```

### Step 2: Detailed Feature Planning

Next, switch to the Feature Planner agent for detailed requirements:

```bash
$ crules agent select
# Select "Feature Planner Agent"

User: I need to plan a "report" command that will generate a summary of all rules in a project.

Agent: Here's a detailed plan for the "report" command:

## Requirements
- Command: `crules report [--format=<format>] [--output=<file>]`
- Formats: text (default), json, markdown
- Information to include:
  - Number of rules
  - Rule types breakdown
  - File sizes
  - Last modified dates
  - Content summary

## Implementation Tasks
1. Create a new file `internal/cli/report.go`
2. Add report command to root command
3. Implement rule scanning logic
4. Create formatters for each output type
5. Implement file output option
```

### Step 3: Implementation

Now, switch to the Implementer agent to write the code:

```bash
$ crules agent select
# Select "Implementer Agent"

User: I need to implement the "report" command as planned. Please start with the basic structure in internal/cli/report.go.

Agent: I'll implement the report command. First, let's create the file structure:

[Agent provides implementation code for report.go with command structure, flags, and basic reporting logic]
```

### Step 4: Testing

Use the Runner agent to test the implementation:

```bash
$ crules agent select
# Select "Runner Agent"

User: I need to test the new report command.

Agent: Let's run some tests for the new report command:

1. First, let's build the binary:
   ```
   go build -o crules ./cmd/crules
   ```

2. Now, let's run basic functionality tests:
   ```
   ./crules report
   ./crules report --format=json
   ./crules report --format=markdown --output=report.md
   ```

[Agent provides analysis of the test results]
```

### Step 5: Documentation

Finally, use the Documentation agent to document the new feature:

```bash
$ crules agent select
# Select "Documentation Agent"

User: I need to document the new report command.

Agent: I'll add documentation for the new report command:

[Agent provides documentation updates for the command, including usage examples, parameter descriptions, and output formats]
```

## Example 2: Fixing a Bug

This workflow demonstrates how to use agents to diagnose and fix a bug.

### Step 1: Analyze with Fix Planner

Start with the Fix Planner agent to analyze the issue:

```bash
$ crules agent select
# Select "Fix Planner Agent"

User: The sync command isn't correctly preserving file permissions. Files end up with 644 permissions after sync.

Agent: Let's analyze this file permission issue:

## Problem Analysis
- The sync command is changing file permissions to 644
- This suggests we're not preserving the original file permissions during copy
- Likely location: internal/projects/sync.go

## Potential Causes
1. Using simple file copy that doesn't preserve permissions
2. Explicitly setting permissions to 644
3. Not using appropriate OS functions

## Fix Strategy
1. Check how files are copied in the sync function
2. Ensure we use functions that preserve metadata (like os.Copy with appropriate flags)
3. Add tests to verify permission preservation
```

### Step 2: Implementation with Implementer

Switch to the Implementer agent to fix the code:

```bash
$ crules agent select
# Select "Implementer Agent"

User: I need to fix the file permission issue in the sync command as analyzed by the Fix Planner.

Agent: I'll fix the file permission issue in the sync command. Let's look at the current implementation first.

[Agent reviews the code, identifies the issue, and provides a fix that preserves file permissions]
```

### Step 3: Verification with Runner

Use the Runner agent to verify the fix:

```bash
$ crules agent select
# Select "Runner Agent"

User: I need to test the fix for the file permissions issue.

Agent: Let's test the fix for the file permissions issue:

1. First, let's build the binary:
   ```
   go build -o crules ./cmd/crules
   ```

2. Now, let's create a test file with specific permissions:
   ```
   touch testfile
   chmod 755 testfile
   ```

3. Let's run the sync command and check if permissions are preserved:
   ```
   ./crules sync
   ls -la testfile
   ```

[Agent analyzes the results and confirms that permissions are now preserved]
```

## Example 3: Code Review

This example shows how to use agents for code review.

### Step 1: Review with Technical Wizard

Use the Technical Wizard agent to review code at a high level:

```bash
$ crules agent select
# Select "Technical Wizard Agent"

User: Can you review this new implementation of the agent selection UI?

Agent: I'll review the agent selection UI implementation:

## High-Level Review
- The approach using a terminal UI library is solid
- The separation of UI logic from business logic is good
- There are a few architectural concerns to address

## Strengths
- Clean component-based design
- Good error handling
- Responsive UI with keyboard shortcuts

## Areas for Improvement
- Consider extracting the UI theme into a configuration
- The agent loading might benefit from background caching
- Some error messages could be more user-friendly

Overall, this is a well-implemented feature but with room for some architectural refinements.
```

### Step 2: Detailed Implementation Feedback

Switch to the Implementer agent for detailed code feedback:

```bash
$ crules agent select
# Select "Implementer Agent"

User: Can you provide detailed implementation feedback on the agent selection UI?

Agent: Here's my detailed implementation feedback:

[Agent provides specific code-level feedback, including suggestions for refactoring, optimizing performance, and improving error handling]
```

## Best Practices for Agent Workflows

1. **Use specialized agents for specific tasks**: Each agent is optimized for certain types of tasks.

2. **Chain agents in a logical sequence**: Start with planning, move to implementation, then testing and documentation.

3. **Provide context when switching agents**: When you select a new agent, briefly summarize what previous agents have done.

4. **Be specific in your requests**: Clearly state what you need from each agent to get the most relevant help.

5. **Review agent suggestions critically**: While agents provide expert guidance, always review their suggestions critically.

## See Also

- [Agent System Documentation](../user-guide/agents.md)
- [Command Reference](../user-guide/commands.md)
- [Extending Agents](../developer-guide/extending-agents.md) 