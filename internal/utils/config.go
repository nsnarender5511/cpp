package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// Default configuration values - centralized here
const (
	// DefaultAgentsDirName is the default directory name for agent rules
	DefaultAgentsDirName = "vibe"

	// DefaultConfigFileName is the default configuration file name
	DefaultConfigFileName = "config.json"

	// DefaultRegistryFileName is the default registry file name
	DefaultRegistryFileName = "registry.json"

	// DefaultRulesDirName is the default directory name for rules - this is not to be changed ever
	DefaultRulesDirName = ".cursor/rules"

	// DefaultSourceFolder is the name of the folder to copy from the cloned repo
	DefaultSourceFolder = "default"

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
	SourceFolder      string      `json:"sourceFolder"`
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
			SourceFolder:      DefaultSourceFolder,
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
		SourceFolder:      cm.config.SourceFolder,
	}
}

// SetConfig updates the configuration
func (cm *ConfigManager) SetConfig(config *Config) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.config = config
}
