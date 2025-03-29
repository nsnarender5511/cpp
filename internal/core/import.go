package core

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"crules/internal/ui"
	"crules/internal/utils"
)

// ImportRules handles the end-to-end process of importing rules from a URL
func ImportRules(sourceURL string, force bool, webMode bool) error {
	// Validate URL
	if _, err := url.ParseRequestURI(sourceURL); err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	utils.Info("Importing rules from URL | url=" + sourceURL)
	ui.Info("Fetching rules from %s...", sourceURL)

	// Fetch content
	content, err := utils.FetchFromURL(sourceURL)
	if err != nil {
		return fmt.Errorf("failed to fetch content: %w", err)
	}

	ui.Info("Parsing rules...")

	// Parse rules
	rules, err := ParseRules(content, sourceURL, webMode)
	if err != nil {
		return fmt.Errorf("failed to parse rules: %w", err)
	}

	if len(rules) == 0 {
		return fmt.Errorf("no valid rules found in the content")
	}

	// Display found rules
	ui.Header("Found %d rules to import:", len(rules))
	for i, rule := range rules {
		ui.Plain("  %d. %s - %s", i+1, rule.Name, rule.Description)
	}

	ui.Plain("")
	if !ui.PromptYesNo("Do you want to import these rules?") {
		return fmt.Errorf("import cancelled by user")
	}

	// Store rules to main location
	ui.Info("Storing rules to main location...")
	if err := StoreRules(rules, force); err != nil {
		return fmt.Errorf("failed to store rules to main location: %w", err)
	}

	// Check if we also need to store to current directory
	currentDir, err := os.Getwd()
	if err != nil {
		ui.Warning("Could not access current directory: %v", err)
	} else {
		config := utils.LoadConfig()
		currentRulesDir := filepath.Join(currentDir, config.RulesDirName)

		// If rules dir exists in current directory, also store there
		if utils.DirExists(currentRulesDir) {
			ui.Info("Also storing rules to current directory...")
			if err := StoreRulesToPath(rules, currentRulesDir, force); err != nil {
				ui.Warning("Failed to store rules to current directory: %v", err)
			} else {
				ui.Success("Rules also stored to current directory")
			}
		}
	}

	ui.Success("Successfully imported %d rules", len(rules))
	return nil
}
