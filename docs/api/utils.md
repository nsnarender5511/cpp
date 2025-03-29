# Utils API Reference

> 📌 **Summary**: Utility functions and helpers for the crules system.

## Overview

The utils package provides utility functions and helpers for the crules system, including file operations, path manipulation, validation, and other common tasks.

## File Operations

Functions for working with files and directories.

```go
// FileExists checks if a file exists at the given path
func FileExists(path string) bool

// DirExists checks if a directory exists at the given path
func DirExists(path string) bool

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error

// CopyDir copies a directory from src to dst recursively
func CopyDir(src, dst string) error

// ReadFile reads the contents of a file as a string
func ReadFile(path string) (string, error)

// WriteFile writes a string to a file
func WriteFile(path, content string) error
```

## Path Manipulation

Functions for manipulating file paths.

```go
// GetAbsPath converts a relative path to an absolute path
func GetAbsPath(path string) (string, error)

// GetRelPath converts an absolute path to a relative path
func GetRelPath(path, basePath string) (string, error)

// JoinPath joins path elements into a single path
func JoinPath(elements ...string) string

// SplitPath splits a path into its directory and file name components
func SplitPath(path string) (dir, file string)
```

## Validation

Functions for validating various inputs.

```go
// IsValidPath checks if a path is valid
func IsValidPath(path string) bool

// IsValidDirectoryName checks if a directory name is valid
func IsValidDirectoryName(name string) bool

// IsValidFileName checks if a file name is valid
func IsValidFileName(name string) bool
```

## Miscellaneous

Other utility functions.

```go
// GetHomeDir returns the user's home directory
func GetHomeDir() (string, error)

// GetConfigDir returns the user's configuration directory
func GetConfigDir() (string, error)

// FindProjectRoot attempts to find the root directory of a project
func FindProjectRoot(startDir string) (string, error)

// GenerateRandomID generates a random ID string
func GenerateRandomID(length int) string
```

## Example Usage

```go
// Check if a file exists
if utils.FileExists("/path/to/file.txt") {
    // File exists, do something
}

// Get the absolute path
absPath, err := utils.GetAbsPath("./relative/path")
if err != nil {
    log.Fatalf("Failed to get absolute path: %v", err)
}

// Copy a directory
if err := utils.CopyDir("/path/to/src", "/path/to/dst"); err != nil {
    log.Fatalf("Failed to copy directory: %v", err)
}

// Get the user's home directory
homeDir, err := utils.GetHomeDir()
if err != nil {
    log.Fatalf("Failed to get home directory: %v", err)
}
```
