package service

import (
	"errors"
	"reflect"
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
func (s *Get) Execute(param string, paramType string, output port.InsertOutputDto) error {
	if param == "" {
		s.log.Info("bad request: blank param - param: " + param + " - paramType: " + paramType)
		return errors.New("bad request: blank param")
	}
	ok, err := s.load(param, paramType)
	if err != nil {
		return err
	}
	if !ok {
		s.log.Info("no content: client not found - param: " + param + " - paramType: " + paramType)
		return errors.New("no content: client not found")
	}
	s.prepareOutput(output)
	s.log.Info("get: " + param)
	return nil
}

// loadClient loads a client from the repository
func (s *Get) load(param string, paramType string) (bool, error) {
	maps := map[string]interface{}{
		"id":       s.client.LoadById,
		"nickname": s.client.LoadByNick,
		"email":    s.client.LoadByEmail,
		"document": s.client.LoadByDoc,
		"phone":    s.client.LoadByPhone,
	}
	f := maps[paramType]
	if f == nil {
		s.log.Info("bad request: invalid param type")
		return false, errors.New("bad request: invalid param type - " + paramType + " - " + param)
	}
	if reflect.TypeOf(f) == reflect.TypeOf(s.client.LoadByDoc) {
		iparam, err := strconv.ParseUint(param, 10, 64)
		if err != nil {
			s.log.Info("bad request: param should be a number - param" + param + " - paramType: " + paramType)
			return false, errors.New("bad request: param should be a number")
		}
		return f.(func(uint64) (bool, error))(iparam)
	}
	return f.(func(string) (bool, error))(param)
}

// prepareOutput prepares the output data
func (s *Get) prepareOutput(output port.InsertOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
