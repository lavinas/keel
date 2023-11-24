package usecase

import (
	"net/http"

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

func (s *UseCase) Register(domain port.Domain, result port.DefaultResult) {
	domain.SetBusinessID(s.config.Get(BUSINNESS_ID))
	if err := domain.Validate(); err != nil {
		s.logger.Info("xxxx client - create domain")
		result.Set(http.StatusInternalServerError, "internal error")
		return
	}
	if err := s.repo.Add(domain); err != nil {
		if s.repo.IsDuplicatedError(err) {
			s.logger.Info("xxxx client - add client")
			result.Set(http.StatusConflict, "client id already exists")
			return
		}
		s.logger.Info("xxxxx client - add client")
		result.Set(http.StatusInternalServerError, "internal error")
		return
	}
	s.logger.Info("xxxxx client - add client")
	result.Set(http.StatusCreated, "client created")
}

