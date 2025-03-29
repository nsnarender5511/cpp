# Advanced Usage Examples

> 🚀 This guide covers advanced usage patterns and techniques for power users.

## Overview

These advanced examples demonstrate how to leverage crules for more complex scenarios, customization, and integration with other tools. They're designed for users who are already familiar with basic crules operations.

## Automating Rule Management

### Example 1: Git Hooks Integration

You can use Git hooks to automate rule synchronization between projects.

#### Implementation

1. Create a `post-merge` Git hook in your main rules repository:

```bash
#!/bin/bash
# File: ~/main-rules/.git/hooks/post-merge

# Check if crules is available
if ! command -v crules &> /dev/null; then
    echo "crules not found, skipping sync"
    exit 0
fi

# Run sync to all projects after pulling changes
crules sync --all

echo "✓ Successfully synced rules to all projects"
```

2. Make the hook executable:

```bash
chmod +x ~/main-rules/.git/hooks/post-merge
```

3. Create a `post-commit` hook in your project repositories:

```bash
#!/bin/bash
# File: ~/projects/my-project/.git/hooks/post-commit

# Check if there were changes to .cursor/rules
if git diff-tree -r --name-only HEAD@{1} HEAD | grep -q ".cursor/rules"; then
    # Check if crules is available
    if ! command -v crules &> /dev/null; then
        echo "crules not found, skipping merge"
        exit 0
    fi

    # Run merge to update main rules
    crules merge

    echo "✓ Successfully merged rule changes to main location"
fi
```

4. Make the hook executable:

```bash
chmod +x ~/projects/my-project/.git/hooks/post-commit
```

#### Benefits

- Rules are automatically synced when you pull updates to your main rules repository
- Changes to rules in projects are automatically merged back to the main location
- Ensures consistent rules across all projects without manual intervention

### Example 2: CI/CD Integration

You can integrate crules into your CI/CD pipeline to ensure all developers use the same rules.

#### Implementation

Add crules to your CI configuration (example for GitHub Actions):

```yaml
# File: .github/workflows/rules-sync.yml
name: Sync Cursor Rules

on:
  push:
    paths:
      - '.cursor/rules/**'
  schedule:
    - cron: '0 0 * * *' # Daily at midnight

jobs:
  sync-rules:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Install crules
        run: |
          curl -fsSL https://get.cursor.sh/rules | bash
      
      - name: Sync rules
        run: |
          crules init --main ./rules-main
          crules sync --all
```

#### Benefits

- Ensures rule consistency across team members
- Automates rule synchronization on a schedule
- Triggers updates when rules change

## Advanced Rule Creation

### Example 3: Dynamic Agent Selection

Create a meta-agent that selects the appropriate specialized agent based on context.

#### Implementation

Create a `dynamic-selector.mdc` file:

```markdown
# 🎭 Dynamic Selector

## 🎯 Role:
Intelligent agent selector that chooses the best agent for the current task

## 📝 Description:
This agent analyzes the context of the task and selects the most appropriate
specialized agent to handle it effectively.

## ✅ Capabilities:

### ✅ Context Analysis
Examines the task description, codebase, and user intent to determine
the best-suited agent.

### ✅ Agent Selection
Recommends the most appropriate agent from the available pool based on
task requirements.

### ✅ Explanation
Provides reasoning for agent selection to help users understand the choice.

## 📋 Instructions:
When presented with a task, I will:
1. Analyze the task requirements and context
2. Consider the strengths of available agents
3. Select and recommend the most appropriate agent
4. If multiple agents could be helpful, I will suggest a sequence of agents
5. When no single agent is clearly best, I will recommend the Technical Wizard

## 📊 Selection Logic:
- Code change requests → Implementer
- Bug reports → Fix Planner followed by Implementer
- New feature requests → Feature Planner followed by Implementer
- Code quality concerns → Code Reviewer
- Architecture questions → Architecture Planner
- Documentation needs → Documentation Agent
- API design → Architecture Planner
- Testing strategies → Test Planner
- General queries → Technical Wizard
```

#### Usage

