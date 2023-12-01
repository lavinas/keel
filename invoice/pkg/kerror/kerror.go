package kerror

const (
	Internal   = "INTERNAL"
	BadRequest = "BAD_REQUEST"
	Conflict   = "CONFLICT"
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
	return e.Message
}

// Is checks if the error is the same type
func (e *KError) ErrorType () string {
	return e.Type
}