package domain

import (
	"net/mail"
	"strings"

	"github.com/lavinas/keel/internal/invoice/core/port"
	"github.com/lavinas/keel/pkg/cpf_cnpj"
	"github.com/lavinas/keel/pkg/kerror"
	"github.com/lavinas/keel/pkg/phone"
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

// SetCreate set information for create a new client
func (c *Client) SetCreate(business_id string) {
	c.Base.SetCreate(business_id)
	c.Fit()
}

// Validate validates the client
func (c *Client) Validate(repo port.Repository) *kerror.KError {
	execOrder := []func(repo port.Repository) *kerror.KError{
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
		if c.Phone, _ = phone.ParseUint(c.PhoneStr, cr); c.Phone != 0 {
			break
		}
	}
}

// ValidateName validates the name of the client
func (c *Client) ValidateName(repo port.Repository) *kerror.KError {
	if c.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrClientNameIsRequired)
	}
	if len(strings.Split(c.Name, " ")) < 2 {
		return kerror.NewKError(kerror.BadRequest, ErrClientNameLength)
	}
	return nil
}

// ValidateEmail validates the email of the client
func (c *Client) ValidateEmail(repo port.Repository) *kerror.KError {
	if c.Email == "" {
		return kerror.NewKError(kerror.BadRequest, ErrClientEmailIsRequired)
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrClientEmailIsInvalid)
	}
	return nil
}

// ValidateDocument validates the document of the client
func (c *Client) ValidateDocument(repo port.Repository) *kerror.KError {
	if c.DocumentStr == "" {
		return nil
	}
	if !cpf_cnpj.ValidateCPFOrCNPJ(c.DocumentStr) {
		return kerror.NewKError(kerror.BadRequest, ErrClientDocumentIsInvalid)
	}
	return nil
}

// ValidatePhone validates the phone of the client
func (c *Client) ValidatePhone(repo port.Repository) *kerror.KError {
	if c.PhoneStr == "" {
		return nil
	}
	for _, cr := range countries {
		r := phone.Parse(c.PhoneStr, cr)
		if r != "" {
			return nil
		}
	}
	return kerror.NewKError(kerror.BadRequest, ErrClientPhoneIsInvalid)
}

// ValidateDuplicity validates the duplicity of the model
func (b *Client) ValidateDuplicity(repo port.Repository) *kerror.KError {
	return b.Base.ValidateDuplicity(b, repo)
}

// TableName returns the table name for gorm
func (b *Client) TableName() string {
	return "client"
}
