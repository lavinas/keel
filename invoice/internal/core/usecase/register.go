package usecase

import (
	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Service is the usecase handler for the application
type Service struct {
	repo   port.Repository
	logger port.Logger
	config port.Config
}

func NewService(repo port.Repository, logger port.Logger, config port.Config) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
		config: config,
	}
}

func (s *Service) RegisterClient(dto *dto.RegisterInvoiceClient) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	return nil

}
