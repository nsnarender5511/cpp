package agent

import (
	"context"
	"fmt"

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

// LoadAgent loads an agent by ID
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

	// Create agent context
	context := CreateAgentContext(definition.ID, definition.Type, definition.Version, nil)

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
func (l *Loader) LoadAgentWithContext(id string, agentCtx AgentContext) (*Agent, error) {
	return l.LoadAgentWithContextCancellation(context.Background(), id, agentCtx)
}

// LoadAgentWithContextCancellation loads an agent with both context awareness and cancellation support
func (l *Loader) LoadAgentWithContextCancellation(ctx context.Context, id string, agentCtx AgentContext) (*Agent, error) {
	utils.Debug("Loading agent with context | id=" + id)
	l.reportProgress("load_context_start", id)

	// Check for context cancellation
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent loading canceled: %w", ctx.Err())
	default:
		// Continue if not canceled
	}

	// Get agent definition from registry
	definition, err := l.registry.GetAgent(id)
	if err != nil {
		utils.Error("Failed to get agent | error=" + err.Error())
		l.reportProgress("load_error", "Failed to get agent "+id+": "+err.Error())
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}

	// Check for context cancellation again
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent loading canceled: %w", ctx.Err())
	default:
		// Continue if not canceled
	}

	// Initialize agent with definition and provided context
	agent := &Agent{
		Definition: definition,
		Context:    agentCtx,
	}

	utils.Info("Agent loaded successfully with context | id=" + id + ", name=" + definition.Name)
	l.reportProgress("load_success", id)
	return agent, nil
}
