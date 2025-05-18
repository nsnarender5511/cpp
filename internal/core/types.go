package core

import (
	"context"
	"time"
)


type RuleMetadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


type CursorRule struct {
	Metadata  RuleMetadata        `json:"metadata"`
	Patterns  []string            `json:"patterns"`
	Templates map[string]Template `json:"templates"`
}


type Template struct {
	Content    string            `json:"content"`
	Variables  map[string]string `json:"variables"`
	IsRequired bool              `json:"is_required"`
}


type RuleStorage interface {
	SaveRule(rule *CursorRule) error
	GetRule(name string) (*CursorRule, error)
	ListRules() ([]*CursorRule, error)
	DeleteRule(name string) error
}


type RuleParser interface {
	Parse(path string) (*CursorRule, error)
	ParseContent(content []byte) (*CursorRule, error)
}


type RegistryService interface {
	Sync(ctx context.Context) error
	GetRule(name string) (*CursorRule, error)
	ListRules() ([]*CursorRule, error)
}
