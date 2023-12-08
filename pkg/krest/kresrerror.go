package krest

import "github.com/lavinas/keel/pkg/kerror"

const (
	TypeAddress = "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html"
)

// Error represents an error on keel system
type KRestError struct {
	Etype  string `json:"type"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// NewError creates a new Error
func NewKRestError(err *kerror.KError) *KRestError {
	return &KRestError{
		Etype:  TypeAddress,
		Status: err.GetHTTPCode(),
		Title:  err.GetHTTPTitle(),
		Detail: err.Error(),
	}
}
