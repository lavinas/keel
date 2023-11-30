package domain

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/cpf_cnpj"
	"github.com/lavinas/keel/invoice/pkg/phone"
)

var (
	countries = []string{"BR"}
)

// Client represents a client that send or receive a invoice
type Client struct {
	Base
	Name        string `json:"name"     gorm:"type:varchar(100); not null"`
	Email       string `json:"email"    gorm:"type:varchar(100); not null"`
	DocumentStr string `json:"document" gorm:"-"`
	Document    uint64 `json:"-"        gorm:"type:numeric(20)"`
	PhoneStr    string `json:"phone"    gorm:"-"`
	Phone       uint64 `json:"-"        gorm:"type:numeric(20)"`
}

// Validate validates the client
func (c *Client) Validate(repo port.Repository) error {
	execOrder := []func(repo port.Repository) error{
		c.Base.Validate,
		c.ValidateName,
		c.ValidateEmail,
		c.ValidateDocument,
		c.ValidatePhone,
		c.ValidateDuplicity,
	}
	return ValidateLoop(execOrder, repo)
}

// Fit fits the client information received
func (c *Client) Fit() {
	c.Base.Fit()
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(c.Email)
	c.Document, _ = cpf_cnpj.ParseUint(c.DocumentStr)
	for _, cr := range countries {
		if c.PhoneStr = phone.Parse(c.PhoneStr, cr); c.PhoneStr != "" {
			break
		}
	}
}

// ValidateName validates the name of the client
func (c *Client) ValidateName(repo port.Repository) error {
	if c.Name == "" {
		return errors.New(ErrClientNameIsRequired)
	}
	if len(strings.Split(c.Name, " ")) < 2 {
		return errors.New(ErrClientNameLength)
	}
	return nil
}

// ValidateEmail validates the email of the client
func (c *Client) ValidateEmail(repo port.Repository) error {
	if c.Email == "" {
		return errors.New(ErrClientEmailIsRequired)
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New(ErrClientEmailIsInvalid)
	}
	return nil
}

// ValidateDocument validates the document of the client
func (c *Client) ValidateDocument(repo port.Repository) error {
	if c.DocumentStr == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.DocumentStr) {
		return errors.New(ErrClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *Client) ValidatePhone(repo port.Repository) error {
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

// ValidateDuplicity validates the duplicity of the model
func (b *Client) ValidateDuplicity(repo port.Repository) error {
	return b.Base.ValidateDuplicity(b, repo)
}
