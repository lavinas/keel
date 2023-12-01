package kerror

import (
	"strings"
)

const (
	// Error types
	Internal   = "INTERNAL"
	BadRequest = "BAD_REQUEST"
	Conflict   = "CONFLICT"
	None	   = "NONE"

	messageSeparator = "; "
)

var (
	// orderType is the order of the error types
	orderType = map[string]int{
		Internal:   1,
		BadRequest: 2,
		Conflict:   3,
		None:		4,
		
	}

	// httpCode is the http code of the error types
	httpCode = map[string]int{
		Internal:   500,
		BadRequest: 400,
		Conflict:   409,
		None:		200,
	}
)

// KError represents an error on keel system
type KError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// NewKError creates a new KError
func NewKError(t string, m string) *KError {
	return &KError{
		Type:    t,
		Message: m,
	}
}

// Error returns the error message
func (e *KError) Error() string {
	return strings.TrimLeft(e.Message, messageSeparator)
}

// Is checks if the error is the same type
func (e *KError) ErrorType() string {
	return e.Type
}

// Join joins a message to the error
func (e *KError) Join(t string, m string) {
	if orderType[t] < orderType[e.Type] {
		e.Type = t
	}
	e.Message = e.Message + messageSeparator + m
}

// IsEmpty checks if the error is empty
func (e *KError) IsEmpty() bool {
	return e.Type == None
}

// GetHTTPCode returns the http code of the error
func (e *KError) GetHTTPCode() int {
	return httpCode[e.Type]
}
