package service

import (
	"errors"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/dto"
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
func (s *Service) Create(input dto.CreateInputDto) (*dto.CreateOutputDto, error) {
	if _, m := s.util.ValidateAll(input.Name, input.Nickname, input.Document, input.Phone, input.Email); m != "" {
		s.log.Infof(input, "bad request: " + m)
		return nil, errors.New("bad request: " + m)
	}
	name, nick, doc, phone, email, _ := s.util.ClearAll(input.Name, input.Nickname, input.Document, input.Phone, input.Email)
	client := domain.NewClient(name, nick, doc, phone, email)
	if err := s.repo.Create(client); err != nil {
		s.log.Errorf(input, err)
		return nil, errors.New("internal server error: ")
	}
	s.log.Infof(input, "created")
	return &dto.CreateOutputDto{
		Id:       client.ID,
		Name:     client.Name,
		Nickname: client.Nickname,
		Document: client.Document,
		Phone:    client.Phone,
		Email:    client.Email,
	}, nil
}

// ListAll list all clients
func (l *Service) ListAll() (*dto.ListAllOutputDto, error) {
	c := dto.CreateOutputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: 12321232222,
		Phone:    11999999999,
		Email:    "test@test.com.br",
	}
	r := dto.ListAllOutputDto{
		Clients: []dto.CreateOutputDto{c},
	}
	return &r, nil
}
