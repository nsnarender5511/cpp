package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Common gitignore entries to include in new files
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

// EnsureGitIgnoreEntry ensures that a specific entry exists in the gitignore file
// If the gitignore file doesn't exist, it will be created with default entries
func EnsureGitIgnoreEntry(dir string, entry string) error {
	gitignorePath := filepath.Join(dir, ".gitignore")

	// Check if .gitignore exists
	if !FileExists(gitignorePath) {
		return CreateGitIgnoreWithDefaults(dir, entry)
	}

	// Check if entry already exists
	exists, err := GitIgnoreHasEntry(gitignorePath, entry)
	if err != nil {
		return err
	}

	if exists {
		Debug("GitIgnore entry already exists | entry=" + entry)
		return nil
	}

	// Append entry to existing file
	return AppendToGitIgnore(gitignorePath, entry)
}

// GitIgnoreHasEntry checks if a gitignore file already has a specific entry
func GitIgnoreHasEntry(gitignorePath string, entry string) (bool, error) {
	file, err := os.Open(gitignorePath)
	if err != nil {
		Error("Failed to open gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return false, fmt.Errorf("failed to open gitignore file: %w", err)
	}
	defer file.Close()

	// Normalize the entry for comparison (trim whitespace and trailing slashes)
	normalizedEntry := strings.TrimSpace(entry)
	normalizedEntry = strings.TrimSuffix(normalizedEntry, "/")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Normalize the line for comparison
		normalizedLine := strings.TrimSuffix(line, "/")

		// Check various formats that would match the entry
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

// AppendToGitIgnore appends an entry to an existing gitignore file
func AppendToGitIgnore(gitignorePath, entry string) error {
	Debug("Appending entry to gitignore | path=" + gitignorePath + ", entry=" + entry)

	// Read the content to check if there's a trailing newline
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		Error("Failed to read gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to read gitignore file: %w", err)
	}

	// Open file for appending
	file, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Error("Failed to open gitignore file for writing | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to open gitignore file for writing: %w", err)
	}
	defer file.Close()

	// Add a newline if the file doesn't end with one
	if len(content) > 0 && !strings.HasSuffix(string(content), "\n") {
		if _, err := file.WriteString("\n"); err != nil {
			Error("Failed to write newline to gitignore | path=" + gitignorePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write newline to gitignore: %w", err)
		}
	}

	// Write the entry with a newline
	if _, err := file.WriteString(entry + "\n"); err != nil {
		Error("Failed to write entry to gitignore | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to write entry to gitignore: %w", err)
	}

	Debug("Successfully appended entry to gitignore | path=" + gitignorePath + ", entry=" + entry)
	return nil
}

// CreateGitIgnoreWithDefaults creates a new gitignore file with default entries plus the specified entry
func CreateGitIgnoreWithDefaults(dir string, entry string) error {
	gitignorePath := filepath.Join(dir, ".gitignore")
	Debug("Creating new gitignore file | path=" + gitignorePath)

	file, err := os.Create(gitignorePath)
	if err != nil {
		Error("Failed to create gitignore file | path=" + gitignorePath + ", error=" + err.Error())
		return fmt.Errorf("failed to create gitignore file: %w", err)
	}
	defer file.Close()

	// Write default entries
	for _, defaultEntry := range defaultGitIgnoreEntries {
		if _, err := file.WriteString(defaultEntry + "\n"); err != nil {
			Error("Failed to write default entry to gitignore | path=" + gitignorePath + ", error=" + err.Error())
			return fmt.Errorf("failed to write default entry to gitignore: %w", err)
		}
	}

	// Add the specific entry if it's not already in the defaults
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

// FileExists checks if a file exists and is not a directory
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && !info.IsDir()
}
