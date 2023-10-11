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

// Insert is orquestration of Creating a new client
func (s *Service) Insert(input port.InsertInputDto, output port.InsertOutputDto) error {
	service_client := NewInsert(s.log, s.domain.GetClient(), input, output)
	return service_client.Execute()
}

// Find is orquestration of Updating a client
func (s *Service) Find(input port.FindInputDto, output port.FindOutputDto) error {
	service_client := NewFind(s.config, s.log, s.domain.GetClientSet(), input, output)
	return service_client.Execute()
}

// Update is orquestration of Updating a client
func (s *Service) Update(id string, input port.UpdateInputDto, output port.UpdateOutputDto) error {
	service_client := NewUpdate(s.log, s.domain.GetClient(), id, input, output)
	return service_client.Execute()
}

// Get is orquestration of Getting a client
func (s *Service) Get(param string, input port.InsertInputDto, output port.InsertOutputDto) error {
	service_client := NewGet(s.log, s.domain.GetClient(), param, output)
	return service_client.Execute()
}
