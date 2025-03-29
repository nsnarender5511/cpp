package core

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Rule represents a parsed agent rule
type Rule struct {
	Name        string
	Description string
	Content     string
	Format      string
	Source      string
}

// ParseRules processes content into structured rules
func ParseRules(content []byte, source string, webMode bool) ([]Rule, error) {
	contentFormat := detectFormat(content)

	// Force HTML format if webMode is enabled
	if webMode {
		contentFormat = "html"
	}

	switch contentFormat {
	case "markdown", "mdx":
		return parseMarkdown(content, source)
	case "json":
		return parseJSON(content, source)
	case "html":
		return parseHTML(content, source)
	default:
		// Default to treating as a single text rule
		return []Rule{{
			Name:        getNameFromSource(source),
			Description: "Imported from " + source,
			Content:     string(content),
			Format:      "text",
			Source:      source,
		}}, nil
	}
}

// detectFormat attempts to determine the content format
func detectFormat(content []byte) string {
	contentStr := string(content)

	// Check for JSON format
	if strings.HasPrefix(strings.TrimSpace(contentStr), "{") ||
		strings.HasPrefix(strings.TrimSpace(contentStr), "[") {
		var js interface{}
		if json.Unmarshal(content, &js) == nil {
			return "json"
		}
	}

	// Check for HTML format
	if strings.Contains(contentStr, "<!DOCTYPE html>") ||
		strings.Contains(contentStr, "<html") ||
		strings.Contains(contentStr, "<body") {
		return "html"
	}

	// Check for Markdown format
	if strings.Contains(contentStr, "# ") ||
		strings.Contains(contentStr, "## ") ||
		strings.Contains(contentStr, "```") {
		if strings.Contains(contentStr, "```jsx") ||
			strings.Contains(contentStr, "```tsx") {
			return "mdx"
		}
		return "markdown"
	}

	// Default to text
	return "text"
}

// parseHTML extracts rule content from HTML pages
func parseHTML(content []byte, source string) ([]Rule, error) {
	// Create a reader from the content
	reader := bytes.NewReader(content)

	// Create a new goquery document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var rules []Rule

	// Check if this is cursor.directory specifically
	isCursorDirectory := strings.Contains(source, "cursor.directory")

	// Special handling for cursor.directory
	if isCursorDirectory {
		// Look for code elements with specific class
		codeElements := doc.Find("code.text-sm.block")
		if codeElements.Length() > 0 {
			codeElements.Each(func(i int, s *goquery.Selection) {
				codeText := strings.TrimSpace(s.Text())
				if codeText != "" {
					// Extract the rule name from the URL
					urlPath := strings.TrimSuffix(filepath.Base(source), filepath.Ext(source))
					name := toTitle(strings.ReplaceAll(urlPath, "-", " "))
					description := "Cursor rule imported from " + source

					rules = append(rules, Rule{
						Name:        name,
						Description: description,
						Content:     codeText,
						Format:      "text",
						Source:      source,
					})
				}
			})
			return rules, nil
		}
	}

	// General extraction from code elements
	codeElements := doc.Find("code")
	if codeElements.Length() > 0 {
		codeElements.Each(func(i int, s *goquery.Selection) {
			codeText := strings.TrimSpace(s.Text())
			if codeText != "" {
				// Generate a name from source if possible
				name := getNameFromSource(source)
				if i > 0 {
					name = name + "-" + string(rune('a'+i))
				}

				// Try to extract a title from the code content
				lines := strings.Split(codeText, "\n")
				description := "Imported from " + source
				if len(lines) > 0 && lines[0] != "" {
					name = lines[0]
					if len(lines) > 1 {
						description = lines[1]
					}
				}

				rules = append(rules, Rule{
					Name:        name,
					Description: description,
					Content:     codeText,
					Format:      "text",
					Source:      source,
				})
			}
		})
	}

	// Extract from an article or prose content if no code elements found
	if len(rules) == 0 {
		// Try common selectors for content
		selectors := []string{
			".prose", ".markdown-body", ".content",
			"article", ".article", "main",
			".rule-content", ".text-content",
		}

		for _, selector := range selectors {
			contentEl := doc.Find(selector)
			if contentEl.Length() > 0 {
				contentText := strings.TrimSpace(contentEl.Text())
				if contentText != "" {
					name := getNameFromSource(source)

					// Try to extract a title from page
					title := doc.Find("title").Text()
					if title != "" {
						name = title
					} else {
						// Try h1 as title
						h1 := doc.Find("h1").First().Text()
						if h1 != "" {
							name = h1
						}
					}

					rules = append(rules, Rule{
						Name:        name,
						Description: "Imported from " + source,
						Content:     contentText,
						Format:      "text",
						Source:      source,
					})
					break
				}
			}
		}
	}

	// If still no rules, extract everything from body as last resort
	if len(rules) == 0 {
		bodyText := strings.TrimSpace(doc.Find("body").Text())
		if bodyText != "" {
			name := getNameFromSource(source)

			// Try to get title
			title := doc.Find("title").Text()
			if title != "" {
				name = title
			}

			rules = append(rules, Rule{
				Name:        name,
				Description: "Imported from " + source,
				Content:     bodyText,
				Format:      "text",
				Source:      source,
			})
		}
	}

	// Handle list pages with multiple rule entries
	if len(rules) == 0 {
		ruleEntries := doc.Find(".rule-entry, .card, .rule-card, article, .rule")
		if ruleEntries.Length() > 0 {
			ruleEntries.Each(func(i int, s *goquery.Selection) {
				// Try to find rule name from heading
				name := s.Find("h1, h2, h3, h4, .title, .rule-title").First().Text()
				name = strings.TrimSpace(name)

				// Find description
				description := s.Find("p, .description, .summary").First().Text()
				description = strings.TrimSpace(description)
				if description == "" {
					description = "Imported from " + source
				}

				// Find content
				content := s.Find("code, pre, .content, .rule-content").Text()
				content = strings.TrimSpace(content)
				if content == "" {
					// If no specific content found, use the whole entry text
					content = s.Text()
				}

				if name != "" && content != "" {
					rules = append(rules, Rule{
						Name:        name,
						Description: description,
						Content:     content,
						Format:      "text",
						Source:      source,
					})
				}
			})
		}
	}

	return rules, nil
}

