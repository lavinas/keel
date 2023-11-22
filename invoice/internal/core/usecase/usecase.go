package usecase

import (
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
