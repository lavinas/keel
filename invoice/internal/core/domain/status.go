package domain

import (
	"time"
)

const (
	InvoiceStatusNone      = "none"
	InvoiceStatusCreated   = "created"
)

// InvoiceStatus represents a status of the invoice
type InvoiceStatus struct {
	InvoiceID   string    `json:"-"           gorm:"primaryKey;type:varchar(50); not null"`
	ID          string    `json:"id"          gorm:"primaryKey;type:varchar(50); not null"`
	CreatedAt   time.Time `json:"created_at"  gorm:"type:timestamp; not null"`
}

// NewInvoiceStatus creates a new invoice status
func NewInvoiceStatus(invoiceID string) *InvoiceStatus {
	return &InvoiceStatus{
		InvoiceID:   invoiceID,
		ID:          InvoiceStatusNone,
	}
}

// SetCreated sets the status to created
func (i *InvoiceStatus) SetCreated() {
	id := "created"
	i.ID = id
}

// TableName returns the table name for gorm
func (i *InvoiceStatus) TableName() string {
	return InvoiceStatusCreated
}
