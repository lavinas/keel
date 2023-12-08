package tools

import (
	log "github.com/lavinas/keel/pkg/klog"
)

// Logger is the logger handler for the application
type Logger struct {
	log.Klog
}

// NewLogger creates a new logger handler
func NewLogger(component string, info bool) (*Logger, error) {
	klog, err := log.NewKlog(component, info)
	if err != nil {
		return nil, err
	}
	return &Logger{*klog}, nil
}
