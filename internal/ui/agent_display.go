package ui

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"vibe/internal/agent"
)

// AgentDisplayOptions contains configuration options for displaying agents
type AgentDisplayOptions struct {
	TermWidth       int
	GroupByCategory bool
	CompactMode     bool
	SelectedAgentID string
}

// DefaultAgentDisplayOptions returns default display options
func DefaultAgentDisplayOptions() AgentDisplayOptions {
	width := GetTerminalWidth()
	if width <= 0 {
		width = 80 // Default fallback
	}

	return AgentDisplayOptions{
		TermWidth:       width,
		GroupByCategory: true,
		CompactMode:     width < 100,
		SelectedAgentID: "",
	}
}

// GetTerminalWidth returns the width of the terminal
func GetTerminalWidth() int {
	// Try using the stty command to get terminal size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err == nil {
		parts := strings.Split(strings.TrimSpace(string(out)), " ")
		if len(parts) == 2 {
			width, err := strconv.Atoi(parts[1])
			if err == nil && width > 0 {
				return width
			}
		}
	}

	// If everything fails, return a default width
	return 80
}

// detectAgentCategory attempts to determine the category of an agent
func detectAgentCategory(agent *agent.AgentDefinition) string {
	id := strings.ToLower(agent.ID)

	// Check ID for category hints
	if strings.Contains(id, "review") {
		return "Code Review"
	} else if strings.Contains(id, "test") {
		return "Testing"
	} else if strings.Contains(id, "doc") {
		return "Documentation"
	} else if strings.Contains(id, "git") || strings.Contains(id, "commit") {
		return "Git & Version Control"
	} else if strings.Contains(id, "refactor") {
		return "Refactoring"
	} else if strings.Contains(id, "architect") || strings.Contains(id, "plan") {
		return "Architecture & Planning"
	} else if strings.Contains(id, "debug") || strings.Contains(id, "fix") {
		return "Debugging & Fixing"
	} else if strings.Contains(id, "web") {
		return "Web Tools"
	} else if strings.Contains(id, "answer") || strings.Contains(id, "quick") {
		return "Quick Help"
	}

	// Check description as backup
	if agent.Description != "" {
		descLower := strings.ToLower(agent.Description)
		if strings.Contains(descLower, "code review") {
			return "Code Review"
		} else if strings.Contains(descLower, "test") {
			return "Testing"
		} else if strings.Contains(descLower, "document") {
			return "Documentation"
		} else if strings.Contains(descLower, "git") || strings.Contains(descLower, "commit") {
			return "Git & Version Control"
		}
	}

	// Default category
	return "General"
}

// DisplayAgentListEnhanced displays the list of agents with enhanced formatting
func DisplayAgentListEnhanced(agents []*agent.AgentDefinition, options AgentDisplayOptions) error {
	if len(agents) == 0 {
		Warning("No agents found")
		return nil
	}

	fmt.Println()
	Header("Available Agents")

	if options.GroupByCategory {
		// Group agents by category
		categories := make(map[string][]*agent.AgentDefinition)
		for _, a := range agents {
			category := detectAgentCategory(a)
			categories[category] = append(categories[category], a)
		}

		// Sort categories
		categoryNames := make([]string, 0, len(categories))
		for category := range categories {
			categoryNames = append(categoryNames, category)
		}
		sort.Strings(categoryNames)

		// Print agents by category
		for _, category := range categoryNames {
			groupAgents := categories[category]

			// Skip empty categories
			if len(groupAgents) == 0 {
				continue
			}

			// Print category header
			fmt.Println()
			Header(category)

			// Sort agents in this category by name
			sort.Slice(groupAgents, func(i, j int) bool {
				return groupAgents[i].Name < groupAgents[j].Name
			})

			// Print agents
			if options.CompactMode {
				displayCompactAgentGroup(groupAgents, options)
			} else {
				displayDetailedAgentGroup(groupAgents, options)
			}
		}
	} else {
		// Sort agents by ID
		sort.Slice(agents, func(i, j int) bool {
			return agents[i].ID < agents[j].ID
		})

		// Print all agents without categories
		if options.CompactMode {
			displayCompactAgentGroup(agents, options)
		} else {
			displayDetailedAgentGroup(agents, options)
		}
	}

	fmt.Println()
	Info("To select an agent, use: vibe agent select")
	Info("To get more info about an agent, use: vibe agent info <agent_id>")
	fmt.Println()

	return nil
}

// displayCompactAgentGroup displays a group of agents in compact form
func displayCompactAgentGroup(agents []*agent.AgentDefinition, options AgentDisplayOptions) {
	for i, a := range agents {
		prefix := "  "
		if a.ID == options.SelectedAgentID {
			prefix = "• "
		}

		// Use ID as selector index
		idStr := a.ID

		// Format name with optional version
		nameStr := a.Name
		if a.Version != "" && a.Version != "1.0" {
			nameStr = fmt.Sprintf("%s (%s)", a.Name, a.Version)
		}

		// Print in compact format
		fmt.Printf("%s%-20s %s\n", prefix, idStr, nameStr)

		// Add separator except after last item
		if i < len(agents)-1 {
			fmt.Println()
		}
	}
}

