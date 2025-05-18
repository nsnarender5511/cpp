package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"vibe/internal/core"
	"vibe/internal/ui"
	"vibe/internal/utils"
)

func cleanDirectory(dirPath string) error {
	var filesCleaned int
	var filesFailed int

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			if err := cleanFile(path); err != nil {
				ui.Error(fmt.Sprintf("Failed to clean %s: %v", path, err))
				filesFailed++
			} else {
				filesCleaned++
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path %s: %w", dirPath, err)
	}

	ui.Success(fmt.Sprintf("Successfully cleaned %d Go file(s).", filesCleaned))
	if filesFailed > 0 {
		ui.Warning(fmt.Sprintf("Failed to clean %d Go file(s).", filesFailed))
	}
	return nil
}

func cleanFile(filePath string) error {
	src, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	cleanedSrc, err := core.RemoveComments(src)
	if err != nil {
		return fmt.Errorf("failed to remove comments from %s: %w", filePath, err)
	}

	if err := ioutil.WriteFile(filePath, cleanedSrc, 0644); err != nil {
		return fmt.Errorf("failed to write cleaned content to %s: %w", filePath, err)
	}
	utils.Debug(fmt.Sprintf("Cleaned comments from %s", filePath))
	return nil
}
