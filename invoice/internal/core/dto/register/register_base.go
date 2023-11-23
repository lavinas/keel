package register

import (
	"errors"
	"strings"
)

// RegisterBase is the base dto for register
type RegisterBase struct {
	ID string `json:"id"`
}

// NewRegisterBase creates a new register base
func NewRegisterBase(id string) *RegisterBase {
	return &RegisterBase{
		ID: id,
	}
}

// Validate validates the register base
func (r *RegisterBase) ValidateBase() error {
	if r.ID == "" {
		return errors.New(ErrRegisterClientIDIsRequired)
	}
	if len(strings.Split(r.ID, " ")) > 1 {
		return errors.New(ErrRegisterClientIDLength)
	}
	if strings.ToLower(r.ID) != r.ID {
		return errors.New(ErrRegisterClientIDLower)
	}
	return nil
}

// Get returns the ID
func (r *RegisterBase) Get() string {
	return r.ID
}
