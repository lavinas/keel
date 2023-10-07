package service

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

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
	if err := s.load(); err != nil {
		return err
	}
	s.prepareOutput(s.client, s.output)
	s.log.Infof(s.input, "get")
	return nil
}

// loadClient loads a client from the repository
func (s *ClientGet) load() error {
	maps := map[string]func(string) (bool, error){
		"id":       s.client.LoadById,
		"nickname": s.client.LoadByNick,
		"email":    s.client.LoadByEmail,
	}
	for _, value := range maps {
		found, err := value(s.param)
		if err != nil {
			return err
		}
		if found {
			return nil
		}
	}
	maps2 := map[string]func(uint64) (bool, error) {
		"document": s.client.LoadByDoc,
		"phone":    s.client.LoadByPhone,
	}
	iparam, err := toNumber(s.param)
	if err != nil {
		return nil
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
	s.log.Infof(s.input, "not found")
	return errors.New("not found")
}

// prepareOutput prepares the output data
func (s *ClientGet) prepareOutput(client port.Client, output port.ClientCreateOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}

// toNumber converts a string to uint64
func toNumber(number string) (uint64, error) {
	n := strings.Trim(number, " ")
	re := regexp.MustCompile(`[^0-9]`)
	n = re.ReplaceAllString(n, "")
	return strconv.ParseUint(n, 10, 64)
}