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


type AgentDisplayOptions struct {
	TermWidth       int
	GroupByCategory bool
	CompactMode     bool
	SelectedAgentID string
}


func DefaultAgentDisplayOptions() AgentDisplayOptions {
	width := GetTerminalWidth()
	if width <= 0 {
		width = 80 
	}

	return AgentDisplayOptions{
		TermWidth:       width,
		GroupByCategory: true,
		CompactMode:     width < 100,
		SelectedAgentID: "",
	}
}


func GetTerminalWidth() int {
	
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

	
	return 80
}


func detectAgentCategory(agent *agent.AgentDefinition) string {
	id := strings.ToLower(agent.ID)

	
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

	
	return "General"
}


func DisplayAgentListEnhanced(agents []*agent.AgentDefinition, options AgentDisplayOptions) error {
	if len(agents) == 0 {
		Warning("No agents found")
		return nil
	}

	fmt.Println()
	Header("Available Agents")

	if options.GroupByCategory {
		
		categories := make(map[string][]*agent.AgentDefinition)
		for _, a := range agents {
			category := detectAgentCategory(a)
			categories[category] = append(categories[category], a)
		}

		
		categoryNames := make([]string, 0, len(categories))
		for category := range categories {
			categoryNames = append(categoryNames, category)
		}
		sort.Strings(categoryNames)

		
		for _, category := range categoryNames {
			groupAgents := categories[category]

			
			if len(groupAgents) == 0 {
				continue
			}

			
			fmt.Println()
			Header(category)

			
			sort.Slice(groupAgents, func(i, j int) bool {
				return groupAgents[i].Name < groupAgents[j].Name
			})

			
			if options.CompactMode {
				displayCompactAgentGroup(groupAgents, options)
			} else {
				displayDetailedAgentGroup(groupAgents, options)
			}
		}
	} else {
		
		sort.Slice(agents, func(i, j int) bool {
			return agents[i].ID < agents[j].ID
		})

		
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


func displayCompactAgentGroup(agents []*agent.AgentDefinition, options AgentDisplayOptions) {
	for i, a := range agents {
		prefix := "  "
		if a.ID == options.SelectedAgentID {
			prefix = "• "
		}

		
		idStr := a.ID

		
		nameStr := a.Name
		if a.Version != "" && a.Version != "1.0" {
			nameStr = fmt.Sprintf("%s (%s)", a.Name, a.Version)
		}

		
		fmt.Printf("%s%-20s %s\n", prefix, idStr, nameStr)

		
		if i < len(agents)-1 {
			fmt.Println()
		}
	}
}


func displayDetailedAgentGroup(agents []*agent.AgentDefinition, options AgentDisplayOptions) {
	for i, a := range agents {
		prefix := "  "
		if a.ID == options.SelectedAgentID {
			prefix = "• "
		}

		
		idStr := a.ID

		
		nameStr := a.Name
		if a.Version != "" && a.Version != "1.0" {
			nameStr = fmt.Sprintf("%s (%s)", a.Name, a.Version)
		}

		
		fmt.Printf("%s%-20s %s\n", prefix, idStr, nameStr)

		
		if a.Description != "" {
			shortDesc := truncateText(a.Description, options.TermWidth-30)
			fmt.Printf("    %s\n", shortDesc)
		}

		
		if len(a.Templates) > 0 {
			fmt.Printf("    Templates: %s\n", strings.Join(a.Templates, ", "))
		}

		
		if i < len(agents)-1 {
			fmt.Println()
		}
	}
}


func truncateText(text string, maxWidth int) string {
	if len(text) <= maxWidth {
		return text
	}

	
	return text[:maxWidth-3] + "..."
}


func DisplayAgentInfoEnhanced(agent *agent.AgentDefinition, verbose bool) error {
	if agent == nil {
		return fmt.Errorf("agent is nil")
	}

	fmt.Println()
	Header("Agent Information")
	fmt.Println()

	
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

	
	if agent.Description != "" {
		Header("Description")
		fmt.Println()
		Plain("  %s", agent.Description)
		fmt.Println()
	}

	
	if len(agent.Templates) > 0 {
		Header("Templates")
		fmt.Println()
		for _, tmpl := range agent.Templates {
			Plain("  • %s", tmpl)
		}
		fmt.Println()
	}

	
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





































































