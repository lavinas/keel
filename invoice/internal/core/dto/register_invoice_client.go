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

type RegisterInvoiceClient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}

// Validate validates the client
func (c *RegisterInvoiceClient) Validate() error {
	return ValidateLoop([]func() error{
		c.ValidateName,
		c.ValidateEmail,
		c.ValidateDocument,
		c.ValidatePhone,
	})
}

// ValidateName validates the name of the client
func (c *RegisterInvoiceClient) ValidateName() error {
	if c.Name == "" {
		return errors.New(ErrRegisterInvoiceClientNameIsRequired)
	}
	if len(strings.Split(c.Name, " ")) < 2 {
		return errors.New(ErrRegisterInvoiceClientNameLength)
	}
	return nil
}

// ValidateEmail validates the email of the client
func (c *RegisterInvoiceClient) ValidateEmail() error {
	if c.Email == "" {
		return errors.New(ErrRegisterInvoiceClientEmailIsRequired)
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New(ErrRegisterInvoiceClientEmailIsInvalid)
	}
	return nil
}

// ValidateDocument validates the document of the client
func (c *RegisterInvoiceClient) ValidateDocument() error {
	if c.Document == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.Document) {
		return errors.New(ErrRegisterInvoiceClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *RegisterInvoiceClient) ValidatePhone() error {
	if c.Phone == "" {
		return nil
	}
	for _, cr := range countries {
		r := phone.Parse(c.Phone, cr)
		if r != "" {
			return nil
		}
	}
	return errors.New(ErrRegisterInvoiceClientPhoneIsInvalid)
}
