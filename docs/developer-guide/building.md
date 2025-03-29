# Building from Source

> 🔨 This guide explains how to build crules from source, set up a development environment, and contribute to the project.

## Prerequisites

Before building crules, ensure you have the following prerequisites installed:

- **Go** (version 1.19 or later)
- **Git**
- **Make** (optional, but recommended)

### Installing Go

If you don't have Go installed, follow these instructions:

- **macOS**:
  ```bash
  brew install go
  ```

- **Linux**:
  ```bash
  # Ubuntu/Debian
  sudo apt-get update
  sudo apt-get install golang-go
  
  # Fedora
  sudo dnf install golang
  ```

- **Windows**:
  Download and run the installer from [golang.org/dl](https://golang.org/dl/)

### Verifying Go Installation

Verify that Go is installed correctly:

```bash
go version
```

Ensure you're using Go 1.19 or later.

## Getting the Source Code

Clone the crules repository:

```bash
git clone https://github.com/org/crules.git
cd crules
```

## Project Structure

The project follows a standard Go project layout:

```
crules/
├── cmd/                 # Command-line application entry points
│   └── crules/         # Main application
├── internal/            # Private application code
│   ├── agent/          # Agent system implementation
│   ├── core/           # Core functionality
│   ├── git/            # Git operations
│   ├── ui/             # Terminal UI components
│   ├── utils/          # Utility functions
│   └── version/        # Version information
├── pkg/                 # Libraries that can be used by external applications
├── docs/                # Documentation
├── test/                # Test data and scripts
└── scripts/             # Build and maintenance scripts
```

## Building the Application

### Using Make

If you have Make installed, you can use the following commands:

```bash
# Build the application
make build

# Run tests
make test

# Clean build artifacts
make clean

# Install locally
make install
```

### Using Go Commands

If you don't have Make installed, you can use Go commands directly:

```bash
# Build the application
go build -o crules ./cmd/crules

# Run tests
go test ./...

# Install locally
go install ./cmd/crules
```

## Development Environment Setup

### Go Modules

crules uses Go modules for dependency management. To initialize modules:

```bash
# Initialize dependencies (this is usually not needed as go.mod exists)
go mod tidy
```

### IDE Setup

#### VSCode

1. Install the Go extension for VSCode
2. Open the project folder in VSCode
3. Install recommended Go tools when prompted
4. Configure settings.json:
   ```json
   {
     "go.useLanguageServer": true,
     "go.lintTool": "golangci-lint",
     "go.formatTool": "goimports",
     "go.testFlags": ["-v"]
   }
   ```

#### GoLand/IntelliJ

1. Open the project in GoLand
2. Ensure the GOROOT is set correctly in Settings > Languages & Frameworks > Go > GOROOT
3. Enable Go modules integration in Settings > Languages & Frameworks > Go > Go Modules

### Setting Up Pre-commit Hooks

To ensure code quality, set up pre-commit hooks:

```bash
# Copy the hooks to your .git directory
cp scripts/pre-commit .git/hooks/
chmod +x .git/hooks/pre-commit
```

## Building for Different Platforms

### Cross-Compilation

Go supports cross-compilation for different platforms:

```bash
# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o crules-darwin-amd64 ./cmd/crules

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o crules-linux-amd64 ./cmd/crules

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o crules-windows-amd64.exe ./cmd/crules
```

### Building Release Binaries

To build optimized release binaries:

```bash
# Optimized build with stripped symbols
go build -ldflags="-s -w" -o crules ./cmd/crules
```

## Running and Testing

### Running Locally

After building, you can run the application:

```bash
# Run with the built binary
./crules --version

# Or use go run
go run ./cmd/crules --version
```

### Running Tests

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Manual Testing

For manual testing:

1. Build the application
2. Create a test directory
3. Run initialization in the test directory
4. Test various commands

```bash
mkdir -p test/testproject
cd test/testproject
../../crules init
../../crules agent list
```

## Debugging

### Using Delve

[Delve](https://github.com/go-delve/delve) is a powerful debugger for Go:

```bash
# Install Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug the application
dlv debug ./cmd/crules

# Set a breakpoint
(dlv) break internal/core/sync.go:100

# Run to breakpoint
(dlv) continue
```

### Verbose Logging

Enable verbose and debug logging for troubleshooting:

```bash
./crules --verbose sync
./crules --debug agent select
```

## Dependencies

crules has the following key dependencies:

| Dependency | Purpose |
|------------|---------|
| github.com/jedib0t/go-pretty/v6 | Terminal tables and text formatting |
| github.com/manifoldco/promptui | Interactive terminal prompts |
| github.com/spf13/cobra | Command-line interface |
| github.com/fatih/color | Terminal color output |
| github.com/PuerkitoBio/goquery | HTML parsing for web imports |

To update dependencies:

```bash
go get -u ./...
go mod tidy
```

## Common Issues

### Build Errors

If you encounter build errors:

1. Ensure your Go version is up to date
2. Run `go mod tidy` to refresh dependencies
3. Check for missing dependencies with `go mod verify`

### Permission Issues

If you encounter permission issues when running or building:

```bash
# Make sure the build directory is writable
chmod -R u+w .

# Make sure the binary is executable
chmod +x crules
```

### Module Not Found Errors

If Go cannot find a module:

```bash
# Verify modules
go mod verify

# Clear module cache if needed
go clean -modcache
```

## Creating a Release

To create a release build:

1. Update version information in `internal/version/version.go`
2. Update CHANGELOG.md with the latest changes
3. Build with release flags
4. Create a release tag

```bash
# Update version
git tag v0.1.0
git push origin v0.1.0

# Build release binaries
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o crules-darwin-amd64 ./cmd/crules
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o crules-linux-amd64 ./cmd/crules
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o crules-windows-amd64.exe ./cmd/crules
```

## Contributing

Follow these steps to contribute to the project:

1. Fork the repository
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes
4. Run tests:
   ```bash
   go test ./...
   ```
5. Commit your changes using conventional commit messages:
   ```bash
   git commit -m "feat: add new feature"
   ```
6. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
7. Create a pull request

## See Also

- [Code Structure](./code-structure.md) for details on the codebase organization
- [Testing](./testing.md) for information on testing the project
- [Contributing](./contributing.md) for contribution guidelines

## Navigation

- Previous: [Testing](./testing.md)
- Next: [Code Structure](./code-structure.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md) 