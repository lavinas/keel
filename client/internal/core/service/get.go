package service

import (
	"errors"
	"strconv"

	"github.com/lavinas/keel/client/internal/core/port"
)

// Get is the service for getting a client
type Get struct {
	log    port.Log
	client port.Client
}

// NewGet creates a new client get service
func NewGet(log port.Log, client port.Client) *Get {
	return &Get{
		log:    log,
		client: client,
	}
}

// Execute executes the service
func (s *Get) Execute(param string, output port.InsertOutputDto) error {
	if param == "" {
		s.log.Info("bad request: blank param")
		return errors.New("bad request: blank param")
	}
	if err := s.load(param); err != nil {
		return err
	}
	s.prepareOutput(output)
	s.log.Info("get: " + param)
	return nil
}

// loadClient loads a client from the repository
func (s *Get) load(param string) error {
	maps := map[string]func(string) (bool, error){
		"id":       s.client.LoadById,
		"nickname": s.client.LoadByNick,
		"email":    s.client.LoadByEmail,
	}
	for _, funct := range maps {
		found, err := funct(param)
		if err != nil {
			return err
		}
		if found {
			return nil
		}
	}
	maps2 := map[string]func(uint64) (bool, error){
		"document": s.client.LoadByDoc,
		"phone":    s.client.LoadByPhone,
	}

	iparam, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		s.log.Info("not found x: " + param)
		return errors.New("not found x: " + param)
	}
	for _, value := range maps2 {
		found, err := value(iparam)
		if err != nil {
			return err
		}
		if found {
			return nil
		}
	}
	s.log.Info("no content: " + param + " not found")
	return errors.New("no content: " + param + " not found")
}

// prepareOutput prepares the output data
func (s *Get) prepareOutput(output port.InsertOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
