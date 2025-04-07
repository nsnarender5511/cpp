package core

import (
	"bytes"
	"context"
	"vibe/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ParserConfig holds configuration for rule parsers
type ParserConfig struct {
	HTTPTimeout     time.Duration
	MaxResponseSize int64
	MaxFileSize     int64
	MaxRetries      int
}

// DefaultParserConfig returns default parser configuration
func DefaultParserConfig() ParserConfig {
	return ParserConfig{
		HTTPTimeout:     30 * time.Second,
		MaxResponseSize: 10 * 1024 * 1024, // 10MB
		MaxFileSize:     10 * 1024 * 1024, // 10MB
		MaxRetries:      3,
	}
}

// ParsedRule represents a parsed agent rule before conversion to CursorRule
type ParsedRule struct {
	Name        string
	Description string
	Content     string
	Format      string
	Source      string
}

// ParseRules processes content into structured rules
func ParseRules(content []byte, source string, isWeb bool) ([]ParsedRule, error) {
	if !isWeb {
		return nil, wrapValidationError("isWeb", "non-web content parsing not implemented")
	}
	return parseCursorDirectory(content, source)
}

// extractDomain gets the domain from a URL
func extractDomain(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", wrapParseError(urlStr, err, 0)
	}
	return u.Host, nil
}

// findContent searches for content using multiple selectors
func findContent(doc *goquery.Document, selectors []string) (string, string, bool) {
	for _, selector := range selectors {
		if el := doc.Find(selector); el.Length() > 0 {
			if text := strings.TrimSpace(el.Text()); text != "" {
				return text, selector, true
			}
		}
	}
	return "", "", false
}

// parseCursorDirectory extracts rules from cursor.directory content
func parseCursorDirectory(content []byte, source string) ([]ParsedRule, error) {
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, wrapOpError("parseCursorDirectory", source, err, "failed to parse HTML")
	}

	var rules []ParsedRule

	// Look for code elements with specific class first
	codeElements := doc.Find("code.text-sm.block")
	if codeElements.Length() > 0 {
		codeElements.Each(func(i int, s *goquery.Selection) {
			if codeText := strings.TrimSpace(s.Text()); codeText != "" {
				urlPath := strings.TrimSuffix(filepath.Base(source), filepath.Ext(source))
				name := toTitle(strings.ReplaceAll(urlPath, "-", " "))
				rules = append(rules, ParsedRule{
					Name:        name,
					Description: "Cursor rule imported from " + source,
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
			if codeText := strings.TrimSpace(s.Text()); codeText != "" {
				name := getNameFromSource(source)
				if i > 0 {
					name = name + "-" + string(rune('a'+i))
				}

				lines := strings.Split(codeText, "\n")
				description := "Imported from " + source
				if len(lines) > 0 && lines[0] != "" {
					name = lines[0]
					if len(lines) > 1 {
						description = lines[1]
					}
				}

				rules = append(rules, ParsedRule{
					Name:        name,
					Description: description,
					Content:     codeText,
					Format:      "text",
					Source:      source,
				})
			}
		})
	}

	// If no rules found, try common content selectors
	if len(rules) == 0 {
		selectors := []string{
			".prose", ".markdown-body", ".content",
			"article", ".article", "main",
			".rule-content", ".text-content",
		}

		if contentText, _, found := findContent(doc, selectors); found {
			name := getNameFromSource(source)

			// Try to extract a title from page
			if title := doc.Find("title").Text(); title != "" {
				name = title
			} else if h1 := doc.Find("h1").First().Text(); h1 != "" {
				name = h1
			}

			rules = append(rules, ParsedRule{
				Name:        name,
				Description: "Imported from " + source,
				Content:     contentText,
				Format:      "text",
				Source:      source,
			})
		}
	}

	if len(rules) == 0 {
		return nil, wrapNotFoundError("rules", source)
	}

	return rules, nil
}

// ToCursorRule converts a ParsedRule to a CursorRule
func (p *ParsedRule) ToCursorRule() *CursorRule {
	now := time.Now()
	return &CursorRule{
		Metadata: RuleMetadata{
			Name:        p.Name,
			Description: p.Description,
			Version:     "1.0.0",
			Author:      "cursor.directory",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Patterns: []string{},
		Templates: map[string]Template{
			"default": {
				Content:    p.Content,
				Variables:  make(map[string]string),
				IsRequired: true,
			},
		},
	}
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
	base := filepath.Base(source)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	return toTitle(name)
}

// WebRuleParser implements RuleParser for web-based rules
type WebRuleParser struct {
	client *http.Client
	config ParserConfig
}

// NewWebRuleParser creates a new WebRuleParser
func NewWebRuleParser(config *ParserConfig) *WebRuleParser {
	cfg := DefaultParserConfig()
	if config != nil {
		cfg = *config
	}

	return &WebRuleParser{
		client: &http.Client{
			Timeout: cfg.HTTPTimeout,
		},
		config: cfg,
	}
}

// Parse implements RuleParser.Parse
func (p *WebRuleParser) Parse(ctx context.Context, path string) (*CursorRule, error) {
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		return nil, wrapValidationError("url", "invalid URL format")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to create request")
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to fetch URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, wrapOpError("Parse", path, fmt.Errorf("HTTP %d", resp.StatusCode), "failed to fetch URL")
	}

	// Limit response size
	body := http.MaxBytesReader(nil, resp.Body, p.config.MaxResponseSize)
	content, err := io.ReadAll(body)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to read response body")
	}

	// Parse the content
	parsedRules, err := ParseRules(content, path, true)
	if err != nil {
		return nil, err
	}

	if len(parsedRules) == 0 {
		return nil, &NotFoundError{
			Resource: "rules",
			ID:       path,
		}
	}

	// Convert the first parsed rule to a CursorRule
	return parsedRules[0].ToCursorRule(), nil
}

// ParseContent implements RuleParser.ParseContent
func (p *WebRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	// Parse the content as a web-based rule
	parsedRules, err := ParseRules(content, "inline-content", true)
	if err != nil {
		return nil, &OpError{
			Op:   "ParseContent",
			Path: "inline-content",
			Err:  err,
		}
	}

	if len(parsedRules) == 0 {
		return nil, &NotFoundError{
			Resource: "rules",
			ID:       "inline-content",
		}
	}

	// Convert the first parsed rule to a CursorRule
	return parsedRules[0].ToCursorRule(), nil
}

// FileRuleParser implements RuleParser for file-based rules
type FileRuleParser struct {
	config ParserConfig
}

// NewFileRuleParser creates a new FileRuleParser
func NewFileRuleParser(config *ParserConfig) *FileRuleParser {
	cfg := DefaultParserConfig()
	if config != nil {
		cfg = *config
	}
	return &FileRuleParser{config: cfg}
}

// Parse implements RuleParser.Parse for FileRuleParser
func (p *FileRuleParser) Parse(path string) (*CursorRule, error) {
	// Validate path
	if !filepath.IsAbs(path) && strings.Contains(path, "..") {
		return nil, wrapValidationError("path", "invalid file path: must be absolute or not contain parent references")
	}

	// Check file stats
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, &NotFoundError{
				Resource: "rule file",
				ID:       path,
			}
		}
		return nil, wrapOpError("Parse", path, err, "failed to stat file")
	}

	// Validate file size
	if info.Size() > p.config.MaxFileSize {
		return nil, wrapValidationError("size", fmt.Sprintf("file too large: %d bytes (max %d)", info.Size(), p.config.MaxFileSize))
	}

	// Read and parse file
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to read file")
	}

	utils.Debug(fmt.Sprintf("Parsing file rule | path=%s size=%d", path, len(content)))

	// Parse the content
	parsedRules, err := ParseRules(content, path, false)
	if err != nil {
		return nil, err
	}

	if len(parsedRules) == 0 {
		return nil, wrapValidationError("content", "no rules found in file")
	}

	// Convert to CursorRule
	return parsedRules[0].ToCursorRule(), nil
}

