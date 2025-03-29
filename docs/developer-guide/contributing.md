# Contributing to crules

> 🤝 This guide explains how to contribute to the crules project, including development setup, coding standards, and the contribution workflow.

## Getting Started

Thank you for considering contributing to crules! This guide will help you get started with development and explain our processes.

### Prerequisites

To contribute to crules, you'll need:

1. Go 1.20 or later
2. Git
3. A GitHub account
4. Basic knowledge of Go programming

### Development Environment Setup

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/your-username/crules.git
cd crules
```

3. Add the upstream repository as a remote:

```bash
git remote add upstream https://github.com/original-owner/crules.git
```

4. Install dependencies:

```bash
go mod download
```

5. Build the tool:

```bash
go build -o crules ./cmd/crules
```

6. Run tests to ensure everything is working:

```bash
go test ./...
```

## Development Workflow

### Creating a New Feature

1. **Create a branch** from the `main` branch:

```bash
git checkout main
git pull upstream main
git checkout -b feature/your-feature-name
```

2. **Implement your changes**, following the coding standards
3. **Write tests** for your changes
4. **Update documentation** if needed
5. **Commit your changes** with clear, descriptive commit messages
6. **Push your branch** to your fork:

```bash
git push -u origin feature/your-feature-name
```

7. **Create a Pull Request** against the `main` branch of the upstream repository

### Fixing a Bug

1. **Create a branch** from the `main` branch:

```bash
git checkout main
git pull upstream main
git checkout -b fix/bug-description
```

2. **Implement your fix**, following the coding standards
3. **Write tests** to verify the fix and prevent regression
4. **Commit your changes** with clear, descriptive commit messages
5. **Push your branch** to your fork:

```bash
git push -u origin fix/bug-description
```

6. **Create a Pull Request** against the `main` branch of the upstream repository

## Coding Standards

### Go Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Add comments to exported functions, types, and variables

### Package Structure

- Place related functionality in the same package
- Use internal packages for code that shouldn't be imported by other projects
- Keep package names short and descriptive

### Error Handling

- Return errors rather than using panic
- Use descriptive error messages
- Wrap errors with context using `fmt.Errorf("doing something: %w", err)`

### Testing

- Write unit tests for all new functionality
- Aim for high test coverage, especially for critical paths
- Use table-driven tests where appropriate
- Use mocks for external dependencies

## Pull Request Process

1. **Update your branch** with the latest changes from the main branch
2. **Ensure all tests pass** on your branch
3. **Submit your pull request** with a clear description of the changes
4. **Address review comments** and make requested changes
5. **Update your pull request** with the changes

Your pull request will be reviewed by maintainers, who may request changes or suggest improvements. Once your changes are approved, they will be merged into the main branch.

## Commit Message Guidelines

Follow these guidelines for commit messages:

1. **Start with a type prefix**:
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation changes
   - `style:` for formatting changes
   - `refactor:` for code refactoring
   - `test:` for adding or updating tests
   - `chore:` for maintenance tasks

2. **Keep the first line short** (under 72 characters)
3. **Use the imperative mood** ("Add feature" not "Added feature")
4. **Describe what was changed and why**, not how

Example:
```
feat: add interactive agent selection

Add an interactive terminal UI for selecting agents using the bubbletea library.
This improves user experience by providing a visual interface for browsing
available agents.
```

## Documentation

When contributing new features or making significant changes, please update the relevant documentation:

1. **Update the README.md** if needed
2. **Update or create documentation** in the `docs/` directory
3. **Add examples** if appropriate
4. **Update command help text** for CLI changes

## Release Process

The release process is managed by maintainers, who will:

1. Decide when to create a new release
2. Update version numbers
3. Create a release tag
4. Build and publish binaries

## Community and Communication

- **GitHub Issues**: Use for bug reports, feature requests, and questions
- **Pull Requests**: Use for contributing code changes
- **Discussions**: Use for general discussions about the project

## Code of Conduct

Please follow our [Code of Conduct](../CODE_OF_CONDUCT.md) in all your interactions with the project.

## License

By contributing to crules, you agree that your contributions will be licensed under the project's license.

## See Also

- [Architecture](./architecture.md)
- [Code Structure](./code-structure.md)
- [Extending Agents](./extending-agents.md)
