package domain

import (
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/kerror"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Description string `json:"description"`
}

// Validate validates the product
func (i *Product) Validate(repo port.Repository) *kerror.KError {
	return ValidateLoop([]func(repo port.Repository) *kerror.KError{
		i.Base.Validate,
		i.ValidateDescription,
		i.ValidateDuplicity,
	}, repo)
}

// Fit fits the product information received
func (i *Product) Fit() {
	i.Base.Fit()
	i.Description = strings.TrimSpace(i.Description)
}

// Validate Description validates the description of the product
func (i *Product) ValidateDescription(repo port.Repository) *kerror.KError {
	if i.Description == "" {
		return kerror.NewKError(kerror.BadRequest, ErrProductDescriptionIsRequired)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (b *Product) ValidateDuplicity(repo port.Repository) *kerror.KError {
	return b.Base.ValidateDuplicity(b, repo)
}
