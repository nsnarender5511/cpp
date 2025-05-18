package agent

import (
	"time"
)


type AgentType string


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


type AgentContextWriter interface {
	IncrementExecutionCount()
	IncrementErrorCount()
	SetMetadata(key string, value interface{})
	Set(key ContextKey, value interface{}) error
	SetData(data map[string]interface{})
}


type AgentContext interface {
	AgentContextReader
	AgentContextWriter
}


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


type AgentConfig struct {
	Metadata  AgentMetadata          `json:"metadata"`
	Settings  map[string]interface{} `json:"settings"`
	Templates []string               `json:"templates"`
}


type AgentRegistry interface {
	LoadAgents() error
	GetAgent(id string) (*AgentConfig, error)
	ListAgents() []*AgentConfig
	AddAgent(config *AgentConfig) error
	RemoveAgent(id string) error
}


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


type Agent struct {
	Definition *AgentDefinition
	Context    AgentContext
}


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


type AgentManager interface {
	ListAgents() ([]*AgentDefinition, error)
	GetAgent(id string) (*AgentDefinition, bool)
	LoadAgent(id string) (*Agent, error)
}
