# Commands Reference

> 🖥️ **Detailed descriptions of all available crules commands**

## Overview

crules provides several commands to manage Cursor rules across multiple projects. This reference guide describes each command, its purpose, usage, and examples.

## Basic Usage

```bash
crules [OPTIONS] <command>
```

Where `OPTIONS` include:
- `--verbose`: Show informational messages on console
- `--debug`: Show debug messages on console (implies verbose)

## Available Commands

### init

Initializes the current directory with rules from the main location.

```bash
crules init
```

**Purpose**: Use this command when you want to set up a new project with Cursor rules from your main location. This command will:
1. Create the rules directory in the current project
2. Copy rules from the main location to the project
3. Register the project in the registry

**Example**:
```bash
cd /path/to/your/project
crules init
```

**Exit Codes**:
- 0: Success
- 10: Init error

---

### merge

Merges current rules to the main location and syncs to all registered locations.

```bash
crules merge
```

**Purpose**: Use this command when you've made changes to rules in the current project and want to merge them back to the main location and sync those changes to all other registered projects. This command will:
1. Copy rules from the current project to the main location
2. Sync the updated rules to all other registered projects

**Example**:
```bash
# After making changes to rules
crules merge
```

**Exit Codes**:
- 0: Success
- 11: Merge error

---

### sync

Forces sync from main location to the current directory, overwriting local changes.

```bash
crules sync
```

**Purpose**: Use this command when you want to update the current project with the latest rules from the main location, discarding any local changes. This command will:
1. Copy rules from the main location to the current project, overwriting any existing rules

**Example**:
```bash
# To discard local changes and get latest rules
crules sync
```

**Exit Codes**:
- 0: Success
- 12: Sync error

---

### list

Displays all registered projects.

```bash
crules list
```

**Purpose**: Use this command to see a list of all projects that have been registered with crules. This command will:
1. Display all registered projects
2. Indicate which projects exist and which are missing

**Example**:
```bash
crules list
```

**Output Example**:
```
Registered projects (3):
  1. /path/to/project1
  2. /path/to/project2
  3. /path/to/project3 (not found)

1 project(s) could not be found. Run 'crules clean' to remove them.
```

**Exit Codes**:
- 0: Success
- 13: List error

---

### clean

Removes non-existent projects from the registry.

```bash
crules clean
```

**Purpose**: Use this command to clean up the registry by removing entries for projects that no longer exist on disk. This command will:
1. Check each registered project to see if it exists
2. Remove entries for projects that don't exist
3. Report how many projects were removed

**Example**:
```bash
crules clean
```

**Output Example**:
```
Successfully removed 1 non-existent project(s) from registry.
```

**Exit Codes**:
- 0: Success
- 14: Clean error

## Command Flow Examples

### Typical Workflow

1. Initialize a new project:
   ```bash
   cd /path/to/new/project
   crules init
   ```

2. After making changes to rules in the project:
   ```bash
   crules merge
   ```

3. In another project, pull in the latest rules:
   ```bash
   cd /path/to/another/project
   crules sync
   ```

4. View all registered projects:
   ```bash
   crules list
   ```

5. Clean up registry if needed:
   ```bash
   crules clean
   ```

## Related Documentation

- [Configuration](configuration.md) - Learn about configuring crules
- [Basic Usage Examples](../examples/basic-usage.md) - See more examples
