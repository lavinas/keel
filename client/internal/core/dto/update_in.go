// Package dto is the package that defines the DTOs for the client adapter
package dto

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/client/pkg/ktools"
)

// UpdateInputDto is the input DTO used to create a client
type UpdateInputDto struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// IsBlank checks if the input DTO is blank
func (c *UpdateInputDto) IsBlank() bool {
	return strings.Trim(c.Name, " ") == "" && strings.Trim(c.Nickname, " ") == "" &&
		strings.Trim(c.Document, " ") == "" && strings.Trim(c.Phone, " ") == "" &&
		strings.Trim(c.Email, " ") == ""
}

// Validate validates the input DTO fields (name, nickname, document, phone, email)
func (c *UpdateInputDto) Validate() error {
	if c.IsBlank() {
		return errors.New("at least one field must be filled")
	}
	msg := ""
	if strings.Trim(c.Name, " ") != "" {
		if _, err := ktools.FormatName(c.Name); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Document, " ") != "" {
		if _, err := ktools.FormatDocument(c.Document); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Phone, " ") != "" {
		if _, err := ktools.FormatPhone(c.Phone); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Email, " ") != "" {
		if _, err := ktools.FormatEmail(c.Email); err != nil {
			msg += err.Error() + " | "
		}
	}
	if msg == "" {
		return nil
	}
	msg = strings.Trim(msg, " |")
	return errors.New(msg)
}

// FormatUpdate formats all fields (name, nickname, document, phone, email) for update values
func (c *UpdateInputDto) Format() error {
	if c.IsBlank() {
		return errors.New("at least one field must be filled")
	}
	var err error
	var name, doc, phone, email string
	if strings.Trim(c.Name, " ") != "" {
		if name, err = ktools.FormatName(c.Name); err != nil {
			return err
		}
		c.Name = name
	}
	if strings.Trim(c.Nickname, " ") != "" {
		c.Nickname, _ = ktools.FormatNickname(c.Nickname)
	}
	if strings.Trim(c.Document, " ") != "" {
		if doc, err = ktools.FormatDocument(c.Document); err != nil {
			return err
		}
		c.Document = doc
	}
	if strings.Trim(c.Phone, " ") != "" {
		if phone, err = ktools.FormatPhone(c.Phone); err != nil {
			return err
		}
		c.Phone = phone
	}
	if strings.Trim(c.Email, " ") != "" {
		if email, err = ktools.FormatEmail(c.Email); err != nil {
			return err
		}
		c.Email = email
	}
	return nil
}

// Get returns all fields (name, nickname, document, phone, email)
func (c *UpdateInputDto) Get() (string, string, string, string, string) {
	return c.Name, c.Nickname, c.Document, c.Phone, c.Email
}
