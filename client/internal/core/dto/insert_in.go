// Package dto is the package that defines the DTOs for the client adapter
package dto

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/client/pkg/ktools"
)

// InsertInputDto is the input DTO used to create a client
type InsertInputDto struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// IsBlank checks if the input DTO is blank
func (c *InsertInputDto) IsBlank() bool {
	return strings.Trim(c.Name, " ") == "" && strings.Trim(c.Nickname, " ") == "" &&
		strings.Trim(c.Document, " ") == "" && strings.Trim(c.Phone, " ") == "" &&
		strings.Trim(c.Email, " ") == ""
}

// Validate validates the input DTO
func (c *InsertInputDto) Validate() error {
	msg := ""
	if _, err := ktools.FormatName(c.Name); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := ktools.FormatNickname(c.Nickname); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := ktools.FormatDocument(c.Document); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := ktools.FormatPhone(c.Phone); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := ktools.FormatEmail(c.Email); err != nil {
		msg += err.Error() + " | "
	}
	if msg == "" {
		return nil
	}
	msg = strings.Trim(msg, " |")
	return errors.New(msg)
}

// Format formats all fields (name, nickname, document, phone, email)
func (c *InsertInputDto) Format() error {
	var err error
	var name, nick, doc, phone, email string
	if name, err = ktools.FormatName(c.Name); err != nil {
		return err
	}
	if nick, err = ktools.FormatNickname(c.Nickname); err != nil {
		return err
	}
	if doc, err = ktools.FormatDocument(c.Document); err != nil {
		return err
	}
	if phone, err = ktools.FormatPhone(c.Phone); err != nil {
		return err
	}
	if email, err = ktools.FormatEmail(c.Email); err != nil {
		return err
	}
	c.Name = name
	c.Nickname = nick
	c.Document = doc
	c.Phone = phone
	c.Email = email
	return nil
}

// Get returns all fields (name, nickname, document, phone, email)
func (c *InsertInputDto) Get() (string, string, string, string, string) {
	return c.Name, c.Nickname, c.Document, c.Phone, c.Email
}
