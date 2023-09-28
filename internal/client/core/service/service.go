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
}

// NewCreate creates a new Create service
func NewService(log port.Log, repo port.Repo) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

// Create creates a new client
func (s *Service) Create(input dto.CreateInputDto) (*port.CreateOutputDto, error) {
	if err := input.Validate(); err != nil {
		s.log.Infof(input, "bad request: " + err.Error())
		return nil, errors.New("bad request: " + err.Error())
	}
	input.Format()
	name, nick, doc, phone, email := input.GetName(), input.GetNickname(), input.GetDocument(), input.GetPhone(), input.GetEmail()
	client, _ := domain.NewClient(name, nick, doc, phone, email)
	if err := s.repo.Create(client); err != nil {
		s.log.Errorf(input, err)
		return nil, errors.New("internal server error")
	}
	s.log.Infof(input, "created")
	var output port.CreateOutputDto
	output.Fill(client.ID, name, nick, doc, phone, email)
	return &output, nil

}

