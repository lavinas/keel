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
func (i *Product) Validate(p interface{}) error {
	return ValidateLoop([]func(p interface{}) error{
		i.Base.Validate,
		i.ValidateDescription,
	}, p)
}

// Validate Description validates the description of the product
func (i *Product) ValidateDescription(p interface{}) error {
	if i.Description == "" {
		return errors.New(ErrProductDescriptionIsRequired)
	}
	return nil
}
