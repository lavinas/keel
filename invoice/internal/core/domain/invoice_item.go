package domain

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// InvoiceItem represents a item in the invoice
type InvoiceItem struct {
	ID           string  `json:"id"          gorm:"type:varchar(50);primaryKey;not null"`
	InvoiceID    string  `json:"invoice_id"  gorm:"type:varchar(50); not null"`
	Description  string  `json:"description" gorm:"type:varchar(255)"`
	QuantityStr  string  `json:"quantity"    gorm:"-"`
	Quantity     int     `json:"-"           gorm:"type:int; not null"`
	UnitPriceStr string  `json:"unit_price"  gorm:"-"`
	UnitPrice    float64 `json:"-"           gorm:"type:decimal(20, 2); not null"`
}

// Validate validates the invoice item
func (i *InvoiceItem) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	}, repo)
}

// Marshal marshals the invoice item
func (i *InvoiceItem) Marshal() error {
	i.ID = uuid.New().String()
	var err error
	i.Quantity, err = strconv.Atoi(i.QuantityStr)
	if err != nil {
		return err
	}
	i.UnitPrice, err = strconv.ParseFloat(i.UnitPriceStr, 64)
	if err != nil {
		return err
	}
	return nil
}

// ValidateQuantity validates the quantity of the invoice item
func (c *InvoiceItem) ValidateQuantity(repo port.Repository) error {
	if c.QuantityStr == "" {
		return errors.New(ErrInvoiceItemQuantityRequired)
	} else if q, err := strconv.Atoi(c.QuantityStr); err != nil {
		return errors.New(ErrInvoiceItemQuantityInvalid)
	} else if q <= 0 {
		return errors.New(ErrInvoiceItemQuantityLessOrEqualZero)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item
func (c *InvoiceItem) ValidateUnitPrice(repo port.Repository) error {
	if c.UnitPriceStr == "" {
		return errors.New(ErrInvoiceItemPriceRequired)
	} else if p, err := strconv.ParseFloat(c.UnitPriceStr, 64); err != nil {
		return errors.New(ErrInvoiceItemPriceInvalid)
	} else if p <= 0 {
		return errors.New(ErrInvoiceItemPriceLessOrEqualZero)
	}
	return nil
}

// GetAmount returns the amount of the invoice item
func (c *InvoiceItem) GetAmount() float64 {
	return float64(c.Quantity) * c.UnitPrice
}
