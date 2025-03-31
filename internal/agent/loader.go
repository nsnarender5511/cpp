package agent

import (
	"fmt"
	"os"

	"cursor++/internal/utils"
)

// Loader handles loading and initializing agents
type Loader struct {
	registry     *Registry
	config       *utils.Config
	progressFunc func(event string, message string)
}

// NewLoader creates a new agent loader
func NewLoader(registry *Registry, config *utils.Config) *Loader {
	return &Loader{
		registry:     registry,
		config:       config,
		progressFunc: nil,
	}
}

// SetProgressCallback sets a callback function to report progress during agent loading
func (l *Loader) SetProgressCallback(progressFunc func(event string, message string)) {
	l.progressFunc = progressFunc
}

// reportProgress reports progress to the callback if available
func (l *Loader) reportProgress(event string, message string) {
	if l.progressFunc != nil {
		l.progressFunc(event, message)
	}
}

// LoadAgent loads and initializes an agent by ID
func (l *Loader) LoadAgent(id string) (*Agent, error) {
	utils.Debug("Loading agent | id=" + id)
	l.reportProgress("load_start", id)

	// Get agent definition from registry
	definition, err := l.registry.GetAgent(id)
	if err != nil {
		utils.Error("Failed to get agent | error=" + err.Error())
		l.reportProgress("load_error", "Failed to get agent "+id+": "+err.Error())
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}

	// Read agent definition content if not already loaded
	if definition.Content == "" {
		utils.Debug("Reading agent definition content | path=" + definition.DefinitionPath)
		l.reportProgress("reading_content", id)
		content, err := os.ReadFile(definition.DefinitionPath)
		if err != nil {
			utils.Error("Failed to read agent definition | path=" + definition.DefinitionPath +
				", error=" + err.Error())
			l.reportProgress("read_error", "Failed to read agent "+id+": "+err.Error())
			return nil, fmt.Errorf("failed to read agent definition: %w", err)
		}
		definition.Content = string(content)
	}

	// Create agent context
	context := NewAgentContext()

	// Initialize agent with definition and context
	agent := &Agent{
		Definition: definition,
		Context:    context,
	}

	utils.Info("Agent loaded successfully | id=" + id + ", name=" + definition.Name)
	l.reportProgress("load_success", id)
	return agent, nil
}

// LoadAgentWithContext loads an agent and initializes it with the provided context
func (l *Loader) LoadAgentWithContext(id string, context *AgentContext) (*Agent, error) {
	utils.Debug("Loading agent with context | id=" + id)
	l.reportProgress("load_context_start", id)

	// Get agent definition from registry
	definition, err := l.registry.GetAgent(id)
	if err != nil {
		utils.Error("Failed to get agent | error=" + err.Error())
		l.reportProgress("load_error", "Failed to get agent "+id+": "+err.Error())
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}

	// Read agent definition content if not already loaded
	if definition.Content == "" {
		utils.Debug("Reading agent definition content | path=" + definition.DefinitionPath)
		l.reportProgress("reading_content", id)
		content, err := os.ReadFile(definition.DefinitionPath)
		if err != nil {
			utils.Error("Failed to read agent definition | path=" + definition.DefinitionPath +
				", error=" + err.Error())
			l.reportProgress("read_error", "Failed to read agent "+id+": "+err.Error())
			return nil, fmt.Errorf("failed to read agent definition: %w", err)
		}
		definition.Content = string(content)
	}

	// Initialize agent with definition and provided context
	agent := &Agent{
		Definition: definition,
		Context:    context,
	}

	utils.Info("Agent loaded successfully with context | id=" + id + ", name=" + definition.Name)
	l.reportProgress("load_success", id)
	return agent, nil
}
