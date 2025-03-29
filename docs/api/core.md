# Core API Reference

> 📌 **Summary**: Core functionality API documentation for the crules system.

## Overview

The core package provides the fundamental functionality of the crules system, handling rule management, synchronization, and project tracking.

## Core Components

### Configuration Management

The configuration system handles loading, saving, and validating the crules configuration file.

```go
// LoadConfig loads the configuration from the default location
func LoadConfig() (*Config, error)

// SaveConfig saves the configuration to the default location
func SaveConfig(config *Config) error
```

### Rule Management

These functions handle the initialization, synchronization, and merging of rules across projects.

```go
// InitRules initializes rules in the current directory
func InitRules(config *Config) error

// SyncRules synchronizes rules from the main location to the current directory
func SyncRules(config *Config) error

// MergeRules merges rules from the current directory to the main location
func MergeRules(config *Config) error
```

### Project Registry

Functions for managing the registry of projects that use crules.

```go
// RegisterProject adds a project to the registry
func RegisterProject(config *Config, path string) error

// ListProjects returns a list of all registered projects
func ListProjects(config *Config) ([]string, error)

// CleanProjects removes non-existent projects from the registry
func CleanProjects(config *Config) error
```

## Error Handling

The core package defines several error types for specific failure scenarios:

- `ErrConfigNotFound`: Configuration file not found
- `ErrMainLocationNotSet`: Main rules location not set
- `ErrProjectNotRegistered`: Project not found in registry
- `ErrRulesNotInitialized`: Rules not initialized in the directory

## Example Usage

```go
// Initialize a new configuration
config := &core.Config{
    MainLocation: "/path/to/main/rules",
    Projects: []string{},
}

// Save the configuration
if err := core.SaveConfig(config); err != nil {
    log.Fatalf("Failed to save config: %v", err)
}

// Register a project
if err := core.RegisterProject(config, "/path/to/project"); err != nil {
    log.Fatalf("Failed to register project: %v", err)
}
```
