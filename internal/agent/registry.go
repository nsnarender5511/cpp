package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"crules/internal/utils"
)

// Registry manages the collection of available agents
type Registry struct {
	agents   map[string]*AgentDefinition
	rulesDir string
	config   *utils.Config
}

// NewRegistry creates a new agent registry
func NewRegistry(config *utils.Config, rulesDir string) (*Registry, error) {
	utils.Debug("Creating new agent registry | rulesDir=" + rulesDir)

	registry := &Registry{
		agents:   make(map[string]*AgentDefinition),
		rulesDir: rulesDir,
		config:   config,
	}

	// Scan for agent definitions
	if err := registry.scanAgents(); err != nil {
		utils.Error("Failed to scan for agents | error=" + err.Error())
		return nil, fmt.Errorf("failed to scan for agents: %v", err)
	}

	return registry, nil
}

// scanAgents discovers agent definition files in the rules directory
func (r *Registry) scanAgents() error {
	utils.Debug("Scanning for agent definitions | rulesDir=" + r.rulesDir)

	// Ensure rules directory exists
	if err := os.MkdirAll(r.rulesDir, r.config.DirPermission); err != nil {
		utils.Error("Failed to create rules directory | rulesDir=" + r.rulesDir + ", error=" + err.Error())
		return fmt.Errorf("failed to create rules directory: %v", err)
	}

	// Walk through the directory and find .mdc files
	mdcFiles := make([]string, 0)
	err := filepath.Walk(r.rulesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process regular files with .mdc extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mdc") {
			utils.Debug("Found agent definition file | path=" + path)
			mdcFiles = append(mdcFiles, path)
		}

		return nil
	})

	if err != nil {
		utils.Error("Error walking rules directory | error=" + err.Error())
		return fmt.Errorf("error walking rules directory: %v", err)
	}

	utils.Info("Found agent definition files | count=" + fmt.Sprintf("%d", len(mdcFiles)))

	// Process each .mdc file
	for _, path := range mdcFiles {
		if err := r.processAgentFile(path); err != nil {
			utils.Warn("Failed to process agent file | path=" + path + ", error=" + err.Error())
			// Continue with other files, don't stop on error
		}
	}

	return nil
}

// processAgentFile parses an agent definition file and adds it to the registry
func (r *Registry) processAgentFile(path string) error {
	utils.Debug("Processing agent file | path=" + path)

	// Read file
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Extract metadata from the file
	// For simplicity, we're assuming the first line after "# " is the agent name
	// and the content after "## 🎯 Role:" is the description
	// In a real implementation, proper parsing would be needed

	lines := strings.Split(string(content), "\n")
	var name, description string
	var capabilities []string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "# ") && name == "" {
			name = strings.TrimPrefix(trimmed, "# ")
		} else if strings.HasPrefix(trimmed, "## 🎯 Role:") && i+1 < len(lines) {
			// Take the next line as description
			if i+1 < len(lines) {
				description = strings.TrimSpace(lines[i+1])
			}
		} else if strings.HasPrefix(trimmed, "### ✅") {
			// Extract capability from this line
			capability := strings.TrimPrefix(trimmed, "### ✅")
			capability = strings.TrimSpace(capability)
			if capability != "" {
				capabilities = append(capabilities, capability)
			}
		}

		// For simplicity, we'll only process the first few lines
		if i > 50 {
			break
		}
	}

	// If we couldn't extract a name, use the filename
	if name == "" {
		base := filepath.Base(path)
		name = strings.TrimSuffix(base, ".mdc")
	}

	// Create short ID from filename
	base := filepath.Base(path)
	id := strings.TrimSuffix(base, ".mdc")

	// Create agent definition
	agent := &AgentDefinition{
		ID:             id,
		Name:           name,
		Description:    description,
		Capabilities:   capabilities,
		Version:        "1.0", // Default version
		DefinitionPath: path,
	}

	// Add to registry
	r.agents[id] = agent
	utils.Debug("Added agent to registry | id=" + id + ", name=" + name)

	return nil
}

// GetAgent returns an agent by ID
func (r *Registry) GetAgent(id string) (*AgentDefinition, bool) {
	agent, exists := r.agents[id]
	return agent, exists
}

// ListAgents returns all available agents
func (r *Registry) ListAgents() []*AgentDefinition {
	agents := make([]*AgentDefinition, 0, len(r.agents))
	for _, agent := range r.agents {
		agents = append(agents, agent)
	}
	return agents
}

// AgentExists checks if an agent with the given ID exists
func (r *Registry) AgentExists(id string) bool {
	_, exists := r.agents[id]
	return exists
}
