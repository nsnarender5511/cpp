package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"vibe/internal/utils"
)

// Registry keeps track of all projects using vibe
type Registry struct {
	Projects []string `json:"projects"`
	path     string   // path to registry file
	config   *utils.Config
	mutex    *sync.RWMutex
}

// LoadRegistry loads or creates the registry
func LoadRegistry(registryPath string, config *utils.Config) (*Registry, error) {
	utils.Debug("Loading registry | path=" + registryPath)

	registry := &Registry{
		Projects: []string{},
		path:     registryPath,
		config:   config,
	}

	// Create new if doesn't exist
	if _, err := os.Stat(registryPath); os.IsNotExist(err) {
		utils.Debug("Registry file does not exist, creating new | path=" + registryPath)

		// Ensure directory exists
		registryDir := filepath.Dir(registryPath)
		if err := utils.EnsureDirExists(registryDir, config.DirPermission); err != nil {
			return nil, wrapOpError("LoadRegistry", registryDir, err, "failed to create registry directory")
		}

		return registry, registry.save()
	}

	// Read existing
	utils.Debug("Reading existing registry | path=" + registryPath)
	data, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, wrapOpError("LoadRegistry", registryPath, err, "failed to read registry file")
	}

	if err := json.Unmarshal(data, registry); err != nil {
		return nil, wrapParseError(registryPath, err, 0)
	}

	registry.config = config
	registry.path = registryPath
	utils.Debug("Registry loaded successfully | projects=" + strconv.Itoa(len(registry.Projects)))
	return registry, nil
}

// AddProject adds a project to registry
func (r *Registry) AddProject(projectPath string) error {
	utils.Debug("Adding project to registry | project=" + projectPath)

	// Validate project path
	if !utils.DirExists(projectPath) {
		return wrapValidationError("projectPath", "directory does not exist")
	}

	// Check if already registered
	for _, p := range r.Projects {
		if p == projectPath {
			utils.Debug("Project already registered, skipping | project=" + projectPath)
			return nil // Already registered
		}
	}

	r.Projects = append(r.Projects, projectPath)
	utils.Debug("Project added to registry | project=" + projectPath)
	return r.save()
}

// GetProjects returns all registered projects
func (r *Registry) GetProjects() []string {
	utils.Debug("Getting registered projects | count=" + strconv.Itoa(len(r.Projects)))
	return r.Projects
}

// CleanProjects removes projects that no longer exist
func (r *Registry) CleanProjects() (int, error) {
	utils.Debug("Cleaning registry of non-existent projects")

	originalCount := len(r.Projects)
	validProjects := make([]string, 0, originalCount)

	for _, project := range r.Projects {
		if utils.DirExists(project) {
			validProjects = append(validProjects, project)
		} else {
			utils.Debug("Removing non-existent project | project=" + project)
		}
	}

	r.Projects = validProjects
	removedCount := originalCount - len(validProjects)

	if removedCount > 0 {
		if err := r.save(); err != nil {
			return 0, wrapOpError("CleanProjects", r.path, err, "failed to save registry after cleaning")
		}
	}

	utils.Info("Registry cleaned | removed=" + strconv.Itoa(removedCount))
	return removedCount, nil
}

// save writes registry to disk
func (r *Registry) save() error {
	utils.Debug("Saving registry | path=" + r.path)
	data, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return wrapOpError("save", r.path, err, "failed to marshal registry")
	}

	if err := os.WriteFile(r.path, data, r.config.FilePermission); err != nil {
		return wrapOpError("save", r.path, err, "failed to write registry file")
	}

	utils.Debug("Registry saved successfully | path=" + r.path)
	return nil
}

// GetProjectCount returns the number of projects in the registry
func (r *Registry) GetProjectCount() int {
	if r.mutex != nil {
		r.mutex.RLock()
		defer r.mutex.RUnlock()
	}
	return len(r.Projects)
}
