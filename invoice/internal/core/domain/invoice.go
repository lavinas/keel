package domain

import (
	"errors"
	"time"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	Business      *Client        `json:"business"`
	Customer      *Client        `json:"customer"`
	Date          time.Time      `json:"date"`
	Due           time.Time      `json:"due"`
	Amount        float64        `json:"amount"`
	Items         []*InvoiceItem `json:"item"`
	Instruction   *Instruction   `json:"instruction"`
	InvoiceStatus *InvoiceStatus `json:"invoice_status"`
	PaymentStatus *PaymentStatus `json:"payment_status"`
}

// Validate validates the invoice
func (p *Invoice) Validate() error {
	return ValidateLoop([]func() error{
		p.Base.Validate,
		p.ValidateBusiness,
		p.ValidateCustomer,
		p.ValidateDate,
		p.ValidateDue,
		p.ValidateAmount,
		p.ValidateItems,
		p.ValidateInstruction,
		p.ValidateInvoiceStatus,
		p.ValidatePaymentStatus,
	})
}

// ValidateBusiness validates the business of the invoice
func (p *Invoice) ValidateBusiness() error {
	if p.Business == nil {
		return errors.New(ErrInvoiceBusinessIsRequired)
	}
	return p.Business.Validate()
}

// ValidateCustomer validates the custumer of the invoice
func (p *Invoice) ValidateCustomer() error {
	if p.Customer == nil {
		return errors.New(ErrInvoiceCustomerIsRequired)
	}
	return p.Customer.Validate()
}

// ValidateDate validates the Date of the invoice
func (p *Invoice) ValidateDate() error {
	if p.Date.IsZero() {
		return errors.New(ErrInvoiceDateIsRequired)
	}
	return nil
}

// ValidateDue validates the Due Date of the invoice
func (p *Invoice) ValidateDue() error {
	if p.Due.IsZero() {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice
func (p *Invoice) ValidateAmount() error {
	itemAmount := p.GetItemsAmount()
	if itemAmount > 0 && p.Amount != itemAmount {
		return errors.New(ErrInvoiceAmountNotMatch)
	}
	if p.Amount <= 0 {
		return errors.New(ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// Validate validates the invoice
func (p *Invoice) GetItemsAmount() float64 {
	sum := float64(0)
	for _, item := range p.Items {
		sum += item.GetAmount()
	}
	return sum
}

// ValidateItems validates the items of the invoice
func (p *Invoice) ValidateItems() error {
	for _, item := range p.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// ValidateInstruction validates the instruction of the invoice
func (p *Invoice) ValidateInstruction() error {
	if p.Instruction == nil {
		return nil
	}
	return p.Instruction.Validate()
}

// ValidateInvoiceStatus validates the invoice
func (p *Invoice) ValidateInvoiceStatus() error {
	if p.InvoiceStatus == nil {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	return p.InvoiceStatus.Validate()
}

// ValidatePaymentStatus validates the invoice
func (p *Invoice) ValidatePaymentStatus() error {
	if p.PaymentStatus == nil {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	return p.PaymentStatus.Validate()
}
