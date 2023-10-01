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
	if err := createDomain(log, domain, input); err != nil {
		return err
	}
	if err := duplicity(log, domain, input); err != nil {
		return err
	}
	if err := store(log, domain, input); err != nil {
		return err
	}
	prepareOutput(domain, output)
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
func createDomain(log port.Log, domain port.Domain, input port.CreateInputDto) error {
	input.Format()
	name, nick, doc, phone, email := input.GetName(), input.GetNickname(), input.GetDocument(), input.GetPhone(), input.GetEmail()
	_, err := domain.ClientInit(name, nick, doc, phone, email)
	if err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// duplicity checks if a document or email is already registered
func duplicity(log port.Log, domain port.Domain, input port.CreateInputDto) error {
	message := ""
	b, err := domain.ClientDocumentDuplicity()
	if err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	if b {
		message += "document already registered | "
	}
	e, err := domain.ClientEmailDuplicity()
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
func store(log port.Log, domain port.Domain, input port.CreateInputDto) error {
	// Store client
	if err := domain.ClientSave(); err != nil {
		log.Errorf(input, err)
		return errors.New("internal server error")
	}
	return nil
}

// prepareOutput prepares output data of Create service
func prepareOutput(domain port.Domain, output port.CreateOutputDto) {
	id, name, nick, doc, phone, email := domain.ClientGetFormatted()
	output.Fill(id, name, nick, doc, phone, email)

}
