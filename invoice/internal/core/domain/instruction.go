package domain

import (
	"errors"
)

// Instruction represents a instruction for be showed in the invoice
type Instruction struct {
	Base
	Business    *Client `json:"business"`
	Description string  `json:"description"`
}

// Validate validates the instruction
func (i *Instruction) Validate() error {
	return ValidateLoop([]func() error{
		i.Base.Validate,
		i.ValidateBusiness,
	})
}

// ValidateBusiness validates the business of the product
func (p *Instruction) ValidateBusiness() error {
	if p.Business == nil {
		return errors.New(ErrInstructionBusinessIsRequired)
	}
	return p.Business.Validate()
}
