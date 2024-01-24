package handler

import "github.com/lavinas/keel/pkg/kerror"

const (
	TypeAddress = "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html"
)

// Error represents an error on keel system
type RestError struct {
	Etype  string `json:"type"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// NewError creates a new Error
func NewError(err *kerror.KError) *RestError {
	return &RestError{
		Etype:  TypeAddress,
		Status: err.GetHTTPCode(),
		Title:  err.GetHTTPTitle(),
		Detail: err.Error(),
	}
}
