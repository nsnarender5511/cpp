package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)


type ContextPersistence struct {
	contextDir string
}


func NewContextPersistence(contextDir string) *ContextPersistence {
	return &ContextPersistence{
		contextDir: contextDir,
	}
}


func (p *ContextPersistence) SaveContext(ctx AgentContext) error {
	if ctx == nil {
		return fmt.Errorf("cannot save nil context")
	}

	
	contextImpl, ok := ctx.(*AgentContextImpl)
	if !ok {
		return fmt.Errorf("context is not of type *AgentContextImpl")
	}

	data := contextImpl.ToContextData()
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal context data: %w", err)
	}

	if err := os.MkdirAll(p.contextDir, 0755); err != nil {
		return fmt.Errorf("failed to create context directory: %w", err)
	}

	filename := filepath.Join(p.contextDir, fmt.Sprintf("%s.json", ctx.GetAgentID()))
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write context file: %w", err)
	}

	return nil
}


func (p *ContextPersistence) LoadContext(agentID string) (AgentContext, error) {
	filename := filepath.Join(p.contextDir, fmt.Sprintf("%s.json", agentID))

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil 
		}
		return nil, fmt.Errorf("failed to read context file: %w", err)
	}

	var data ContextData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal context data: %w", err)
	}

	
	ctx := &AgentContextImpl{}
	ctx.FromData(&data)
	return ctx, nil
}
