package git

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// GitManager handles Git operations for rule repositories
type GitManager struct {
	service GitService
}

// NewGitManager creates a new GitManager
func NewGitManager(service GitService) *GitManager {
	if service == nil {
		service = NewGitCommandService(nil)
	}
	return &GitManager{service: service}
}

// CloneOrPull clones a repository if it doesn't exist, or pulls if it does
func (m *GitManager) CloneOrPull(ctx context.Context, url string, destDir string) error {
	// Ensure destination directory exists
	if err := os.MkdirAll(filepath.Dir(destDir), 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Check if repository already exists
	if _, err := os.Stat(filepath.Join(destDir, ".git")); os.IsNotExist(err) {
		// Clone repository
		if err := m.service.Clone(ctx, url, destDir); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}
	} else {
		// Pull latest changes
		if err := m.service.Pull(ctx, destDir); err != nil {
			return fmt.Errorf("failed to pull repository: %w", err)
		}
	}

	return nil
}

// CheckoutRef checks out a specific reference in a repository
func (m *GitManager) CheckoutRef(ctx context.Context, repoPath string, ref string) error {
	if err := m.service.Checkout(ctx, repoPath, ref); err != nil {
		return fmt.Errorf("failed to checkout ref %s: %w", ref, err)
	}
	return nil
}
