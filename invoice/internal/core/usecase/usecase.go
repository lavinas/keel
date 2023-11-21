package usecase

import (
	"time"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
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

// RegisterClient registers a new client
func (s *UseCase) RegisterClient(dto port.RegisterClient) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	id, name, email, document, phone := dto.Get()
	client := domain.NewClient(id, name, email, document, phone, time.Time{}, time.Time{})
	if err := client.Validate(); err != nil {
		return err
	}
	if err := s.repo.AddClient(client); err != nil {
		return err
	}
	return nil
}
