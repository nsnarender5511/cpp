package agent

import (
	"context"
	"fmt"

	"vibe/internal/utils"
)


type Loader struct {
	registry     *Registry
	config       *utils.Config
	progressFunc func(event string, message string)
}


func NewLoader(registry *Registry, config *utils.Config) *Loader {
	return &Loader{
		registry:     registry,
		config:       config,
		progressFunc: nil,
	}
}


func (l *Loader) SetProgressCallback(progressFunc func(event string, message string)) {
	l.progressFunc = progressFunc
}


func (l *Loader) reportProgress(event string, message string) {
	if l.progressFunc != nil {
		l.progressFunc(event, message)
	}
}


func (l *Loader) LoadAgent(id string) (*Agent, error) {
	utils.Debug("Loading agent | id=" + id)
	l.reportProgress("load_start", id)

	
	definition, err := l.registry.GetAgent(id)
	if err != nil {
		utils.Error("Failed to get agent | error=" + err.Error())
		l.reportProgress("load_error", "Failed to get agent "+id+": "+err.Error())
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}

	
	context := CreateAgentContext(definition.ID, definition.Type, definition.Version, nil)

	
	agent := &Agent{
		Definition: definition,
		Context:    context,
	}

	utils.Info("Agent loaded successfully | id=" + id + ", name=" + definition.Name)
	l.reportProgress("load_success", id)
	return agent, nil
}


func (l *Loader) LoadAgentWithContext(id string, agentCtx AgentContext) (*Agent, error) {
	return l.LoadAgentWithContextCancellation(context.Background(), id, agentCtx)
}


func (l *Loader) LoadAgentWithContextCancellation(ctx context.Context, id string, agentCtx AgentContext) (*Agent, error) {
	utils.Debug("Loading agent with context | id=" + id)
	l.reportProgress("load_context_start", id)

	
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent loading canceled: %w", ctx.Err())
	default:
		
	}

	
	definition, err := l.registry.GetAgent(id)
	if err != nil {
		utils.Error("Failed to get agent | error=" + err.Error())
		l.reportProgress("load_error", "Failed to get agent "+id+": "+err.Error())
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}

	
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent loading canceled: %w", ctx.Err())
	default:
		
	}

	
	agent := &Agent{
		Definition: definition,
		Context:    agentCtx,
	}

	utils.Info("Agent loaded successfully with context | id=" + id + ", name=" + definition.Name)
	l.reportProgress("load_success", id)
	return agent, nil
}
