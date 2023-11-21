package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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
func (s *UseCase) RegisterClient(dto port.RegisterClient, result port.DefaultResult) {
	if !s.registerClientValidateDto(dto, result) {
		return
	}
	ok, client := s.registerClientCreateDomain(dto, result)
	if !ok {
		return
	}
	if !s.registerClientAddClient(dto, client, result) {
		return
	}
	message := fmt.Sprintf("client registered: %s", client.ID)
	s.logObj("info", "register client - success", message, dto)
	result.Set(http.StatusCreated, "client registered")
}


// registerClientValidateDto validates the dto of register client usecase
func (s *UseCase) registerClientValidateDto(dto port.RegisterClient, result port.DefaultResult) bool {
	if err := dto.Validate(); err != nil {
		message := fmt.Sprintf("bad request: %s", err.Error())
		s.logObj("info", "register client - validate dto", message, dto)
		result.Set(http.StatusBadRequest, err.Error())
		return false
	}
	return true
}

// registerClientCreateDomain creates the domain of register client usecase
func (s *UseCase) registerClientCreateDomain(dto port.RegisterClient, result port.DefaultResult) (bool, *domain.Client){
	id, name, email, document, phone := dto.Get()
	client := domain.NewClient(id, name, email, s.strToUint64(document), s.strToUint64(phone),
		time.Time{}, time.Time{})
	if err := client.Validate(); err != nil {
		message := fmt.Sprintf("internal error on validate: %s", err.Error())
		s.logObj("error", "register client - create domain", message, dto)
		result.Set(http.StatusInternalServerError, "internal error")
		return false, nil
	}
	return true, client
}

// registerClientAddClient adds the client of register client usecase
func (s *UseCase) registerClientAddClient(dto port.RegisterClient, client *domain.Client, result port.DefaultResult) bool{
	if err := s.repo.AddClient(client); err != nil {
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

// strToUint64 converts a string to uint64
func (s *UseCase) strToUint64(str string) uint64 {
	re := regexp.MustCompile(`[^0-9]`)
	str = re.ReplaceAllString(str, "")
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

// logInfoObj logs an info message with an object
func (s *UseCase) logObj(logType string, prefix string, message string, obj any) {
	dto_log, _ := json.Marshal(obj)
	dto_str := string(dto_log)
	if logType == "info" {
		s.logger.Infof("%s | %s | %s", prefix, message, dto_str)
	} else if logType == "error" {
		s.logger.Errorf("%s: %s", message, dto_str)
	}
}

