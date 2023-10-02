package service

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/internal/client/core/port"
)

// Orquestration of Creating a new client
func ServiceClientCreate(log port.Log, domain port.Domain, input port.CreateInputDto, output port.CreateOutputDto) error {
	if err := validateInput(log, input); err != nil {
		return err
	}
	client, err := createDomain(log, domain, input)
	if err != nil {
		return err
	}
	if err := duplicity(log, client, input); err != nil {
		return err
	}
	if err := store(log, client, input); err != nil {
		return err
	}
	prepareOutput(client, output)
	log.Infof(input, "created")
	return nil
}

// validateInput validates input data of Create service
func validateInput(log port.Log, input port.CreateInputDto) error {
	if err := input.Validate(); err != nil {
		log.Infof(input, "bad request: "+err.Error())
		return errors.New("bad request: " + err.Error())
	}
	return nil
}

// createDomain creates a new client domain
func createDomain(log port.Log, domain port.Domain, input port.CreateInputDto) (port.Client, error) {
	input.Format()
	return domain.GetClient(input)
}

// duplicity checks if a document or email is already registered
func duplicity(log port.Log, client port.Client, input port.CreateInputDto) error {
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
func store(log port.Log, client port.Client, input port.CreateInputDto) error {
	// Store client
	if err := client.Save(); err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Create service
func prepareOutput(client port.Client, output port.CreateOutputDto) {
	id, name, nick, doc, phone, email := client.GetFormatted()
	output.Fill(id, name, nick, doc, phone, email)
}
