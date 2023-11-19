package dto

import (
	"errors"
	"strconv"
	"strings"
	"time"
)


// InvoiceCreate is the DTO for creating an invoice.
type InvoiceCreate struct {
	BusinessID    string        `json:"business_id"`
	CustomerID    string        `json:"customer_id"`
	Number        string        `json:"number"`
	Date          string        `json:"date"`
	Due           string        `json:"due"`
	Amount        string        `json:"amount"`
	Items         []*InvoiceItemCreate `json:"items"`
	InstructionID string        `json:"instruction_id"`
}

// Validate validates the invoice create DTO
func (i *InvoiceCreate) Validate() error {
	return ValidateLoop([]func() error{
		i.ValidateBusinessID,
		i.ValidateCustomerID,
		i.ValidateNumber,
		i.ValidateDate,
		i.ValidateDue,
		i.ValidateAmount,
		i.ValidateItems,
	})

}

// ValidateBusinessID validates the business id of the invoice create DTO
func (i *InvoiceCreate) ValidateBusinessID() error {
	if i.BusinessID == "" {
		return errors.New(ErrInvoiceCreateDtoBusinessIDRequired)
	}
	return nil
}

// ValidateCustomerID validates the customer id of the invoice create DTO
func (i *InvoiceCreate) ValidateCustomerID() error {
	if i.CustomerID == "" {
		return errors.New(ErrInvoiceCreateDtoCustomerIDRequired)
	}
	return nil
}

// ValidateNumber validates the number of the invoice create DTO
func (i *InvoiceCreate) ValidateNumber() error {
	if i.Number == "" {
		return errors.New(ErrInvoiceCreateDtoNumberRequired)
	}
	if len(strings.Split(i.Number, " ")) > 1 {
		return errors.New(ErrInvoiceCreateDtoNumberInvalid)
	}
	return nil
}

// ValidateDate validates the date of the invoice create DTO
func (i *InvoiceCreate) ValidateDate() error {
	if i.Date == "" {
		return errors.New(ErrInvoiceCreateDtoDateRequired)
	}
	dt, err := time.Parse(time.DateOnly, i.Date)
	if err != nil {
		return errors.New(ErrInvoiceCreateDtoDateInvalid)
	}
	if dt.IsZero() {
		return errors.New(ErrInvoiceCreateDtoDateInvalid)
	}
	return nil
}

// ValidateDue validates the due of the invoice create DTO
func (i *InvoiceCreate) ValidateDue() error {
	if i.Due == "" {
		return errors.New(ErrInvoiceCreateDtoDueRequired)
	}
	dt, err := time.Parse(time.DateOnly, i.Due)
	if err != nil {
		return errors.New(ErrInvoiceCreateDtoDueInvalid)
	}
	if dt.IsZero() {
		return errors.New(ErrInvoiceCreateDtoDueInvalid)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice create DTO
func (i *InvoiceCreate) ValidateAmount() error {
	if i.Amount == "" {
		return errors.New(ErrInvoiceCreateDtoAmountRequired)
	}
	v, err := strconv.ParseFloat(i.Amount, 64)
	if err != nil {
		return errors.New(ErrInvoiceCreateDtoAmountInvalid)
	}
	if v <= 0 {
		return errors.New(ErrInvoiceCreateDtoAmountInvalid)
	}
	return nil
}

// ValidateItems validates the items of the invoice create DTO
func (i *InvoiceCreate) ValidateItems() error {
	if len(i.Items) == 0 {
		return nil
	}
	sum := float64(0)
	for _, item := range i.Items {
		if err := item.Validate(); err != nil {
			return err
		}
		sum += item.GetAmount()
	}
	amount, err := strconv.ParseFloat(i.Amount, 64)
	if err != nil {
		return nil
	}
	if sum != amount {
		return errors.New(ErrInvoiceCreateDtoAmountInvalid)
	}
	return nil
}
