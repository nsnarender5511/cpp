package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Default configuration values - centralized here
const (
	DefaultAgentsDirName    = "rules"
	DefaultRegistryFileName = "registry.json"
	DefaultRulesDirName     = ".cursor"
	DefaultDirPermission    = 0755
	DefaultFilePermission   = 0644
)

// Config holds application configuration
type Config struct {
	AgentsDirName     string
	RegistryFileName  string
	RulesDirName      string
	DirPermission     os.FileMode
	FilePermission    os.FileMode
	MultiAgentEnabled bool
	LastSelectedAgent string
}

// LoadConfig loads configuration from environment with validation
func LoadConfig() *Config {
	// Default values
	config := &Config{
		AgentsDirName:     DefaultAgentsDirName,
		RegistryFileName:  DefaultRegistryFileName,
		RulesDirName:      DefaultRulesDirName,
		DirPermission:     DefaultDirPermission,
		FilePermission:    DefaultFilePermission,
		MultiAgentEnabled: true,
		LastSelectedAgent: "",
	}

	// Override with environment variables if they exist
	if val := os.Getenv("AGENTS_DIR_NAME"); val != "" {
		// Validate directory path for potential security issues
		if validateDirPath(val) {
			config.AgentsDirName = val
		} else {
			Debug("Invalid AGENTS_DIR_NAME value, using default: " + val)
		}
	}

	if val := os.Getenv("REGISTRY_FILE_NAME"); val != "" {
		// Validate filename
		if validateFileName(val) {
			config.RegistryFileName = val
		} else {
			Debug("Invalid REGISTRY_FILE_NAME value, using default: " + val)
		}
	}

	if val := os.Getenv("RULES_DIR_NAME"); val != "" {
		// Validate directory path
		if validateDirPath(val) {
			config.RulesDirName = val
		} else {
			Debug("Invalid RULES_DIR_NAME value, using default: " + val)
		}
	}

	if val := os.Getenv("DIR_PERMISSION"); val != "" {
		// Validate directory permissions
		if perm, err := strconv.ParseUint(val, 8, 32); err == nil {
			// Ensure permissions are reasonable (at least read + execute)
			if perm >= 0500 && perm <= 0777 {
				config.DirPermission = os.FileMode(perm)
			} else {
				Debug(fmt.Sprintf("Directory permission value out of reasonable range: %s", val))
			}
		} else {
			Debug("Invalid DIR_PERMISSION value, using default: " + val)
		}
	}

	if val := os.Getenv("FILE_PERMISSION"); val != "" {
		// Validate file permissions
		if perm, err := strconv.ParseUint(val, 8, 32); err == nil {
			// Ensure permissions are reasonable (at least readable)
			if perm >= 0400 && perm <= 0777 {
				config.FilePermission = os.FileMode(perm)
			} else {
				Debug(fmt.Sprintf("File permission value out of reasonable range: %s", val))
			}
		} else {
			Debug("Invalid FILE_PERMISSION value, using default: " + val)
		}
	}

	if val := os.Getenv("MULTI_AGENT_ENABLED"); val != "" {
		// Validate boolean value
		config.MultiAgentEnabled = val == "true" || val == "1" || val == "yes"
	}

	if val := os.Getenv("LAST_SELECTED_AGENT"); val != "" {
		// Validate agent ID format
		if validateFileName(val) {
			config.LastSelectedAgent = val
		} else {
			Debug("Invalid LAST_SELECTED_AGENT value, using default: " + val)
		}
	}

	// Log final configuration
	Info(fmt.Sprintf("Configuration loaded | AgentsDirName=%s, RulesDirName=%s, MultiAgentEnabled=%v, LastSelectedAgent=%s",
		config.AgentsDirName, config.RulesDirName, config.MultiAgentEnabled, config.LastSelectedAgent))

	return config
}

// validateDirPath checks if a directory path is valid and safe
func validateDirPath(path string) bool {
	// Check for directory traversal attempts
	if strings.Contains(path, "..") {
		Warn("Security warning: path contains directory traversal pattern: " + path)
		return false
	}

	// Check for absolute paths (which might be acceptable in some cases)
	if filepath.IsAbs(path) {
		Debug("Directory path is absolute: " + path)
		// Consider the context - if absolute paths are not desired, return false
	}

	// Check for suspicious path elements
	suspicious := []string{"/tmp", "/dev", "/proc", "/sys", "/var/run"}
	for _, suspect := range suspicious {
		if strings.Contains(path, suspect) {
			Warn("Security warning: path contains suspicious elements: " + path)
			return false
		}
	}

	return true
}

// validateFileName checks if a filename is valid and safe
func validateFileName(name string) bool {
	// Check for directory traversal attempts
	if strings.Contains(name, "..") || strings.Contains(name, "/") || strings.Contains(name, "\\") {
		Warn("Security warning: invalid filename: " + name)
		return false
	}

	// Check for empty or overly long filenames
	if name == "" || len(name) > 255 {
		Warn("Invalid filename length: " + name)
		return false
	}

	return true
}

// SaveConfig saves the configuration to the default location
func SaveConfig(config *Config) error {
	// Get app paths
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = DefaultAppName
	}
	appPaths := GetAppPaths(appName)

	// Ensure config directory exists
	if err := EnsureDirExists(appPaths.ConfigDir, config.DirPermission); err != nil {
		Error("Cannot create config directory | path=" + appPaths.ConfigDir + ", error=" + err.Error())
		return fmt.Errorf("cannot create config directory: %v", err)
	}

	// Build config content
	content := fmt.Sprintf("AGENTS_DIR_NAME=%s\n", config.AgentsDirName)
	content += fmt.Sprintf("REGISTRY_FILE_NAME=%s\n", config.RegistryFileName)
	content += fmt.Sprintf("RULES_DIR_NAME=%s\n", config.RulesDirName)
	content += fmt.Sprintf("DIR_PERMISSION=%o\n", config.DirPermission)
	content += fmt.Sprintf("FILE_PERMISSION=%o\n", config.FilePermission)
	content += fmt.Sprintf("MULTI_AGENT_ENABLED=%t\n", config.MultiAgentEnabled)

	// Only add LastSelectedAgent if it's set
	if config.LastSelectedAgent != "" {
		content += fmt.Sprintf("LAST_SELECTED_AGENT=%s\n", config.LastSelectedAgent)
	}

	// Write to config.env file
	configFile := filepath.Join(appPaths.ConfigDir, "config.env")
	if err := os.WriteFile(configFile, []byte(content), config.FilePermission); err != nil {
		Error("Failed to write config file | path=" + configFile + ", error=" + err.Error())
		return fmt.Errorf("failed to write config file: %v", err)
	}

	Info("Configuration saved successfully | path=" + configFile)
	return nil
}
