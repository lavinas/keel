package service

import (
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Service are services to orchestrate invoice domain
type Service struct {
	create *Create
}

// NewCreate creates a new Create service
func NewService(repo port.Repo, log port.Log, consumer port.RestConsumer, domain port.Domain) *Service {
	return &Service{
		create: NewCreate(repo, log, consumer, domain.GetInvoice()),
	}
}

// Create is orquestration of Creating a new invoice
func (s *Service) Create(input port.CreateInputDto, output port.CreateOutputDto) error {
	return s.create.Execute(input, output)
}
