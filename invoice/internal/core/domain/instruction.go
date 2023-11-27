package domain

import (
	"errors"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Description string `json:"description"`
}

// Validate validates the instruction
func (i *Instruction) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateDescription,
	}, repo)
}

func (i *Instruction) ValidateDescription(repo port.Repository) error {
	if i.Description == "" {
		return errors.New(ErrInstructionDescriptionIsRequired)
	}
	return nil
}
