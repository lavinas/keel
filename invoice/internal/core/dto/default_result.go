package dto

// DefaultResult is the default result dto
type DefaultResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewDefaultResult(code int, message string) *DefaultResult {
	return &DefaultResult{
		Code:    code,
		Message: message,
	}
}

// Get returns the code and message
func (d *DefaultResult) Get() (int, string) {
	return d.Code, d.Message
}

// Set sets the code and message
func (d *DefaultResult) Set(code int, message string) {
	d.Code = code
	d.Message = message
}
