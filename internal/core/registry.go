package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"cursor++/internal/utils"
)

// Registry keeps track of all projects using cursor++
type Registry struct {
	Projects []string `json:"projects"`
	path     string   // path to registry file
	config   *utils.Config
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
			utils.Error("Failed to create registry directory | path=" + registryDir + ", error=" + err.Error())
			return nil, fmt.Errorf("cannot create registry directory: %v", err)
		}

		return registry, registry.save()
	}

	// Read existing
	utils.Debug("Reading existing registry | path=" + registryPath)
	data, err := os.ReadFile(registryPath)
	if err != nil {
		utils.Error("Failed to read registry file | path=" + registryPath + ", error=" + err.Error())
		return nil, err
	}

	if err := json.Unmarshal(data, registry); err != nil {
		utils.Error("Failed to parse registry file | path=" + registryPath + ", error=" + err.Error())
		return nil, err
	}

	registry.config = config
	registry.path = registryPath
	utils.Debug("Registry loaded successfully | projects=" + fmt.Sprintf("%d", len(registry.Projects)))
	return registry, nil
}

// AddProject adds a project to registry
func (r *Registry) AddProject(projectPath string) error {
	utils.Debug("Adding project to registry | project=" + projectPath)

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
	utils.Debug("Getting registered projects | count=" + fmt.Sprintf("%d", len(r.Projects)))
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
			utils.Error("Failed to save registry after cleaning | error=" + err.Error())
			return 0, err
		}
	}

	utils.Info("Registry cleaned | removed=" + fmt.Sprintf("%d", removedCount))
	return removedCount, nil
}

// save writes registry to disk
func (r *Registry) save() error {
	utils.Debug("Saving registry | path=" + r.path)
	data, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		utils.Error("Failed to marshal registry | error=" + err.Error())
		return err
	}

	if err := os.WriteFile(r.path, data, r.config.FilePermission); err != nil {
		utils.Error("Failed to write registry | path=" + r.path + ", error=" + err.Error())
		return err
	}

	utils.Debug("Registry saved successfully | path=" + r.path)
	return nil
}
