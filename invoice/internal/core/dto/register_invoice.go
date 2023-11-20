package dto

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// RegisterInvoice is the DTO for creating an invoice.
type RegisterInvoice struct {
	BusinessID             string                 `json:"business_id"`
	Business               *RegisterClient        `json:"business"`
	CustomerID             string                 `json:"customer_id"`
	Customer               *RegisterClient        `json:"customer"`
	Number                 string                 `json:"number"`
	Date                   string                 `json:"date"`
	Due                    string                 `json:"due"`
	Amount                 string                 `json:"amount"`
	Items                  []*RegisterInvoiceItem `json:"items"`
	InstructionID          string                 `json:"instruction_id"`
	InstructionDescription string                 `json:"instruction_description"`
}

// Validate validates the invoice create DTO
func (i *RegisterInvoice) Validate() error {
	return ValidateLoop([]func() error{
		i.ValidateBusiness,
		i.ValidateCustomer,
		i.ValidateNumber,
		i.ValidateDate,
		i.ValidateDue,
		i.ValidateAmount,
		i.ValidateItems,
	})

}

// ValidateBusinessID validates the business id of the invoice create DTO
func (i *RegisterInvoice) ValidateBusiness() error {
	if i.BusinessID == "" && i.Business == nil {
		return errors.New(ErrRegisterInvoiceDtoBusinessRequired)
	}
	if i.BusinessID != "" && i.Business != nil {
		return errors.New(ErrRegisterInvoiceDtoBusinessDuplicity)
	}
	if i.Business != nil {
		return i.Business.Validate()
	}
	return nil
}

// ValidateCustomerID validates the customer id of the invoice create DTO
func (i *RegisterInvoice) ValidateCustomer() error {
	if i.CustomerID == "" && i.Customer == nil {
		return errors.New(ErrRegisterInvoiceDtoCustomerRequired)
	}
	if i.CustomerID != "" && i.Customer != nil {
		return errors.New(ErrRegisterInvoiceDtoCustomerIDDuplicity)
	}
	if i.Customer != nil {
		return i.Customer.Validate()
	}
	return nil
}

// ValidateNumber validates the number of the invoice create DTO
func (i *RegisterInvoice) ValidateNumber() error {
	if i.Number == "" {
		return errors.New(ErrRegisterInvoiceDtoNumberRequired)
	}
	if len(strings.Split(i.Number, " ")) > 1 {
		return errors.New(ErrRegisterInvoiceDtoNumberInvalid)
	}
	return nil
}

// ValidateDate validates the date of the invoice create DTO
func (i *RegisterInvoice) ValidateDate() error {
	if i.Date == "" {
		return errors.New(ErrRegisterInvoiceDtoDateRequired)
	}
	dt, err := time.Parse(time.DateOnly, i.Date)
	if err != nil {
		return errors.New(ErrRegisterInvoiceDtoDateInvalid)
	}
	if dt.IsZero() {
		return errors.New(ErrRegisterInvoiceDtoDateInvalid)
	}
	return nil
}

// ValidateDue validates the due of the invoice create DTO
func (i *RegisterInvoice) ValidateDue() error {
	if i.Due == "" {
		return errors.New(ErrRegisterInvoiceDtoDueRequired)
	}
	dt, err := time.Parse(time.DateOnly, i.Due)
	if err != nil {
		return errors.New(ErrRegisterInvoiceDtoDueInvalid)
	}
	if dt.IsZero() {
		return errors.New(ErrRegisterInvoiceDtoDueInvalid)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice create DTO
func (i *RegisterInvoice) ValidateAmount() error {
	if i.Amount == "" {
		return errors.New(ErrRegisterInvoiceDtoAmountRequired)
	}
	v, err := strconv.ParseFloat(i.Amount, 64)
	if err != nil {
		return errors.New(ErrRegisterInvoiceDtoAmountInvalid)
	}
	if v <= 0 {
		return errors.New(ErrRegisterInvoiceDtoAmountInvalid)
	}
	return nil
}

// ValidateItems validates the items of the invoice create DTO
func (i *RegisterInvoice) ValidateItems() error {
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
		return errors.New(ErrRegisterInvoiceDtoAmountInvalid)
	}
	return nil
}
