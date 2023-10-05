package service

import (
	"errors"
	"fmt"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientList is the service used to list all clients
type ClientList struct {
	log     port.Log
	clients port.ClientSet
	output  port.ClientListOutputDto
}

// NewClientList creates a new ClientList
func NewClientList(log port.Log, clients port.ClientSet, output port.ClientListOutputDto) *ClientList {
	return &ClientList{
		log:     log,
		clients: clients,
		output:  output,
	}
}

// Execute executes the service to list all clients
func (s *ClientList) Execute() error {
	if err := s.clients.Load(); err != nil {
		s.log.Error("Error loading clients")
		return errors.New("internal error")
	}
	s.clients.SetOutput(s.output)
	s.log.Info(fmt.Sprintf("Listing %d clients", s.output.Count()))
	return nil
}
