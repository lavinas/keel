package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lavinas/keel/client/internal/core/port"
)

// Update is the service for updating a client
type Update struct {
	log    port.Log
	client port.Client
}

// NewUpdate creates a new client update service
func NewUpdate(log port.Log, client port.Client) *Update {
	return &Update{
		log:    log,
		client: client,
	}
}

// Execute executes the service
func (s *Update) Execute(id string, input port.UpdateInputDto, output port.UpdateOutputDto) error {
	if err := s.validateId(id, input); err != nil {
		return err
	}
	if err := s.validateInput(input); err != nil {
		return err
	}
	if err := s.loadClient(id, input); err != nil {
		return err
	}
	if err := s.duplicity(input); err != nil {
		return err
	}
	if err := s.update(id, input); err != nil {
		return err
	}
	s.prepareOutput(output)
	s.log.Infof(input, "updated")
	return nil
}

// validateId validates id of Update service
func (s *Update) validateId(id string, input port.UpdateInputDto) error {
	if id == "" {
		s.log.Infof(input, "bad request: blank id")
		return errors.New("bad request: blank id")
	}
	if _, err := uuid.Parse(id); err != nil {
		s.log.Infof(input, "bad request: invalid id "+id)
		return errors.New("bad request: invalid id")
	}
	return nil
}

// validateInput validates input data of Update service
func (s *Update) validateInput(input port.UpdateInputDto) error {
	if input.IsBlank() {
		s.log.Infof(input, "bad request: blank input ")
		return errors.New("bad request: blank input")
	}
	if err := input.Validate(); err != nil {
		s.log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	if err := input.Format(); err != nil {
		s.log.Infof(input, err.Error())
		return errors.New("internal error")
	}
	return nil
}

// loadClient loads a client from repository
func (s *Update) loadClient(id string, input port.UpdateInputDto) error {
	result, err := s.client.LoadById(id)
	if err != nil {
		s.log.Infof(input, err.Error())
		return errors.New("internal error")
	}
	if !result {
		s.log.Infof(input, "not found")
		return errors.New("not found")
	}
	return nil
}

// duplicity checks if a document or email is already registered
func (s *Update) duplicity(input port.UpdateInputDto) error {
	message := ""
	_, nick, doc, _, email := input.Get()
	if strings.Trim(doc, " ") != "" {
		m, err := s.duplicityDocument(input)
		if err != nil {
			return err
		}
		message += m
	}
	if strings.Trim(email, " ") != "" {
		m, err := s.duplicityEmail(input)
		if err != nil {
			return err
		}
		message += m
	}
	if strings.Trim(nick, " ") != "" {
		m, err := s.duplicityNick(input)
		if err != nil {
			return err
		}
		message += m
	}
	if message != "" {
		message = strings.Trim(message, " |")
		s.log.Infof(input, "conflict: "+message)
		return errors.New("conflict: " + message)
	}
	return nil
}

// duplicityDocument treats the document duplicity
func (s *Update) duplicityDocument(input port.UpdateInputDto) (string, error) {
	b, err := s.client.DocumentDuplicity()
	if err != nil {
		s.log.Errorf(input, err)
		return "", errors.New("internal server error")
	}
	if b {
		return "document already registered | ", nil
	}
	return "", nil
}

// duplicityEmail treats the email duplicity
func (s *Update) duplicityEmail(input port.UpdateInputDto) (string, error) {
	e, err := s.client.EmailDuplicity()
	if err != nil {
		s.log.Errorf(input, err)
		return "", errors.New("internal server error |")
	}
	if e {
		return "email already registered", nil
	}
	return "", nil
}

// duplicityNick treats the nickname duplicity
func (s *Update) duplicityNick(input port.UpdateInputDto) (string, error) {
	n, err := s.client.NickDuplicity()
	if err != nil {
		s.log.Errorf(input, err)
		return "", errors.New("internal server error |")
	}
	if n {
		return "nickname already registered", nil
	}
	return "", nil
}

// update updates a client
func (s *Update) update(id string, input port.UpdateInputDto) error {
	_, uname, unick, udoc, uphone, uemail := s.client.Get()
	name, nick, doc, phone, email := input.Get()
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
	s.client.Load(id, uname, unick, udoc, uphone, uemail)
	if err := s.client.Update(); err != nil {
		s.log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Update service
func (s *Update) prepareOutput(output port.UpdateOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
