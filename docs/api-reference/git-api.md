---
version: v0.1.0
last_updated: 2023-03-29
applies_to: cursor++ v0.1.0+
---

# Git API Reference

> 🔄 API reference for the Git integration utilities of cursor++.

## Overview

The Git API provides utilities for interacting with Git repositories in cursor++. It includes functions for Git repository detection, .gitignore management, and Git operations.

This reference covers the most important types and functions in the Git package located at `/internal/git`.

## Repository Detection

### `IsGitRepository` Function

```go
func IsGitRepository(path string) bool
```

Determines if the specified path is within a Git repository.

#### Parameters

- `path string`: The path to check

#### Returns

- `bool`: `true` if the path is within a Git repository, `false` otherwise

#### Example

```go
if git.IsGitRepository(projectPath) {
    fmt.Println("Project is in a Git repository")
} else {
    fmt.Println("Project is not in a Git repository")
}
```

### `GetRepositoryRoot` Function

```go
func GetRepositoryRoot(path string) (string, error)
```

Gets the root directory of a Git repository containing the specified path.

#### Parameters

- `path string`: The path to check

#### Returns

- `string`: The absolute path to the repository root
- `error`: An error if the path is not in a Git repository or if there was a problem

#### Example

```go
root, err := git.GetRepositoryRoot(projectPath)
if err != nil {
    return err
}
fmt.Printf("Repository root: %s\n", root)
```

## .gitignore Management

### `UpdateGitignore` Function

```go
func UpdateGitignore(repoPath string, entries []string) error
```

Updates the .gitignore file in the repository to include the specified entries.

#### Parameters

- `repoPath string`: The path to the repository
- `entries []string`: The entries to add to .gitignore

#### Returns

- `error`: An error if there was a problem updating the .gitignore file

#### Example

```go
err := git.UpdateGitignore(projectPath, []string{".cursor/", "*.log"})
if err != nil {
    return err
}
fmt.Println("Updated .gitignore")
```

### `IsPathIgnored` Function

```go
func IsPathIgnored(repoPath string, path string) (bool, error)
```

Checks if a path is ignored according to Git's ignore rules.

#### Parameters

- `repoPath string`: The path to the repository
- `path string`: The path to check

#### Returns

- `bool`: `true` if the path is ignored, `false` otherwise
- `error`: An error if there was a problem checking the path

#### Example

```go
ignored, err := git.IsPathIgnored(projectPath, ".cursor/rules")
if err != nil {
    return err
}
if ignored {
    fmt.Println("Path is ignored by Git")
} else {
    fmt.Println("Path is not ignored by Git")
}
```

## Git Operations

### `Commit` Function

```go
func Commit(repoPath string, message string, files []string) error
```

Commits changes to the specified files with the given message.

#### Parameters

- `repoPath string`: The path to the repository
- `message string`: The commit message
- `files []string`: The files to commit

#### Returns

- `error`: An error if there was a problem committing the changes

#### Example

```go
err := git.Commit(projectPath, "Update rules", []string{".cursor/rules/wizard.mdc"})
if err != nil {
    return err
}
fmt.Println("Changes committed")
```

### `GetChangedFiles` Function

```go
func GetChangedFiles(repoPath string) ([]string, error)
```

Gets a list of files that have been changed but not committed.

#### Parameters

- `repoPath string`: The path to the repository

#### Returns

- `[]string`: The list of changed files
- `error`: An error if there was a problem getting the changed files

#### Example

```go
files, err := git.GetChangedFiles(projectPath)
if err != nil {
    return err
}
fmt.Printf("Changed files: %v\n", files)
```

### `GetCurrentBranch` Function

```go
func GetCurrentBranch(repoPath string) (string, error)
```

Gets the name of the current branch.

#### Parameters

- `repoPath string`: The path to the repository

#### Returns

- `string`: The name of the current branch
- `error`: An error if there was a problem getting the branch name

#### Example

```go
branch, err := git.GetCurrentBranch(projectPath)
if err != nil {
    return err
}
fmt.Printf("Current branch: %s\n", branch)
```

## Repository Information

### `GetRepositoryInfo` Function

```go
func GetRepositoryInfo(repoPath string) (*RepositoryInfo, error)
```

Gets information about a Git repository.

#### Parameters

- `repoPath string`: The path to the repository

#### Returns

- `*RepositoryInfo`: A struct containing information about the repository
- `error`: An error if there was a problem getting the repository information

