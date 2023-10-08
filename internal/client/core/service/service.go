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

// NewInsert creates a new Insert service
func NewService(domain port.Domain, config port.Config, log port.Log, repo port.Repo) *Service {
	return &Service{
		domain: domain,
		config: config,
		log:    log,
		repo:   repo,
	}
}

// ClientInsert is orquestration of Creating a new client
func (s *Service) ClientInsert(input port.ClientInsertInputDto, output port.ClientInserOutputDto) error {
	service_client := NewClientInsert(s.log, s.domain.GetClient(), input, output)
	return service_client.Execute()
}

// ClientList is orquestration of Updating a client
func (s *Service) ClientList(input port.ClientListInputDto, output port.ClientListOutputDto) error {
	service_client := NewClientList(s.config, s.log, s.domain.GetClientSet(), input, output)
	return service_client.Execute()
}

// ClientUpdate is orquestration of Updating a client
func (s *Service) ClientUpdate(id string, input port.ClientInsertInputDto, output port.ClientInserOutputDto) error {
	service_client := NewClientUpdate(s.log, s.domain.GetClient(), id, input, output)
	return service_client.Execute()
}

// ClientGet is orquestration of Getting a client
func (s *Service) ClientGet(param string, input port.ClientInsertInputDto, output port.ClientInserOutputDto) error {
	service_client := NewClientGet(s.log, s.domain.GetClient(), param, output)
	return service_client.Execute()
}
