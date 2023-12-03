package usecase

import (
	"net/http"
	"reflect"

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
func (s *UseCase) Create(domain port.Domain, result port.DefaultResult) {
	// Prepare domain
	name := "Creating " + reflect.TypeOf(domain).String()
	domain.SetCreate(s.config.Get(BUSINNESS_ID))
	// Validate
	if err := domain.Validate(s.repo); err != nil {
		s.logger.Infof("%s - %s", name, err.Error())
		result.Set(err.GetHTTPCode(), err.Error())
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
