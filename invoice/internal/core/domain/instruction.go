package domain

import (
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/kerror"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Description string `json:"description"`
}

// SetCreate set information for create a new instruction
func (i *Instruction) SetCreate(business_id string) {
	i.Base.SetCreate(business_id)
	i.Fit()
}

// Validate validates the instruction
func (i *Instruction) Validate(repo port.Repository) *kerror.KError {
	return ValidateLoop([]func(repo port.Repository) *kerror.KError{
		i.Base.Validate,
		i.ValidateDescription,
		i.ValidateDuplicity,
	}, repo)
}

// Fit fits the instruction information received
func (i *Instruction) Fit() {
	i.Base.Fit()
	i.Description = strings.TrimSpace(i.Description)
}

// ValidateDescription validates the description of the instruction
func (i *Instruction) ValidateDescription(repo port.Repository) *kerror.KError {
	if i.Description == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInstructionDescriptionIsRequired)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (b *Instruction) ValidateDuplicity(repo port.Repository) *kerror.KError {
	return b.Base.ValidateDuplicity(b, repo)
}

// TableName returns the table name for gorm
func (b *Instruction) TableName() string {
	return "instruction"
}
