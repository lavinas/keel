package register

import (
	"errors"
	"strconv"
)

type RegisterInvoiceItem struct {
	ProductID          string `json:"product_id"`
	ProductDescription string `json:"product_description"`
	Description        string `json:"description"`
	Quantity           string `json:"quantity"`
	UnitPrice          string `json:"unit_price"`
}

// Validate validates the invoice item dto
func (i *RegisterInvoiceItem) Validate() error {
	return ValidateLoop([]func() error{
		i.ValidateProductID,
		i.ValidateDescription,
		i.ValidateQuantity,
		i.ValidateUnitPrice,
	})
}

// ValidateProductID validates the product id of the invoice item dto
func (i *RegisterInvoiceItem) ValidateProductID() error {
	if i.ProductID == "" {
		return errors.New(ErrRegisterInvoiceItemDtoProductIDRequired)
	}
	return nil
}

// ValidateProductDescription validates the product description of the invoice item dto
func (i *RegisterInvoiceItem) ValidateProductDescription() error {
	if i.ProductDescription == "" {
		return errors.New(ErrRegisterInvoiceItemDtoProductDescriptionRequired)
	}
	return nil
}

// ValidateDescription validates the description of the invoice item dto
func (i *RegisterInvoiceItem) ValidateDescription() error {
	if i.Description == "" {
		return errors.New(ErrRegisterInvoiceItemDtoDescriptionRequired)
	}
	return nil
}

// ValidateQuantity validates the quantity of the invoice item dto
func (i *RegisterInvoiceItem) ValidateQuantity() error {
	if i.Quantity == "" {
		return errors.New(ErrRegisterInvoiceItemDtoQuantityRequired)
	}
	if _, err := strconv.Atoi(i.Quantity); err != nil {
		return errors.New(ErrRegisterInvoiceItemDtoQuantityInvalid)
	}
	return nil
}

// ValidateUnitPrice validates the unit price of the invoice item dto
func (i *RegisterInvoiceItem) ValidateUnitPrice() error {
	if i.UnitPrice == "" {
		return errors.New(ErrRegisterInvoiceItemDtoUnitPriceRequired)
	}
	if _, err := strconv.ParseFloat(i.UnitPrice, 64); err != nil {
		return errors.New(ErrRegisterInvoiceItemDtoUnitPriceInvalid)
	}
	return nil
}

// GetAmount returns the amount of the invoice item dto
func (i *RegisterInvoiceItem) GetAmount() float64 {
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
