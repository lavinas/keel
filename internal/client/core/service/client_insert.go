package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientInsert is the service for creating a new client
type ClientInsert struct {
	log    port.Log
	client port.Client
	input  port.ClientInsertInputDto
	output port.ClientInserOutputDto
}

// NewClientInsert creates a new client create service
func NewClientInsert(log port.Log, client port.Client, input port.ClientInsertInputDto, output port.ClientInserOutputDto) *ClientInsert {
	return &ClientInsert{
		log:    log,
		client: client,
		input:  input,
		output: output,
	}
}

// Execute executes the service
func (s *ClientInsert) Execute() error {
	if err := s.validateInput(); err != nil {
		return err
	}
	if err := s.loadClient(); err != nil {
		return err
	}
	if err := s.duplicity(); err != nil {
		return err
	}
	if err := s.store(); err != nil {
		return err
	}
	s.prepareOutput()
	s.log.Infof(s.input, "created")
	return nil
}

// loadClient loads a client from the input dto
func (s *ClientInsert) loadClient() error {
	s.input.Format()
	name, nick, doc, phone, email := s.input.Get()
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	iphone, _ := strconv.ParseUint(phone, 10, 64)
	if err := s.client.Insert(name, nick, idoc, iphone, email); err != nil {
		return err
	}
	return nil
}

// validateInput validates input data of Insert service
func (s *ClientInsert) validateInput() error {
	if err := s.input.Validate(); err != nil {
		s.log.Infof(s.input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// duplicity checks if a document or email is already registered
func (s *ClientInsert) duplicity() error {
	message := ""
	m, err := s.duplicityDocument()
	if err != nil {
		return err
	}
	message += m
	m, err = s.duplicityEmail()
	if err != nil {
		return err
	}
	message += m
	m, err = s.duplicityNick()
	if err != nil {
		return err
	}
	message += m
	if message != "" {
		message = strings.Trim(message, " |")
		s.log.Infof(s.input, "conflict: "+message)
		return errors.New("conflict: " + message)
	}
	return nil
}

// duplicityDocument treats the document duplicity
func (s *ClientInsert) duplicityDocument() (string, error) {
	b, err := s.client.DocumentDuplicity()
	if err != nil {
		s.log.Errorf(s.input, err)
		return "", errors.New("internal server error")
	}
	if b {
		return "document already registered | ", nil
	}
	return "", nil
}

// duplicityEmail treats the email duplicity
func (s *ClientInsert) duplicityEmail() (string, error) {
	e, err := s.client.EmailDuplicity()
	if err != nil {
		s.log.Errorf(s.input, err)
		return "", errors.New("internal server error |")
	}
	if e {
		return "email already registered | ", nil
	}
	return "", nil
}

func (s *ClientInsert) duplicityNick() (string, error) {
	n, err := s.client.NickDuplicity()
	if err != nil {
		s.log.Errorf(s.input, err)
		return "", errors.New("internal server error |")
	}
	if n {
		return "nickname already registered | ", nil
	}
	return "", nil
}

// store stores a new client
func (s *ClientInsert) store() error {
	// Store client
	if err := s.client.Save(); err != nil {
		s.log.Errorf(s.input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Insert service
func (s *ClientInsert) prepareOutput() {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}
