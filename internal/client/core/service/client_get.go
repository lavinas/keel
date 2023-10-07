package service

import (
	"errors"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientGet is the service for getting a client
type ClientGet struct {
	log    port.Log
	client port.Client
	param  string
	input  port.ClientCreateInputDto
	output port.ClientCreateOutputDto
}

// NewClientGet creates a new client get service
func NewClientGet(log port.Log, client port.Client, param string, input port.ClientCreateInputDto, output port.ClientCreateOutputDto) *ClientGet {
	return &ClientGet{
		log:    log,
		client: client,
		param:  param,
		input:  input,
		output: output,
	}
}

// Execute executes the service
func (s *ClientGet) Execute() error {
	if s.param == "" {
		s.log.Infof(s.input, "bad request: blank param")
		return errors.New("bad request: blank param")
	}
	if err := s.loadClient(); err != nil {
		return err
	}
	s.prepareOutput(s.client, s.output)
	s.log.Infof(s.input, "get")
	return nil
}

// loadClient loads a client from the repository
func (s *ClientGet) loadClient() error {
	if err := s.client.LoadById(s.param); err != nil {
		s.log.Errorf(s.input, err)
		return err
	}
	return nil
}

// prepareOutput prepares the output data
func (s *ClientGet) prepareOutput(client port.Client, output port.ClientCreateOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}
