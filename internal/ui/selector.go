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


type AgentSelector struct {
	agents []*agent.AgentDefinition
}


func NewAgentSelector(agents []*agent.AgentDefinition) *AgentSelector {
	return &AgentSelector{
		agents: agents,
	}
}


func (s *AgentSelector) Run() (*agent.AgentDefinition, error) {
	if len(s.agents) == 0 {
		return nil, fmt.Errorf("no agents available")
	}

	
	fmt.Println()
	Header("Select an Agent")
	fmt.Println()

	
	categories := make(map[string][]*agent.AgentDefinition)
	for _, a := range s.agents {
		category := detectAgentCategory(a)
		categories[category] = append(categories[category], a)
	}

	
	indexToAgent := make(map[int]*agent.AgentDefinition)
	index := 1

	
	for category, agents := range categories {
		
		if len(agents) == 0 {
			continue
		}

		
		fmt.Println()
		Header(category)

		
		for _, agent := range agents {
			
			indexToAgent[index] = agent

			
			nameStr := agent.Name
			if agent.Version != "" && agent.Version != "1.0" {
				nameStr = fmt.Sprintf("%s (%s)", agent.Name, agent.Version)
			}

			
			fmt.Printf(" %2d. %-20s %s\n", index, agent.ID, nameStr)

			
			if agent.Description != "" {
				
				shortDesc := truncateText(agent.Description, 70)
				fmt.Printf("     %s\n", shortDesc)
			}

			index++
		}
	}

	
	fmt.Println()
	Prompt("Enter agent number (1-%d): ", index-1)

	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	
	input = strings.TrimSpace(input)

	
	selectedIndex, err := strconv.Atoi(input)
	if err != nil {
		return nil, fmt.Errorf("invalid selection, please enter a number")
	}

	
	if selectedIndex < 1 || selectedIndex >= index {
		return nil, fmt.Errorf("invalid selection: %d is out of range", selectedIndex)
	}

	
	selectedAgent := indexToAgent[selectedIndex]
	if selectedAgent == nil {
		return nil, fmt.Errorf("internal error: agent not found at index %d", selectedIndex)
	}

	fmt.Println()
	Success("Selected agent: %s", selectedAgent.Name)

	return selectedAgent, nil
}


func (s *AgentSelector) RunWithContext(ctx context.Context) (*agent.AgentDefinition, error) {
	if len(s.agents) == 0 {
		return nil, fmt.Errorf("no agents available")
	}

	
	resultCh := make(chan struct {
		agent *agent.AgentDefinition
		err   error
	}, 1)

	
	go func() {
		agentDef, err := s.Run()
		resultCh <- struct {
			agent *agent.AgentDefinition
			err   error
		}{agentDef, err}
	}()

	
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("agent selection canceled: %w", ctx.Err())
	case result := <-resultCh:
		return result.agent, result.err
	}
}
