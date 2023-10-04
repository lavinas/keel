package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientCreate is the service for creating a new client
type ClientCreate struct {
	log    port.Log
	client port.Client
	input  port.ClientCreateInputDto
	output port.ClientCreateOutputDto
}

// NewClientCreate creates a new client create service
func NewClientCreate(log port.Log, client port.Client, input port.ClientCreateInputDto, output port.ClientCreateOutputDto) *ClientCreate {
	return &ClientCreate{
		log:    log,
		client: client,
		input:  input,
		output: output,
	}
}

// Execute executes the service
func (s *ClientCreate) Execute() error {
	if err := validateInput(s.log, s.input); err != nil {
		return err
	}
	if err := loadClient(s.input, s.client); err != nil {
		return err
	}
	if err := duplicity(s.log, s.client, s.input); err != nil {
		return err
	}
	if err := store(s.log, s.client, s.input); err != nil {
		return err
	}
	prepareOutput(s.client, s.output)
	s.log.Infof(s.input, "created")
	return nil
}

// loadClient loads a client from the input dto
func loadClient(input port.ClientCreateInputDto, client port.Client) error {
	input.Format()
	name, nick, doc, phone, email := input.Get()
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	iphone, _ := strconv.ParseUint(phone, 10, 64)
	if err := client.Create(name, nick, idoc, iphone, email); err != nil {
		return err
	}
	return nil
}

// validateInput validates input data of Create service
func validateInput(log port.Log, input port.ClientCreateInputDto) error {
	if err := input.Validate(); err != nil {
		log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// duplicity checks if a document or email is already registered
func duplicity(log port.Log, client port.Client, input port.ClientCreateInputDto) error {
	message := ""
	m, err := duplicityDocument(log, client, input)
	if err != nil {
		return err
	}
	message += m
	m, err = duplicityEmail(log, client, input)
	if err != nil {
		return err
	}
	message += m
	if message != "" {
		message = strings.Trim(message, " |")
		log.Infof(input, "conflict: "+message)
		return errors.New("conflict: " + message)
	}
	return nil
}

// duplicityDocument treats the document duplicity
func duplicityDocument(log port.Log, client port.Client, input port.ClientCreateInputDto) (string, error) {
	b, err := client.DocumentDuplicity()
	if err != nil {
		log.Errorf(input, err)
		return "", errors.New("internal server error")
	}
	if b {
		return "document already registered | ", nil
	}
	return "", nil
}

// duplicityEmail treats the email duplicity
func duplicityEmail(log port.Log, client port.Client, input port.ClientCreateInputDto) (string, error) {
	e, err := client.EmailDuplicity()
	if err != nil {
		log.Errorf(input, err)
		return "", errors.New("internal server error |")
	}
	if e {
		return "email already registered", nil
	}
	return "", nil
}

// store stores a new client
func store(log port.Log, client port.Client, input port.ClientCreateInputDto) error {
	// Store client
	if err := client.Save(); err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Create service
func prepareOutput(client port.Client, output port.ClientCreateOutputDto) {
	id, name, nick, doc, phone, email := client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
