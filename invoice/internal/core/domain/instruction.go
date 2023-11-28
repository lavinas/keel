package domain

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Description string `json:"description"`
}

// Fit fits the instruction information received
func (i *Instruction) Fit() {
	i.Base.Fit()
	i.Description = strings.TrimSpace(i.Description)
}

// Validate validates the instruction
func (i *Instruction) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateDescription,
	}, repo)
}

// ValidateDescription validates the description of the instruction
func (i *Instruction) ValidateDescription(repo port.Repository) error {
	if i.Description == "" {
		return errors.New(ErrInstructionDescriptionIsRequired)
	}
	return nil
}
