package service

import (
	"github.com/lavinas/keel/internal/invoice/core/port"
)

// Service are services to orchestrate invoice domain
type Service struct {
	log    port.Log
	domain port.Domain
}

// NewCreate creates a new Create service
func NewService(log port.Log, domain port.Domain) *Service {
	return &Service{
		log:    log,
		domain: domain,
	}
}

// Create is orquestration of Creating a new invoice
func (s *Service) Create(input port.CreateInputDto, output port.CreateOutputDto) error {
	service_invoice := NewCreate(s.log, s.domain.GetInvoice(), input, output)
	return service_invoice.Execute()
}




