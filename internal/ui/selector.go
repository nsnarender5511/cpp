package ui

import (
	"fmt"
	"strings"

	"crules/internal/agent"
)

// AgentSelector provides an interactive terminal-based UI for agent selection
type AgentSelector struct {
	agents []*agent.AgentDefinition
}

// NewAgentSelector creates a new agent selector
func NewAgentSelector(agents []*agent.AgentDefinition) *AgentSelector {
	return &AgentSelector{
		agents: agents,
	}
}

// Run displays the interactive selection menu and returns the selected agent
func (s *AgentSelector) Run() (*agent.AgentDefinition, error) {
	if len(s.agents) == 0 {
		Error("No agents available for selection")
		return nil, fmt.Errorf("no agents available")
	}

	Header("Select an Agent:")

	// Display all agents with numbers
	for i, agent := range s.agents {
		// Display the agent name
		Plain("  %d. %s", i+1, agent.Name)

		// Display a truncated description if available
		if agent.Description != "" {
			// Get first sentence or truncate to 80 chars
			shortDesc := truncateDescription(agent.Description, 80)
			Plain("     %s", shortDesc)
		}
		Plain("")
	}

	// Get user selection
	selection := 0
	maxAttempts := 3
	attempts := 0

	for attempts < maxAttempts {
		Prompt("Enter agent number (1-%d): ", len(s.agents))

		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			Warning("Invalid input. Please enter a number.")
			attempts++
			continue
		}

		if input < 1 || input > len(s.agents) {
			Warning("Invalid selection. Please enter a number between 1 and %d.", len(s.agents))
			attempts++
			continue
		}

		selection = input - 1 // Convert to 0-based index
		break
	}

	if attempts == maxAttempts {
		Error("Maximum number of attempts reached. Exiting.")
		return nil, fmt.Errorf("maximum number of attempts reached")
	}

	selectedAgent := s.agents[selection]
	Success("Selected agent: %s", selectedAgent.Name)

	// Show agent details
	Header("Agent Details:")
	Plain("  ID: %s", selectedAgent.ID)
	Plain("  Name: %s", selectedAgent.Name)
	Plain("  Version: %s", selectedAgent.Version)
	Plain("  Description:")

	// Display full description with proper line wrapping
	if selectedAgent.Description != "" {
		// Split description into paragraphs and display with indentation
		paragraphs := strings.Split(selectedAgent.Description, "\n\n")
		for _, paragraph := range paragraphs {
			// Wrap long paragraphs to fit terminal (assuming 80 columns)
			wrappedText := wrapText(paragraph, 70, "    ")
			Plain("%s", wrappedText)
			Plain("") // Empty line between paragraphs
		}
	} else {
		Plain("    No description available.")
	}

	if len(selectedAgent.Capabilities) > 0 {
		Plain("\n  Capabilities:")
		for _, capability := range selectedAgent.Capabilities {
			Plain("    - %s", capability)
		}
	}

	return selectedAgent, nil
}

// truncateDescription returns a shortened version of the description
func truncateDescription(description string, maxLength int) string {
	// Find first sentence
	endIdx := strings.Index(description, ".")
	if endIdx > 0 && endIdx < maxLength {
		return description[:endIdx+1]
	}

	// Or truncate to maxLength
	if len(description) > maxLength {
		return description[:maxLength] + "..."
	}

	return description
}

// wrapText wraps text to fit within the specified width with given prefix
func wrapText(text string, width int, prefix string) string {
	if text == "" {
		return prefix
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return prefix
	}

	var result strings.Builder
	result.WriteString(prefix)

	lineLength := len(prefix)
	for i, word := range words {
		wordLength := len(word)

		if i > 0 {
			// Check if adding this word would exceed the width
			if lineLength+wordLength+1 > width {
				// Start a new line
				result.WriteString("\n")
				result.WriteString(prefix)
				lineLength = len(prefix)
			} else {
				// Add a space before the word
				result.WriteString(" ")
				lineLength++
			}
		}

		result.WriteString(word)
		lineLength += wordLength
	}

	return result.String()
}
