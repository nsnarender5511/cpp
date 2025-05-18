package git

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)


type GitManager struct {
	service GitService
}


func NewGitManager(service GitService) *GitManager {
	if service == nil {
		service = NewGitCommandService(nil)
	}
	return &GitManager{service: service}
}


func (m *GitManager) CloneOrPull(ctx context.Context, url string, destDir string) error {
	
	if err := os.MkdirAll(filepath.Dir(destDir), 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	
	if _, err := os.Stat(filepath.Join(destDir, ".git")); os.IsNotExist(err) {
		
		if err := m.service.Clone(ctx, url, destDir); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}
	} else {
		
		if err := m.service.Pull(ctx, destDir); err != nil {
			return fmt.Errorf("failed to pull repository: %w", err)
		}
	}

	return nil
}


func (m *GitManager) CheckoutRef(ctx context.Context, repoPath string, ref string) error {
	if err := m.service.Checkout(ctx, repoPath, ref); err != nil {
		return fmt.Errorf("failed to checkout ref %s: %w", ref, err)
	}
	return nil
}
