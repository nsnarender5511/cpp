---
version: v0.1.0
last_updated: 2023-03-29
applies_to: cursor++ v0.1.0+
---

# Contributing to cursor++

> 🤝 Guidelines for contributing to cursor++.

## Introduction

Thank you for considering contributing to cursor++! This document provides guidelines and instructions for contributing to the project. By following these guidelines, you help maintain the quality and consistency of the codebase.

## Code of Conduct

Please read and follow our [Code of Conduct](./code-of-conduct.md). We expect all contributors to adhere to these guidelines to ensure a positive and inclusive environment for everyone.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git

### Setting Up the Development Environment

1. Fork the repository on GitHub
2. Clone your fork to your local machine:
   ```bash
   git clone https://github.com/YOUR-USERNAME/cursor++.git
   cd cursor++
   ```
3. Add the original repository as an upstream remote:
   ```bash
   git remote add upstream https://github.com/nsnarender5511/vibe.git
   ```
4. Install dependencies:
   ```bash
   go mod tidy
   ```

## Development Workflow

### Branching Strategy

- `main`: The primary branch containing stable code
- `develop`: The development branch for ongoing work
- Feature branches: Create from `develop` with the format `feature/your-feature-name`
- Bugfix branches: Create from `develop` with the format `bugfix/issue-description`

### Creating a Feature Branch

```bash
# Ensure you're on the develop branch
git checkout develop

# Pull the latest changes
git pull upstream develop

# Create a new feature branch
git checkout -b feature/your-feature-name
```

### Making Changes

When making changes:

1. Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) for style guidance
2. Write tests for new functionality
3. Ensure your code passes all existing tests
4. Update documentation as needed

### Running Tests

Run tests with:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

### Committing Changes

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types include:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code changes that neither fix bugs nor add features
- `test`: Adding or modifying tests
- `chore`: Changes to the build process or tools

Example:
```
feat(agent): add support for selection by index
```

### Creating a Pull Request

1. Push your branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
2. Open a pull request against the `develop` branch of the original repository
3. Fill out the pull request template with all required information
4. Wait for the CI checks to pass
5. Address any feedback from reviewers

## Coding Standards

### Go Formatting

- Use `gofmt` or `goimports` to format your code
- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines

### Code Organization

- Keep packages focused and cohesive
- Use descriptive names for packages, types, and functions
- Organize imports in groups: standard library, external dependencies, internal packages

### Documentation

- Document all exported types, functions, and methods
- Add examples where appropriate
- Keep documentation up-to-date when changing code

### Documentation Contributions

#### Documentation Structure
- All documentation is located in the `docs/` directory
- User-facing documentation is in `docs/user-guide/`
- Developer documentation is in `docs/developer-guide/`
- API reference is in `docs/api-reference/`
- Examples are in `docs/examples/`

#### Documentation Standards
- Use Markdown for all documentation files
- Include version indicators at the top of each documentation file:
  ```markdown
  ---
  version: v0.1.0
  last_updated: YYYY-MM-DD
  applies_to: cursor++ v0.1.0+
  ---
  ```
- Include a table of contents for files longer than 100 lines
- Use relative links for cross-references within documentation
- Provide code examples for technical concepts
- Add visual elements (diagrams, screenshots) for complex ideas
- Use consistent formatting across all documentation files
- Follow the structure of existing documentation for new files

#### Documentation Testing
- Verify all links are valid before submitting documentation changes
- Ensure code examples are up-to-date with the current API
- Test documentation rendering in GitHub to ensure proper formatting
- Verify that all referenced files and images exist
- Check for spelling and grammatical errors

#### Documentation Workflow
1. Identify documentation gaps or improvements
2. Create a branch for your documentation changes
3. Make your changes following the documentation standards
4. Test your changes as described above
5. Submit a pull request with a clear description of your changes

### Error Handling

- Always check errors
- Use descriptive error messages
- Wrap errors with additional context using `fmt.Errorf("failed to do something: %w", err)`
- Define custom error types for specific error conditions

### Testing

- Write unit tests for all new functions and methods
- Aim for high test coverage (at least 80%)
- Use table-driven tests where appropriate
- Mock dependencies to test edge cases and error conditions

## Project Structure

The project follows a standard Go layout:

```
cursor++/
├── cmd/           # Command-line applications
├── internal/      # Private application code
│   ├── agent/     # Agent-related code
│   ├── core/      # Core functionality
│   ├── git/       # Git integration
│   ├── ui/        # Terminal UI components
│   ├── utils/     # Utility functions
│   └── version/   # Version-related code
├── docs/          # Documentation
├── test/          # Additional test files
└── scripts/       # Build and utility scripts
```

## Reporting Issues

When reporting issues:

1. Use the issue template if available
2. Provide a clear and descriptive title
3. Include steps to reproduce the issue
4. Include the expected and actual behavior
5. Include version information and environment details
6. Attach logs or screenshots if relevant

## Feature Requests

When requesting features:

1. Use the feature request template if available
2. Describe the feature and its benefits
3. Provide use cases
4. Consider edge cases and potential issues
5. Suggest an implementation approach if possible

## Release Process

The project follows [Semantic Versioning](https://semver.org/). The release process includes:

1. Merging feature branches into `develop`
2. Creating a release branch from `develop`
3. Testing and fixing issues on the release branch
4. Merging the release branch into `main`
5. Tagging the release with a version number
6. Building and publishing artifacts

## License

By contributing to this project, you agree that your contributions will be licensed under the project's [license](../../LICENSE).

## Contact

If you have questions or need help, you can:

- Open an issue on GitHub
- Join our community chat
- Email the maintainers

Thank you for contributing to cursor++! Your help makes the project better for everyone.

## See Also

- [Code of Conduct](./code-of-conduct.md) for community guidelines
- [Testing](./testing.md) for information on testing the project
- [Building from Source](./building.md) for instructions on building the project

## Navigation

- Previous: [Extending Agents](./extending-agents.md)
- Next: [Code of Conduct](./code-of-conduct.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md)
