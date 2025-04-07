package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"vibe/internal/agent"
)

// AgentSelector is a CLI for selecting agents
type AgentSelector struct {
	agents []*agent.AgentDefinition
}

// NewAgentSelector creates a new selector
func NewAgentSelector(agents []*agent.AgentDefinition) *AgentSelector {
	return &AgentSelector{
		agents: agents,
	}
}

// Run starts the selector CLI and returns the selected agent or error
func (s *AgentSelector) Run() (*agent.AgentDefinition, error) {
	if len(s.agents) == 0 {
		return nil, fmt.Errorf("no agents available")
	}

	// Display agents with numbers
	fmt.Println()
	Header("Select an Agent")
	fmt.Println()

	// Group agents by category for better organization
	categories := make(map[string][]*agent.AgentDefinition)
	for _, a := range s.agents {
		category := detectAgentCategory(a)
		categories[category] = append(categories[category], a)
	}

	// Create a map of index to agent for selection
	indexToAgent := make(map[int]*agent.AgentDefinition)
	index := 1

	// Display agents by category
	for category, agents := range categories {
		// Skip empty categories
		if len(agents) == 0 {
			continue
		}

		// Print category header
		fmt.Println()
		Header(category)

		// Display agents in this category
		for _, agent := range agents {
			// Store the agent at this index
			indexToAgent[index] = agent

			// Format name with optional version
			nameStr := agent.Name
			if agent.Version != "" && agent.Version != "1.0" {
				nameStr = fmt.Sprintf("%s (%s)", agent.Name, agent.Version)
			}

			// Print agent with index
			fmt.Printf(" %2d. %-20s %s\n", index, agent.ID, nameStr)

			// Add description if available
			if agent.Description != "" {
				// Use the truncateText function from agent_display.go
				shortDesc := truncateText(agent.Description, 70)
				fmt.Printf("     %s\n", shortDesc)
			}

			index++
		}
	}

	// Prompt for selection
	fmt.Println()
	Prompt("Enter agent number (1-%d): ", index-1)

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	// Trim whitespace
	input = strings.TrimSpace(input)

	// Parse as number
	selectedIndex, err := strconv.Atoi(input)
	if err != nil {
		return nil, fmt.Errorf("invalid selection, please enter a number")
	}

	// Validate selection
	if selectedIndex < 1 || selectedIndex >= index {
		return nil, fmt.Errorf("invalid selection: %d is out of range", selectedIndex)
	}

	// Get selected agent
	selectedAgent := indexToAgent[selectedIndex]
	if selectedAgent == nil {
		return nil, fmt.Errorf("internal error: agent not found at index %d", selectedIndex)
	}

	fmt.Println()
	Success("Selected agent: %s", selectedAgent.Name)

	return selectedAgent, nil
}

// RunWithContext starts the selector CLI with context awareness and returns the selected agent or error
func (s *AgentSelector) RunWithContext(ctx context.Context) (*agent.AgentDefinition, error) {
	if len(s.agents) == 0 {
		return nil, fmt.Errorf("no agents available")
	}

	// Create a channel for user input
	resultCh := make(chan struct {
		agent *agent.AgentDefinition
		err   error
	}, 1)

	// Run the selector in a goroutine
	go func() {
		agentDef, err := s.Run()
		resultCh <- struct {
			agent *agent.AgentDefinition
			err   error
		}{agentDef, err}
	}()

	// Wait for either context cancellation or selection completion
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent selection canceled: %w", ctx.Err())
	case result := <-resultCh:
		return result.agent, result.err
	}
}
