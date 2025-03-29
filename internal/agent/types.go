package agent

import (
	"time"
)

// AgentDefinition represents a single agent's definition and metadata
type AgentDefinition struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Capabilities   []string `json:"capabilities"`
	Version        string   `json:"version"`
	DefinitionPath string   `json:"-"` // Path to the .mdc file
	Content        string   `json:"-"` // The actual content of the agent definition
}

// Agent represents a loaded and initialized agent
type Agent struct {
	Definition *AgentDefinition
	Context    *AgentContext
}

// AgentContext represents shared context between agents
type AgentContext struct {
	Data        map[string]interface{}
	LastUpdated time.Time
}

// NewAgentContext creates a new agent context
func NewAgentContext() *AgentContext {
	return &AgentContext{
		Data:        make(map[string]interface{}),
		LastUpdated: time.Now(),
	}
}

// Set stores a value in the context
func (c *AgentContext) Set(key string, value interface{}) {
	c.Data[key] = value
	c.LastUpdated = time.Now()
}

// Get retrieves a value from the context
func (c *AgentContext) Get(key string) (interface{}, bool) {
	value, exists := c.Data[key]
	return value, exists
}

// AgentManager interface defines methods for agent management
type AgentManager interface {
	ListAgents() ([]*AgentDefinition, error)
	GetAgent(id string) (*AgentDefinition, bool)
	LoadAgent(id string) (*Agent, error)
}
