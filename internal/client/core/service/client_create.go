package service

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

type ClientCreate struct {
	log    port.Log
	client port.Client
	input  port.ClientCreateInputDto
	output port.ClientCreateOutputDto
}

func NewClientCreate(log port.Log, client port.Client, input port.ClientCreateInputDto, output port.ClientCreateOutputDto) *ClientCreate {
	return &ClientCreate{
		log:    log,
		client: client,
		input:  input,
		output: output,
	}
}

func (s *ClientCreate) Execute() error {
	if err := validateInput(s.log, s.input); err != nil {
		return err
	}
	s.input.Format()

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
	b, err := client.DocumentDuplicity()
	if err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	if b {
		message += "document already registered | "
	}
	e, err := client.EmailDuplicity()
	if err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error |")
	}
	if e {
		message += "email already registered"
	}
	if message != "" {
		message = strings.Trim(message, " |")
		log.Infof(input, "conflict: "+message)
		return errors.New("conflict: " + message)
	}
	return nil
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