// parseMarkdown parses markdown-formatted rule content
func parseMarkdown(content []byte, source string) ([]Rule, error) {
	contentStr := string(content)
	var rules []Rule

	// Simple regex to find markdown headings that might be rule names
	ruleRegex := regexp.MustCompile(`(?m)^#+\s+(.+)$`)
	matches := ruleRegex.FindAllStringSubmatchIndex(contentStr, -1)

	if len(matches) == 0 {
		// If no headings found, treat entire content as one rule
		filename := filepath.Base(source)
		name := strings.TrimSuffix(filename, filepath.Ext(filename))

		rules = append(rules, Rule{
			Name:        name,
			Description: "Imported from " + source,
			Content:     contentStr,
			Format:      "markdown",
			Source:      source,
		})
	} else {
		// Process each section as a potential rule
		for i, match := range matches {
			nameStart, nameEnd := match[2], match[3]
			name := contentStr[nameStart:nameEnd]

			// Get content until next heading or end of file
			contentStart := match[0] // Include the heading
			contentEnd := len(contentStr)
			if i < len(matches)-1 {
				contentEnd = matches[i+1][0]
			}

			ruleContent := contentStr[contentStart:contentEnd]

			// Extract description from first paragraph if possible
			description := name
			descRegex := regexp.MustCompile(`(?m)^#+\s+.+\n+(.+)$`)
			if descMatch := descRegex.FindStringSubmatch(ruleContent); len(descMatch) > 1 {
				description = descMatch[1]
			}

			rules = append(rules, Rule{
				Name:        name,
				Description: description,
				Content:     ruleContent,
				Format:      "markdown",
				Source:      source,
			})
		}
	}

	return rules, nil
}

// parseJSON parses JSON-formatted rule content
func parseJSON(content []byte, source string) ([]Rule, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(content, &jsonData); err != nil {
		// Try as array
		var jsonArray []interface{}
		if err := json.Unmarshal(content, &jsonArray); err != nil {
			// Not valid JSON, return empty rules
			return nil, err
		}

		// Process JSON array
		var rules []Rule
		for i, item := range jsonArray {
			if itemMap, ok := item.(map[string]interface{}); ok {
				rule := processJSONItem(itemMap, source, i)
				rules = append(rules, rule)
			}
		}
		return rules, nil
	}

	// Handle different JSON structures based on what's detected
	var rules []Rule

	// Try to detect if this is a collection or single rule
	if ruleName, ok := jsonData["name"].(string); ok {
		// Looks like a single rule
		description := ""
		if desc, ok := jsonData["description"].(string); ok {
			description = desc
		}

		rules = append(rules, Rule{
			Name:        ruleName,
			Description: description,
			Content:     string(content),
			Format:      "json",
			Source:      source,
		})
	} else if rulesArr, ok := jsonData["rules"].([]interface{}); ok {
		// Looks like a collection
		for i, r := range rulesArr {
			if ruleObj, ok := r.(map[string]interface{}); ok {
				rule := processJSONItem(ruleObj, source, i)
				rules = append(rules, rule)
			}
		}
	} else {
		// Unknown structure, use filename as rule name
		filename := filepath.Base(source)
		name := strings.TrimSuffix(filename, filepath.Ext(filename))

		rules = append(rules, Rule{
			Name:        name,
			Description: "Imported from " + source,
			Content:     string(content),
			Format:      "json",
			Source:      source,
		})
	}

	return rules, nil
}

// processJSONItem converts a JSON map to a Rule
func processJSONItem(item map[string]interface{}, source string, index int) Rule {
	name := ""
	if n, ok := item["name"].(string); ok && n != "" {
		name = n
	} else {
		// Fallback name based on source and index
		filename := filepath.Base(source)
		baseName := strings.TrimSuffix(filename, filepath.Ext(filename))
		name = baseName + "-" + string(rune('a'+index))
	}

	description := ""
	if d, ok := item["description"].(string); ok {
		description = d
	} else {
		description = "Imported from " + source
	}

	// Convert item to JSON string
	content, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		// Fallback if marshaling fails
		content = []byte("{}")
	}

	return Rule{
		Name:        name,
		Description: description,
		Content:     string(content),
		Format:      "json",
		Source:      source,
	}
}

// getNameFromSource extracts a name from the source URL
func getNameFromSource(source string) string {
	filename := filepath.Base(source)

	// Remove extension if present
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// If name is still empty, use "imported-rule"
	if name == "" {
		name = "imported-rule"
	}

	return name
}

// toTitle converts a string to title case
func toTitle(s string) string {
	// Convert to lowercase first
	s = strings.ToLower(s)

	// Use regex to find word boundaries
	re := regexp.MustCompile(`\b\w`)

	// Replace first letter of each word with uppercase
	return re.ReplaceAllStringFunc(s, func(match string) string {
		return strings.ToUpper(match)
	})
}
