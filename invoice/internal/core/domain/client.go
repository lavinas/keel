package domain

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

// Client represents a client that send or receive a invoice
type Client struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100)"`
	Document string `json:"document" gorm:"type:decimal(20)"`
	Phone    string `json:"phone" gorm:"type:varchar(20)"`
}

// Validate validates the client
func (c *Client) Validate() error {
	return ValidateLoop([]func() error{
		c.Base.Validate,
		c.ValidateName,
		c.ValidateEmail,
		c.ValidateDocument,
		c.ValidatePhone,
	})
}

// ValidateName validates the name of the client
func (c *Client) ValidateName() error {
	if c.Name == "" {
		return errors.New(ErrClientNameIsRequired)
	}
	if len(strings.Split(c.Name, " ")) < 2 {
		return errors.New(ErrClientNameLength)
	}
	return nil
}

// ValidateEmail validates the email of the client
func (c *Client) ValidateEmail() error {
	if c.Email == "" {
		return errors.New(ErrClientEmailIsRequired)
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New(ErrClientEmailIsInvalid)
	}
	return nil
}

// ValidateDocument validates the document of the client
func (c *Client) ValidateDocument() error {
	if c.Document == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.Document) {
		return errors.New(ErrClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *Client) ValidatePhone() error {
	if c.Phone == "" {
		return nil
	}
	for _, cr := range countries {
		r := phone.Parse(c.Phone, cr)
		if r != "" {
			return nil
		}
	}
	return errors.New(ErrClientPhoneIsInvalid)
}
