package handler


import (
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Rest is the rest handler for the application
type Rest struct {
	Log port.Logger
}

// NewRest creates a new rest handler
func NewRest(logger port.Logger) *Rest {
	return &Rest{
		Log: logger,
	}
}







