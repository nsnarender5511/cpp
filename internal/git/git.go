package git

import (
	"fmt"
	"os"
	"os/exec"

	"cursor++/internal/utils"
)

// Clone clones a git repository to the specified local path
func Clone(url, destination string) error {
	utils.Debug("Cloning git repository | url=" + url + ", destination=" + destination)

	// Ensure the destination directory exists (parent directory)
	if err := os.MkdirAll(destination, 0755); err != nil {
		utils.Error("Failed to create destination directory | destination=" + destination + ", error=" + err.Error())
		return fmt.Errorf("failed to create destination directory: %v", err)
	}

	// Prepare the git clone command
	cmd := exec.Command("git", "clone", url, destination)

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := string(output)
		utils.Error("Git clone failed | error=" + err.Error() + ", output=" + errMsg)
		return fmt.Errorf("git clone failed: %v\n%s", err, errMsg)
	}

	utils.Info("Git repository cloned successfully | url=" + url + ", destination=" + destination)
	return nil
}

// IsValidRepo checks if a URL is a valid git repository
func IsValidRepo(url string) bool {
	utils.Debug("Checking if URL is a valid git repository | url=" + url)

	// The ls-remote command checks if the repository exists without cloning it
	cmd := exec.Command("git", "ls-remote", url)

	err := cmd.Run()
	isValid := err == nil

	utils.Debug("Git repository validation result | url=" + url + ", isValid=" + fmt.Sprintf("%v", isValid))
	return isValid
}

// CleanupOnFailure removes the destination directory if cloning fails
func CleanupOnFailure(destination string) {
	utils.Debug("Cleaning up failed git clone | destination=" + destination)

	// Only remove if it's not empty (something was partially cloned)
	if utils.DirExists(destination) {
		// Check if it's a git repository by looking for .git directory
		gitDir := destination + "/.git"
		if utils.DirExists(gitDir) {
			utils.Debug("Removing partial git clone | destination=" + destination)
			if err := os.RemoveAll(destination); err != nil {
				utils.Warn("Failed to clean up git clone | destination=" + destination + ", error=" + err.Error())
			}
		}
	}
}
