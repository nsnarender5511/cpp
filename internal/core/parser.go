package core

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
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

// ParserFunc defines the function signature for all domain handlers
type ParserFunc func(content []byte, source string) ([]Rule, error)

// DomainRegistry maps domains to their specific parser functions
var DomainRegistry = map[string]ParserFunc{
	"cursor.directory": parseCursorDirectory,
	// More domains can be added here later
}

// ParseRules processes content into structured rules
func ParseRules(content []byte, source string, isWeb bool) ([]Rule, error) {
	// Only handle web content in this simplified version
	if !isWeb {
		return nil, fmt.Errorf("non-web content parsing not implemented in this version")
	}

	// Extract domain from source URL
	domain, err := extractDomain(source)
	if err != nil {
		return nil, fmt.Errorf("invalid source URL: %w", err)
	}

	// Find a handler for this domain
	for registeredDomain, handler := range DomainRegistry {
		if strings.Contains(domain, registeredDomain) {
			return handler(content, source)
		}
	}

	// No handler found for this domain
	return nil, fmt.Errorf("unsupported domain: %s", domain)
}

// extractDomain gets the domain from a URL
func extractDomain(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// parseCursorDirectory extracts rules from cursor.directory content
func parseCursorDirectory(content []byte, source string) ([]Rule, error) {
	// Create HTML document from content
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var rules []Rule

	// Look for code elements with specific class first
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
	} else {
		// Fall back to generic code elements
		codeElements = doc.Find("code")
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

	// If no rules found, try common content selectors as last resort
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

	if len(rules) == 0 {
		return nil, fmt.Errorf("no rules found in content from %s", source)
	}

	return rules, nil
}

// Helper functions

// toTitle converts a string to title case
func toTitle(s string) string {
	words := []string{}
	for _, word := range strings.Fields(strings.ToLower(s)) {
		if len(word) == 0 {
			continue
		}
		if len(word) == 1 {
			words = append(words, strings.ToUpper(word))
			continue
		}
		words = append(words, strings.ToUpper(word[:1])+word[1:])
	}
	return strings.Join(words, " ")
}

// getNameFromSource extracts a readable name from a source URL or file
func getNameFromSource(source string) string {
	// Extract filename without extension
	base := filepath.Base(source)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// Replace hyphens, underscores with spaces and convert to title case
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")

	return toTitle(name)
}
