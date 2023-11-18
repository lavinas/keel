package model

import (
	"errors"
)

const (
	ErrInvoiceItemIDLength = "invoice item id must have only one word"
	ErrInvoiceItemIDLower  = "invoice item id must be lower case"
	ErrInvoiceItemQuantity = "invoice item quantity must be greater than 0"
	ErrInvoiceItemPrice    = "invoice item price must be greater than 0"
)

// InvoiceItem represents a item in the invoice
type InvoiceItem struct {
	Base
	Product     *Product `json:"service"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	UnitPrice   float64  `json:"unit_price"`
}

// Validate validates the invoice item
func (i *InvoiceItem) Validate() error {
	return i.ValidateLoop([]func() error{
		i.Base.Validate,
		i.ValidateProduct,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	})
}

// ValidateProduct validates the product of the invoice item
func (c *InvoiceItem) ValidateProduct() error {
	if c.Product == nil {
		return errors.New(ErrProductBusinessIsRequired)
	}
	return c.Product.Validate()
}

// ValidateQuantity validates the quantity of the invoice item
func (c *InvoiceItem) ValidateQuantity() error {
	if c.Quantity <= 0 {
		return errors.New(ErrInvoiceItemQuantity)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item
func (c *InvoiceItem) ValidateUnitPrice() error {
	if c.UnitPrice <= 0 {
		return errors.New(ErrInvoiceItemPrice)
	}
	return nil
}

// GetAmount returns the amount of the invoice item
func (c *InvoiceItem) GetAmount() float64 {
	return float64(c.Quantity) * c.UnitPrice
}
