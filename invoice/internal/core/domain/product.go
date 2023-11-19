package domain

import (
	"errors"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Business    *Client `json:"business"`
	Description string  `json:"description"`
}

// Validate validates the product
func (p *Product) Validate() error {
	return ValidateLoop([]func() error{
		p.Base.Validate,
		p.ValidateBusiness,
	})
}

// ValidateBusiness validates the business of the product
func (p *Product) ValidateBusiness() error {
	if p.Business == nil {
		return errors.New(ErrProductBusinessIsRequired)
	}
	return p.Business.Validate()
}
