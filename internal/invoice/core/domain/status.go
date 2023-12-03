package domain

import (
	"time"
)

const (
	// Status sets
	InvoiceSet = "invoice"
	PaymentSet = "payment"
	// Common status values
	None = "none"
	// Invoice status values
	Created = "created"
	// Payment status values
	Open = "open"
)

// InvoiceStatus represents a status of the invoice
type Status struct {
	InvoiceID string    `json:"-"           gorm:"primaryKey;type:varchar(50); not null"`
	ID        string    `json:"id"          gorm:"primaryKey;type:varchar(50); not null"`
	Set       string    `json:"-"           gorm:"type:varchar(50); not null"`
	CreatedAt time.Time `json:"created_at"  gorm:"type:timestamp; not null"`
}

// NewInvoiceStatus creates a new invoice status
func NewInvoiceStatus(invoiceID string, set string) *Status {
	return &Status{
		InvoiceID: invoiceID,
		Set:       set,
		ID:        None,
	}
}

// SetCreated sets the status to created
func (i *Status) SetCreated() {
	if i.Set == InvoiceSet {
		i.ID = Created
	}
}

// SetOpen sets the status to open
func (i *Status) SetOpen() {
	if i.Set == PaymentSet {
		i.ID = Open
	}
}

// TableName returns the table name for gorm
func (i *Status) TableName() string {
	return "status"
}
