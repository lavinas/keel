package dto

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lavinas/keel/invoice/pkg/ktools"
)

const (
	// ErrItemReferenceEmpty is the error message for an empty reference
	ErrItemReferenceEmpty = "item reference is empty"
	// ErrItemDescriptionEmpty is the error message for an empty description
	ErrItemDescriptionEmpty = "item description is empty"
	// ErrItemQuantityEmpty is the error message for an empty quantity
	ErrItemQuantityEmpty = "item quantity is empty"
	// ErrItemQuantityNotNumeric is the error message for a quantity not numeric
	ErrItemQuantityNotNumeric = "item quantity is not numeric"
	// ErrItemQuantidyZeroNegative is the error message for a quantity zero or negative
	ErrItemQuantidyZeroNegative = "item quantity is zero or negative"
	// ItemQuantityMax is the maximum quantity allowed
	ItemQuantityMax = 1000
	// ErrItemQuantityMax is the error message for a quantity too big
	ErrItemQuantityMax = "item quantity is too big - limit is 1000"
	// ErrItemPriceEmpty is the error message for an empty price
	ErrItemPriceEmpty = "item price is empty"
	// ErrItemPriceNotNumeric is the error message for a price not numeric
	ErrItemPriceNotNumeric = "item price is not numeric"
	// ErrItemPriceZeroNegative is the error message for a price zero or negative
	ErrItemPriceZeroNegative = "item price is zero or negative"
	// ItemPriceMax is the maximum price allowed
	ItemPriceMax = 900000
	// ErrItemPriceMax is the error message for a price too big
	ErrItemPriceMax = "item price is too big - limit is 900000"
)

// CreateInputProductDto is the DTO for a Product Item of a new invoice creation
type CreateInputItemDto struct {
	Reference   string `json:"reference"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

// Validate validates the InsertInputDto
func (i CreateInputItemDto) Validate() error {
	validationMap := map[string]func() error{
		"reference":   i.ValidateReference,
		"description": i.ValidateDescription,
		"quantity":    i.ValidateQuantity,
		"price":       i.ValidatePrice,
	}
	var errs []error
	for _, value := range validationMap {
		errs = append(errs, value())
	}
	if err := ktools.MergeError(errs...); err != nil {
		return err
	}
	return nil
}

// GetReference returns the reference
func (i CreateInputItemDto) GetReference() string {
	return i.Reference
}

// GetDescription returns the description
func (i CreateInputItemDto) GetDescription() string {
	return i.Description
}

// GetQuantity returns the quantity
func (i CreateInputItemDto) GetQuantity() (uint64, error) {
	quantity, err := strconv.ParseUint(i.Quantity, 10, 64)
	if err != nil {
		return 0, errors.New(ErrItemQuantityNotNumeric)
	}
	return quantity, nil
}

// GetPrice returns the price
func (i CreateInputItemDto) GetPrice() (float64, error) {
	price, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return 0, errors.New(ErrItemPriceNotNumeric)
	}
	return price, nil
}

// ValidateReference validates the reference
func (i CreateInputItemDto) ValidateReference() error {
	i.Reference = strings.Trim(i.Reference, " ")
	if i.Reference == "" {
		return errors.New(ErrItemReferenceEmpty)
	}
	return nil
}

// ValidateDescription validates the description
func (i CreateInputItemDto) ValidateDescription() error {
	i.Description = strings.Trim(i.Description, " ")
	if i.Description == "" {
		return errors.New(ErrItemDescriptionEmpty)
	}
	return nil
}

// ValidateQuantity validates the quantity
func (i CreateInputItemDto) ValidateQuantity() error {
	i.Quantity = strings.Trim(i.Quantity, " ")
	if i.Quantity == "" {
		return errors.New(ErrItemQuantityEmpty)
	}
	quantity, err := strconv.Atoi(i.Quantity)
	if err != nil {
		return errors.New(ErrItemQuantityNotNumeric)
	}
	if quantity <= 0 {
		return errors.New(ErrItemQuantidyZeroNegative)
	}
	if quantity > ItemQuantityMax {
		return errors.New(ErrItemQuantityMax)
	}
	return nil
}

// ValidatePrice validates the price
func (i CreateInputItemDto) ValidatePrice() error {
	i.Price = strings.Trim(i.Price, " ")
	if i.Price == "" {
		return errors.New(ErrItemPriceEmpty)
	}
	price, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return errors.New(ErrItemPriceNotNumeric)
	}
	if price <= 0 {
		return errors.New(ErrItemPriceZeroNegative)
	}
	if price > ItemPriceMax {
		return errors.New(ErrItemPriceMax)
	}
	return nil
}
