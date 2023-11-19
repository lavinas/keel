package dto

import (
	"errors"
	"strconv"
)

type InvoiceItemCreate struct {
	ProductID   string  `json:"product_id"`
	Description string  `json:"description"`
	Quantity    string     `json:"quantity"`
	UnitPrice   string `json:"unit_price"`
}

// Validate validates the invoice item dto
func (i *InvoiceItemCreate) Validate() error {
	return ValidateLoop([]func() error{
		i.ValidateProductID,
		i.ValidateDescription,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	})
}

// ValidateProductID validates the product id of the invoice item dto
func (i *InvoiceItemCreate) ValidateProductID() error {
	if i.ProductID == "" {
		return errors.New(ErrInvoiceItemCreateDtoProductIDRequired)
	}
	return nil
}

// ValidateDescription validates the description of the invoice item dto
func (i *InvoiceItemCreate) ValidateDescription() error {
	if i.Description == "" {
		return errors.New(ErrInvoiceItemCreateDtoDescriptionRequired)
	}
	return nil
}

// ValidateQuantity validates the quantity of the invoice item dto
func (i *InvoiceItemCreate) ValidateQuantity() error {
	if i.Quantity == "" {
		return errors.New(ErrInvoiceItemCreateDtoQuantityRequired)
	}
	if _, err := strconv.Atoi(i.Quantity); err != nil {
		return errors.New(ErrInvoiceItemCreateDtoQuantityInvalid)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item dto
func (i *InvoiceItemCreate) ValidateUnitPrice() error {
	if i.UnitPrice == "" {
		return errors.New(ErrInvoiceItemCreateDtoUnitPriceRequired)
	}
	if _, err := strconv.ParseFloat(i.UnitPrice, 64); err != nil {
		return errors.New(ErrInvoiceItemCreateDtoUnitPriceInvalid)
	}
	return nil
}

// GetAmount returns the amount of the invoice item dto
func (i *InvoiceItemCreate) GetAmount() float64 {
	quantity, err := strconv.Atoi(i.Quantity)
	if err != nil {
		return 0
	}
	unitPrice, err := strconv.ParseFloat(i.UnitPrice, 64)
	if err != nil {
		return 0
	}
	return float64(quantity) * unitPrice
}


