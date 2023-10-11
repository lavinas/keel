// Package dto is the package that defines the DTOs for the client adapter
package dto

import (
	"errors"
	"strconv"
)

// FindInputDto is the input DTO used to list all clients
type FindInputDto struct {
	Page     string `form:"page"`
	PerPage  string `form:"per_page"`
	Name     string `form:"name"`
	Nickname string `form:"nickname"`
	Document string `form:"document"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
}

// Validate validates the input DTO for client list
func (c *FindInputDto) Validate() error {
	msg := ""
	if err := validatePage(c.Page); err != nil {
		msg += err.Error() + " | "
	}
	if err := validatePerPage(c.PerPage); err != nil {
		msg += err.Error() + " | "
	}
	if err := validateDocument(c.Document); err != nil {
		msg += err.Error() + " | "
	}
	if err := validatePhone(c.Phone); err != nil {
		msg += err.Error() + " | "
	}
	if msg == "" {
		return nil
	}
	msg = msg[:len(msg)-3]
	return errors.New(msg)
}

// Get returns the input DTO values
func (c *FindInputDto) Get() (string, string, string, string, string, string, string) {
	return c.Page, c.PerPage, c.Name, c.Nickname, c.Document, c.Phone, c.Email
}

// validatePage validates the page field for client list
func validatePage(page string) error {
	if page != "" {
		if _, err := strconv.ParseUint(page, 10, 64); err != nil {
			return errors.New("page must be a number")
		}
	}
	return nil
}

// validatePerPage validates the perPage field for client list
func validatePerPage(perPage string) error {
	if perPage != "" {
		if _, err := strconv.ParseUint(perPage, 10, 64); err != nil {
			return errors.New("perPage must be a number")
		}
	}
	return nil
}

// validateDocument validates the document field for client list
func validateDocument(document string) error {
	if document != "" {
		if _, err := strconv.ParseUint(document, 10, 64); err != nil {
			return errors.New("document must be a number")
		}
	}
	return nil
}

// validateDocument validates the document field for client list
func validatePhone(phone string) error {
	if phone != "" {
		if _, err := strconv.ParseUint(phone, 10, 64); err != nil {
			return errors.New("phone must be a number")
		}
	}
	return nil
}
