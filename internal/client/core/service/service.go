package service

import (
	"encoding/json"
	"errors"

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
	if _, m := s.util.ValidateAll(input.Name, input.Nickname, input.Document, input.Phone, input.Email); m != "" {
		return nil, errors.New("invalid input: " + m)
	}
	name, nick, doc, phone, email, _ := s.util.ClearAll(input.Name, input.Nickname, input.Document, input.Phone, input.Email)
	client := domain.NewClient(name, nick, doc, phone, email)
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
