package dto

import (
	"errors"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// RegisterInstruction is the DTO for registering an instruction
type RegisterInstruction struct {
	RegisterBase
	Description string `json:"description"`
}

// GetDomain returns the domain of the client
func (c *RegisterInstruction) GetDomain(businnes_id string) port.Domain {
	return domain.NewInstruction(businnes_id, c.ID, c.Description, time.Time{}, time.Time{})
}

// Get returns the ID and Description
func (r *RegisterInstruction) Get() (string, string) {
	return r.ID, r.Description
}

// Validate validates the instruction
func (r *RegisterInstruction) Validate() error {
	return ValidateLoop([]func() error{
		r.ValidateBase,
		r.ValidateDescription,
	})
}

// ValidateDescription validates the description of the instruction
func (r *RegisterInstruction) ValidateDescription() error {
	if r.Description == "" {
		return errors.New(ErrRegisterInstructionDescriptionIsRequired)
	}
	return nil
}
