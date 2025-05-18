package utils

import "fmt"


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


type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s", e.Field, e.Message)
}


func wrapOpError(op, path string, err error, msg string) error {
	return &OpError{
		Op:      op,
		Path:    path,
		Err:     err,
		Message: msg,
	}
}








