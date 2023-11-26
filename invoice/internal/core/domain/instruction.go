package domain

import (
	"errors"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Description string `json:"description"`
}

// Validate validates the instruction
func (i *Instruction) Validate() error {
	return ValidateLoop([]func() error{
		i.Base.Validate,
		i.ValidateDescription,
	})
}

func (i *Instruction) ValidateDescription() error {
	if i.Description == "" {
		return errors.New(ErrInstructionDescriptionIsRequired)
	}
	return nil
}
