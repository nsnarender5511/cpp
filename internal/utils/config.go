package utils

import (
	"os"
	"strconv"
)

// Default configuration values - centralized here
const (
	DefaultRulesDirName     = ".cursor/rules"
	DefaultRegistryFileName = "registry.json"
	DefaultDirPermission    = 0755
	DefaultFilePermission   = 0644
)

// Config holds application configuration
type Config struct {
	RulesDirName     string
	RegistryFileName string
	DirPermission    os.FileMode
	FilePermission   os.FileMode
}

// LoadConfig loads configuration from environment
func LoadConfig() *Config {
	// Default values
	config := &Config{
		RulesDirName:     DefaultRulesDirName,
		RegistryFileName: DefaultRegistryFileName,
		DirPermission:    DefaultDirPermission,
		FilePermission:   DefaultFilePermission,
	}

	// Override with environment variables if they exist
	if val := os.Getenv("RULES_DIR_NAME"); val != "" {
		config.RulesDirName = val
	}

	if val := os.Getenv("REGISTRY_FILE_NAME"); val != "" {
		config.RegistryFileName = val
	}

	if val := os.Getenv("DIR_PERMISSION"); val != "" {
		if perm, err := strconv.ParseUint(val, 8, 32); err == nil {
			config.DirPermission = os.FileMode(perm)
		}
	}

	if val := os.Getenv("FILE_PERMISSION"); val != "" {
		if perm, err := strconv.ParseUint(val, 8, 32); err == nil {
			config.FilePermission = os.FileMode(perm)
		}
	}

	return config
}
