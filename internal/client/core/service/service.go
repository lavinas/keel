package service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	log port.Log
	repo port.Repo
	util port.Util
}

// NewCreate creates a new Create service
func NewService(log port.Log, repo port.Repo, util port.Util) *Service {
	return &Service{
		log: log,
		repo: repo,
		util: util,
	}
}

// Create creates a new client
func (s *Service) Create(input domain.CreateInputDto) (*domain.CreateOutputDto, error) {
	if !s.util.ValidateDocument(input.Document) {
		logInfo(s.log, input, "bad request: invalid document")		
		return nil, errors.New("bad request: invalid document")
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

func logError(log port.Log, input any, err error) {
	b, _ := json.Marshal(input)
	log.Error (err.Error() + " | " + string(b))
}

func logInfo(log port.Log, input any, message string) {
	b, _ := json.Marshal(input)
	log.Info (message + " | " + string(b))
}
