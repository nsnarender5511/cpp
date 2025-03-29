# Contributing Guide

> 🤝 **Guidelines for contributing to the crules project**

## Getting Started

Thank you for considering contributing to crules! This document outlines the process for contributing to the project and provides guidelines to help you get started.

## Development Environment Setup

1. **Fork the Repository**
   
   Start by forking the repository on GitHub.

2. **Clone Your Fork**
   
   ```bash
   git clone https://github.com/yourusername/crules.git
   cd crules
   ```

3. **Set Up Remote**
   
   ```bash
   git remote add upstream https://github.com/originalowner/crules.git
   ```

4. **Install Dependencies**
   
   ```bash
   go mod download
   ```

5. **Set Up Environment Variables**
   
   Copy the example environment file and modify as needed:
   
   ```bash
   cp .env.example .env
   ```

## Development Workflow

1. **Create a Branch**
   
   Always create a new branch for your changes:
   
   ```bash
   git checkout -b feature/your-feature-name
   ```
   
   Use a prefix that indicates the type of change:
   - `feature/` for new features
   - `fix/` for bug fixes
   - `docs/` for documentation
   - `refactor/` for refactoring existing code
   - `test/` for adding or updating tests

2. **Make Your Changes**
   
   Develop and test your changes locally.

3. **Follow Code Style**
   
   - Use Go's standard formatting with `gofmt`
   - Follow best practices from [Effective Go](https://golang.org/doc/effective_go)
   - Run linting tools to ensure code quality

4. **Write Tests**
   
   - Add tests for new functionality
   - Ensure existing tests pass with your changes
   - Run the full test suite before submitting your changes

5. **Commit Your Changes**
   
   Use clear and descriptive commit messages:
   
   ```bash
   git commit -m "feat: add support for XYZ"
   ```
   
   Follow [Conventional Commits](https://www.conventionalcommits.org/) format:
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation changes
   - `test:` for adding tests
   - `refactor:` for refactoring code
   - `chore:` for maintenance tasks

6. **Update Your Branch**
   
   Before submitting, update your branch with the latest changes from upstream:
   
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

7. **Push Your Changes**
   
   ```bash
   git push origin feature/your-feature-name
   ```

8. **Create a Pull Request**
   
   - Go to the GitHub repository page
   - Click "New Pull Request"
   - Select your branch to compare
   - Fill out the PR template with details about your changes

## Pull Request Guidelines

1. **Description**
   
   Provide a clear description of what your PR does. Link to related issues if applicable.

2. **Changes**
   
   Explain the changes you've made and why they're necessary.

3. **Testing**
   
   Describe how you tested your changes and what tests were added or updated.

4. **Screenshots**
   
   Include screenshots if your changes affect the UI or UX.

5. **Documentation**
   
   Update documentation if your changes affect user-facing functionality.

## Code Review Process

1. Maintainers will review your PR
2. Address any feedback or requested changes
3. Once approved, a maintainer will merge your PR
4. Your contribution will be part of the next release

## Development Guidelines

### Code Structure

- Follow the existing code organization
- Keep packages focused on a single responsibility
- Use interfaces to define clear contracts between components

### Error Handling

- Always check error returns
- Provide context in error messages
- Use consistent error handling patterns

### Logging

- Use the application's logging system
- Log at appropriate levels
- Include relevant context in log messages

### Testing

- Write unit tests for all new functionality
- Aim for high test coverage
- Create integration tests where appropriate

## Building and Testing

### Build the Application

```bash
go build -o crules cmd/main.go
```

### Run Tests

```bash
go test ./...
```

### Run with Verbose Output

```bash
./crules --verbose <command>
```

### Run with Debug Output

```bash
./crules --debug <command>
```

## Reporting Issues

If you find a bug or have a feature request:

1. Check if it's already reported in the issue tracker
2. Create a new issue with a clear title and description
3. Include steps to reproduce, expected behavior, and actual behavior
4. Add relevant system information and logs

## Community Guidelines

- Be respectful and inclusive
- Keep discussions focused on the project
- Help others when you can
- Follow the code of conduct

## License

By contributing to this project, you agree that your contributions will be licensed under the project's license.

## Questions?

If you have any questions about contributing, feel free to open an issue for discussion.

## Related Documentation

- [Code Structure](code-structure.md)
- [Testing Guide](testing.md)
