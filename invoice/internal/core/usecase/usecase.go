package usecase

import (
	"net/http"
	"reflect"
	"time"

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

// Register registers a domain
func (s *UseCase) Register(domain port.Domain, result port.DefaultResult) {
	// Prepare domain
	name := "Register " + reflect.TypeOf(domain).String()
	domain.SetBusinessID(s.config.Get(BUSINNESS_ID))
	now := time.Now().In(time.Local)
	domain.SetCreatedAt(now)
	domain.SetUpdatedAt(now)
	domain.Fit()
	// Validate
	if err := domain.Validate(s.repo); err != nil {
		s.logger.Infof("%s - %s", name, err.Error())
		result.Set(http.StatusBadRequest, err.Error())
		return
	}
	// Add to repository
	if err := s.repo.Add(domain); err != nil {
		s.logger.Infof("%s - %s", name, err.Error())
		result.Set(http.StatusInternalServerError, "internal error")
		return
	}
	// Done
	s.logger.Infof("%s - %s", name, "Done")
	result.Set(http.StatusCreated, "created")
}
