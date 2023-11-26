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
	Name        string `json:"name"     gorm:"type:varchar(100)"`
	Email       string `json:"email"    gorm:"type:varchar(100)"`
	DocumentStr string `json:"document" gorm:"-"`
	DocumentNum uint64 `json:"-"        gorm:"type:numeric(20)"`
	PhoneStr    string `json:"phone"    gorm:"-"`
	PhoneNum    uint64 `json:"-"        gorm:"type:numeric(20)"`
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

// Marshal marshals the client
func (c *Client) Marshal() {
	c.DocumentNum = cpf_cnpj.ParseUint(c.DocumentStr)
	for _, cr := range countries {
		if c.PhoneNum = phone.ParseUint(c.PhoneStr, cr); c.PhoneNum != 0 {
			break
		}
	}
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
	if c.DocumentStr == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.DocumentStr) {
		return errors.New(ErrClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *Client) ValidatePhone() error {
	if c.PhoneStr == "" {
		return nil
	}
	for _, cr := range countries {
		r := phone.Parse(c.PhoneStr, cr)
		if r != "" {
			return nil
		}
	}
	return errors.New(ErrClientPhoneIsInvalid)
}
