package domain

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/invoice/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

// Item represents a item in the invoice
type Item struct {
	Base
	InvoiceID    string   `json:"invoice_id"  gorm:"type:varchar(50); not null"`
	ProductID    string   `json:"product_id"  gorm:"type:varchar(50); not null"`
	Product      *Product `json:"product"     gorm:"foreignKey:BusinessID,ProductID;associationForeignKey:BusinessID,ID"`
	Description  string   `json:"description" gorm:"type:varchar(255)"`
	QuantityStr  string   `json:"quantity"    gorm:"-"`
	Quantity     int      `json:"-"           gorm:"type:int; not null"`
	UnitPriceStr string   `json:"unit_price"  gorm:"-"`
	UnitPrice    float64  `json:"-"           gorm:"type:decimal(20, 2); not null"`
}

// SetCreate set information for create a new invoice item
func (i *Item) SetCreate(business_id string) {
	i.Base.SetCreate(business_id)
	i.Fit()
	if i.Product != nil {
		i.Product.SetCreate(i.BusinessID)
	}
}

// Validate validates the invoice item
func (i *Item) Validate(repo port.Repository) *kerror.KError {
	return ValidateLoop([]func(repo port.Repository) *kerror.KError{
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
func (c *Item) ValidateProduct(repo port.Repository) *kerror.KError {
	if c.ProductID == "" && c.Product == nil {
		return kerror.NewKError(kerror.BadRequest, ErrItemProductRequired)
	} else if c.ProductID != "" && c.Product != nil {
		return kerror.NewKError(kerror.BadRequest, ErrItemProductConflict)
	} else if c.ProductID != "" {
		return c.ValidateProductID(repo)
	} else if err := c.Product.Validate(repo); err != nil {
		err.SetPrefix("product ")
		return err
	}
	return nil
}

// ValidateProductID validates the product id of the invoice item
func (c *Item) ValidateProductID(repo port.Repository) *kerror.KError {
	if c.ProductID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrItemProductRequired)
	}
	var product Product
	product.ID = c.ProductID
	if err := product.ValidateID(repo); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrItemProductInvalid)
	}
	if exists, err := repo.Exists(&product, c.BusinessID, c.ProductID); err != nil {
		return kerror.NewKError(kerror.Internal, ErrItemProductInvalid)
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrItemProductNotFound)
	}
	return nil
}

// ValidateQuantity validates the quantity of the invoice item
func (c *Item) ValidateQuantity(repo port.Repository) *kerror.KError {
	if c.QuantityStr == "" {
		return kerror.NewKError(kerror.BadRequest, ErrItemQuantityRequired)
	} else if q, err := strconv.Atoi(c.QuantityStr); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrItemQuantityInvalid)
	} else if q <= 0 {
		return kerror.NewKError(kerror.BadRequest, ErrItemQuantityLessOrEqualZero)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item
func (c *Item) ValidateUnitPrice(repo port.Repository) *kerror.KError {
	if c.UnitPriceStr == "" {
		return kerror.NewKError(kerror.BadRequest, ErrItemPriceRequired)
	} else if p, err := strconv.ParseFloat(c.UnitPriceStr, 64); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrItemPriceInvalid)
	} else if p <= 0 {
		return kerror.NewKError(kerror.BadRequest, ErrItemPriceLessOrEqualZero)
	}
	return nil
}

// GetAmount returns the amount of the invoice item
func (c *Item) GetAmount() float64 {
	return float64(c.Quantity) * c.UnitPrice
}

// ValidateDuplicity validates the duplicity of the invoice item
func (c *Item) ValidateDuplicity(repo port.Repository) *kerror.KError {
	return c.Base.ValidateDuplicity(c, repo)
}

// TableName returns the table name for gorm
func (c *Item) TableName() string {
	return "item"
}
