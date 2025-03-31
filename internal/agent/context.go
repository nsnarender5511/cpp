package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"cursor++/internal/utils"
)

// ContextData represents the serializable part of agent context
type ContextData struct {
	Data        map[string]interface{} `json:"data"`
	LastUpdated time.Time              `json:"last_updated"`
}

// ContextPersistence handles saving and loading agent context
type ContextPersistence struct {
	config   *utils.Config
	appPaths utils.AppPaths
}

// NewContextPersistence creates a new context persistence handler
func NewContextPersistence(config *utils.Config, appPaths utils.AppPaths) *ContextPersistence {
	return &ContextPersistence{
		config:   config,
		appPaths: appPaths,
	}
}

// SaveContext saves the agent context to disk
func (p *ContextPersistence) SaveContext(context *AgentContext) error {
	utils.Debug("Saving agent context")

	// Create data to save
	data := &ContextData{
		Data:        context.Data,
		LastUpdated: context.LastUpdated,
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		utils.Error("Failed to marshal context data | error=" + err.Error())
		return fmt.Errorf("failed to marshal context data: %v", err)
	}

	// Create context file path
	contextDir := filepath.Join(p.appPaths.DataDir, "agent_context")
	if err := os.MkdirAll(contextDir, p.config.DirPermission); err != nil {
		utils.Error("Failed to create context directory | error=" + err.Error())
		return fmt.Errorf("failed to create context directory: %v", err)
	}

	contextPath := filepath.Join(contextDir, "context.json")
	utils.Debug("Saving context to file | path=" + contextPath)

	// Write to file
	if err := os.WriteFile(contextPath, jsonData, p.config.FilePermission); err != nil {
		utils.Error("Failed to write context file | error=" + err.Error())
		return fmt.Errorf("failed to write context file: %v", err)
	}

	utils.Info("Agent context saved successfully | path=" + contextPath)
	return nil
}

// LoadContext loads the agent context from disk
func (p *ContextPersistence) LoadContext() (*AgentContext, error) {
	utils.Debug("Loading agent context")

	// Create context file path
	contextDir := filepath.Join(p.appPaths.DataDir, "agent_context")
	contextPath := filepath.Join(contextDir, "context.json")

	// Check if file exists
	if _, err := os.Stat(contextPath); os.IsNotExist(err) {
		utils.Debug("Context file does not exist, creating new context | path=" + contextPath)
		return NewAgentContext(), nil
	}

	// Read file
	jsonData, err := os.ReadFile(contextPath)
	if err != nil {
		utils.Error("Failed to read context file | error=" + err.Error())
		return nil, fmt.Errorf("failed to read context file: %v", err)
	}

	// Parse JSON
	var data ContextData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		utils.Error("Failed to parse context data | error=" + err.Error())
		return nil, fmt.Errorf("failed to parse context data: %v", err)
	}

	// Create context
	context := &AgentContext{
		Data:        data.Data,
		LastUpdated: data.LastUpdated,
	}

	utils.Info("Agent context loaded successfully | path=" + contextPath)
	return context, nil
}
