package model

import (
	"errors"
)

const (
	ErrProductIDLength           = "product id must have only one word"
	ErrProductIDLower            = "product id must be lower case"
	ErrProductBusinessIsRequired = "product business is required"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Business    *Client `json:"business"`
	Description string  `json:"description"`
}

// Validate validates the product
func (p *Product) Validate() error {
	return p.ValidateLoop([]func() error{
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
