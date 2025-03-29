package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"crules/internal/ui"
	"crules/internal/utils"
)

// StoreRules saves parsed rules to the filesystem
func StoreRules(rules []Rule, forceOverwrite bool) error {
	// Get main rules location from config
	config := utils.LoadConfig()

	// Get app paths
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = utils.DefaultAppName
	}

	appPaths := utils.GetAppPaths(appName)
	baseDir := appPaths.GetRulesDir(config.RulesDirName)

	return StoreRulesToPath(rules, baseDir, forceOverwrite)
}

// StoreRulesToPath saves parsed rules to a specific directory path
func StoreRulesToPath(rules []Rule, baseDir string, forceOverwrite bool) error {
	// Ensure the directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Store each rule
	savedCount := 0
	for _, rule := range rules {
		filename := sanitizeFilename(rule.Name) + getExtensionForFormat(rule.Format)
		fullPath := filepath.Join(baseDir, filename)

		// Check if file already exists
		if _, err := os.Stat(fullPath); err == nil && !forceOverwrite {
			// File exists and force flag is not set
			if !ui.PromptYesNo(fmt.Sprintf("Rule '%s' already exists in %s. Overwrite?", rule.Name, baseDir)) {
				// Generate alternative name
				timestamp := fmt.Sprintf("%d", time.Now().Unix())
				filename = sanitizeFilename(rule.Name) + "-" + timestamp + getExtensionForFormat(rule.Format)
				fullPath = filepath.Join(baseDir, filename)
			}
		}

		// Add metadata header if it's a markdown file
		content := rule.Content
		if rule.Format == "markdown" || rule.Format == "mdx" {
			// Add metadata as a comment if not already present
			if !strings.Contains(content, "Description:") {
				metaHeader := fmt.Sprintf("<!-- \nName: %s\nDescription: %s\nSource: %s\nImported: %s\n-->\n\n",
					rule.Name,
					rule.Description,
					rule.Source,
					time.Now().Format(time.RFC3339),
				)
				content = metaHeader + content
			}
		}

		// Write the file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write rule file: %w", err)
		}

		utils.Info(fmt.Sprintf("Saved rule %s to %s", rule.Name, fullPath))
		savedCount++
	}

	if savedCount == 0 {
		return fmt.Errorf("no rules were saved")
	}

	return nil
}

// sanitizeFilename creates a valid filename from a rule name
func sanitizeFilename(name string) string {
	// Remove characters that aren't allowed in filenames
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name

	for _, char := range invalid {
		result = strings.ReplaceAll(result, char, "-")
	}

	// Convert to lowercase and replace spaces with hyphens
	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, " ", "-")

	// Ensure it's not too long
	if len(result) > 100 {
		result = result[:100]
	}

	return result
}

// getExtensionForFormat returns the appropriate file extension for a rule format
func getExtensionForFormat(format string) string {
	switch format {
	case "markdown", "mdx":
		return ".mdc"
	case "json":
		return ".json"
	default:
		return ".txt"
	}
}
