package service

import (
	"errors"
	"strconv"

	"github.com/lavinas/keel/internal/client/core/port"
)

// Get is the service for getting a client
type Get struct {
	log    port.Log
	client port.Client
	param  string
	output port.InsertOutputDto
}

// NewGet creates a new client get service
func NewGet(log port.Log, client port.Client, param string, output port.InsertOutputDto) *Get {
	return &Get{
		log:    log,
		client: client,
		param:  param,
		output: output,
	}
}

// Execute executes the service
func (s *Get) Execute() error {
	if s.param == "" {
		s.log.Info("bad request: blank param")
		return errors.New("bad request: blank param")
	}
	if err := s.load(); err != nil {
		return err
	}
	s.prepareOutput()
	s.log.Info("get: " + s.param)
	return nil
}

// loadClient loads a client from the repository
func (s *Get) load() error {
	maps := map[string]func(string) (bool, error){
		"id":       s.client.LoadById,
		"nickname": s.client.LoadByNick,
		"email":    s.client.LoadByEmail,
	}
	for _, funct := range maps {
		found, err := funct(s.param)
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

	iparam, err := strconv.ParseUint(s.param, 10, 64)
	if err != nil {
		s.log.Info("not found: " + s.param)
		return errors.New("not found: " + s.param)
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
	s.log.Info("no content: " + s.param + " not found")
	return errors.New("no content: " + s.param + " not found")
}

// prepareOutput prepares the output data
func (s *Get) prepareOutput() {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}
