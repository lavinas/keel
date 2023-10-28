package service

import (
	"errors"

	"github.com/lavinas/keel/internal/invoice/core/port"
)

// Create is a service that creates a new invoice
type Create struct {
	log port.Log
	invoice port.Invoice
	input port.CreateInputDto
	output port.CreateOutputDto
}

// NewCreate is a factory that creates a new Create service
func NewCreate(log port.Log, invoice port.Invoice, input port.CreateInputDto, output port.CreateOutputDto) *Create {
	return &Create{
		log: log,
		invoice: invoice,
		input: input,
		output: output,
	}
}

// Execute is a method that executes the service
func (s *Create) Execute() error {
	execMap := map[string]func() error{
		"validate": s.valiedateInput,
		"load": s.loadDomain,
		"save": s.saveDomain,
		"output": s.createOutput,
	}
	for _, v := range execMap {
		if err := v(); err != nil {
			return err
		}
	}
	return nil 
}

// valiedateInput is a method that validates the input for the service
func (s *Create) valiedateInput() error {
	if err := s.input.Validate(); err != nil {
		err = errors.New("bad request: " + err.Error())
		s.log.Infof(s.input, "validate: " + err.Error())
		s.output.Load(err.Error(), "")
		return err
	}
	return nil
}

// loadDomain is a method that loads the domain for the service
func (s *Create) loadDomain() error {
	// load invoice
	if err := s.invoice.Load(s.input); err != nil {
		err = errors.New("internal error: " + err.Error())
		s.log.Infof(s.input, "load: " + err.Error())
		s.output.Load(err.Error(), "")
		return err
	}
	return nil
}

// saveDomain is a method that saves the domain for the service
func (s *Create) saveDomain() error {
	if err := s.invoice.Save(); err != nil {
		err = errors.New("internal error: " + err.Error())
		s.log.Infof(s.input, "save: " + err.Error())
		s.output.Load(err.Error(), "")
		return err
	}
	return nil
}

// createOutput is a method that creates the output for the service
func (s *Create) createOutput() error {
	s.output.Load("created", s.invoice.GetReference())
	return nil
}
