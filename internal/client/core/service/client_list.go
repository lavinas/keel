package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientList is the service used to list all clients
type ClientList struct {
	config  port.Config
	log     port.Log
	clients port.ClientSet
	input   port.ClientListInputDto
	output  port.ClientListOutputDto
}

// NewClientList creates a new ClientList
func NewClientList(config port.Config, log port.Log, clients port.ClientSet, input port.ClientListInputDto, output port.ClientListOutputDto) *ClientList {
	return &ClientList{
		config:  config,
		log:     log,
		clients: clients,
		input:   input,
		output:  output,
	}
}

// Execute executes the service to list all clients
func (s *ClientList) Execute() error {
	if err := s.validateInput(s.log, s.input); err != nil {
		return err
	}
	page, perPage := s.getPage(s.config, s.input)
	if err := s.clients.Load(page, perPage); err != nil {
		s.log.Error("Error loading clients")
		return errors.New("internal error")
	}
	s.clients.SetOutput(s.output)
	s.log.Info(fmt.Sprintf("Listing %d clients", s.output.Count()))
	return nil
}

func (s *ClientList) validateInput(log port.Log, input port.ClientListInputDto) error {
	if err := input.Validate(); err != nil {
		log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// getPage returns the page and perPage values from the input dto
func (s *ClientList) getPage(config port.Config, input port.ClientListInputDto) (uint64, uint64) {
	page, perPage, _, _, _, _, _ := input.Get()
	if page == "" {
		page = "1"
	}
	p, _ := strconv.ParseUint(page, 10, 64)
	if perPage == "" {
		var err error
		perPage, err = config.GetField("rest", "per_page")
		if err != nil {
			perPage = "10"
		}
	}
	pp, err := strconv.ParseUint(perPage, 10, 64)
	if err != nil {
		pp = 10
	}
	return p, pp
}
