package kerror

import (
	"strings"
)

const (
	// Error types
	Internal   = "INTERNAL"
	BadRequest = "BAD_REQUEST"
	Conflict   = "CONFLICT"
	None       = "NONE"

	messageSeparator = "; "
)

var (
	// orderEtype is the order of the error types
	orderEtype = map[string]int{
		Internal:   1,
		BadRequest: 2,
		Conflict:   3,
		None:       4,
	}

	// httpCode is the http code of the error types
	httpCode = map[string]int{
		Internal:   500,
		BadRequest: 400,
		Conflict:   409,
		None:       200,
	}
)

// KError represents an error on keel system
type KError struct {
	etype   string
	message string
	prefix  string
}

// NewKError creates a new KError
func NewKError(t string, m string) *KError {
	return &KError{
		etype:   t,
		message: m,
	}
}

// Error returns the error message
func (e *KError) Error() string {
	m := strings.TrimLeft(e.message, messageSeparator)
	if e.prefix != "" {
		m = strings.Replace(m, messageSeparator, messageSeparator+e.prefix, -1)
		m = e.prefix + m
	}
	return m
}

// Is checks if the error is the same type
func (e *KError) GetType() string {
	return e.etype
}

// Join joins a message to the error
func (e *KError) Join(t string, m string) {
	if orderEtype[t] < orderEtype[e.etype] {
		e.etype = t
	}
	e.message = e.message + messageSeparator + m
}

// IsEmpty checks if the error is empty
func (e *KError) IsEmpty() bool {
	return e.etype == None
}

// GetHTTPCode returns the http code of the error
func (e *KError) GetHTTPCode() int {
	return httpCode[e.etype]
}

// SetPrefix sets the prefix of the error
func (e *KError) SetPrefix(prefix string) {
	e.prefix = prefix
}
