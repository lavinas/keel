package dto

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/pkg/formatter"
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
	msg := ""
	if strings.Trim(c.Name, " ") != "" {
		if _, err := formatter.FormatName(c.Name); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Nickname, " ") != "" {
		if _, err := formatter.FormatNickname(c.Nickname); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Document, " ") != "" {
		if _, err := formatter.FormatDocument(c.Document); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Phone, " ") != "" {
		if _, err := formatter.FormatPhone(c.Phone); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Email, " ") != "" {
		if _, err := formatter.FormatEmail(c.Email); err != nil {
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
	var err error
	var name, nick, doc, phone, email string
	if strings.Trim(c.Name, " ") != "" {
		if name, err = formatter.FormatName(c.Name); err != nil {
			return err
		}
	}
	if strings.Trim(c.Nickname, " ") != "" {
		if nick, err = formatter.FormatNickname(c.Nickname); err != nil {
			return err
		}
	}
	if strings.Trim(c.Document, " ") != "" {
		if doc, err = formatter.FormatDocument(c.Document); err != nil {
			return err
		}
	}
	if strings.Trim(c.Phone, " ") != "" {
		if phone, err = formatter.FormatPhone(c.Phone); err != nil {
			return err
		}
	}
	if strings.Trim(c.Email, " ") != "" {
		if email, err = formatter.FormatEmail(c.Email); err != nil {
			return err
		}
	}
	c.Name = name
	c.Nickname = nick
	c.Document = doc
	c.Phone = phone
	c.Email = email
	return nil
}

// Get returns all fields (name, nickname, document, phone, email)
func (c *UpdateInputDto) Get() (string, string, string, string, string) {
	return c.Name, c.Nickname, c.Document, c.Phone, c.Email
}
