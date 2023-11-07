package service

import (
	"errors"

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
		s.valiedateInput,
		s.loadDomain,
		s.checkDuplicity,
		s.saveDomain,
		s.updateInvoiceClients,
		s.createOutput,
	}
	for _, v := range execOrder {
		if err := v(); err != nil {
			return err
		}
	}
	s.log.Infof(s.input, "created")
	return nil
}

// valiedateInput is a method that validates the input
//
//	for the service
func (s *Create) valiedateInput() error {
	if err := s.input.Validate(); err != nil {
		err = errors.New("bad request: " + err.Error())
		s.log.Infof(s.input, "validate: "+err.Error())
		s.output.Load(err.Error(), "")
		return err
	}
	return nil
}

// loadDomain is a method that loads the domain for the service
func (s *Create) loadDomain() error {
	// load invoice
	if err := s.invoice.Load(s.input); err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "load: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
	}
	return nil
}

// checkDuplicity is a method that checks the duplicity
func (s *Create) checkDuplicity() error {
	if duplicated, err := s.invoice.IsDuplicated(); err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "load: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
	} else if duplicated {
		err = errors.New("conflict: duplicated invoice reference")
		s.log.Infof(s.input, "load: "+err.Error())
		s.output.Load(err.Error(), "")
		return err
	}
	return nil
}

// saveDomain is a method that saves the domain for the service
func (s *Create) saveDomain() error {
	if err := s.invoice.Save(); err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "save: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
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

// updateClients updates the clients of invoice after call the client service
func (s *Create) updateInvoiceClients() error {
	updateMap := map[string]func(port.GetClientByNicknameInputDto) error{
		s.input.GetBusinessNickname(): s.invoice.LoadBusiness,
		s.input.GetCustomerNickname(): s.invoice.LoadCustomer,
	}
	for k, f := range updateMap {
		dto, err := s.getInvoiceClientDto(k)
		if err != nil {
			return err
		}
		if err := s.loadInvoiceClient(f, dto); err != nil {
			return err
		}
	}
	if err := s.updateInvoice(); err != nil {
		return err
	}
	return nil
}

// updateClientInvoice  updates the client invoice after consulting the external service
func (s *Create) getInvoiceClientDto(nickname string) (port.GetClientByNicknameInputDto, error) {
	dto := dto.NewGetClientByNicknameInputDto()
	ok, err := s.consumer.GetClientByNickname(nickname, dto)
	if err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "getting "+nickname+": "+err.Error())
		s.output.Load(rerr.Error(), "")
		return nil, rerr
	}
	if !ok {
		return nil, nil
	}
	return dto, nil
}

// loadInvoice loads the invoice client
func (s *Create) loadInvoiceClient(f func(port.GetClientByNicknameInputDto) error, dto port.GetClientByNicknameInputDto) error {
	if err := f(dto); err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "update consumer: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
	}
	return nil
}

// updateInvoice updates the invoice
func (s *Create) updateInvoice() error {
	if err := s.invoice.Update(); err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "update invoice: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
	}
	return nil
}
