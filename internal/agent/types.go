package agent

import (
	"time"
)

// AgentType represents the type of an agent
type AgentType string

// AgentContextReader defines read-only operations for agent context
type AgentContextReader interface {
	GetAgentID() string
	GetAgentType() string
	GetAgentVersion() string
	GetLastExecution() time.Time
	GetExecutionCount() int
	GetErrorCount() int
	GetMetadata(key string) (interface{}, bool)
	Get(key ContextKey) (interface{}, bool)
	GetString(key ContextKey) (string, error)
	GetInt(key ContextKey) (int, error)
	GetData() map[string]interface{}
	GetLastUpdated() time.Time
}

// AgentContextWriter defines write operations for agent context
type AgentContextWriter interface {
	IncrementExecutionCount()
	IncrementErrorCount()
	SetMetadata(key string, value interface{})
	Set(key ContextKey, value interface{}) error
	SetData(data map[string]interface{})
}

// AgentContext combines read and write operations
type AgentContext interface {
	AgentContextReader
	AgentContextWriter
}

// AgentMetadata represents metadata about an agent
type AgentMetadata struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Version     string                 `json:"version"`
	Description string                 `json:"description"`
	Author      string                 `json:"author"`
	Tags        []string               `json:"tags"`
	Properties  map[string]interface{} `json:"properties"`
	LastUpdated time.Time              `json:"last_updated"`
}

// AgentConfig represents configuration for an agent
type AgentConfig struct {
	Metadata  AgentMetadata          `json:"metadata"`
	Settings  map[string]interface{} `json:"settings"`
	Templates []string               `json:"templates"`
}

// AgentRegistry represents a registry of available agents
type AgentRegistry interface {
	LoadAgents() error
	GetAgent(id string) (*AgentConfig, error)
	ListAgents() []*AgentConfig
	AddAgent(config *AgentConfig) error
	RemoveAgent(id string) error
}

// AgentDefinition represents a single agent's definition and metadata
type AgentDefinition struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	Version        string                 `json:"version"`
	Type           string                 `json:"type"`
	Config         map[string]interface{} `json:"config"`
	Templates      []string               `json:"templates"`
	LastUpdated    time.Time              `json:"last_updated"`
	Content        string                 `json:"content,omitempty"`
	DefinitionPath string                 `json:"definition_path,omitempty"`
}

// Agent represents a loaded and initialized agent
type Agent struct {
	Definition *AgentDefinition
	Context    AgentContext
}

// ContextData represents the serializable part of agent context
type ContextData struct {
	AgentID        string                 `json:"agent_id"`
	AgentType      string                 `json:"agent_type"`
	AgentVersion   string                 `json:"agent_version"`
	LastExecution  time.Time              `json:"last_execution"`
	ExecutionCount int                    `json:"execution_count"`
	ErrorCount     int                    `json:"error_count"`
	Metadata       map[string]interface{} `json:"metadata"`
	CustomData     map[string]interface{} `json:"custom_data"`
	Data           map[string]interface{} `json:"data"`
	LastUpdated    time.Time              `json:"last_updated"`
}

// AgentManager interface defines methods for agent management
type AgentManager interface {
	ListAgents() ([]*AgentDefinition, error)
	GetAgent(id string) (*AgentDefinition, bool)
	LoadAgent(id string) (*Agent, error)
}
