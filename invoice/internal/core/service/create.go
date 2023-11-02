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
	execMap := map[string]func() error{
		"validate":       s.valiedateInput,
		"load":           s.loadDomain,
		"duplicity":      s.checkDuplicity,
		"save":           s.saveDomain,
		"updateBusiness": s.updateBusinnes,
		"updateConsumer": s.updateConsumer,
		"output":         s.createOutput,
	}
	for _, v := range execMap {
		if err := v(); err != nil {
			return err
		}
	}
	s.log.Infof(s.input, "created")
	return nil
}

// valiedateInput is a method that validates the input for the service
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

// createOutput is a method that creates the output for the service
func (s *Create) createOutput() error {
	s.output.Load("created", s.invoice.GetReference())
	return nil
}

// updateBusinnes updates the business client invoice after consulting the external service
func (s *Create) updateBusinnes() error {
	return s.updateClient(s.invoice.GetBusiness())
}

// updateConsumer updates the consumer client invoice after consulting the external service
func (s *Create) updateConsumer() error {
	return s.updateClient(s.invoice.GetConsumer())
}

// updateClientInvoice is a method that updates the client invoice after consulting the external service
func (s *Create) updateClient(client port.InvoiceClient) error {
	dto := dto.NewGetClientByNicknameInputDto()
	nickname := client.GetNickname()
	ok, err := s.consumer.GetClientByNickname(nickname, dto)
	if err != nil {
		rerr := errors.New("internal error")
		s.log.Infof(s.input, "update: "+err.Error())
		s.output.Load(rerr.Error(), "")
		return rerr
	}
	if !ok {
		return nil
	}
	client.LoadGetClientNicknameDto(dto)
	return nil
}
