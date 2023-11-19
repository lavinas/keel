package usecase

import (
	"time"

	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/internal/core/domain"
)

// Service is the usecase handler for the application
type UseCase struct {
	repo   port.Repository
	logger port.Logger
	config port.Config
}

// NewService creates a new usecase service
func NewService(repo port.Repository, logger port.Logger, config port.Config) *UseCase {
	return &UseCase{
		repo:   repo,
		logger: logger,
		config: config,
	}
}

// RegisterClient registers a new client
func (s *UseCase) RegisterClient(dto *dto.RegisterInvoiceClient) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	client := domain.NewClient(dto.ID, dto.Name, dto.Email, dto.Document, dto.Phone, time.Time{}, time.Time{})
	if err := client.Validate(); err != nil {
		return err
	}
	if err := s.repo.AddClient(client); err != nil {
		return err
	}
	return nil
}
