# Testing Guide

> 🧪 **Testing strategies and practices for the crules project**

## Overview

This guide outlines the testing approach for the crules project, including test organization, writing effective tests, and running the test suite.

## Testing Philosophy

The crules project follows these testing principles:

1. **Test-driven development** is encouraged for new features
2. **High test coverage** is a goal, especially for core functionality
3. **Maintainable tests** are as important as maintainable code
4. **Integration tests** ensure components work together correctly
5. **Mock external dependencies** to enable reliable unit testing

## Test Types

### Unit Tests

Unit tests focus on testing individual functions, methods, or classes in isolation.

**Naming Convention**:
- Files: `*_test.go` in the same package as the code being tested
- Functions: `TestFunctionName` for testing the function named `FunctionName`

**Example**:
```go
// sync_test.go
package core

import "testing"

func TestSync(t *testing.T) {
    // Test setup
    manager, err := NewSyncManager()
    if err != nil {
        t.Fatalf("Failed to create sync manager: %v", err)
    }
    
    // Test execution
    err = manager.Sync()
    
    // Assertions
    if err != nil {
        t.Errorf("Sync failed: %v", err)
    }
    
    // More assertions...
}
```

### Integration Tests

Integration tests verify that components work together correctly.

**Naming Convention**:
- Directory: `integration` subdirectory
- Files: `*_integration_test.go`
- Functions: `TestIntegrationScenarioName`

**Example**:
```go
// integration/sync_integration_test.go
package integration

import "testing"

func TestIntegrationFullSyncFlow(t *testing.T) {
    // Test the full flow of initialization, sync, and merge
    
    // Setup test environment
    
    // Run init command
    
    // Verify results
    
    // Run sync command
    
    // Verify results
    
    // Run merge command
    
    // Verify final state
}
```

### End-to-End Tests

End-to-end tests verify the entire application works as expected from a user's perspective.

**Naming Convention**:
- Directory: `e2e` subdirectory
- Files: `*_e2e_test.go`
- Functions: `TestE2EUserWorkflow`

## Test Organization

```
crules/
├── internal/
│   ├── core/
│   │   ├── registry.go
│   │   ├── registry_test.go  # Unit tests for registry
│   │   ├── sync.go
│   │   └── sync_test.go      # Unit tests for sync
│   └── ...
├── tests/
│   ├── integration/          # Integration tests
│   │   └── ...
│   ├── e2e/                  # End-to-end tests
│   │   └── ...
│   └── fixtures/             # Test fixtures
│       └── ...
└── ...
```

## Writing Effective Tests

### Test Structure

Follow the AAA pattern:

1. **Arrange**: Set up the test environment and inputs
2. **Act**: Execute the code being tested
3. **Assert**: Verify the results match expectations

```go
func TestExample(t *testing.T) {
    // Arrange
    input := "test input"
    expected := "expected output"
    
    // Act
    actual := FunctionUnderTest(input)
    
    // Assert
    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}
```

### Mocking

Use interfaces and dependency injection to enable mocking external dependencies.

```go
// Define an interface for the dependency
type FileSystem interface {
    Exists(path string) bool
    Read(path string) ([]byte, error)
    Write(path string, data []byte) error
}

// Implement a mock for testing
type MockFileSystem struct {
    ExistsFunc func(path string) bool
    ReadFunc   func(path string) ([]byte, error)
    WriteFunc  func(path string, data []byte) error
}

func (m *MockFileSystem) Exists(path string) bool {
    return m.ExistsFunc(path)
}

func (m *MockFileSystem) Read(path string) ([]byte, error) {
    return m.ReadFunc(path)
}

func (m *MockFileSystem) Write(path string, data []byte) error {
    return m.WriteFunc(path, data)
}

// In your tests
func TestWithMock(t *testing.T) {
    mockFS := &MockFileSystem{
        ExistsFunc: func(path string) bool {
            return path == "/expected/path"
        },
        // Configure other methods...
    }
    
    // Inject the mock into the code under test
    sut := NewSyncManager(mockFS)
    
    // Test with the mock...
}
```

### Test Fixtures

Store test data in the `tests/fixtures` directory to keep tests clean and reusable.

```go
func loadFixture(t *testing.T, name string) []byte {
    data, err := ioutil.ReadFile(filepath.Join("../../tests/fixtures", name))
    if err != nil {
        t.Fatalf("Failed to load fixture %s: %v", name, err)
    }
    return data
}

func TestWithFixture(t *testing.T) {
    data := loadFixture(t, "sample_registry.json")
    // Use the fixture data in your test
}
```

## Running Tests

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test ./... -cover
```

### Generate Coverage Report

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run Specific Tests

```bash
# Run tests in a specific package
go test ./internal/core

# Run a specific test
go test ./internal/core -run TestSyncManager_Init

# Run tests matching a pattern
go test ./... -run "TestRegistry.*"
```

### Debugging Tests

```bash
# Run tests with verbose output
go test -v ./...

# Run tests with logging
go test -v ./... -args -debug
```

## Continuous Integration

Tests are automatically run in CI for:
- Pull requests
- Merges to main branch
- Release tags

The CI pipeline runs:
1. Unit tests
2. Integration tests
3. End-to-end tests
4. Code coverage reporting

## Testing Environment Variables

Set up environment variables for testing by creating a `.env.test` file:

```
APP_NAME=crules_test
RULES_DIR_NAME=.cursor/rules_test
REGISTRY_FILE_NAME=registry_test.json
LOG_LEVEL=debug
```

## Related Documentation

- [Contributing Guide](contributing.md) - How to contribute to the codebase
- [Code Structure](code-structure.md) - Detailed code organization
