package domain

import (
	"time"
)

const (
	// Email status values
	None      = "none"
	Created   = "created"
	Sending   = "sending"
	Sent      = "sent"
	SendError = "send_error"
)

// InvoiceStatus represents a status of the invoice
type Status struct {
	EmailID   string    `json:"-"           gorm:"primaryKey;type:varchar(50); not null"`
	ID        string    `json:"id"          gorm:"primaryKey;type:varchar(50); not null"`
	CreatedAt time.Time `json:"created_at"  gorm:"type:timestamp; not null"`
	Note      string    `json:"note"        gorm:"type:varchar(250); not null"`
}

// NewInvoiceStatus creates a new invoice status
func NewStatus(EmailID string) *Status {
	return &Status{
		EmailID:   EmailID,
		ID:        None,
		CreatedAt: time.Now(),
		Note:      "",
	}
}

// SetCreated sets the status to created
func (i *Status) SetCreated(note string) {
	i.ID = Created
	i.Note = note
}

// SetSending sets the status to sending
func (i *Status) SetSending(note string) {
	i.ID = Sending
	i.Note = note
}

// SetSent sets the status to sent
func (i *Status) SetSent(note string) {
	i.ID = Sent
	i.Note = note
}

// SetSendError sets the status to send error
func (i *Status) SetSendError(note string) {
	i.ID = SendError
	i.Note = note
}

// TableName returns the table name for gorm
func (b *Status) TableName() string {
	return "status"
}
