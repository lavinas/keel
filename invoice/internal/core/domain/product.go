package domain

import (
	"errors"
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Description string `json:"description"`
}

// Validate validates the product
func (i *Product) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateDescription,
	}, repo)
}

// Fit fits the product information received
func (i *Product) Fit() {
	i.Base.Fit()
	i.Description = strings.TrimSpace(i.Description)
}

// Validate Description validates the description of the product
func (i *Product) ValidateDescription(repo port.Repository) error {
	if i.Description == "" {
		return errors.New(ErrProductDescriptionIsRequired)
	}
	return nil
}
