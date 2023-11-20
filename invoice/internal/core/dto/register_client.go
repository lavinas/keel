package dto

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/lavinas/keel/invoice/pkg/cpf_cnpj"
	"github.com/lavinas/keel/invoice/pkg/phone"
)

var (
	countries = []string{"BR"}
)

// RegisterClient is the dto for registering a new client
type RegisterClient struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}

// Validate validates the client
func (c *RegisterClient) Validate() error {
	return ValidateLoop([]func() error{
		c.ValidateID,
		c.ValidateName,
		c.ValidateEmail,
		c.ValidateDocument,
		c.ValidatePhone,
	})
}

// Get returns the client ID, Name, Email, Document and Phone
func (c *RegisterClient) Get() (string, string, string, string, string) {
	return c.ID, c.Name, c.Email, c.Document, c.Phone
}

// ValidateID validates the id of the client
func (c *RegisterClient) ValidateID() error {
	if c.ID == "" {
		return nil
	}
	if len(strings.Split(c.ID, " ")) > 1 {
		return errors.New(ErrRegisterClientIDLength)
	}
	if strings.ToLower(c.ID) != c.ID {
		return errors.New(ErrRegisterClientIDLower)
	}
	return nil
}

// ValidateName validates the name of the client
func (c *RegisterClient) ValidateName() error {
	if c.Name == "" {
		return errors.New(ErrRegisterClientNameIsRequired)
	}
	if len(strings.Split(c.Name, " ")) < 2 {
		return errors.New(ErrRegisterClientNameLength)
	}
	return nil
}

// ValidateEmail validates the email of the client
func (c *RegisterClient) ValidateEmail() error {
	if c.Email == "" {
		return errors.New(ErrRegisterClientEmailIsRequired)
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New(ErrRegisterClientEmailIsInvalid)
	}
	return nil
}

// ValidateDocument validates the document of the client
func (c *RegisterClient) ValidateDocument() error {
	if c.Document == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.Document) {
		return errors.New(ErrRegisterClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *RegisterClient) ValidatePhone() error {
	if c.Phone == "" {
		return nil
	}
	for _, cr := range countries {
		r := phone.Parse(c.Phone, cr)
		if r != "" {
			return nil
		}
	}
	return errors.New(ErrRegisterClientPhoneIsInvalid)
}