// displayDetailedAgentGroup displays a group of agents with detailed information
func displayDetailedAgentGroup(agents []*agent.AgentDefinition, options AgentDisplayOptions) {
	for i, a := range agents {
		prefix := "  "
		if a.ID == options.SelectedAgentID {
			prefix = "• "
		}

		// Use ID as selector
		idStr := a.ID

		// Format name with optional version
		nameStr := a.Name
		if a.Version != "" && a.Version != "1.0" {
			nameStr = fmt.Sprintf("%s (%s)", a.Name, a.Version)
		}

		// Print detailed format
		fmt.Printf("%s%-20s %s\n", prefix, idStr, nameStr)

		// Print a short description if available
		if a.Description != "" {
			shortDesc := truncateText(a.Description, options.TermWidth-30)
			fmt.Printf("    %s\n", shortDesc)
		}

		// Print tags if available
		if len(a.Templates) > 0 {
			fmt.Printf("    Templates: %s\n", strings.Join(a.Templates, ", "))
		}

		// Add separator except after last item
		if i < len(agents)-1 {
			fmt.Println()
		}
	}
}

// truncateText shortens text to fit within maxWidth characters
func truncateText(text string, maxWidth int) string {
	if len(text) <= maxWidth {
		return text
	}

	// Return truncated string with ellipsis
	return text[:maxWidth-3] + "..."
}

// DisplayAgentInfoEnhanced shows detailed information about an agent with enhanced formatting
func DisplayAgentInfoEnhanced(agent *agent.AgentDefinition, verbose bool) error {
	if agent == nil {
		return fmt.Errorf("agent is nil")
	}

	fmt.Println()
	Header("Agent Information")
	fmt.Println()

	// Create info box
	Plain("  ID:        %s", agent.ID)
	Plain("  Name:      %s", agent.Name)

	if agent.Version != "" {
		Plain("  Version:   %s", agent.Version)
	}

	if agent.Type != "" {
		Plain("  Type:      %s", agent.Type)
	}

	Plain("  Updated:   %s", agent.LastUpdated.Format("2006-01-02 15:04:05"))
	fmt.Println()

	// Display description
	if agent.Description != "" {
		Header("Description")
		fmt.Println()
		Plain("  %s", agent.Description)
		fmt.Println()
	}

	// Display templates
	if len(agent.Templates) > 0 {
		Header("Templates")
		fmt.Println()
		for _, tmpl := range agent.Templates {
			Plain("  • %s", tmpl)
		}
		fmt.Println()
	}

	// Display configuration
	if len(agent.Config) > 0 && verbose {
		Header("Configuration")
		fmt.Println()
		for k, v := range agent.Config {
			Plain("  %s: %v", k, v)
		}
		fmt.Println()
	}

	return nil
}

// cleanPromptContent removes excessive whitespace from prompt content
// func cleanPromptContent(content string) string {
// 	// Remove repeated empty lines
// 	re := regexp.MustCompile(`\\n{3,}`)
// 	content = re.ReplaceAllString(content, "\\n\\n")
//
// 	// Trim leading/trailing whitespace
// 	return strings.TrimSpace(content)
// }
//
// // formatFileSize returns a human-readable file size
// func formatFileSize(size int64) string {
// 	const unit = 1024
// 	if size < unit {
// 		return fmt.Sprintf("%d B", size)
// 	}
// 	div, exp := int64(unit), 0
// 	for n := size / unit; n >= unit; n /= unit {
// 		div *= unit
// 		exp++
// 	}
// 	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
// }
//
// // createCodeBox creates a formatted code box with the given content
// func createCodeBox(content string, width int, scrolling bool) string {
// 	if content == "" {
// 		return ""
// 	}
//
// 	lines := strings.Split(content, "\\n")
// 	maxContentWidth := width - 6 // Account for box decorations
//
// 	// Prepare box decorations based on terminal width
// 	topBorder := "┌" + strings.Repeat("─", width-2) + "┐"
// 	bottomBorder := "└" + strings.Repeat("─", width-2) + "┘"
//
// 	var result strings.Builder
// 	result.WriteString(topBorder + "\\n")
//
// 	// Process each line
// 	for _, line := range lines {
// 		if len(line) > maxContentWidth {
// 			if scrolling {
// 				// For scrollable content, wrap the text
// 				for len(line) > 0 {
// 					if len(line) <= maxContentWidth {
// 						result.WriteString("│ " + line + strings.Repeat(" ", maxContentWidth-len(line)) + " │\\n")
// 						break
// 					}
//
// 					result.WriteString("│ " + line[:maxContentWidth] + " │\\n")
// 					line = line[maxContentWidth:]
// 				}
// 			} else {
// 				// For non-scrollable, truncate with ellipsis
// 				truncated := line[:maxContentWidth-3] + "..."
// 				result.WriteString("│ " + truncated + strings.Repeat(" ", maxContentWidth-len(truncated)) + " │\\n")
// 			}
// 		} else {
// 			padding := strings.Repeat(" ", maxContentWidth-len(line))
// 			result.WriteString("│ " + line + padding + " │\\n")
// 		}
// 	}
//
// 	result.WriteString(bottomBorder)
// 	return result.String()
// }
