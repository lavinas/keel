package usecase

import (
	"net/http"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Register registers a input dto
func (s *UseCase) Register(dto port.Register, result port.DefaultResult) {
	if err := dto.Validate(); err != nil {
		s.logger.Info("register client - validate dto")
		result.Set(http.StatusBadRequest, err.Error())
		return
	}
	domain := dto.GetDomain(s.config.Get(BUSINNESS_ID))
	if err := domain.Validate(); err != nil {
		s.logger.Info("register client - create domain")
		result.Set(http.StatusInternalServerError, "internal error")
		return
	}
	if err := s.repo.Add(domain); err != nil {
		if s.repo.IsDuplicatedError(err) {
			s.logger.Info("register client - add client")
			result.Set(http.StatusConflict, "client id already exists")
			return
		}
		s.logger.Info("register client - add client")
		result.Set(http.StatusInternalServerError, "internal error")
		return
	}
}
