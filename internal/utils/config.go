package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Default configuration values - centralized here
const (
	// DefaultAgentsDirName is the default directory name for agent rules
	DefaultAgentsDirName = "cursor-rules"

	// DefaultConfigFileName is the default configuration file name
	DefaultConfigFileName = "config.json"

	// DefaultRegistryFileName is the default registry file name
	DefaultRegistryFileName = "registry.json"

	// DefaultRulesDirName is the default directory name for rules - this is not to be changed ever
	DefaultRulesDirName = ".cursor/rules"

	// DefaultDirPermission is the default permission for directories
	DefaultDirPermission = 0755

	// DefaultFilePermission is the default permission for files
	DefaultFilePermission = 0644
)

// Config represents the application configuration
type Config struct {
	RulesDirName      string      `json:"rulesDirName"`
	RegistryFileName  string      `json:"registryFileName"`
	DirPermission     os.FileMode `json:"dirPermission"`
	FilePermission    os.FileMode `json:"filePermission"`
	MultiAgentEnabled bool        `json:"multiAgentEnabled"`
	AgentsDirName     string      `json:"agentsDirName"`
	LastSelectedAgent string      `json:"lastSelectedAgent"`
}

// ConfigValidator defines a validation function for config values
type ConfigValidator func(value string) error

// ConfigManager manages application configuration
type ConfigManager struct {
	config     *Config
	mu         sync.RWMutex
	validators map[string]ConfigValidator
}

// ConfigProvider defines an interface for accessing configuration
type ConfigProvider interface {
	GetConfig() *Config
}

// NewConfigManager creates a new ConfigManager
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config: &Config{
			RulesDirName:      DefaultRulesDirName,
			RegistryFileName:  DefaultRegistryFileName,
			DirPermission:     DefaultDirPermission,
			FilePermission:    DefaultFilePermission,
			AgentsDirName:     DefaultAgentsDirName,
			MultiAgentEnabled: false,
			LastSelectedAgent: "",
		},
		validators: make(map[string]ConfigValidator),
	}
}

// RegisterValidator adds a new validator for a config field
func (cm *ConfigManager) RegisterValidator(field string, validator ConfigValidator) {
	cm.validators[field] = validator
}

// Load loads the configuration from file
func (cm *ConfigManager) Load() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	configPath := filepath.Join(GetAppPaths(DefaultAgentsDirName).ConfigDir, DefaultConfigFileName)

	// If config file doesn't exist, use defaults
	if !FileExists(configPath) {
		Debug("Config file not found, using defaults | path=" + configPath)
		return nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return wrapOpError("Load", configPath, err, "failed to read config file")
	}

	if err := json.Unmarshal(data, cm.config); err != nil {
		return wrapOpError("Load", configPath, err, "failed to parse config file")
	}

	Debug("Loaded configuration from file | path=" + configPath)
	return nil
}

// Save saves the configuration to file
func (cm *ConfigManager) Save() error {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	configPath := filepath.Join(GetAppPaths(DefaultAgentsDirName).ConfigDir, DefaultConfigFileName)

	data, err := json.MarshalIndent(cm.config, "", "  ")
	if err != nil {
		return wrapOpError("Save", configPath, err, "failed to marshal config")
	}

	if err := os.MkdirAll(filepath.Dir(configPath), DefaultDirPermission); err != nil {
		return wrapOpError("Save", configPath, err, "failed to create config directory")
	}

	if err := os.WriteFile(configPath, data, DefaultFilePermission); err != nil {
		return wrapOpError("Save", configPath, err, "failed to write config file")
	}

	Debug("Saved configuration to file | path=" + configPath)
	return nil
}

// GetConfig returns a copy of the current configuration
func (cm *ConfigManager) GetConfig() *Config {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	// Return a copy to prevent external modification
	return &Config{
		RulesDirName:      cm.config.RulesDirName,
		RegistryFileName:  cm.config.RegistryFileName,
		DirPermission:     cm.config.DirPermission,
		FilePermission:    cm.config.FilePermission,
		MultiAgentEnabled: cm.config.MultiAgentEnabled,
		AgentsDirName:     cm.config.AgentsDirName,
		LastSelectedAgent: cm.config.LastSelectedAgent,
	}
}

// SetConfig updates the configuration
func (cm *ConfigManager) SetConfig(config *Config) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.config = config
}

// validateDirPath checks if a directory path is valid and safe
func validateDirPath(path string) error {
	// Check for directory traversal attempts
	if strings.Contains(path, "..") {
		Warn("Security warning: path contains directory traversal pattern: " + path)
		return fmt.Errorf("path contains directory traversal")
	}

	// Check for absolute paths (which might be acceptable in some cases)
	if filepath.IsAbs(path) {
		Debug("Directory path is absolute: " + path)
		// Consider the context - if absolute paths are not desired, return error
		return fmt.Errorf("absolute paths are not allowed")
	}

	// Check for suspicious path elements
	suspicious := []string{"/tmp", "/dev", "/proc", "/sys", "/var/run"}
	for _, suspect := range suspicious {
		if strings.Contains(path, suspect) {
			Warn("Security warning: path contains suspicious elements: " + path)
			return fmt.Errorf("path contains suspicious elements")
		}
	}

	return nil
}

// validateFileName checks if a filename is valid and safe
func validateFileName(name string) error {
	// Check for directory traversal attempts
	if strings.Contains(name, "..") || strings.Contains(name, "/") || strings.Contains(name, "\\") {
		Warn("Security warning: invalid filename: " + name)
		return fmt.Errorf("filename contains directory traversal")
	}

	// Check for empty or overly long filenames
	if name == "" || len(name) > 255 {
		Warn("Invalid filename length: " + name)
		return fmt.Errorf("invalid filename length")
	}

	return nil
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
	content := fmt.Sprintf("AGENTS_DIR_NAME=%s\n", config.RulesDirName)
	content += fmt.Sprintf("REGISTRY_FILE_NAME=%s\n", config.RegistryFileName)
	content += fmt.Sprintf("DIR_PERMISSION=%o\n", config.DirPermission)
	content += fmt.Sprintf("FILE_PERMISSION=%o\n", config.FilePermission)

	// Write to config.env file
	configFile := filepath.Join(appPaths.ConfigDir, "config.env")
	if err := os.WriteFile(configFile, []byte(content), config.FilePermission); err != nil {
		Error("Failed to write config file | path=" + configFile + ", error=" + err.Error())
		return fmt.Errorf("failed to write config file: %v", err)
	}

	Info("Configuration saved successfully | path=" + configFile)
	return nil
}