// ParseContent implements RuleParser.ParseContent
func (p *FileRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	// Try to parse as JSON first
	var rule CursorRule
	if err := json.Unmarshal(content, &rule); err == nil {
		// Validate required fields
		if rule.Metadata.Name == "" {
			return nil, &ValidationError{
				Field:   "metadata.name",
				Message: "name is required",
			}
		}
		return &rule, nil
	}

	// If not JSON, try to parse as a text rule
	text := strings.TrimSpace(string(content))
	if text == "" {
		return nil, &ValidationError{
			Field:   "content",
			Message: "rule content cannot be empty",
		}
	}

	// Extract name from first line if possible
	lines := strings.Split(text, "\n")
	name := filepath.Base(lines[0])
	description := "File-based rule"
	if len(lines) > 1 {
		description = strings.TrimSpace(lines[1])
	}

	// Create a basic rule
	now := time.Now()
	return &CursorRule{
		Metadata: RuleMetadata{
			Name:        name,
			Description: description,
			Version:     "1.0.0",
			Author:      "local",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Patterns: []string{},
		Templates: map[string]Template{
			"default": {
				Content:    text,
				Variables:  make(map[string]string),
				IsRequired: true,
			},
		},
	}, nil
}

// CompositeRuleParser combines multiple parsers
type CompositeRuleParser struct {
	fileParser *FileRuleParser
	webParser  *WebRuleParser
	config     ParserConfig
}

// NewCompositeRuleParser creates a new CompositeRuleParser
func NewCompositeRuleParser(config *ParserConfig) *CompositeRuleParser {
	cfg := DefaultParserConfig()
	if config != nil {
		cfg = *config
	}
	return &CompositeRuleParser{
		fileParser: NewFileRuleParser(&cfg),
		webParser:  NewWebRuleParser(&cfg),
		config:     cfg,
	}
}

// Parse implements RuleParser.Parse for CompositeRuleParser
func (p *CompositeRuleParser) Parse(path string) (*CursorRule, error) {
	utils.Debug(fmt.Sprintf("Attempting to parse rule | path=%s", path))

	// Check if path is a URL
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return p.webParser.Parse(context.Background(), path)
	}

	// Try file parser first
	rule, fileErr := p.fileParser.Parse(path)
	if fileErr == nil {
		return rule, nil
	}

	// If file parsing fails and path looks like a URL, try web parser
	if strings.Contains(path, "://") {
		rule, webErr := p.webParser.Parse(context.Background(), path)
		if webErr != nil {
			// Both parsers failed
			return nil, wrapOpError("Parse", path, fmt.Errorf("file: %v, web: %v", fileErr, webErr), "all parsers failed")
		}
		return rule, nil
	}

	// If path doesn't look like a URL, return the file parser error
	return nil, fileErr
}

// ParseContent implements RuleParser.ParseContent for CompositeRuleParser
func (p *CompositeRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	utils.Debug(fmt.Sprintf("Attempting to parse content | size=%d", len(content)))

	// Try file parser first
	rule, fileErr := p.fileParser.ParseContent(content)
	if fileErr == nil {
		return rule, nil
	}

	// Try web parser if file parser fails
	rule, webErr := p.webParser.ParseContent(content)
	if webErr != nil {
		// Both parsers failed
		return nil, wrapOpError("ParseContent", "composite", fmt.Errorf("file: %v, web: %v", fileErr, webErr), "all parsers failed")
	}

	return rule, nil
}
