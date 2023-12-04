package domain

import (
	"time"
)

const (
	// Status sets
	EmailSet = "email"
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
	Set       string    `json:"-"           gorm:"type:varchar(50); not null"`
	CreatedAt time.Time `json:"created_at"  gorm:"type:timestamp; not null"`
	Note      string    `json:"description" gorm:"type:varchar(250); not null"`
}

// NewInvoiceStatus creates a new invoice status
func NewInvoiceStatus(EmailID string, set string) *Status {
	return &Status{
		EmailID: EmailID,
		Set:       set,
		ID:        None,
	}
}

// SetCreated sets the status to created
func (i *Status) SetCreated(note string) {
	if i.Set == EmailSet {
		i.ID = Created
		i.Note = note
	}
}

// SetSending sets the status to sending
func (i *Status) SetSending(note string) {
	if i.Set == EmailSet {
		i.ID = Sending
		i.Note = note
	}
}

// SetSent sets the status to sent
func (i *Status) SetSent(note string) {
	if i.Set == EmailSet {
		i.ID = Sent
		i.Note = note
	}
}

// SetSendError sets the status to send error
func (i *Status) SetSendError(note string) {
	if i.Set == EmailSet {
		i.ID = SendError
		i.Note = note
	}
}