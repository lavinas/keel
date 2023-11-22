package usecase

import (
	"encoding/json"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

const (
	BUSINNESS_ID = "BUSINNESS_ID"
)

// Service is the usecase handler for the application
type UseCase struct {
	repo   port.Repository
	logger port.Logger
	config port.Config
}

// NewService creates a new usecase service
func NewUseCase(config port.Config, logger port.Logger, repo port.Repository) *UseCase {
	return &UseCase{
		repo:   repo,
		logger: logger,
		config: config,
	}
}

// logInfoObj logs an info message with an object
func (s *UseCase) logObj(logType string, prefix string, message string, obj any) {
	dto_log, _ := json.Marshal(obj)
	dto_str := string(dto_log)
	if logType == "info" {
		s.logger.Infof("%s | %s | %s", prefix, message, dto_str)
	} else if logType == "error" {
		s.logger.Errorf("%s: %s", message, dto_str)
	}
}
