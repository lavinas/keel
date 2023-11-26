package domain

import (
	"errors"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Description string `json:"description"`
}

// Validate validates the product
func (p *Product) Validate() error {
	return ValidateLoop([]func() error{
		p.Base.Validate,
		p.ValidateDescription,
	})
}

// Validate Description validates the description of the product
func (p *Product) ValidateDescription() error {
	if p.Description == "" {
		return errors.New(ErrProductDescriptionIsRequired)
	}
	return nil
}
