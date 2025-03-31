package utils

import "fmt"

// OpError represents an operational error
type OpError struct {
	Op      string
	Path    string
	Err     error
	Message string
}

func (e *OpError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s: %v", e.Op, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s: %v", e.Op, e.Path, e.Err)
}

func (e *OpError) Unwrap() error {
	return e.Err
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s", e.Field, e.Message)
}

// wrapOpError wraps an error with operation context
func wrapOpError(op, path string, err error, msg string) error {
	return &OpError{
		Op:      op,
		Path:    path,
		Err:     err,
		Message: msg,
	}
}

// wrapValidationError creates a validation error
func wrapValidationError(field, msg string) error {
	return &ValidationError{
		Field:   field,
		Message: msg,
	}
}
