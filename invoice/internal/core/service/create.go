package service

import (
	"errors"
	"strings"
	"fmt"

	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Create is a service that creates a new invoice
type Create struct {
	log      port.Log
	consumer port.RestConsumer
	invoice  port.Invoice
	input    port.CreateInputDto
	output   port.CreateOutputDto
}

// NewCreate is a factory that creates a new Create service
func NewCreate(log port.Log, consumer port.RestConsumer, invoice port.Invoice, input port.CreateInputDto, output port.CreateOutputDto) *Create {
	return &Create{
		log:      log,
		consumer: consumer,
		invoice:  invoice,
		input:    input,
		output:   output,
	}
}

// Execute is a method that executes the service
func (s *Create) Execute() error {
	execOrder := []func() error{
		s.formatInput,
		s.loadDomain,
		s.checkDuplicity,
		s.saveDomain,
		s.updateInvoiceBusiness,
		s.updateInvoiceCustomer,
		s.createOutput,
	}
	for _, v := range execOrder {
		if err := v(); err != nil {
			s.log.Errorf(s.input, err)
			if strings.Contains(err.Error(), "internal") {
				err = errors.New("internal error")
			}
			s.output.Load(err.Error(), "")
			return err
		}
	}
	s.log.Infof(s.input, "created")
	return nil
}

// valiedateInput is a method that validates the input
//
//	for the service
func (s *Create) formatInput() error {
	if err := s.input.Format(); err != nil {
		return fmt.Errorf("bad request: %w", err)
	}
	return nil
}

// loadDomain is a method that loads the domain for the service
func (s *Create) loadDomain() error {
	// load invoice
	if err := s.invoice.Load(s.input); err != nil {
		return fmt.Errorf("internal error: %w", err)
	}
	return nil
}

// checkDuplicity is a method that checks the duplicity
func (s *Create) checkDuplicity() error {
	if duplicated, err := s.invoice.IsDuplicated(); err != nil {
		return fmt.Errorf("internal error: %w", err)
	} else if duplicated {
		return fmt.Errorf("conflict: duplicated invoice reference")
	}
	return nil
}

// saveDomain is a method that saves the domain for the service
func (s *Create) saveDomain() error {
	if err := s.invoice.Save(); err != nil {
		return fmt.Errorf("internal error: %w", err)
	}
	return nil
}

// createOutput is a method that creates the output
//
//	for the service
func (s *Create) createOutput() error {
	s.output.Load("created", s.invoice.GetReference())
	return nil
}

// updateInvoiceBusiness updates the business client of invoice after call the client service
func (s *Create) updateInvoiceBusiness() error {
	if err := s.updateClient(s.invoice.GetBusiness()); err != nil {
		return err
	}
	return nil
}

// updateInvoiceCustomer updates the customer client of invoice after call the client service
func (s *Create) updateInvoiceCustomer() error {
	if err := s.updateClient(s.invoice.GetCustomer()); err != nil {
		return err
	}
	return nil
}

// updateClient updates the client after consulting the external service
func (s *Create) updateClient(client port.InvoiceClient) error {
	if !client.IsNew() {
		return nil
	}
	dto, err := s.getInvoiceClientDto(client.GetNickname())
	if err != nil {
		return err
	}
	if err := client.LoadGetClientNicknameDto(dto); err != nil {
		return fmt.Errorf("internal error: %w", err)
	}
	if err := client.Update(); err != nil {
		return fmt.Errorf("internal error: %w", err)
	}
	return nil
}

// updateClientInvoice  updates the client invoice after consulting the external service
func (s *Create) getInvoiceClientDto(nickname string) (port.GetClientByNicknameInputDto, error) {
	dto := dto.NewGetClientByNicknameInputDto()
	ok, err := s.consumer.GetClientByNickname(nickname, dto)
	if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	if !ok {
		return nil, nil
	}
	return dto, nil
}