package service

import (
	"github.com/lavinas/keel/client/internal/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	domain port.Domain
	log    port.Log
	insert Insert
	update Update
	find   Find
	get    Get
}

// NewInsert creates a new Insert service
func NewService(domain port.Domain, config port.Config, log port.Log) *Service {
	return &Service{
		domain: domain,

		log:    log,
		insert: *NewInsert(log, domain.GetClient()),
		update: *NewUpdate(log, domain.GetClient()),
		find:  *NewFind(config, log, domain.GetClientSet()),
		get:   *NewGet(log, domain.GetClient()),
	}
}

// Insert is orquestration of Creating a new client
func (s *Service) Insert(input port.InsertInputDto, output port.InsertOutputDto) error {
	return s.insert.Execute(input, output)
}

// Find is orquestration of Updating a client
func (s *Service) Find(input port.FindInputDto, output port.FindOutputDto) error {
	return s.find.Execute(input, output)
}

// Update is orquestration of Updating a client
func (s *Service) Update(id string, input port.UpdateInputDto, output port.UpdateOutputDto) error {
	return s.update.Execute(id, input, output)
}

// Get is orquestration of Getting a client
func (s *Service) Get(param string, paramType string, output port.InsertOutputDto) error {
	return s.get.Execute(param, paramType, output)
}
