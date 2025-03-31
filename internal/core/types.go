package core

import (
	"context"
	"time"
)

// RuleMetadata represents metadata about a rule
type RuleMetadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CursorRule represents a complete cursor rule definition
type CursorRule struct {
	Metadata  RuleMetadata        `json:"metadata"`
	Patterns  []string            `json:"patterns"`
	Templates map[string]Template `json:"templates"`
}

// Template represents a rule template
type Template struct {
	Content    string            `json:"content"`
	Variables  map[string]string `json:"variables"`
	IsRequired bool              `json:"is_required"`
}

// RuleStorage defines the interface for rule persistence
type RuleStorage interface {
	SaveRule(rule *CursorRule) error
	GetRule(name string) (*CursorRule, error)
	ListRules() ([]*CursorRule, error)
	DeleteRule(name string) error
}

// RuleParser defines the interface for parsing rules
type RuleParser interface {
	Parse(path string) (*CursorRule, error)
	ParseContent(content []byte) (*CursorRule, error)
}

// RegistryService defines the interface for rule registry operations
type RegistryService interface {
	Sync(ctx context.Context) error
	GetRule(name string) (*CursorRule, error)
	ListRules() ([]*CursorRule, error)
}
