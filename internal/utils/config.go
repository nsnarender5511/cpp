package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)


const (
	
	DefaultAgentsDirName = "vibe"

	
	DefaultConfigFileName = "config.json"

	
	DefaultRegistryFileName = "registry.json"

	
	DefaultRulesDirName = ".cursor/rules"

	
	DefaultSourceFolder = "default"

	
	DefaultDirPermission = 0755

	
	DefaultFilePermission = 0644
)


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


type ConfigValidator func(value string) error


type ConfigManager struct {
	config     *Config
	mu         sync.RWMutex
	validators map[string]ConfigValidator
}


type ConfigProvider interface {
	GetConfig() *Config
}


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


func (cm *ConfigManager) RegisterValidator(field string, validator ConfigValidator) {
	cm.validators[field] = validator
}


func (cm *ConfigManager) Load() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	configPath := filepath.Join(GetAppPaths(DefaultAgentsDirName).ConfigDir, DefaultConfigFileName)

	
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


func (cm *ConfigManager) GetConfig() *Config {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	
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


func (cm *ConfigManager) SetConfig(config *Config) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.config = config
}
