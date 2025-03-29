# Basic Usage Examples

> 🧩 **Simple examples to get started with crules**

## Initial Setup

This example demonstrates setting up a new project with Cursor rules.

```bash
# Navigate to your project directory
cd /path/to/your/project

# Initialize the project with rules from your main location
crules init

# You should see output confirming successful initialization
# The rules are now available in your project's .cursor/rules directory
```

## Syncing Rules

This example shows how to update your project with the latest rules from the main location.

```bash
# Navigate to your project directory
cd /path/to/your/project

# Sync rules from the main location to your project
crules sync

# Any local changes to rules will be overwritten with the latest version
# from the main location
```

## Making and Merging Changes

This example demonstrates making changes to rules in a project and merging them back to the main location.

```bash
# Navigate to your project directory
cd /path/to/your/project

# Make changes to your rules in .cursor/rules
# For example, edit .cursor/rules/my-rule.md

# Once you're satisfied with your changes, merge them back
crules merge

# This will update the main location with your changes
# and sync them to all other registered projects
```

## Listing Registered Projects

This example shows how to view all projects that are registered with crules.

```bash
# List all registered projects
crules list

# You should see output similar to:
# Registered projects (3):
#   1. /path/to/project1
#   2. /path/to/project2
#   3. /path/to/project3 (not found)
```

## Cleaning the Registry

This example demonstrates how to clean up the registry by removing entries for non-existent projects.

```bash
# Remove non-existent projects from the registry
crules clean

# You should see output confirming how many projects were removed
```

## Using Debug Mode

This example shows how to run commands with debug output for troubleshooting.

```bash
# Run any command with debug output
crules --debug init

# You'll see detailed output about what the command is doing
```

## Using Verbose Mode

This example demonstrates using verbose mode for more informative output.

```bash
# Run any command with verbose output
crules --verbose sync

# You'll see more information about the command's execution
```

## Complete Workflow Example

This example shows a complete workflow for managing Cursor rules across multiple projects.

```bash
# Step 1: Initialize first project
cd /path/to/project1
crules init

# Step 2: Initialize second project
cd /path/to/project2
crules init

# Step 3: Make changes to rules in first project
cd /path/to/project1
# Edit rules...

# Step 4: Merge changes from first project
crules merge

# Step 5: Verify changes were synced to second project
cd /path/to/project2
# The rules should now include the changes from project1

# Step 6: List all registered projects
crules list

# Step 7: If any projects no longer exist, clean the registry
crules clean
```

## Working with Custom Configuration

This example demonstrates using custom configuration with environment variables.

```bash
# Create .env file in your project with custom settings
echo "LOG_LEVEL=debug" > .env
echo "RULES_DIR_NAME=.my-rules" >> .env

# Now when you run crules, it will use your custom configuration
crules init

# The rules will be stored in .my-rules instead of .cursor/rules
# And logging will be at debug level
```

## Related Documentation

- [Commands Reference](../user-guide/commands.md) - Detailed information about all commands
- [Configuration Guide](../user-guide/configuration.md) - How to configure crules
- [Advanced Usage Examples](advanced-usage.md) - More complex usage patterns
