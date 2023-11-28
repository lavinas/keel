package domain

import (
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Item represents a item in the invoice
type Item struct {
	Base
	InvoiceID    string  `json:"invoice_id"  gorm:"type:varchar(50); not null"`
	Description  string  `json:"description" gorm:"type:varchar(255)"`
	QuantityStr  string  `json:"quantity"    gorm:"-"`
	Quantity     int     `json:"-"           gorm:"type:int; not null"`
	UnitPriceStr string  `json:"unit_price"  gorm:"-"`
	UnitPrice    float64 `json:"-"           gorm:"type:decimal(20, 2); not null"`
}

// Validate validates the invoice item
func (i *Item) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	}, repo)
}

// Fit fits the invoice item information received
func (i *Item) Fit() {
	i.Base.Fit()
	if i.Base.ID == "" {
		i.Base.ID = uuid.New().String()
	}
	i.Description = strings.TrimSpace(i.Description)
	i.Quantity, _ = strconv.Atoi(i.QuantityStr)
	i.UnitPrice, _ = strconv.ParseFloat(i.UnitPriceStr, 64)
}

// ValidateQuantity validates the quantity of the invoice item
func (c *Item) ValidateQuantity(repo port.Repository) error {
	if c.QuantityStr == "" {
		return errors.New(ErrItemQuantityRequired)
	} else if q, err := strconv.Atoi(c.QuantityStr); err != nil {
		return errors.New(ErrItemQuantityInvalid)
	} else if q <= 0 {
		return errors.New(ErrItemQuantityLessOrEqualZero)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item
func (c *Item) ValidateUnitPrice(repo port.Repository) error {
	if c.UnitPriceStr == "" {
		return errors.New(ErrItemPriceRequired)
	} else if p, err := strconv.ParseFloat(c.UnitPriceStr, 64); err != nil {
		return errors.New(ErrItemPriceInvalid)
	} else if p <= 0 {
		return errors.New(ErrItemPriceLessOrEqualZero)
	}
	return nil
}

// GetAmount returns the amount of the invoice item
func (c *Item) GetAmount() float64 {
	return float64(c.Quantity) * c.UnitPrice
}
