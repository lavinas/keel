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
func (i *Instruction) Validate(p interface{}) error {
	return ValidateLoop([]func(interface{}) error{
		i.Base.Validate,
		i.ValidateDescription,
	}, p)
}

func (i *Instruction) ValidateDescription(p interface{}) error {
	if i.Description == "" {
		return errors.New(ErrInstructionDescriptionIsRequired)
	}
	return nil
}
