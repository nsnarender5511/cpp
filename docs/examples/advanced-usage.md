# Advanced Documentation Techniques

> 🔍 This guide provides advanced techniques for working with the Document Syncer agent in cursor++.

## Overview

This guide covers advanced features and techniques for using the Document Syncer agent effectively. These techniques build on the basics covered in the [Basic Usage Guide](./basic-usage.md).

## Advanced Context Analysis

The Document Syncer agent can perform sophisticated analysis of your codebase and documentation. By providing clear context, you can get more precise results.

### Providing Project Context

When working on large codebases, provide the agent with project context:

```
@doc-syncer.mdc I'm working on a Go project with multiple packages. The main documentation is in the docs/ directory, and we have API documentation in docs/api-reference/. Please help me synchronize the API docs with our current implementation.
```

This gives the agent important information about:
- The programming language (Go)
- Project structure (multiple packages)
- Documentation location (docs/ directory)
- Specific task (synchronize API documentation)

### Path-Based Context

Direct the agent to specific parts of your codebase:

```
@doc-syncer.mdc Check for inconsistencies between the code in internal/agent/ and the documentation in docs/user-guide/agents.md
```

This narrows the scope to make the analysis more targeted and efficient.

## Customizing Synchronization Behavior

### Selective Updates

You can instruct the agent to focus on specific aspects of the documentation:

```
@doc-syncer.mdc Update only the command syntax in all documentation files, keeping all other content unchanged.
```

### Handling Conflicts

When there are conflicts between code and documentation, specify how to resolve them:

```
@doc-syncer.mdc When synchronizing, if there's a conflict between the code and documentation, prioritize the code implementation and update the docs accordingly.
```

## Cross-Reference Maintenance

One of the Document Syncer agent's strengths is maintaining cross-references between documentation files.

### Link Verification

Ask the agent to verify that all links in your documentation are valid:

```
@doc-syncer.mdc Verify that all internal links in the documentation are working properly.
```

### Updating Cross-References

When reorganizing documentation, the agent can update all references:

```
@doc-syncer.mdc I've moved the installation guide from docs/installation.md to docs/user-guide/installation.md. Please update all references to this file throughout the documentation.
```

## Multi-Stage Documentation Tasks

Complex documentation tasks can be broken down into stages:

### Stage 1: Analysis

```
@doc-syncer.mdc Analyze the inconsistencies between our code and documentation without making any changes yet.
```

### Stage 2: Planning

```
@doc-syncer.mdc Based on your analysis, create a synchronization plan to update the documentation.
```

### Stage 3: Review

```
@doc-syncer.mdc Show me the specific changes you'll make to each documentation file.
```

### Stage 4: Implementation

```
@doc-syncer.mdc Proceed with implementing the proposed changes.
```

### Stage 5: Verification

```
@doc-syncer.mdc Verify that all documentation is now consistent with the code.
```

## Documentation Format Conversion

The Document Syncer agent can help convert documentation between formats:

```
@doc-syncer.mdc Convert our README.md from Markdown to AsciiDoc format, maintaining all content and structure.
```

## Incorporating Code Examples

Keep code examples in documentation up-to-date:

```
@doc-syncer.mdc Update all code examples in the documentation to match our current implementation.
```

## Integration with Version Control

Reference specific commits or branches:

```
@doc-syncer.mdc Update the documentation to match the code in the 'feature/new-api' branch.
```

## Documentation Structure Analysis

Get insights into your documentation structure:

```
@doc-syncer.mdc Analyze our documentation structure and suggest improvements for better organization.
```

## Interactive Advanced Examples

Let's walk through some advanced scenarios.

### Scenario 1: Major Version Documentation Update

When releasing a new major version, documentation often needs comprehensive updates:

```
User: @doc-syncer.mdc We're preparing to release version 2.0 of our API. I need to update all the documentation to reflect the breaking changes.

Document Syncer: I'll help you update your documentation for the 2.0 release. First, let me analyze what needs to be updated.

[After analysis]

Document Syncer: I've analyzed your codebase and documentation. Here's what needs to be updated for version 2.0:

1. API Endpoints:
   - 3 endpoints have changed URL patterns
   - 5 endpoints have modified request/response formats
   - 2 new endpoints need documentation

2. Configuration Options:
   - 4 configuration parameters are deprecated
   - 7 new configuration options need documentation

3. Example Code:
   - All examples need updating to use new syntax

Would you like me to create a comprehensive update plan, or should we tackle these changes one section at a time?

User: Let's create a comprehensive plan first.

Document Syncer: Here's the comprehensive documentation update plan for version 2.0:

[Details of the update plan]

User: The plan looks good. Let's start with updating the API endpoint documentation first.

Document Syncer: I'll update the API endpoint documentation. Here are the specific changes I'm making:

[Details of changes to API documentation]

Would you like me to apply these changes?

User: Yes, please apply them.

Document Syncer: I've updated all API endpoint documentation to match version a2.0. Let's move on to the configuration options next.
```

### Scenario 2: Documentation Restructuring

When reorganizing your documentation structure:

```
User: @doc-syncer.mdc We're reorganizing our documentation. We want to move from a flat structure to a hierarchical one with user guide, API docs, and developer guide sections.

Document Syncer: I'll help you restructure your documentation. Let me first understand your current structure.

[After analysis]

Document Syncer: I've analyzed your current documentation structure. You have 15 Markdown files in a flat directory structure. Based on your request for a hierarchical organization, I propose the following structure:

1. docs/
   - README.md (overview)
   - documentation-map.md (new file to help navigation)
   
2. docs/user-guide/
   - getting-started.md (from current intro.md)
   - installation.md (from current install.md)
   - configuration.md (from current config.md)
   - commands.md (from current commands.md)
   
3. docs/api-reference/
   - core-api.md (from current api.md)
   - endpoints.md (from current endpoints.md)
   
4. docs/developer-guide/
   - contributing.md (from current contributing.md)
   - architecture.md (from current architecture.md)
   - testing.md (new file from testing sections)

Does this structure meet your requirements?

User: That looks good, but let's also add a troubleshooting guide to the user guide section.

Document Syncer: I'll add a troubleshooting guide to the user guide section. Here's the updated structure:

[Updated structure with troubleshooting.md added]

User: Perfect! Please implement this restructuring.

Document Syncer: I'll implement the documentation restructuring now. This will involve:
1. Creating the new directory structure
2. Moving and renaming files
3. Updating all internal links and references
4. Creating the new documentation-map.md file
5. Adding a new troubleshooting.md file

I'll keep all content intact while updating the structure. Would you like me to proceed?

User: Yes, please proceed.

Document Syncer: I've completed the documentation restructuring. All files have been organized into the hierarchical structure, and all internal links have been updated to point to the new locations.
```

## Performance Optimization Tips

For the best results when working with the Document Syncer agent:

1. **Be Specific**: Clearly define the scope of your documentation task
2. **Provide Context**: Share information about your project structure
3. **Use Staged Approaches**: Break complex tasks into analysis, planning, and implementation steps
4. **Review Carefully**: Always review proposed changes before applying them
5. **Maintain Consistency**: Use consistent terminology and structure across documentation

## Next Steps

As you become more proficient with the Document Syncer agent, consider:

- Creating a dedicated documentation maintenance workflow in your development process
- Implementing regular documentation reviews using the agent
- Building custom documentation templates for your projects

## Navigation

- Previous: [Agent Workflows](./agent-workflows.md)
- Up: [Examples](../README.md#examples)
- Home: [Documentation Home](../README.md)
