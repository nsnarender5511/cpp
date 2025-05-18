package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vibe/internal/utils"

	"github.com/PuerkitoBio/goquery"
)


type ParserConfig struct {
	HTTPTimeout     time.Duration
	MaxResponseSize int64
	MaxFileSize     int64
	MaxRetries      int
}


func DefaultParserConfig() ParserConfig {
	return ParserConfig{
		HTTPTimeout:     30 * time.Second,
		MaxResponseSize: 10 * 1024 * 1024, 
		MaxFileSize:     10 * 1024 * 1024, 
		MaxRetries:      3,
	}
}


type ParsedRule struct {
	Name        string
	Description string
	Content     string
	Format      string
	Source      string
}


func ParseRules(content []byte, source string, isWeb bool) ([]ParsedRule, error) {
	if !isWeb {
		return nil, wrapValidationError("isWeb", "non-web content parsing not implemented")
	}
	return parseCursorDirectory(content, source)
}











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


func parseCursorDirectory(content []byte, source string) ([]ParsedRule, error) {
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, wrapOpError("parseCursorDirectory", source, err, "failed to parse HTML")
	}

	var rules []ParsedRule

	
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

	
	if len(rules) == 0 {
		selectors := []string{
			".prose", ".markdown-body", ".content",
			"article", ".article", "main",
			".rule-content", ".text-content",
		}

		if contentText, _, found := findContent(doc, selectors); found {
			name := getNameFromSource(source)

			
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


func getNameFromSource(source string) string {
	base := filepath.Base(source)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	return toTitle(name)
}


type WebRuleParser struct {
	client *http.Client
	config ParserConfig
}


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

	
	body := http.MaxBytesReader(nil, resp.Body, p.config.MaxResponseSize)
	content, err := io.ReadAll(body)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to read response body")
	}

	
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

	
	return parsedRules[0].ToCursorRule(), nil
}


func (p *WebRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	
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

	
	return parsedRules[0].ToCursorRule(), nil
}


type FileRuleParser struct {
	config ParserConfig
}


func NewFileRuleParser(config *ParserConfig) *FileRuleParser {
	cfg := DefaultParserConfig()
	if config != nil {
		cfg = *config
	}
	return &FileRuleParser{config: cfg}
}


func (p *FileRuleParser) Parse(path string) (*CursorRule, error) {
	
	if !filepath.IsAbs(path) && strings.Contains(path, "..") {
		return nil, wrapValidationError("path", "invalid file path: must be absolute or not contain parent references")
	}

	
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

	
	if info.Size() > p.config.MaxFileSize {
		return nil, wrapValidationError("size", fmt.Sprintf("file too large: %d bytes (max %d)", info.Size(), p.config.MaxFileSize))
	}

	
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, wrapOpError("Parse", path, err, "failed to read file")
	}

	utils.Debug(fmt.Sprintf("Parsing file rule | path=%s size=%d", path, len(content)))

	
	parsedRules, err := ParseRules(content, path, false)
	if err != nil {
		return nil, err
	}

	if len(parsedRules) == 0 {
		return nil, wrapValidationError("content", "no rules found in file")
	}

	
	return parsedRules[0].ToCursorRule(), nil
}


func (p *FileRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	
	var rule CursorRule
	if err := json.Unmarshal(content, &rule); err == nil {
		
		if rule.Metadata.Name == "" {
			return nil, &ValidationError{
				Field:   "metadata.name",
				Message: "name is required",
			}
		}
		return &rule, nil
	}

	
	text := strings.TrimSpace(string(content))
	if text == "" {
		return nil, &ValidationError{
			Field:   "content",
			Message: "rule content cannot be empty",
		}
	}

	
	lines := strings.Split(text, "\n")
	name := filepath.Base(lines[0])
	description := "File-based rule"
	if len(lines) > 1 {
		description = strings.TrimSpace(lines[1])
	}

	
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


type CompositeRuleParser struct {
	fileParser *FileRuleParser
	webParser  *WebRuleParser
	config     ParserConfig
}


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


func (p *CompositeRuleParser) Parse(path string) (*CursorRule, error) {
	utils.Debug(fmt.Sprintf("Attempting to parse rule | path=%s", path))

	
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return p.webParser.Parse(context.Background(), path)
	}

	
	rule, fileErr := p.fileParser.Parse(path)
	if fileErr == nil {
		return rule, nil
	}

	
	if strings.Contains(path, "://") {
		rule, webErr := p.webParser.Parse(context.Background(), path)
		if webErr != nil {
			
			return nil, wrapOpError("Parse", path, fmt.Errorf("file: %v, web: %v", fileErr, webErr), "all parsers failed")
		}
		return rule, nil
	}

	
	return nil, fileErr
}


func (p *CompositeRuleParser) ParseContent(content []byte) (*CursorRule, error) {
	utils.Debug(fmt.Sprintf("Attempting to parse content | size=%d", len(content)))

	
	rule, fileErr := p.fileParser.ParseContent(content)
	if fileErr == nil {
		return rule, nil
	}

	
	rule, webErr := p.webParser.ParseContent(content)
	if webErr != nil {
		
		return nil, wrapOpError("ParseContent", "composite", fmt.Errorf("file: %v, web: %v", fileErr, webErr), "all parsers failed")
	}

	return rule, nil
}
