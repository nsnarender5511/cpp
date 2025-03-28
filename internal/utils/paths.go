package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

// Default application name if not specified
const DefaultAppName = "crules"

// AppPaths holds all relevant paths for the application
type AppPaths struct {
	ConfigDir string // For configuration files
	DataDir   string // For registry and user data
	CacheDir  string // For temporary or cached data
	LogDir    string // For log files
}

// GetAppPaths returns platform-specific paths for the application
func GetAppPaths(appName string) AppPaths {
	// Default application name for paths
	if appName == "" {
		// Try to get app name from environment
		envAppName := os.Getenv("APP_NAME")
		if envAppName != "" {
			appName = envAppName
		} else {
			appName = DefaultAppName
		}
	}

	// Get locations based on OS
	switch runtime.GOOS {
	case "windows":
		return getWindowsAppPaths(appName)
	case "darwin":
		return getDarwinAppPaths(appName)
	default:
		return getUnixAppPaths(appName)
	}
}

// getWindowsAppPaths returns Windows-specific paths
func getWindowsAppPaths(appName string) AppPaths {
	p := AppPaths{}

	// Configuration: %APPDATA%\crules
	p.ConfigDir = filepath.Join(os.Getenv("APPDATA"), appName)

	// Data: %LOCALAPPDATA%\crules
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		localAppData = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local")
	}
	p.DataDir = filepath.Join(localAppData, appName)

	// Cache: %LOCALAPPDATA%\crules\Cache
	p.CacheDir = filepath.Join(localAppData, appName, "Cache")

	// Logs: %LOCALAPPDATA%\crules\Logs
	p.LogDir = filepath.Join(localAppData, appName, "Logs")

	return p
}

// getDarwinAppPaths returns macOS-specific paths
func getDarwinAppPaths(appName string) AppPaths {
	p := AppPaths{}
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback if home directory can't be determined
		home = "."
	}

	// Configuration: ~/Library/Application Support/crules
	p.ConfigDir = filepath.Join(home, "Library", "Application Support", appName)

	// Data: ~/Library/Application Support/crules
	p.DataDir = filepath.Join(home, "Library", "Application Support", appName)

	// Cache: ~/Library/Caches/crules
	p.CacheDir = filepath.Join(home, "Library", "Caches", appName)

	// Logs: ~/Library/Logs/crules
	p.LogDir = filepath.Join(home, "Library", "Logs", appName)

	return p
}

// getUnixAppPaths returns Linux/Unix-specific paths (XDG compliant)
func getUnixAppPaths(appName string) AppPaths {
	p := AppPaths{}
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback if home directory can't be determined
		home = "."
	}

	// Check for XDG environment variables
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = filepath.Join(home, ".config")
	}

	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(home, ".local", "share")
	}

	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome == "" {
		xdgCacheHome = filepath.Join(home, ".cache")
	}

	xdgStateHome := os.Getenv("XDG_STATE_HOME")
	if xdgStateHome == "" {
		xdgStateHome = filepath.Join(home, ".local", "state")
	}

	// Configuration: ~/.config/crules
	p.ConfigDir = filepath.Join(xdgConfigHome, appName)

	// Data: ~/.local/share/crules
	p.DataDir = filepath.Join(xdgDataHome, appName)

	// Cache: ~/.cache/crules
	p.CacheDir = filepath.Join(xdgCacheHome, appName)

	// Logs: ~/.local/state/crules/logs
	p.LogDir = filepath.Join(xdgStateHome, appName, "logs")

	return p
}

// Helper methods for the AppPaths struct

// GetRulesDir returns the path where rules should be stored
func (p AppPaths) GetRulesDir(dirName string) string {
	if dirName == "" {
		dirName = DefaultRulesDirName
	}
	return filepath.Join(p.DataDir, dirName)
}

// GetRegistryFile returns the path to the registry file
func (p AppPaths) GetRegistryFile(fileName string) string {
	if fileName == "" {
		fileName = DefaultRegistryFileName
	}
	return filepath.Join(p.DataDir, fileName)
}

// DefaultLogFileName is the default name for log files
const DefaultLogFileName = "crules.log"

// GetLogFile returns the path to a log file
func (p AppPaths) GetLogFile(logName string) string {
	if logName == "" {
		logName = DefaultLogFileName
	}
	return filepath.Join(p.LogDir, logName)
}

// EnsureDirExists creates directory if it doesn't exist
func EnsureDirExists(dir string, perm os.FileMode) error {
	return os.MkdirAll(dir, perm)
}
