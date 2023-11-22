package dto

import (
	"errors"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/cpf_cnpj"
	"github.com/lavinas/keel/invoice/pkg/phone"
)

var (
	countries = []string{"BR"}
)

// RegisterClient is the dto for registering a new client
type RegisterClient struct {
	RegisterBase
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}

// Validate validates the client
func (c *RegisterClient) Validate() error {
	return ValidateLoop([]func() error{
		c.ValidateBase,
		c.ValidateName,
		c.ValidateEmail,
		c.ValidateDocument,
		c.ValidatePhone,
	})
}

// GetDomain returns the domain of the client
func (c *RegisterClient) GetDomain(businnes_id string) port.Domain {
	return domain.NewClient(businnes_id, c.ID, c.Name, c.Email, c.strToUint64(c.Document), c.strToUint64(c.Phone), time.Time{}, time.Time{})
}

// Get returns the client ID, Name, Email, Document and Phone
func (c *RegisterClient) Get() (string, string, string, string, string) {
	return c.ID, c.Name, c.Email, c.Document, c.Phone
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

// strToUint64 converts a string to uint64
func (s *RegisterClient) strToUint64(str string) uint64 {
	re := regexp.MustCompile(`[^0-9]`)
	str = re.ReplaceAllString(str, "")
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}
