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
	InvoiceID    string   `json:"invoice_id"  gorm:"type:varchar(50); not null"`
	ProductID    string   `json:"product_id"  gorm:"type:varchar(50); not null"`
	Product      *Product `json:"-"           gorm:"foreignKey:BusinessID,ProductID;associationForeignKey:BusinessID,ID"`
	Description  string   `json:"description" gorm:"type:varchar(255)"`
	QuantityStr  string   `json:"quantity"    gorm:"-"`
	Quantity     int      `json:"-"           gorm:"type:int; not null"`
	UnitPriceStr string   `json:"unit_price"  gorm:"-"`
	UnitPrice    float64  `json:"-"           gorm:"type:decimal(20, 2); not null"`
}

// Validate validates the invoice item
func (i *Item) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
		i.ValidateProduct,
		i.ValidateDuplicity,
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

// ValidateProduct validates the product of the invoice item
func (c *Item) ValidateProduct(repo port.Repository) error {
	if c.ProductID == "" && c.Product == nil {
		return errors.New(ErrItemProductRequired)
	} else if c.ProductID != "" && c.Product != nil {
		return errors.New(ErrItemProductConflict)
	} else if c.ProductID != "" {
		return c.ValidateProductID(repo)
	} else if c.Product.Validate(repo) != nil {
		return errors.New(ErrItemProductInvalid)
	}
	return nil
}

// ValidateProductID validates the product id of the invoice item
func (c *Item) ValidateProductID(repo port.Repository) error {
	if c.ProductID == "" {
		return errors.New(ErrItemProductRequired)
	}
	var product Product
	product.ID = c.ProductID
	if err := product.ValidateID(repo); err != nil {
		return errors.New(ErrItemProductInvalid)
	}
	if exists, err := repo.Exists(&product, c.BusinessID, c.ProductID); err != nil {
		return errors.New(ErrItemProductInvalid)
	} else if !exists {
		return errors.New(ErrItemProductNotFound)
	}
	return nil
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

// ValidateDuplicity validates the duplicity of the invoice item
func (c *Item) ValidateDuplicity(repo port.Repository) error {
	return c.Base.ValidateDuplicity(c, repo)
}