When faced with a complex task, use the Dynamic Selector first:

```bash
crules agent select dynamic-selector
```

The agent will then help you choose the best specialized agent for your task.

### Example 4: Creating a Versioned Agent Family

Create a family of agents with different versions for specific frameworks or languages.

#### Implementation

1. Create a base agent template:

```markdown
# 🛠️ ${LANGUAGE} Implementer v${VERSION}

## 🎯 Role:
Specialized implementation agent for ${LANGUAGE} ${VERSION} projects

## 📝 Description:
This agent specializes in implementing code in ${LANGUAGE} ${VERSION},
following best practices and idiomatic patterns.

## ✅ Capabilities:

### ✅ ${LANGUAGE} Implementation
Creates idiomatic ${LANGUAGE} code following version ${VERSION} standards.

### ✅ Testing
Implements appropriate tests using ${TEST_FRAMEWORK}.

### ✅ Documentation
Generates code documentation following ${LANGUAGE} conventions.

## 📋 Instructions:
When implementing ${LANGUAGE} code, I will:
1. Follow ${LANGUAGE} ${VERSION} best practices
2. Use idiomatic patterns and features
3. Implement appropriate error handling
4. Add tests using ${TEST_FRAMEWORK}
5. Include necessary documentation
```

2. Create specific versions for different languages:

```bash
# Create TypeScript implementer
sed -e 's/\${LANGUAGE}/TypeScript/g' \
    -e 's/\${VERSION}/5.0/g' \
    -e 's/\${TEST_FRAMEWORK}/Jest/g' \
    base-template.mdc > typescript-implementer.mdc

# Create Python implementer
sed -e 's/\${LANGUAGE}/Python/g' \
    -e 's/\${VERSION}/3.11/g' \
    -e 's/\${TEST_FRAMEWORK}/pytest/g' \
    base-template.mdc > python-implementer.mdc

# Create Go implementer
sed -e 's/\${LANGUAGE}/Go/g' \
    -e 's/\${VERSION}/1.18/g' \
    -e 's/\${TEST_FRAMEWORK}/testing package/g' \
    base-template.mdc > go-implementer.mdc
```

3. Sync these to all projects:

```bash
crules sync --all
```

#### Benefits

- Language-specific guidance and patterns
- Consistent implementation across different languages
- Easy to update as language versions evolve

## Rule Management at Scale

### Example 5: Team-Wide Rule Distribution

For larger organizations, you can set up a centralized rule distribution system.

#### Implementation

1. Create a central rule repository:

```bash
mkdir -p ~/organization/cursor-rules
cd ~/organization/cursor-rules
git init
```

2. Set up a structure for different teams:

```bash
mkdir -p teams/frontend
mkdir -p teams/backend
mkdir -p teams/devops
mkdir -p shared
```

3. Create a script to sync rules based on team:

```bash
#!/bin/bash
# File: sync-team-rules.sh

# Get team name from first argument
TEAM=$1

if [ -z "$TEAM" ]; then
    echo "Error: Team name required"
    echo "Usage: ./sync-team-rules.sh <team-name>"
    exit 1
fi

# Check if team exists
if [ ! -d "teams/$TEAM" ]; then
    echo "Error: Team not found: $TEAM"
    exit 1
fi

# Create temporary directory for combined rules
TEMP_DIR=$(mktemp -d)

# Copy shared rules
cp -r shared/* $TEMP_DIR/

# Copy team-specific rules (overriding shared if needed)
cp -r teams/$TEAM/* $TEMP_DIR/

# Set as main rules location
crules init --main $TEMP_DIR

# Sync to current project
crules sync

# Clean up
rm -rf $TEMP_DIR

echo "✓ Successfully synced rules for team: $TEAM"
```

4. Make the script executable:

```bash
chmod +x sync-team-rules.sh
```

5. Configure users to run this script from their projects:

```bash
cd ~/projects/my-frontend-project
~/organization/cursor-rules/sync-team-rules.sh frontend
```

#### Benefits

- Central management of rules
- Team-specific customizations
- Shared base rules for consistency
- Controlled rule distribution

