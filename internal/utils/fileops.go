package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

// CopyDirSelective copies a specific subfolder from a source directory to a destination
// If sourceFolderName is empty, it behaves like CopyDir and copies everything
func CopyDirSelective(src, dst, sourceFolderName string) error {
	Debug("Copying directory selectively | source=" + src + ", destination=" + dst + ", sourceFolder=" + sourceFolderName)

	cm := NewConfigManager()
	if err := cm.Load(); err != nil {
		Error("Failed to load configuration: " + err.Error())
		return err
	}
	config := cm.GetConfig()

	// If no source folder specified, use normal copy
	if sourceFolderName == "" {
		Debug("No source folder specified, using standard copy")
		return CopyDir(src, dst)
	}

	// Check if the source subfolder exists
	sourceFolderPath := filepath.Join(src, sourceFolderName)
	if !DirExists(sourceFolderPath) {
		errorMsg := fmt.Sprintf("Source folder does not exist: %s", sourceFolderPath)
		Error(errorMsg)
		return fmt.Errorf(errorMsg)
	}

	// Create destination directory
	if err := os.MkdirAll(dst, config.DirPermission); err != nil {
		Error("Failed to create destination directory | path=" + dst + ", error=" + err.Error())
		return err
	}

	// Now copy from the source subfolder to destination
	Debug("Copying from source subfolder | path=" + sourceFolderPath)
	return CopyDir(sourceFolderPath, dst)
}

// EnsurePathInFile ensures that a specific pattern is present in a file
// If the file doesn't exist, it creates it with the pattern
// If the file exists but doesn't contain the pattern, it appends the pattern
// If the file already contains the pattern, it does nothing
func EnsurePathInFile(filePath, pattern string) error {
	Debug("Ensuring pattern exists in file | path=" + filePath + ", pattern=" + pattern)

	// Check if file exists
	if !FileExists(filePath) {
		Debug("File does not exist, creating new file | path=" + filePath)
		// Create directory if it doesn't exist
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			Error("Failed to create directory for file | path=" + dir + ", error=" + err.Error())
			return fmt.Errorf("failed to create directory for file: %w", err)
		}

		// Create file with pattern
		return os.WriteFile(filePath, []byte(pattern+"\n"), 0644)
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		Error("Failed to read file | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Check if pattern already exists (either as exact match or with newline)
	contentStr := string(content)
	if strings.Contains(contentStr, pattern) ||
		strings.Contains(contentStr, pattern+"\n") ||
		strings.Contains(contentStr, pattern+"\r\n") {
		Debug("Pattern already exists in file, skipping | path=" + filePath + ", pattern=" + pattern)
		return nil
	}

	// Append pattern to file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Error("Failed to open file for writing | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	// Add a newline if file doesn't end with one
	if len(content) > 0 && !strings.HasSuffix(contentStr, "\n") {
		if _, err := file.WriteString("\n"); err != nil {
			Error("Failed to write newline to file | path=" + filePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write newline to file: %w", err)
		}
	}

	// Write pattern
	if _, err := file.WriteString(pattern + "\n"); err != nil {
		Error("Failed to write pattern to file | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to write pattern to file: %w", err)
	}

	Debug("Successfully added pattern to file | path=" + filePath + ", pattern=" + pattern)
	return nil
}
