package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


var defaultGitIgnoreEntries = []string{
	"# OS files",
	".DS_Store",
	"Thumbs.db",
	"*.swp",
	"",
	"# Cursor rules",
	".cursor/",
	"",
}



func EnsureGitIgnoreEntry(dir string, entry string) error {
	gitignorePath := filepath.Join(dir, ".gitignore")

	
	if !FileExists(gitignorePath) {
		return CreateGitIgnoreWithDefaults(dir, entry)
	}

	
	exists, err := GitIgnoreHasEntry(gitignorePath, entry)
	if err != nil {
		return err
	}

	if exists {
		Debug("GitIgnore entry already exists | entry=" + entry)
		return nil
	}

	
	return AppendToGitIgnore(gitignorePath, entry)
}


func GitIgnoreHasEntry(gitignorePath string, entry string) (bool, error) {
	file, err := os.Open(gitignorePath)
	if err != nil {
		Error("Failed to open gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return false, fmt.Errorf("failed to open gitignore file: %w", err)
	}
	defer file.Close()

	
	normalizedEntry := strings.TrimSpace(entry)
	normalizedEntry = strings.TrimSuffix(normalizedEntry, "/")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		
		normalizedLine := strings.TrimSuffix(line, "/")

		
		if normalizedLine == normalizedEntry ||
			normalizedLine == normalizedEntry+"/" ||
			normalizedLine == "/"+normalizedEntry ||
			normalizedLine == "/"+normalizedEntry+"/" {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		Error("Error reading gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return false, fmt.Errorf("error reading gitignore file: %w", err)
	}

	return false, nil
}


func AppendToGitIgnore(gitignorePath, entry string) error {
	Debug("Appending entry to gitignore | path=" + gitignorePath + ", entry=" + entry)

	
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		Error("Failed to read gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to read gitignore file: %w", err)
	}

	
	file, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Error("Failed to open gitignore file for writing | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to open gitignore file for writing: %w", err)
	}
	defer file.Close()

	
	if len(content) > 0 && !strings.HasSuffix(string(content), "\n") {
		if _, err := file.WriteString("\n"); err != nil {
			Error("Failed to write newline to gitignore | path=" + gitignorePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write newline to gitignore: %w", err)
		}
	}

	
	if _, err := file.WriteString(entry + "\n"); err != nil {
		Error("Failed to write entry to gitignore | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to write entry to gitignore: %w", err)
	}

	Debug("Successfully appended entry to gitignore | path=" + gitignorePath + ", entry=" + entry)
	return nil
}


func CreateGitIgnoreWithDefaults(dir string, entry string) error {
	gitignorePath := filepath.Join(dir, ".gitignore")
	Debug("Creating new gitignore file | path=" + gitignorePath)

	file, err := os.Create(gitignorePath)
	if err != nil {
		Error("Failed to create gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to create gitignore file: %w", err)
	}
	defer file.Close()

	
	for _, defaultEntry := range defaultGitIgnoreEntries {
		if _, err := file.WriteString(defaultEntry + "\n"); err != nil {
			Error("Failed to write default entry to gitignore | path=" + gitignorePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write default entry to gitignore: %w", err)
		}
	}

	
	entryExists := false
	for _, defaultEntry := range defaultGitIgnoreEntries {
		if defaultEntry == entry {
			entryExists = true
			break
		}
	}

	if !entryExists {
		if _, err := file.WriteString(entry + "\n"); err != nil {
			Error("Failed to write entry to gitignore | path=" + gitignorePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write entry to gitignore: %w", err)
		}
	}

	Debug("Successfully created gitignore file | path=" + gitignorePath)
	return nil
}


func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && !info.IsDir()
}