### Example 6: Rule Version Control Integration

Integrate rule versioning with semantic versioning for better change management.

#### Implementation

1. Create a version tracking file:

```json
{
  "version": "1.2.3",
  "lastUpdated": "2023-06-15",
  "changes": [
    {
      "version": "1.2.3",
      "date": "2023-06-15",
      "changes": [
        "Added React performance agent",
        "Updated TypeScript implementer to support v5.0",
        "Fixed Go code reviewer linting rules"
      ]
    },
    {
      "version": "1.2.2",
      "date": "2023-06-01",
      "changes": [
        "Added microservices architecture planner",
        "Updated documentation agent instructions"
      ]
    }
  ]
}
```

2. Create a version update script:

```bash
#!/bin/bash
# File: bump-version.sh

# Check for version type argument
VERSION_TYPE=$1

if [ -z "$VERSION_TYPE" ]; then
    echo "Error: Version type required"
    echo "Usage: ./bump-version.sh <major|minor|patch>"
    exit 1
fi

# Read current version
CURRENT_VERSION=$(jq -r '.version' version.json)
MAJOR=$(echo $CURRENT_VERSION | cut -d. -f1)
MINOR=$(echo $CURRENT_VERSION | cut -d. -f2)
PATCH=$(echo $CURRENT_VERSION | cut -d. -f3)

# Bump version based on type
if [ "$VERSION_TYPE" = "major" ]; then
    MAJOR=$((MAJOR + 1))
    MINOR=0
    PATCH=0
elif [ "$VERSION_TYPE" = "minor" ]; then
    MINOR=$((MINOR + 1))
    PATCH=0
elif [ "$VERSION_TYPE" = "patch" ]; then
    PATCH=$((PATCH + 1))
else
    echo "Error: Invalid version type: $VERSION_TYPE"
    echo "Usage: ./bump-version.sh <major|minor|patch>"
    exit 1
fi

NEW_VERSION="$MAJOR.$MINOR.$PATCH"
TODAY=$(date +%Y-%m-%d)

# Update version in file
TMP_FILE=$(mktemp)
jq --arg v "$NEW_VERSION" --arg d "$TODAY" \
   '.version = $v | .lastUpdated = $d' version.json > $TMP_FILE
mv $TMP_FILE version.json

# Add new change entry
echo "Enter comma-separated changes for version $NEW_VERSION:"
read CHANGES_INPUT

# Convert comma-separated input to JSON array
CHANGES_JSON="["
IFS=',' read -ra CHANGES_ARRAY <<< "$CHANGES_INPUT"
for i in "${!CHANGES_ARRAY[@]}"; do
    CHANGES_JSON+="\"${CHANGES_ARRAY[$i]}\""
    if [ $i -lt $((${#CHANGES_ARRAY[@]}-1)) ]; then
        CHANGES_JSON+=","
    fi
done
CHANGES_JSON+="]"

# Add new entry to changes array
TMP_FILE=$(mktemp)
jq --arg v "$NEW_VERSION" --arg d "$TODAY" --argjson c "$CHANGES_JSON" \
   '.changes = [{"version": $v, "date": $d, "changes": $c}] + .changes' version.json > $TMP_FILE
mv $TMP_FILE version.json

echo "✓ Successfully bumped version to $NEW_VERSION"
```

3. Make the script executable:

```bash
chmod +x bump-version.sh
```

4. Use for versioning:

```bash
# After making significant changes
./bump-version.sh minor

# After small fixes
./bump-version.sh patch

# After major changes
./bump-version.sh major
```

#### Benefits

- Clear versioning of rule changes
- Changelog for rule updates
- Easier tracking of rule evolution
- Better communication about rule changes to team members

## API Integration

### Example 7: Integration with External APIs

Integrate crules with external AI APIs for specialized capabilities.

#### Implementation

1. Create a wrapper script for API integration:

