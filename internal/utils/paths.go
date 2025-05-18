package utils

import (
	"os"
	"path/filepath"
	"runtime"
)


const DefaultAppName = "vibe"


type AppPaths struct {
	ConfigDir string 
	DataDir   string 
	CacheDir  string 
	LogDir    string 
}


func GetAppPaths(appName string) AppPaths {
	
	if appName == "" {
		
		envAppName := os.Getenv("APP_NAME")
		if envAppName != "" {
			appName = envAppName
		} else {
			appName = DefaultAppName
		}
	}

	
	switch runtime.GOOS {
	case "windows":
		return getWindowsAppPaths(appName)
	case "darwin":
		return getDarwinAppPaths(appName)
	default:
		return getUnixAppPaths(appName)
	}
}


func getWindowsAppPaths(appName string) AppPaths {
	p := AppPaths{}

	
	p.ConfigDir = filepath.Join(os.Getenv("APPDATA"), appName)

	
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		localAppData = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local")
	}
	p.DataDir = filepath.Join(localAppData, appName)

	
	p.CacheDir = filepath.Join(localAppData, appName, "Cache")

	
	p.LogDir = filepath.Join(localAppData, appName, "Logs")

	return p
}


func getDarwinAppPaths(appName string) AppPaths {
	p := AppPaths{}
	home, err := os.UserHomeDir()
	if err != nil {
		
		home = "."
	}

	
	p.ConfigDir = filepath.Join(home, "Library", "Application Support", appName)

	
	p.DataDir = filepath.Join(home, "Library", "Application Support", appName)

	
	p.CacheDir = filepath.Join(home, "Library", "Caches", appName)

	
	p.LogDir = filepath.Join(home, "Library", "Logs", appName)

	return p
}


func getUnixAppPaths(appName string) AppPaths {
	p := AppPaths{}
	home, err := os.UserHomeDir()
	if err != nil {
		
		home = "."
	}

	
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

	
	p.ConfigDir = filepath.Join(xdgConfigHome, appName)

	
	p.DataDir = filepath.Join(xdgDataHome, appName)

	
	p.CacheDir = filepath.Join(xdgCacheHome, appName)

	
	p.LogDir = filepath.Join(xdgStateHome, appName, "logs")

	return p
}




func (p AppPaths) GetRulesDir(dirName string) string {
	if dirName == "" {
		dirName = DefaultRulesDirName
	}
	return filepath.Join(p.DataDir, dirName)
}


func (p AppPaths) GetRegistryFile(fileName string) string {
	if fileName == "" {
		fileName = DefaultRegistryFileName
	}
	return filepath.Join(p.DataDir, fileName)
}


const DefaultLogFileName = "vibe.log"


func (p AppPaths) GetLogFile(logName string) string {
	if logName == "" {
		logName = DefaultLogFileName
	}
	return filepath.Join(p.LogDir, logName)
}


func EnsureDirExists(dir string, perm os.FileMode) error {
	return os.MkdirAll(dir, perm)
}
