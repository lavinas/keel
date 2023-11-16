package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/client/internal/core/port"
)

// Insert is the service for creating a new client
type Insert struct {
	log    port.Log
	client port.Client
}

// NewInsert creates a new client create service
func NewInsert(log port.Log, client port.Client) *Insert {
	return &Insert{
		log:    log,
		client: client,
	}
}

// Execute executes the service
func (s *Insert) Execute(input port.InsertInputDto, output port.InsertOutputDto) error {
	execOrder := []func(port.InsertInputDto) error{
		s.validateInput,
		s.loadClient,
		s.duplicity,
		s.store,
	}
	for _, f := range execOrder {
		if err := f(input); err != nil {
			return err
		}
	}
	s.prepareOutput(output)
	s.log.Infof(input, "created")
	return nil
}

// loadClient loads a client from the input dto
func (s *Insert) loadClient(input port.InsertInputDto) error {
	input.Format()
	name, nick, doc, phone, email := input.Get()
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	iphone, _ := strconv.ParseUint(phone, 10, 64)
	if err := s.client.Insert(name, nick, idoc, iphone, email); err != nil {
		s.log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// validateInput validates input data of Insert service
func (s *Insert) validateInput(input port.InsertInputDto) error {
	if err := input.Validate(); err != nil {
		s.log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// duplicity checks if a document or email is already registered
func (s *Insert) duplicity(input port.InsertInputDto) error {
	execOrder := []func(port.InsertInputDto) (string, error){
		s.duplicityDocument,
		s.duplicityEmail,
		s.duplicityNick,
	}
	message := ""
	for _, f := range execOrder {
		m, err := f(input)
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
func (s *Insert) duplicityDocument(input port.InsertInputDto) (string, error) {
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
func (s *Insert) duplicityEmail(input port.InsertInputDto) (string, error) {
	e, err := s.client.EmailDuplicity()
	if err != nil {
		s.log.Errorf(input, err)
		return "", errors.New("internal server error |")
	}
	if e {
		return "email already registered | ", nil
	}
	return "", nil
}

func (s *Insert) duplicityNick(input port.InsertInputDto) (string, error) {
	n, err := s.client.NickDuplicity()
	if err != nil {
		s.log.Errorf(input, err)
		return "", errors.New("internal server error |")
	}
	if n {
		return "nickname already registered | ", nil
	}
	return "", nil
}

// store stores a new client
func (s *Insert) store(input port.InsertInputDto) error {
	// Store client
	if err := s.client.Save(); err != nil {
		s.log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Insert service
func (s *Insert) prepareOutput(output port.InsertOutputDto) {
	id, name, nick, doc, phone, email := s.client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
