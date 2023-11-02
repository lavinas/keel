package service

import (
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Service are services to orchestrate invoice domain
type Service struct {
	log      port.Log
	consumer port.RestConsumer
	domain   port.Domain
}

// NewCreate creates a new Create service
func NewService(log port.Log, consumer port.RestConsumer, domain port.Domain) *Service {
	return &Service{
		log:      log,
		consumer: consumer,
		domain:   domain,
	}
}

// Create is orquestration of Creating a new invoice
func (s *Service) Create(input port.CreateInputDto, output port.CreateOutputDto) error {
	service_invoice := NewCreate(s.log, s.consumer, s.domain.GetInvoice(), input, output)
	return service_invoice.Execute()
}
