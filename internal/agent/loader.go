package agent

import (
	"fmt"
	"os"

	"crules/internal/utils"
)

// Loader handles loading and initializing agents
type Loader struct {
	registry *Registry
	config   *utils.Config
}

// NewLoader creates a new agent loader
func NewLoader(registry *Registry, config *utils.Config) *Loader {
	return &Loader{
		registry: registry,
		config:   config,
	}
}

// LoadAgent loads and initializes an agent by ID
func (l *Loader) LoadAgent(id string) (*Agent, error) {
	utils.Debug("Loading agent | id=" + id)

	// Get agent definition from registry
	definition, exists := l.registry.GetAgent(id)
	if !exists {
		utils.Error("Agent not found | id=" + id)
		return nil, fmt.Errorf("agent '%s' not found", id)
	}

	// Read agent definition content if not already loaded
	if definition.Content == "" {
		utils.Debug("Reading agent definition content | path=" + definition.DefinitionPath)
		content, err := os.ReadFile(definition.DefinitionPath)
		if err != nil {
			utils.Error("Failed to read agent definition | path=" + definition.DefinitionPath + ", error=" + err.Error())
			return nil, fmt.Errorf("failed to read agent definition: %v", err)
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
	return agent, nil
}

// LoadAgentWithContext loads an agent and initializes it with the provided context
func (l *Loader) LoadAgentWithContext(id string, context *AgentContext) (*Agent, error) {
	utils.Debug("Loading agent with context | id=" + id)

	// Get agent definition from registry
	definition, exists := l.registry.GetAgent(id)
	if !exists {
		utils.Error("Agent not found | id=" + id)
		return nil, fmt.Errorf("agent '%s' not found", id)
	}

	// Read agent definition content if not already loaded
	if definition.Content == "" {
		utils.Debug("Reading agent definition content | path=" + definition.DefinitionPath)
		content, err := os.ReadFile(definition.DefinitionPath)
		if err != nil {
			utils.Error("Failed to read agent definition | path=" + definition.DefinitionPath + ", error=" + err.Error())
			return nil, fmt.Errorf("failed to read agent definition: %v", err)
		}
		definition.Content = string(content)
	}

	// Initialize agent with definition and provided context
	agent := &Agent{
		Definition: definition,
		Context:    context,
	}

	utils.Info("Agent loaded successfully with context | id=" + id + ", name=" + definition.Name)
	return agent, nil
}
