package service

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	domain port.Domain
	config port.Config
	log    port.Log
	repo   port.Repo
}

// NewCreate creates a new Create service
func NewService(domain port.Domain, config port.Config, log port.Log, repo port.Repo) *Service {
	return &Service{
		domain: domain,
		config: config,
		log:    log,
		repo:   repo,
	}
}

// ClientCreate is orquestration of Creating a new client
func (s *Service) ClientCreate(input port.ClientCreateInputDto, output port.ClientCreateOutputDto) error {
	service_client := NewClientCreate(s.log, s.domain.GetClient(input), input, output)
	return service_client.Execute()
}

// ClientList is orquestration of Updating a client
func (s *Service) ClientList(input port.ClientListInputDto, output port.ClientListOutputDto) error {
	service_client := NewClientList(s.config, s.log, s.domain.GetClientSet(), input, output)
	return service_client.Execute()
}

// ClientUpdate is orquestration of Updating a client
func (s *Service) ClientUpdate(id string, input port.ClientCreateInputDto, output port.ClientCreateOutputDto) error {
	service_client := NewClientUpdate(s.log, s.domain.GetClient(input), id, input, output)
	return service_client.Execute()
}
