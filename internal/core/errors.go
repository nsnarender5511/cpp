package core

import (
	"fmt"
)


type OpError struct {
	Op   string 
	Path string 
	Err  error  
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


type ValidationError struct {
	Field   string 
	Message string 
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}


type NotFoundError struct {
	Resource string 
	ID       string 
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.ID)
}


type ParseError struct {
	Source string 
	Line   int    
	Err    error  
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


type ConfigError struct {
	Key string 
	Err error  
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("configuration error for %s: %v", e.Key, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}


func wrapOpError(op, path string, err error, msg string) error {
	return &OpError{
		Op:   op,
		Path: path,
		Err:  fmt.Errorf("%s: %w", msg, err),
	}
}


func wrapValidationError(field, msg string) error {
	return &ValidationError{
		Field:   field,
		Message: msg,
	}
}


func wrapNotFoundError(resource, id string) error {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}


func wrapParseError(source string, err error, line int) error {
	return &ParseError{
		Source: source,
		Line:   line,
		Err:    err,
	}
}
