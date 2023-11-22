package usecase

import (
	"fmt"
	"net/http"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// RegisterClient registers a new client
func (s *UseCase) Register(dto port.Register, result port.DefaultResult) {
	if !s.registerValidateDto(dto, result) {
		return
	}
	ok, client := s.registerCreateDomain(dto, result)
	if !ok {
		return
	}
	if !s.registerClientAddClient(dto, client, result) {
		return
	}
	s.logObj("info", "register client - success", "client registered", dto)
	result.Set(http.StatusCreated, "client registered")
}

// registerClientValidateDto validates the dto of register client usecase
func (s *UseCase) registerValidateDto(dto port.Register, result port.DefaultResult) bool {
	if err := dto.Validate(); err != nil {
		message := fmt.Sprintf("bad request: %s", err.Error())
		s.logObj("info", "register client - validate dto", message, dto)
		result.Set(http.StatusBadRequest, err.Error())
		return false
	}
	return true
}

// registerClientCreateDomain creates the domain of register client usecase
func (s *UseCase) registerCreateDomain(dto port.Register, result port.DefaultResult) (bool, interface{}) {
	businness_id := s.config.Get(BUSINNESS_ID)
	client := dto.GetDomain(businness_id)
	if err := client.Validate(); err != nil {
		message := fmt.Sprintf("internal error on validate: %s", err.Error())
		s.logObj("error", "register client - create domain", message, dto)
		result.Set(http.StatusInternalServerError, "internal error")
		return false, nil
	}
	return true, client
}

// registerClientAddClient adds the client of register client usecase
func (s *UseCase) registerClientAddClient(dto port.Register, obj interface{}, result port.DefaultResult) bool {
	if err := s.repo.Add(obj); err != nil {
		if s.repo.IsDuplicatedError(err) {
			message := fmt.Sprintf("conflict: %s", err.Error())
			s.logObj("info", "register client - add client", message, dto)
			result.Set(http.StatusConflict, "client id already exists")
			return false
		}
		message := fmt.Sprintf("internal error on add client: %s", err.Error())
		s.logObj("error", "register client - add client", message, dto)
		result.Set(http.StatusInternalServerError, "internal error")
		return false
	}
	return true
}
