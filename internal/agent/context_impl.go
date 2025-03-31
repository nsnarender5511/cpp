package agent

import (
	"encoding/json"
	"fmt"
	"time"
)

// AgentContextImpl provides a concrete implementation of the AgentContext interface
type AgentContextImpl struct {
	agentID        string
	agentType      string
	agentVersion   string
	lastExecution  time.Time
	executionCount int
	errorCount     int
	metadata       map[string]interface{}
	customData     map[ContextKey]interface{}
	data           map[string]interface{}
	lastUpdated    time.Time
}

// CreateAgentContext creates a new agent context with optional data
func CreateAgentContext(agentID string, agentType string, agentVersion string, data map[string]interface{}) AgentContext {
	ctx := &AgentContextImpl{
		agentID:      agentID,
		agentType:    agentType,
		agentVersion: agentVersion,
		metadata:     make(map[string]interface{}),
		customData:   make(map[ContextKey]interface{}),
		data:         make(map[string]interface{}),
		errorCount:   0,
	}

	if data != nil {
		ctx.data = data
	}

	return ctx
}

// NewAgentContext creates a new empty agent context
func NewAgentContext() AgentContext {
	return CreateAgentContext("", "", "", make(map[string]interface{}))
}

// GetAgentID returns the agent ID
func (c *AgentContextImpl) GetAgentID() string {
	return c.agentID
}

// GetAgentType returns the agent type
func (c *AgentContextImpl) GetAgentType() string {
	return c.agentType
}

// GetAgentVersion returns the agent version
func (c *AgentContextImpl) GetAgentVersion() string {
	return c.agentVersion
}

// GetLastExecution returns the last execution time
func (c *AgentContextImpl) GetLastExecution() time.Time {
	return c.lastExecution
}

// GetExecutionCount returns the execution count
func (c *AgentContextImpl) GetExecutionCount() int {
	return c.executionCount
}

// GetErrorCount returns the error count
func (c *AgentContextImpl) GetErrorCount() int {
	return c.errorCount
}

// IncrementExecutionCount increments the execution counter
func (c *AgentContextImpl) IncrementExecutionCount() {
	c.executionCount++
	c.lastExecution = time.Now()
}

// IncrementErrorCount increments the error counter
func (c *AgentContextImpl) IncrementErrorCount() {
	c.errorCount++
}

// SetMetadata sets a metadata value
func (c *AgentContextImpl) SetMetadata(key string, value interface{}) {
	c.metadata[key] = value
}

// GetMetadata retrieves a metadata value
func (c *AgentContextImpl) GetMetadata(key string) (interface{}, bool) {
	value, exists := c.metadata[key]
	return value, exists
}

// Set updates a value in the context using a strongly typed key
func (c *AgentContextImpl) Set(key ContextKey, value interface{}) error {
	switch key {
	case KeyAgentID:
		if str, ok := value.(string); ok {
			c.agentID = str
			return nil
		}
	case KeyAgentType:
		if str, ok := value.(string); ok {
			c.agentType = str
			return nil
		}
	case KeyAgentVersion:
		if str, ok := value.(string); ok {
			c.agentVersion = str
			return nil
		}
	case KeyLastExecution:
		if t, ok := value.(time.Time); ok {
			c.lastExecution = t
			return nil
		}
	case KeyExecutionCount:
		if count, ok := value.(int); ok {
			c.executionCount = count
			return nil
		}
	case KeyErrorCount:
		if count, ok := value.(int); ok {
			c.errorCount = count
			return nil
		}
	case KeyMetadata:
		if meta, ok := value.(map[string]interface{}); ok {
			c.metadata = meta
			return nil
		}
	default:
		c.customData[key] = value
		return nil
	}
	return fmt.Errorf("invalid type for key %s", key)
}

// Get retrieves a value from the context using a strongly typed key
func (c *AgentContextImpl) Get(key ContextKey) (interface{}, bool) {
	switch key {
	case KeyAgentID:
		return c.agentID, true
	case KeyAgentType:
		return c.agentType, true
	case KeyAgentVersion:
		return c.agentVersion, true
	case KeyLastExecution:
		return c.lastExecution, true
	case KeyExecutionCount:
		return c.executionCount, true
	case KeyErrorCount:
		return c.errorCount, true
	case KeyMetadata:
		return c.metadata, true
	default:
		val, ok := c.customData[key]
		return val, ok
	}
}

// GetString retrieves a string value from context
func (c *AgentContextImpl) GetString(key ContextKey) (string, error) {
	value, exists := c.customData[key]
	if !exists {
		return "", fmt.Errorf("key not found: %s", key)
	}
	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("value for key %s is not a string", key)
	}
	return str, nil
}

// GetInt retrieves an integer value from context
func (c *AgentContextImpl) GetInt(key ContextKey) (int, error) {
	value, exists := c.customData[key]
	if !exists {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	num, ok := value.(int)
	if !ok {
		return 0, fmt.Errorf("value for key %s is not an integer", key)
	}
	return num, nil
}

// ToContextData converts the context to serializable data
func (c *AgentContextImpl) ToContextData() ContextData {
	return ContextData{
		AgentID:        c.agentID,
		AgentType:      c.agentType,
		AgentVersion:   c.agentVersion,
		LastExecution:  c.lastExecution,
		ExecutionCount: c.executionCount,
		ErrorCount:     c.errorCount,
		Metadata:       c.metadata,
		CustomData:     serializeCustomData(c.customData),
		Data:           c.data,
		LastUpdated:    c.lastUpdated,
	}
}

// FromData loads context data from serializable format
func (c *AgentContextImpl) FromData(data *ContextData) {
	c.agentID = data.AgentID
	c.agentType = data.AgentType
	c.agentVersion = data.AgentVersion
	c.lastExecution = data.LastExecution
	c.executionCount = data.ExecutionCount
	c.errorCount = data.ErrorCount
	c.metadata = data.Metadata
	c.customData = deserializeCustomData(data.CustomData)
	c.data = data.Data
	c.lastUpdated = data.LastUpdated
}

func serializeCustomData(data map[ContextKey]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		result[string(k)] = v
	}
	return result
}

func deserializeCustomData(data map[string]interface{}) map[ContextKey]interface{} {
	result := make(map[ContextKey]interface{})
	for k, v := range data {
		result[ContextKey(k)] = v
	}
	return result
}

// MarshalJSON implements json.Marshaler
func (c *AgentContextImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToContextData())
}

// UnmarshalJSON implements json.Unmarshaler
func (c *AgentContextImpl) UnmarshalJSON(data []byte) error {
	var contextData ContextData
	if err := json.Unmarshal(data, &contextData); err != nil {
		return err
	}
	c.FromData(&contextData)
	return nil
}

// GetData returns the legacy data map
func (c *AgentContextImpl) GetData() map[string]interface{} {
	return c.data
}

// SetData updates the legacy data map
func (c *AgentContextImpl) SetData(data map[string]interface{}) {
	c.data = data
	c.lastUpdated = time.Now()
}

// GetLastUpdated returns the last update time
func (c *AgentContextImpl) GetLastUpdated() time.Time {
	return c.lastUpdated
}
