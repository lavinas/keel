package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientUpdate is the service for updating a client
type ClientUpdate struct {
	log    port.Log
	client port.Client
	id     string
	input  port.ClientCreateInputDto
	output port.ClientCreateOutputDto
}

// NewClientUpdate creates a new client update service
func NewClientUpdate(log port.Log, client port.Client, id string, input port.ClientCreateInputDto, output port.ClientCreateOutputDto) *ClientUpdate {
	return &ClientUpdate{
		log:    log,
		client: client,
		id:     id,
		input:  input,
		output: output,
	}
}

// Execute executes the service
func (s *ClientUpdate) Execute() error {
	if s.id == "" {
		s.log.Infof(s.input, "bad request: blnk id")
		return errors.New("bad request: blank id")
	}
	if s.input.IsBlank() {
		s.log.Infof(s.input, "bad request: blank input ")
		return errors.New("bad request: blank input")
	}
	if err := s.validateInput(); err != nil {
		return err
	}
	if err := s.loadClient(); err != nil {
		return err
	}
	if err := s.duplicity(); err != nil {
		return err
	}
	if err := s.update(); err != nil {
		return err
	}
	s.prepareOutput()
	s.log.Infof(s.input, "updated")
	return nil
}

// validateInput validates input data of Update service
func (s *ClientUpdate) validateInput() error {
	if err := s.input.ValidateUpdate(); err != nil {
		s.log.Infof(s.input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// loadClient loads a client from repository
func (s *ClientUpdate) loadClient() error {
	if err := s.input.FormatUpdate(); err != nil {
		s.log.Infof(s.input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	if err := s.client.LoadById(s.id); err != nil {
		s.log.Infof(s.input, "not found: "+err.Error())
		return errors.New("not found: " + err.Error())
	}
	return nil
}

// duplicity checks if a document or email is already registered
func (s *ClientUpdate) duplicity() error {
	message := ""
	_, nick, doc, _, email := s.input.Get()
	if strings.Trim(doc, " ") != "" {
		m, err := s.duplicityDocument()
		if err != nil {
			return err
		}
		message += m
	}
	if strings.Trim(email, " ") != "" {
		m, err := s.duplicityEmail()
		if err != nil {
			return err
		}
		message += m
	}
	if strings.Trim(nick, " ") != "" {
		m, err := s.duplicityNick()
		if err != nil {
			return err
		}
		message += m
	}
	if message != "" {
		message = strings.Trim(message, " |")
		s.log.Infof(s.input, "conflict: "+message)
		return errors.New("conflict: " + message)
	}
	return nil
}

// duplicityDocument treats the document duplicity
func (s *ClientUpdate) duplicityDocument() (string, error) {
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
func (s *ClientUpdate) duplicityEmail() (string, error) {
	e, err := s.client.EmailDuplicity()
	if err != nil {
		s.log.Errorf(s.input, err)
		return "", errors.New("internal server error |")
	}
	if e {
		return "email already registered", nil
	}
	return "", nil
}

// duplicityNick treats the nickname duplicity
func (s *ClientUpdate) duplicityNick() (string, error) {
	n, err := s.client.NickDuplicity()
	if err != nil {
		s.log.Errorf(s.input, err)
		return "", errors.New("internal server error |")
	}
	if n {
		return "nickname already registered", nil
	}
	return "", nil
}

// update updates a client
func (s *ClientUpdate) update() error {
	_, uname, unick, udoc, uphone, uemail := s.client.Get()
	name, nick, doc, phone, email := s.input.Get()
	if strings.Trim(name, " ") != "" {
		uname = name
	}
	if strings.Trim(nick, " ") != "" {
		unick = nick
	}
	if strings.Trim(doc, " ") != "" {
		udoc, _ = strconv.ParseUint(doc, 10, 64)
	}
	if strings.Trim(phone, " ") != "" {
		uphone, _ = strconv.ParseUint(phone, 10, 64)
	}
	if strings.Trim(email, " ") != "" {
		uemail = email
	}
	s.client.Load(s.id, uname, unick, udoc, uphone, uemail)
	if err := s.client.Update(); err != nil {
		s.log.Errorf(s.input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Update service
func (s *ClientUpdate) prepareOutput() {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}
