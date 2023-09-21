package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	log  port.Log
	repo port.Repo
	util port.Util
}

// NewCreate creates a new Create service
func NewService(log port.Log, repo port.Repo, util port.Util) *Service {
	return &Service{
		log:  log,
		repo: repo,
		util: util,
	}
}

// Create creates a new client
func (s *Service) Create(input domain.CreateInputDto) (*domain.CreateOutputDto, error) {
	if err := validate(input, s.util, s.log); err != nil {
		return nil, err
	}

	d, _ := strconv.ParseUint(s.util.ClearNumber(input.Document), 10, 64)
	p, _ := strconv.ParseUint(s.util.ClearNumber(input.Phone), 10, 64)

	client := domain.NewClient(input.Name, input.Nickname, p, d, input.Email)
	if err := s.repo.Create(client); err != nil {
		logError(s.log, input, err)
		return nil, errors.New("internal server error: ")
	}
	logInfo(s.log, input, "created")
	return &domain.CreateOutputDto{
		Id:       client.ID,
		Name:     client.Name,
		Nickname: client.Nickname,
		Document: client.Document,
		Phone:    client.Phone,
		Email:    client.Email,
	}, nil
}

// ListAll list all clients
func (l *Service) ListAll() (*domain.ListAllOutputDto, error) {
	c := domain.CreateOutputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: 12321232222,
		Phone:    11999999999,
		Email:    "test@test.com.br",
	}
	r := domain.ListAllOutputDto{
		Clients: []domain.CreateOutputDto{c},
	}
	return &r, nil
}

// logError format a error for logging
func logError(log port.Log, input any, err error) {
	b, _ := json.Marshal(input)
	log.Error(err.Error() + " | " + string(b))
}

// logError format a error for logging
func logInfo(log port.Log, input any, message string) {
	b, _ := json.Marshal(input)
	log.Info(message + " | " + string(b))
}

// validate input data
func validate(input domain.CreateInputDto, util port.Util, log port.Log) error {
	var message string = ""
	message += validateName(input.Name) + " "
	message += validateNickname(input.Nickname) + " "
	message += validateDocument(input.Document, util) + " "
	message += validatePhone(input.Phone) + " "
	message += validateEmail(input.Email, util) + " "

	if message != "" {
		message = "bad request: " + message
		logError(log, input, errors.New(message))
		return errors.New(message)
	}
	return nil
}

func validateDocument(document string, util port.Util) string {
	if strings.Trim(document, " ") == "" {
		return "document is blank"
	}
	if !util.ValidateDocument(document) {
		return "invalid document"
	}
	return ""
}

func validateEmail(email string, util port.Util) string {
	if strings.Trim(email, " ") == "" {
		return "email is blank"
	}
	if !util.ValidateEmail(email) {
		return "invalid email"
	}
	return ""
}

func validateName(name string) string {
	if strings.Trim(name, " ") == "" {
		return "name is blank"
	}
	return ""
}

func validateNickname(nickname string) string {
	if strings.Trim(nickname, " ") == "" {
		return "nickname is blank"
	}
	return ""
}

func validatePhone(phone string) string {
	if strings.Trim(phone, " ") == "" {
		return "phone is blank"
	}
	return ""

}
