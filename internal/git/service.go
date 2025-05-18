package git

import (
	"context"
	"fmt"
	"os/exec"
)


type GitService interface {
	Clone(ctx context.Context, url, dest string) error
	Pull(ctx context.Context, repoPath string) error
	Checkout(ctx context.Context, repoPath, ref string) error
}


type CommandExecutor interface {
	Execute(ctx context.Context, name string, args ...string) ([]byte, error)
}


type ShellCommandExecutor struct{}


func (s *ShellCommandExecutor) Execute(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}


type GitCommandService struct {
	executor CommandExecutor
}


func NewGitCommandService(executor CommandExecutor) GitService {
	if executor == nil {
		executor = &ShellCommandExecutor{}
	}
	return &GitCommandService{executor: executor}
}


func (s *GitCommandService) Clone(ctx context.Context, url, dest string) error {
	output, err := s.executor.Execute(ctx, "git", "clone", url, dest)
	if err != nil {
		return fmt.Errorf("git clone failed: %w\nOutput: %s", err, output)
	}
	return nil
}


func (s *GitCommandService) Pull(ctx context.Context, repoPath string) error {
	output, err := s.executor.Execute(ctx, "git", "-C", repoPath, "pull")
	if err != nil {
		return fmt.Errorf("git pull failed: %w\nOutput: %s", err, output)
	}
	return nil
}


func (s *GitCommandService) Checkout(ctx context.Context, repoPath, ref string) error {
	output, err := s.executor.Execute(ctx, "git", "-C", repoPath, "checkout", ref)
	if err != nil {
		return fmt.Errorf("git checkout failed: %w\nOutput: %s", err, output)
	}
	return nil
}
