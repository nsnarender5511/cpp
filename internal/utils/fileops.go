package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DirExists checks if directory exists
func DirExists(path string) bool {
	_, err := os.Stat(path)
	result := !os.IsNotExist(err)
	if result {
		Debug("Directory exists | path=" + path)
	} else {
		Debug("Directory does not exist | path=" + path)
	}
	return result
}

// ConfirmOverwrite asks user for confirmation
func ConfirmOverwrite(dirName string) bool {
	fmt.Printf("%s already exists. Overwrite? (y/n): ", dirName)
	var response string
	fmt.Scanln(&response)
	result := response == "y"
	if result {
		Debug("User confirmed overwrite")
	} else {
		Debug("User declined overwrite")
	}
	return result
}

// CopyDir copies a directory recursively, filtering to only copy .mdc files
func CopyDir(src, dst string) error {
	Debug("Copying directory | source=" + src + ", destination=" + dst)

	cm := NewConfigManager()
	if err := cm.Load(); err != nil {
		Error("Failed to load configuration: " + err.Error())
		return err
	}
	config := cm.GetConfig()

	// Define directories to skip
	dirsToSkip := map[string]bool{
		".git":         true,
		".cursor":      true,
		".github":      true,
		".vscode":      true,
		"node_modules": true,
	}

	// Create destination directory
	if err := os.MkdirAll(dst, config.DirPermission); err != nil {
		Error("Failed to create destination directory | path=" + dst + ", error=" + err.Error())
		return err
	}

	// Read source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		Error("Failed to read source directory | path=" + src + ", error=" + err.Error())
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := entry.Info()
		if err != nil {
			Error("Failed to get file info | path=" + sourcePath + ", error=" + err.Error())
			return err
		}

		if fileInfo.IsDir() {
			// Skip directories that should be excluded
			if dirsToSkip[entry.Name()] {
				Debug("Skipping excluded directory | path=" + sourcePath)
				continue
			}

			Debug("Copying subdirectory | source=" + sourcePath + ", destination=" + destPath)
			if err := CopyDir(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			// Only copy .mdc files
			if filepath.Ext(entry.Name()) == ".mdc" {
				Debug("Copying .mdc file | source=" + sourcePath + ", destination=" + destPath)
				if err := CopyFile(sourcePath, destPath); err != nil {
					return err
				}
			} else {
				Debug("Skipping non-mdc file | path=" + sourcePath)
			}
		}
	}

	Debug("Directory copied successfully | source=" + src + ", destination=" + dst)
	return nil
}

// CopyFile copies a single file
func CopyFile(src, dst string) error {
	cm := NewConfigManager()
	if err := cm.Load(); err != nil {
		Error("Failed to load configuration: " + err.Error())
		return err
	}
	config := cm.GetConfig()

	sourceFile, err := os.Open(src)
	if err != nil {
		Error("Failed to open source file | path=" + src + ", error=" + err.Error())
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		Error("Failed to create destination file | path=" + dst + ", error=" + err.Error())
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		Error("Failed to copy file contents | source=" + src + ", destination=" + dst + ", error=" + err.Error())
		return err
	}

	if err := os.Chmod(dst, config.FilePermission); err != nil {
		Error("Failed to set file permissions | path=" + dst + ", error=" + err.Error())
		return err
	}

	Debug("File copied successfully | source=" + src + ", destination=" + dst)
	return nil
}
