package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/lavinas/keel/internal/client/core/port"
)

const (
	per_page = "KEEL_PER_PAGE"
)

// Find is the service used to list all clients
type Find struct {
	log     port.Log
	clients port.ClientSet
	input   port.FindInputDto
	output  port.FindOutputDto
}

// NewFind creates a new Find
func NewFind(log port.Log, clients port.ClientSet, input port.FindInputDto, output port.FindOutputDto) *Find {
	return &Find{
		log:     log,
		clients: clients,
		input:   input,
		output:  output,
	}
}

// Execute executes the service to list all clients
func (s *Find) Execute() error {
	if err := s.validateInput(s.log, s.input); err != nil {
		return err
	}
	page, perPage, name, nick, doc, phone, email := s.getAll()
	if err := s.clients.Load(page, perPage, name, nick, doc, phone, email); err != nil {
		s.log.Error("Error loading clients: " + err.Error())
		return errors.New("internal error")
	}
	s.clients.SetOutput(s.output)
	s.log.Info(fmt.Sprintf("Listing %d clients", s.output.Count()))
	return nil
}

func (s *Find) validateInput(log port.Log, input port.FindInputDto) error {
	if err := input.Validate(); err != nil {
		log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// getPage returns the page and perPage values from the input dto
func (s *Find) getAll() (uint64, uint64, string, string, string, string, string) {
	page, perPage, name, nick, doc, phone, email := s.input.Get()
	if page == "" {
		page = "1"
	}
	p, _ := strconv.ParseUint(page, 10, 64)
	if perPage == "" {
		perPage = os.Getenv(per_page)
		if perPage == "" {
			perPage = "10"
		}
	}
	pp, err := strconv.ParseUint(perPage, 10, 64)
	if err != nil {
		pp = 10
	}
	return p, pp, name, nick, doc, phone, email
}
