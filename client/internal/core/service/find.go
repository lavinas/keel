package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lavinas/keel/client/internal/core/port"
)

const (
	per_page = "KEEL_CLIENT_PER_PAGE"
)

// Find is the service used to list all clients
type Find struct {
	config port.Config
	log     port.Log
	clients port.ClientSet
}

// NewFind creates a new Find
func NewFind(config port.Config, log port.Log, clients port.ClientSet) *Find {
	return &Find{
		config: config,
		log:     log,
		clients: clients,
	}
}

// Execute executes the service to list all clients
func (s *Find) Execute(input port.FindInputDto, output port.FindOutputDto) error {
	if err := s.validateInput(s.log, input); err != nil {
		return err
	}
	page, perPage, name, nick, doc, phone, email := s.getAll(input)
	if err := s.clients.Load(page, perPage, name, nick, doc, phone, email); err != nil {
		s.log.Error("Error loading clients: " + err.Error())
		return errors.New("internal error")
	}
	s.clients.SetOutput(output)
	s.log.Info(fmt.Sprintf("Found %d clients", output.Count()))
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
func (s *Find) getAll(input port.FindInputDto) (uint64, uint64, string, string, string, string, string) {
	page, perPage, name, nick, doc, phone, email := input.Get()
	if page == "" {
		page = "1"
	}
	p, _ := strconv.ParseUint(page, 10, 64)
	if perPage == "" {
		perPage = s.config.Get(per_page)
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
