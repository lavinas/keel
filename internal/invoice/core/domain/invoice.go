package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/invoice/core/port"
)

var (
	status_map = map[string]uint{
		"New":       0,
		"Open":      1,
		"Sent":      2,
		"Paid":      3,
		"Cancelled": 4,
	}
)

// Invoice is the domain model for a invoice
type Invoice struct {
	repo         port.Repo
	id           string
	reference    string
	client_id    string
	client_name  string
	client_email string
	client_phone uint64
	value        float64
	paid_value   float64
	date         time.Time
	due          time.Time
	notes        string
	status_id    uint
	status_name  string
	created_at   time.Time
	updated_at   time.Time
}

// NewInvoice creates a new invoice
func NewInvoice(repo port.Repo) *Invoice {
	return &Invoice{
		repo: repo,
	}
}

// Insert loads a invoice
func (i *Invoice) Create(reference, client_id, client_name, client_email string, client_phone uint64,
	value float64, date, due time.Time, notes string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	i.id = id.String()
	i.reference, i.client_id, i.client_name, i.client_email, i.client_phone, i.value, i.date, i.due, i.notes =
		reference, client_id, client_name, client_email, client_phone, value, date, due, notes
	i.paid_value = 0
	i.status_name = "New"
	i.status_id = status_map[i.status_name]
	i.created_at = time.Now()
	i.updated_at = time.Now()
	return nil
}

// Load loads a invoice values
func (i *Invoice) Get() (string, string, string, string, string, uint64, float64, float64, time.Time, time.Time, string) {
	return i.id, i.reference, i.client_id, i.client_name, i.client_email, i.client_phone, i.value, i.paid_value, i.date, i.due, i.notes
}

// GetStatus returns the invoice status id and name
func (i *Invoice) GetStatus() (uint, string) {
	return i.status_id, i.status_name
}

// SetSent sets the invoice as sent
func (i *Invoice) SetSent(sent time.Time) error {
	i.status_name = "sent"
	i.status_id = status_map[i.status_name]
	i.updated_at = time.Now()
	return nil
}

// SetPaid sets the invoice as paid
func (i *Invoice) SetPaid() error {
	i.status_name = "paid"
	i.status_id = status_map[i.status_name]
	i.updated_at = time.Now()
	return nil
}

// SetCancelled sets the invoice as cancelled
func (i *Invoice) SetCancelled() error {
	i.status_name = "cancelled"
	i.status_id = status_map[i.status_name]
	i.updated_at = time.Now()
	return nil
}
