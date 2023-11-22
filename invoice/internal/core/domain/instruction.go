package domain

import (
	"errors"
	"time"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Description string `json:"description"`
}

// NewInstruction creates a new instruction
func NewInstruction(businnes_id, id, description string, created_at, updated_at time.Time) *Instruction {
	return &Instruction{
		Base:        NewBase(businnes_id, id, created_at, updated_at),
		Description: description,
	}
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
