package usecase

import (
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

// Service is the usecase handler for the application
type UseCase struct {
	repo   port.Repository
	logger port.Logger
	config port.Config
}

// NewService creates a new usecase handler
func NewUseCase(repo port.Repository, logger port.Logger, config port.Config) *UseCase {
	return &UseCase{
		repo:   repo,
		logger: logger,
		config: config,
	}
}

// Register registers a domain
func (s *UseCase) Create(dtoIn port.CreateDtoIn, DtoOut port.CreateDtoOut) *kerror.KError {
	if err := dtoIn.Validate(s.repo); err != nil {
		return err
	}
	domain, err := dtoIn.GetDomain()
	if err != nil {
		return err
	}
	if err := domain.SetCreate(s.repo); err != nil {
		return err
	}
	if err := domain.Validate(); err != nil {
		return err
	}
	if err := s.repo.Add(domain); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	}
	if DtoOut.SetDomain(domain); err != nil {
		return err
	}
	return nil
}
