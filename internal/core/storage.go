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


func StoreRules(rules []*CursorRule, forceOverwrite bool) error {
	
	cm := utils.NewConfigManager()
	if err := cm.Load(); err != nil {
		return wrapOpError("StoreRules", "config", err, "failed to load configuration")
	}
	cfg := cm.GetConfig()

	
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = utils.DefaultAgentsDirName
	}

	appPaths := utils.GetAppPaths(appName)
	baseDir := appPaths.GetRulesDir(cfg.RulesDirName)

	return StoreRulesToPath(rules, baseDir, forceOverwrite)
}


func StoreRulesToPath(rules []*CursorRule, baseDir string, forceOverwrite bool) error {
	
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return wrapOpError("StoreRulesToPath", baseDir, err, "failed to create directory")
	}

	
	savedCount := 0
	for _, rule := range rules {
		filename := sanitizeFilename(rule.Metadata.Name) + ".json"
		fullPath := filepath.Join(baseDir, filename)

		
		if _, err := os.Stat(fullPath); err == nil && !forceOverwrite {
			
			if !ui.PromptYesNo(rule.Metadata.Name + " already exists. Overwrite?") {
				
				timestamp := strconv.FormatInt(time.Now().Unix(), 10)
				filename = sanitizeFilename(rule.Metadata.Name) + "-" + timestamp + ".json"
				fullPath = filepath.Join(baseDir, filename)
			}
		}

		
		data, err := json.MarshalIndent(rule, "", "  ")
		if err != nil {
			return wrapOpError("StoreRulesToPath", fullPath, err, "failed to marshal rule")
		}

		
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


func sanitizeFilename(name string) string {
	
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name

	for _, char := range invalid {
		result = strings.ReplaceAll(result, char, "-")
	}

	
	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, " ", "-")

	
	if len(result) > 100 {
		result = result[:100]
	}

	return result
}












