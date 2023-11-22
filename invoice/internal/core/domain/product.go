package domain

import (
	"errors"
	"time"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Description string `json:"description"`
}

// NewProduct creates a new product
func NewProduct(businness_id, id, description string, created_at time.Time, updated_at time.Time) *Product {
	return &Product{
		Base:        NewBase(businness_id, id, created_at, updated_at),
		Description: description,
	}
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
