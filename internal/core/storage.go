package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"vibe/internal/ui"
	"vibe/internal/utils"
)

// StoreRules saves parsed rules to the filesystem
func StoreRules(rules []*CursorRule, forceOverwrite bool) error {
	// Get main rules location from config
	cm := utils.NewConfigManager()
	if err := cm.Load(); err != nil {
		return wrapOpError("StoreRules", "config", err, "failed to load configuration")
	}
	cfg := cm.GetConfig()

	// Get app paths
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = utils.DefaultAgentsDirName
	}

	appPaths := utils.GetAppPaths(appName)
	baseDir := appPaths.GetRulesDir(cfg.RulesDirName)

	return StoreRulesToPath(rules, baseDir, forceOverwrite)
}

// StoreRulesToPath saves parsed rules to a specific directory path
func StoreRulesToPath(rules []*CursorRule, baseDir string, forceOverwrite bool) error {
	// Ensure the directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return wrapOpError("StoreRulesToPath", baseDir, err, "failed to create directory")
	}

	// Store each rule
	savedCount := 0
	for _, rule := range rules {
		filename := sanitizeFilename(rule.Metadata.Name) + ".json"
		fullPath := filepath.Join(baseDir, filename)

		// Check if file already exists
		if _, err := os.Stat(fullPath); err == nil && !forceOverwrite {
			// File exists and force flag is not set
			if !ui.PromptYesNo(rule.Metadata.Name + " already exists. Overwrite?") {
				// Generate alternative name
				timestamp := strconv.FormatInt(time.Now().Unix(), 10)
				filename = sanitizeFilename(rule.Metadata.Name) + "-" + timestamp + ".json"
				fullPath = filepath.Join(baseDir, filename)
			}
		}

		// Marshal rule to JSON
		data, err := json.MarshalIndent(rule, "", "  ")
		if err != nil {
			return wrapOpError("StoreRulesToPath", fullPath, err, "failed to marshal rule")
		}

		// Write the file
		if err := os.WriteFile(fullPath, data, 0644); err != nil {
			return wrapOpError("StoreRulesToPath", fullPath, err, "failed to write rule file")
		}

		utils.Info("Saved rule " + rule.Metadata.Name + " to " + fullPath)
		savedCount++
	}

	if savedCount == 0 {
		return wrapValidationError("rules", "no rules were saved")
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
// func getExtensionForFormat(format string) string {
// 	switch format {
// 	case "markdown", "mdx":
// 		return ".mdc"
// 	case "json":
// 		return ".json"
// 	default:
// 		return ".txt"
// 	}
// }
