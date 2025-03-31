package core

import (
	"fmt"
)

// OpError represents an operational error with context
type OpError struct {
	Op   string // Operation that failed
	Path string // Optional path or identifier related to the error
	Err  error  // Underlying error
}

func (e *OpError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("%s failed for %s: %v", e.Op, e.Path, e.Err)
	}
	return fmt.Sprintf("%s failed: %v", e.Op, e.Err)
}

func (e *OpError) Unwrap() error {
	return e.Err
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string // Field that failed validation
	Message string // Validation error message
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

// NotFoundError represents a resource not found error
type NotFoundError struct {
	Resource string // Type of resource not found
	ID       string // Identifier of the resource
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.ID)
}

// ParseError represents a parsing error
type ParseError struct {
	Source string // Source being parsed
	Line   int    // Optional line number
	Err    error  // Underlying error
}

func (e *ParseError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("parse error in %s at line %d: %v", e.Source, e.Line, e.Err)
	}
	return fmt.Sprintf("parse error in %s: %v", e.Source, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

// ConfigError represents a configuration error
type ConfigError struct {
	Key string // Configuration key
	Err error  // Underlying error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("configuration error for %s: %v", e.Key, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// wrapOpError creates a new OpError with consistent formatting
func wrapOpError(op, path string, err error, msg string) error {
	return &OpError{
		Op:   op,
		Path: path,
		Err:  fmt.Errorf("%s: %w", msg, err),
	}
}

// wrapValidationError creates a new ValidationError with consistent formatting
func wrapValidationError(field, msg string) error {
	return &ValidationError{
		Field:   field,
		Message: msg,
	}
}

// wrapNotFoundError creates a new NotFoundError with consistent formatting
func wrapNotFoundError(resource, id string) error {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

// wrapParseError creates a new ParseError with consistent formatting
func wrapParseError(source string, err error, line int) error {
	return &ParseError{
		Source: source,
		Line:   line,
		Err:    err,
	}
}
