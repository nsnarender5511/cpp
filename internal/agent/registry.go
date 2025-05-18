package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"vibe/internal/utils"
)


type Registry struct {
	agents       map[string]*AgentDefinition
	rulesDir     string
	config       *utils.Config
	progressFunc func(event string, message string)
}


func NewRegistry(config *utils.Config, rulesDir string) (*Registry, error) {
	utils.Debug("Creating new agent registry | rulesDir=" + rulesDir)

	registry := &Registry{
		agents:       make(map[string]*AgentDefinition),
		rulesDir:     rulesDir,
		config:       config,
		progressFunc: nil,
	}

	
	if err := registry.scanAgents(); err != nil {
		utils.Error("Failed to scan for agents | error=" + err.Error())
		return nil, fmt.Errorf("failed to scan for agents: %v", err)
	}

	return registry, nil
}


func (r *Registry) SetProgressCallback(progressFunc func(event string, message string)) {
	r.progressFunc = progressFunc
}


func (r *Registry) reportProgress(event string, message string) {
	if r.progressFunc != nil {
		r.progressFunc(event, message)
	}
}


func (r *Registry) scanAgents() error {
	utils.Debug("Scanning for agent definitions | rulesDir=" + r.rulesDir)
	r.reportProgress("scan_start", "Scanning for agent definitions...")

	
	if err := os.MkdirAll(r.rulesDir, r.config.DirPermission); err != nil {
		utils.Error("Failed to create rules directory | rulesDir=" + r.rulesDir + ", error=" + err.Error())
		r.reportProgress("scan_error", "Failed to create rules directory: "+err.Error())
		return fmt.Errorf("failed to create rules directory: %v", err)
	}

	
	mdcFiles := make([]string, 0)
	err := filepath.Walk(r.rulesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mdc") {
			utils.Debug("Found agent definition file | path=" + path)
			r.reportProgress("file_found", path)
			mdcFiles = append(mdcFiles, path)
		}

		return nil
	})

	if err != nil {
		utils.Error("Error walking rules directory | error=" + err.Error())
		r.reportProgress("scan_error", "Error walking rules directory: "+err.Error())
		return fmt.Errorf("error walking rules directory: %v", err)
	}

	utils.Info("Found agent definition files | count=" + fmt.Sprintf("%d", len(mdcFiles)))
	r.reportProgress("files_count", fmt.Sprintf("%d", len(mdcFiles)))

	
	for _, path := range mdcFiles {
		base := filepath.Base(path)
		id := strings.TrimSuffix(base, ".mdc")
		r.reportProgress("processing_file", id)

		if err := r.processAgentFile(path); err != nil {
			utils.Warn("Failed to process agent file | path=" + path + ", error=" + err.Error())
			r.reportProgress("process_error", id+": "+err.Error())
			
		} else {
			r.reportProgress("process_success", id)
		}
	}

	r.reportProgress("scan_complete", fmt.Sprintf("%d agents loaded", len(r.agents)))
	return nil
}


func (r *Registry) processAgentFile(path string) error {
	
	id := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	
	if !validateAgentID(id) {
		utils.Warn("Invalid agent ID | id=" + id + ", path=" + path)
		return fmt.Errorf("invalid agent ID: %s", id)
	}

	
	name, description, err := r.extractAgentMetadata(path)
	if err != nil {
		utils.Warn("Failed to extract agent metadata | id=" + id + ", path=" + path + ", error=" + err.Error())
		name = id 
		description = "No description available"
	}

	
	templates, err := r.findTemplates(id)
	if err != nil {
		utils.Warn("Failed to find templates | id=" + id + ", error=" + err.Error())
		
	}

	
	agent := &AgentDefinition{
		ID:          id,
		Name:        name,
		Description: description,
		Version:     "1.0", 
		Type:        "ai",  
		Config:      make(map[string]interface{}),
		Templates:   templates,
		LastUpdated: time.Now(),
	}

	
	r.agents[id] = agent
	utils.Debug("Added agent to registry | id=" + id + ", name=" + name)

	return nil
}



func (r *Registry) GetAgent(id string) (*AgentDefinition, error) {
	
	if !validateAgentID(id) {
		utils.Error("Invalid agent ID format | id=" + id)
		return nil, fmt.Errorf("invalid agent ID format: %s", id)
	}

	agent, exists := r.agents[id]
	if !exists {
		utils.Warn("Agent not found | id=" + id)
		return nil, fmt.Errorf("agent not found: %s", id)
	}
	return agent, nil
}


func validateAgentID(id string) bool {
	
	if strings.Contains(id, "/") || strings.Contains(id, "\\") ||
		strings.Contains(id, "..") || strings.Contains(id, ".") ||
		strings.Contains(id, " ") {
		return false
	}

	
	if id == "" || len(id) > 100 {
		return false
	}

	return true
}


func (r *Registry) ListAgents() []*AgentDefinition {
	agents := make([]*AgentDefinition, 0, len(r.agents))
	for _, agent := range r.agents {
		agents = append(agents, agent)
	}
	return agents
}


func (r *Registry) AgentExists(id string) bool {
	
	if !validateAgentID(id) {
		return false
	}

	_, exists := r.agents[id]
	return exists
}


func (r *Registry) ScanAgentsWithAnimation() error {
	
	r.agents = make(map[string]*AgentDefinition)

	
	return r.scanAgents()
}


func (r *Registry) extractAgentMetadata(path string) (string, string, error) {
	
	content, err := os.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("failed to read file: %v", err)
	}

	
	
	
	lines := strings.Split(string(content), "\n")
	var name, description string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "# ") && name == "" {
			name = strings.TrimPrefix(trimmed, "# ")
		} else if strings.HasPrefix(trimmed, "## 🎯 Role:") && i+1 < len(lines) {
			
			if i+1 < len(lines) {
				description = strings.TrimSpace(lines[i+1])
			}
		}

		
		if i > 50 {
			break
		}
	}

	
	if name == "" {
		base := filepath.Base(path)
		name = strings.TrimSuffix(base, ".mdc")
	}

	return name, description, nil
}


func (r *Registry) findTemplates(agentID string) ([]string, error) {
	
	templatesDir := filepath.Join(r.rulesDir, "templates")
	if !utils.DirExists(templatesDir) {
		return nil, nil 
	}

	
	pattern := filepath.Join(templatesDir, agentID+"-*.tmpl")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search for templates: %w", err)
	}

	
	templates := make([]string, len(matches))
	for i, path := range matches {
		
		templates[i] = filepath.Base(path)
	}

	return templates, nil
}


func (r *Registry) GetRulesDir() string {
	return r.rulesDir
}