```python
#!/usr/bin/env python3
# File: api-agent.py

import os
import sys
import json
import requests
import argparse

# Configure argument parser
parser = argparse.ArgumentParser(description='Integrate external APIs with crules')
parser.add_argument('--api', choices=['openai', 'anthropic', 'huggingface'], required=True, help='API to use')
parser.add_argument('--task', choices=['code-gen', 'summarize', 'analyze'], required=True, help='Task to perform')
parser.add_argument('--input', required=True, help='Input file or string')
parser.add_argument('--output', help='Output file (if not specified, prints to stdout)')
args = parser.parse_args()

# Configuration
API_KEYS = {
    'openai': os.environ.get('OPENAI_API_KEY'),
    'anthropic': os.environ.get('ANTHROPIC_API_KEY'),
    'huggingface': os.environ.get('HUGGINGFACE_API_KEY')
}

# Check API key
if not API_KEYS[args.api]:
    print(f"Error: {args.api.upper()}_API_KEY environment variable not set")
    sys.exit(1)

# Read input
input_content = ""
if os.path.exists(args.input):
    with open(args.input, 'r') as f:
        input_content = f.read()
else:
    input_content = args.input

# Prepare API request based on provider
if args.api == 'openai':
    url = "https://api.openai.com/v1/chat/completions"
    headers = {
        "Authorization": f"Bearer {API_KEYS['openai']}",
        "Content-Type": "application/json"
    }
    prompt_map = {
        'code-gen': "Generate code based on the following specification:",
        'summarize': "Summarize the following content:",
        'analyze': "Analyze the following code and provide feedback:"
    }
    data = {
        "model": "gpt-4",
        "messages": [
            {"role": "system", "content": "You are a helpful assistant that specializes in code."},
            {"role": "user", "content": f"{prompt_map[args.task]}\n\n{input_content}"}
        ]
    }
    response = requests.post(url, headers=headers, json=data)
    result = response.json()["choices"][0]["message"]["content"]

elif args.api == 'anthropic':
    url = "https://api.anthropic.com/v1/messages"
    headers = {
        "x-api-key": API_KEYS['anthropic'],
        "Content-Type": "application/json"
    }
    prompt_map = {
        'code-gen': "Generate code based on the following specification:",
        'summarize': "Summarize the following content:",
        'analyze': "Analyze the following code and provide feedback:"
    }
    data = {
        "model": "claude-3-opus-20240229",
        "max_tokens": 4000,
        "messages": [
            {"role": "user", "content": f"{prompt_map[args.task]}\n\n{input_content}"}
        ]
    }
    response = requests.post(url, headers=headers, json=data)
    result = response.json()["content"][0]["text"]

elif args.api == 'huggingface':
    url = "https://api-inference.huggingface.co/models/bigcode/starcoder2-15b"
    headers = {
        "Authorization": f"Bearer {API_KEYS['huggingface']}",
        "Content-Type": "application/json"
    }
    prompt_map = {
        'code-gen': "// Generate code based on the following specification:\n\n",
        'summarize': "// Summarize the following content:\n\n",
        'analyze': "// Analyze the following code and provide feedback:\n\n"
    }
    data = {
        "inputs": f"{prompt_map[args.task]}{input_content}"
    }
    response = requests.post(url, headers=headers, json=data)
    result = response.json()[0]["generated_text"]

# Output result
if args.output:
    with open(args.output, 'w') as f:
        f.write(result)
    print(f"✓ Result written to {args.output}")
else:
    print(result)
```

2. Make the script executable:

```bash
chmod +x api-agent.py
```

3. Use the script with crules:

```bash
# Generate code based on a specification
./api-agent.py --api openai --task code-gen --input spec.md --output implementation.js

# Analyze existing code
./api-agent.py --api anthropic --task analyze --input src/complex-function.js
```

#### Benefits

- Leverage specialized external AI capabilities
- Combine crules with other AI services
- Enhanced code generation and analysis
- Flexibility to choose the best API for specific tasks

## Navigation

- Previous: [Basic Usage](./basic-usage.md)
- Next: [Agent Workflows](./agent-workflows.md)
- Up: [Examples](./README.md)
- Home: [Documentation Home](../README.md)
