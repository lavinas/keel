package domain

import (
	"errors"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// InvoiceItem represents a item in the invoice
type InvoiceItem struct {
	ID          string   `json:"id"         gorm:"type:varchar(50)"`
	InvoiceID   string   `json:"invoice_id" gorm:"type:varchar(50)"`
	Product     *Product `json:"service"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	UnitPrice   float64  `json:"unit_price"`
}

// Validate validates the invoice item
func (i *InvoiceItem) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.ValidateProduct,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	}, repo)
}

// ValidateProduct validates the product of the invoice item
func (c *InvoiceItem) ValidateProduct(repo port.Repository) error {
	if c.Product == nil {
		return errors.New("err")
	}
	return c.Product.Validate(repo)
}

// ValidateQuantity validates the quantity of the invoice item
func (c *InvoiceItem) ValidateQuantity(repo port.Repository) error {
	if c.Quantity <= 0 {
		return errors.New(ErrInvoiceItemQuantity)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item
func (c *InvoiceItem) ValidateUnitPrice(repo port.Repository) error {
	if c.UnitPrice == 0 {
		return errors.New(ErrInvoiceItemPrice)
	}
	return nil
}

// GetAmount returns the amount of the invoice item
func (c *InvoiceItem) GetAmount() float64 {
	return float64(c.Quantity) * c.UnitPrice
}
