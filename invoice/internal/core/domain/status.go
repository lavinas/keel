package domain

import (
	"errors"
)

const (
	InvoiceStatusDraft     = "draft"
	InvoiceStatusNew       = "new"
	InvoiceStatusSent      = "sent"
	InvoiceStatusViewed    = "viewed"
	InvoiceStatusCancelled = "cancelled"
	PaymentStatusUnpaid    = "unpaid"
	PaymentStatusPaid      = "paid"
	PaymentStatusUnderpaid = "underpaid"
	PaymentStatusOverpaid  = "overpaid"
	PaymentStatusReversed  = "reversed"
)

// InvoiceStatus represents a status of the invoice
type InvoiceStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (i *InvoiceStatus) Validate() error {
	if i.ID == "" {
		return errors.New(ErrInvoiceStatusIDIsRequired)
	}
	if i.Name == "" {
		return errors.New(ErrInvoiceStatusNameIsRequired)
	}
	if i.ID != InvoiceStatusDraft && i.ID != InvoiceStatusNew && i.ID != InvoiceStatusSent &&
		i.ID != InvoiceStatusViewed && i.ID != InvoiceStatusCancelled {
		return errors.New(ErrInvoiceStatusIDIsInvalid)
	}
	return nil
}

// PaymentStatus represents a status of the payment of the invoice
type PaymentStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ValidatePaymentStatus validates the payment status
func (p *PaymentStatus) Validate() error {
	if p.ID == "" {
		return errors.New(ErrPaymentStatusIDIsRequired)
	}
	if p.Name == "" {
		return errors.New(ErrPaymentStatusNameIsRequired)
	}
	if p.ID != PaymentStatusUnpaid && p.ID != PaymentStatusPaid && p.ID != PaymentStatusUnderpaid &&
		p.ID != PaymentStatusOverpaid && p.ID != PaymentStatusReversed {
		return errors.New(ErrPaymentStatusIDIsInvalid)
	}
	return nil
}
