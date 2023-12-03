package domain

import (
	"strings"

	"github.com/lavinas/keel/internal/invoice/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

// Product represents a product or service that can be invoiced
type Product struct {
	Base
	Description string `json:"description"`
}

// SetCreate set information for create a new product
func (i *Product) SetCreate(business_id string) {
	i.Base.SetCreate(business_id)
	i.Fit()
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

// TableName returns the table name for gorm
func (b *Product) TableName() string {
	return "product"
}