#### Example

```go
info, err := git.GetRepositoryInfo(projectPath)
if err != nil {
    return err
}
fmt.Printf("Repository URL: %s\n", info.RemoteURL)
fmt.Printf("Current branch: %s\n", info.CurrentBranch)
```

### `RepositoryInfo` Type

```go
type RepositoryInfo struct {
    Path          string
    CurrentBranch string
    RemoteURL     string
    IsClean       bool
}
```

A struct containing information about a Git repository.

- `Path`: The absolute path to the repository
- `CurrentBranch`: The name of the current branch
- `RemoteURL`: The URL of the remote repository
- `IsClean`: Whether the repository has uncommitted changes

## Remote Operations

### `Clone` Function

```go
func Clone(url string, path string) error
```

Clones a Git repository from the specified URL to the specified path.

#### Parameters

- `url string`: The URL of the repository to clone
- `path string`: The path to clone the repository to

#### Returns

- `error`: An error if there was a problem cloning the repository

#### Example

```go
err := git.Clone("https://github.com/user/repo.git", "./repo")
if err != nil {
    return err
}
fmt.Println("Repository cloned")
```

### `Pull` Function

```go
func Pull(repoPath string) error
```

Pulls changes from the remote repository.

#### Parameters

- `repoPath string`: The path to the repository

#### Returns

- `error`: An error if there was a problem pulling changes

#### Example

```go
err := git.Pull(projectPath)
if err != nil {
    return err
}
fmt.Println("Changes pulled from remote")
```

### `Push` Function

```go
func Push(repoPath string) error
```

Pushes changes to the remote repository.

#### Parameters

- `repoPath string`: The path to the repository

#### Returns

- `error`: An error if there was a problem pushing changes

#### Example

```go
err := git.Push(projectPath)
if err != nil {
    return err
}
fmt.Println("Changes pushed to remote")
```

## File Operations

### `AddFile` Function

```go
func AddFile(repoPath string, filePath string) error
```

Adds a file to the Git index.

#### Parameters

- `repoPath string`: The path to the repository
- `filePath string`: The path to the file to add

#### Returns

- `error`: An error if there was a problem adding the file

#### Example

```go
err := git.AddFile(projectPath, ".cursor/rules/wizard.mdc")
if err != nil {
    return err
}
fmt.Println("File added to index")
```

### `GetFileStatus` Function

```go
func GetFileStatus(repoPath string, filePath string) (string, error)
```

Gets the Git status of a file.

#### Parameters

- `repoPath string`: The path to the repository
- `filePath string`: The path to the file to check

#### Returns

- `string`: The status of the file (e.g., "modified", "new", "deleted")
- `error`: An error if there was a problem getting the file status

#### Example

```go
status, err := git.GetFileStatus(projectPath, ".cursor/rules/wizard.mdc")
if err != nil {
    return err
}
fmt.Printf("File status: %s\n", status)
```

## Error Handling

### Common Errors

- `ErrNotGitRepository`: Returned when a path is not within a Git repository
- `ErrGitCommandFailed`: Returned when a Git command fails

### Error Handling Example

```go
root, err := git.GetRepositoryRoot(projectPath)
if err != nil {
    if errors.Is(err, git.ErrNotGitRepository) {
        fmt.Println("The project is not in a Git repository.")
        // Handle the special case
    } else {
        fmt.Printf("An error occurred: %v\n", err)
        // Handle other errors
    }
    return err
}
```

## Best Practices

### Working with Git Repositories

- Always check if a path is within a Git repository before performing Git operations
- Handle the case where a project is not in a Git repository gracefully
- Use `.gitignore` to exclude the `.cursor` directory to avoid committing user-specific AI configurations

### Error Handling

Proper error handling for Git operations:

```go
if !git.IsGitRepository(projectPath) {
    fmt.Println("Project is not in a Git repository, skipping Git operations")
    return nil
}

err := git.UpdateGitignore(projectPath, []string{".cursor/"})
if err != nil {
    return fmt.Errorf("failed to update .gitignore: %w", err)
}
```

### Repository Operations

For operations that modify the repository:

1. Check if the path is in a Git repository
2. Perform the operation
3. Handle errors appropriately
4. Provide user feedback

## Navigation

- Previous: [UI API](./ui-api.md)
- Next: [Version API](./version-api.md)
- Up: [API Reference](./README.md)
- Home: [Documentation Home](../README.md)

---

