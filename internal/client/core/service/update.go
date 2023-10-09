package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

// Update is the service for updating a client
type Update struct {
	log    port.Log
	client port.Client
	id     string
	input  port.InsertInputDto
	output port.InsertOutputDto
}

// NewUpdate creates a new client update service
func NewUpdate(log port.Log, client port.Client, id string, input port.UpdateInputDto, output port.UpdateOutputDto) *Update {
	return &Update{
		log:    log,
		client: client,
		id:     id,
		input:  input,
		output: output,
	}
}

// Execute executes the service
func (s *Update) Execute() error {
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
func (s *Update) validateInput() error {
	if err := s.input.Validate(); err != nil {
		s.log.Infof(s.input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// loadClient loads a client from repository
func (s *Update) loadClient() error {
	if err := s.input.Format(); err != nil {
		s.log.Infof(s.input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	result, err := s.client.LoadById(s.id)
	if err != nil {
		s.log.Infof(s.input, err.Error())
		return errors.New("internal error")
	}
	if !result {
		s.log.Infof(s.input, "not found")
		return errors.New("not found")
	}
	return nil
}

// duplicity checks if a document or email is already registered
func (s *Update) duplicity() error {
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
func (s *Update) duplicityDocument() (string, error) {
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
func (s *Update) duplicityEmail() (string, error) {
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
func (s *Update) duplicityNick() (string, error) {
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
func (s *Update) update() error {
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
func (s *Update) prepareOutput() {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	s.output.Fill(id, name, nick, doc, phone, email)
}
