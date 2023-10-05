package service

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	domain port.Domain
	log    port.Log
	repo   port.Repo
}

// NewCreate creates a new Create service
func NewService(domain port.Domain, log port.Log, repo port.Repo) *Service {
	return &Service{
		domain: domain,
		log:    log,
		repo:   repo,
	}
}

// Orquestration of Creating a new client
func (s *Service) ClientCreate(input port.ClientCreateInputDto, output port.ClientCreateOutputDto) error {
	service_client := NewClientCreate(s.log, s.domain.GetClient(input), input, output)
	return service_client.Execute()
}

func (s *Service) ClientList(output port.ClientListOutputDto) error {
	service_client := NewClientList(s.log, s.domain.GetClientSet(), output)
	return service_client.Execute()
}
