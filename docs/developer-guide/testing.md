---
title: Testing Guide
description: How to test the cursor++ application
applies_to: cursor++ v0.1.0+
---

# Testing Guide

This document explains how to test cursor++ during development.

## Types of Tests

The cursor++ project includes several types of tests:

1. **Unit Tests** - Test individual functions and methods
2. **Integration Tests** - Test the interaction between components
3. **End-to-End Tests** - Test the entire application workflow

## Testing Stack

cursor++ uses the following testing components:

- **Go's built-in testing package**: For unit and integration tests
- **testify**: For assertions and mocking
- **golangci-lint**: For static code analysis
- **GitHub Actions**: For continuous integration testing

## Test Categories

### Unit Tests

Unit tests focus on testing individual functions and methods in isolation. These tests should be:

- Fast to execute
- Independent of the environment
- Focused on a single unit of functionality

```go
func TestFileExists(t *testing.T) {
    // Create a temporary file
    file, err := ioutil.TempFile("", "test_file")
    assert.NoError(t, err)
    defer os.Remove(file.Name())
    
    // Test that the file exists
    assert.True(t, utils.FileExists(file.Name()))
    
    // Test that a non-existent file does not exist
    assert.False(t, utils.FileExists(file.Name() + ".nonexistent"))
}
```

### Integration Tests

Integration tests verify that different components work together correctly. These tests:

- May interact with the file system
- Test multiple components together
- Verify end-to-end workflows

```go
func TestInitAndSync(t *testing.T) {
    // Set up test directories
    mainDir := t.TempDir()
    projectDir := t.TempDir()
    
    // Initialize the main location
    config := &core.Config{MainLocation: mainDir}
    assert.NoError(t, core.SaveConfig(config))
    
    // Create a test rule in the main location
    testRule := path.Join(mainDir, "test.mdc")
    assert.NoError(t, ioutil.WriteFile(testRule, []byte("test content"), 0644))
    
    // Initialize the project
    assert.NoError(t, core.InitRules(config, projectDir))
    
    // Verify the rule was copied to the project
    projectRule := path.Join(projectDir, ".cursor/rules/test.mdc")
    assert.True(t, utils.FileExists(projectRule))
    
    // Modify the rule in the main location
    assert.NoError(t, ioutil.WriteFile(testRule, []byte("updated content"), 0644))
    
    // Sync the rules
    assert.NoError(t, core.SyncRules(config, projectDir))
    
    // Verify the rule was updated in the project
    content, err := ioutil.ReadFile(projectRule)
    assert.NoError(t, err)
    assert.Equal(t, "updated content", string(content))
}
```

### End-to-End Tests

End-to-end tests verify that the entire application works correctly from a user's perspective. These tests:

- Exercise the CLI interface
- Verify actual command behavior
- Test realistic user workflows

```go
func TestCliWorkflow(t *testing.T) {
    // Set up test environment
    testHome := t.TempDir()
    oldHome := os.Getenv("HOME")
    os.Setenv("HOME", testHome)
    defer os.Setenv("HOME", oldHome)
    
    // Run the init command
    cmd := exec.Command("cursor++", "init", "--main", path.Join(testHome, "rules"))
    assert.NoError(t, cmd.Run())
    
    // Verify the configuration was created
    configPath := path.Join(testHome, ".config/cursor++/config.json")
    assert.True(t, utils.FileExists(configPath))
    
    // Additional workflow steps...
}
```

## Running Tests

### Unit Tests

To run all unit tests:

```bash
go test ./...
```

To run tests for a specific package:

```bash
go test ./internal/agent
```

To run a specific test:

```bash
go test ./internal/agent -run TestAgentInitialization
```

### Running Integration Tests

```bash
# Run integration tests
go test ./... -tags=integration

# Run integration tests for a specific package
go test ./internal/core -tags=integration
```

### Code Coverage

```bash
# Run tests with coverage
go test ./... -cover

# Generate coverage profile
go test ./... -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

## Writing Good Tests

### Best Practices

1. **Test behavior, not implementation**: Focus on what the code does, not how it does it.
2. **One assertion per test**: Each test should verify one aspect of behavior.
3. **Use table-driven tests**: For testing multiple inputs and expected outputs.
4. **Mock external dependencies**: Use interfaces and mocks to isolate tests.
5. **Clean up after tests**: Use `defer` to ensure resources are released.
6. **Use descriptive test names**: Names should indicate what the test is verifying.
7. **Test edge cases**: Consider boundary conditions and error scenarios.

### Example of Table-Driven Test

```go
func TestValidatePath(t *testing.T) {
    tests := []struct {
        name     string
        path     string
        expected bool
    }{
        {"Empty path", "", false},
        {"Valid path", "/valid/path", true},
        {"Path with invalid characters", "/invalid/path*", false},
        {"Path with spaces", "/path with spaces", true},
        {"Path with unicode", "/path/with/unicode/🚀", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := utils.IsValidPath(tt.path)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

## Test Mocks

For testing functionality that interacts with external systems (like Git repositories), we use mocks.

The `internal/utils/testutils` package provides helper functions and mock implementations to facilitate testing.

## Continuous Integration

The cursor++ project uses GitHub Actions for continuous integration. The CI pipeline runs all tests on each pull request and commit to the main branch.

See [.github/workflows/test.yml](https://github.com/nsnarender5511/cursor++/blob/main/.github/workflows/test.yml) for the CI configuration.

## Debugging Tests

Tips for debugging failing tests:

1. Use `t.Log` to output debugging information
2. Run tests with `-v` for verbose output
3. Use `t.Errorf` to provide detailed failure information
4. Use `-run` to focus on a specific failing test
5. Check test environment variables and file paths

## Test Fixtures

For tests that require common setup or test data, use test fixtures:

```go
// test/fixtures.go
package test

import (
    "io/ioutil"
    "os"
    "path"
    "testing"
)

// CreateTestRule creates a test rule file in the specified directory
func CreateTestRule(t *testing.T, dir, name, content string) string {
    t.Helper()
    rulePath := path.Join(dir, name)
    err := ioutil.WriteFile(rulePath, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test rule: %v", err)
    }
    return rulePath
}

// SetupTestEnvironment creates a test environment with a main location and project
func SetupTestEnvironment(t *testing.T) (mainDir, projectDir string) {
    t.Helper()
    mainDir = t.TempDir()
    projectDir = t.TempDir()
    return mainDir, projectDir
}
```

## See Also

- [Building from Source](./building.md) for instructions on building the project
- [Contributing Guidelines](./contributing.md) for information on contributing to the project
- [Code Structure](./code-structure.md) for details on the codebase organization

## Navigation

- Previous: [Code of Conduct](./code-of-conduct.md)
- Next: [Building](./building.md)
- Up: [Developer Guide](../README.md#developer-guide)
- Home: [Documentation Home](../README.md)
