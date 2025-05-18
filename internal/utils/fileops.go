package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)


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


func CopyDir(src, dst string) error {
	Debug("Copying directory | source=" + src + ", destination=" + dst)

	cm := NewConfigManager()
	if err := cm.Load(); err != nil {
		Error("Failed to load configuration: " + err.Error())
		return err
	}
	config := cm.GetConfig()

	
	dirsToSkip := map[string]bool{
		".git":         true,
		".cursor":      true,
		".github":      true,
		".vscode":      true,
		"node_modules": true,
	}

	
	if err := os.MkdirAll(dst, config.DirPermission); err != nil {
		Error("Failed to create destination directory | path=" + dst + ", error=" + err.Error())
		return err
	}

	
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
			
			if dirsToSkip[entry.Name()] {
				Debug("Skipping excluded directory | path=" + sourcePath)
				continue
			}

			Debug("Copying subdirectory | source=" + sourcePath + ", destination=" + destPath)
			if err := CopyDir(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			
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



func CopyDirSelective(src, dst, sourceFolderName string) error {
	Debug("Copying directory selectively | source=" + src + ", destination=" + dst + ", sourceFolder=" + sourceFolderName)

	cm := NewConfigManager()
	if err := cm.Load(); err != nil {
		Error("Failed to load configuration: " + err.Error())
		return err
	}
	config := cm.GetConfig()

	
	if sourceFolderName == "" {
		Debug("No source folder specified, using standard copy")
		return CopyDir(src, dst)
	}

	
	sourceFolderPath := filepath.Join(src, sourceFolderName)
	if !DirExists(sourceFolderPath) {
		errorMsg := fmt.Sprintf("Source folder does not exist: %s", sourceFolderPath)
		Error(errorMsg)
		return fmt.Errorf("%s", errorMsg)
	}

	
	if err := os.MkdirAll(dst, config.DirPermission); err != nil {
		Error("Failed to create destination directory | path=" + dst + ", error=" + err.Error())
		return err
	}

	
	Debug("Copying from source subfolder | path=" + sourceFolderPath)
	return CopyDir(sourceFolderPath, dst)
}





func EnsurePathInFile(filePath, pattern string) error {
	Debug("Ensuring pattern exists in file | path=" + filePath + ", pattern=" + pattern)

	
	if !FileExists(filePath) {
		Debug("File does not exist, creating new file | path=" + filePath)
		
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			Error("Failed to create directory for file | path=" + dir + ", error=" + err.Error())
			return fmt.Errorf("failed to create directory for file: %w", err)
		}

		
		return os.WriteFile(filePath, []byte(pattern+"\n"), 0644)
	}

	
	content, err := os.ReadFile(filePath)
	if err != nil {
		Error("Failed to read file | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to read file: %w", err)
	}

	
	contentStr := string(content)
	if strings.Contains(contentStr, pattern) ||
		strings.Contains(contentStr, pattern+"\n") ||
		strings.Contains(contentStr, pattern+"\r\n") {
		Debug("Pattern already exists in file, skipping | path=" + filePath + ", pattern=" + pattern)
		return nil
	}

	
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Error("Failed to open file for writing | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	
	if len(content) > 0 && !strings.HasSuffix(contentStr, "\n") {
		if _, err := file.WriteString("\n"); err != nil {
			Error("Failed to write newline to file | path=" + filePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write newline to file: %w", err)
		}
	}

	
	if _, err := file.WriteString(pattern + "\n"); err != nil {
		Error("Failed to write pattern to file | path=" + filePath + ", error=" + err.Error())
		return fmt.Errorf("failed to write pattern to file: %w", err)
	}

	Debug("Successfully added pattern to file | path=" + filePath + ", pattern=" + pattern)
	return nil
}
