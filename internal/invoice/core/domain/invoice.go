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
	repo      port.Repo
	id        string
	reference string
	business  port.InvoiceClient
	customer  port.InvoiceClient
	amount    float64
	date      time.Time
	due       time.Time
	status_id uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewInvoice creates a new invoice
func NewInvoice(repo port.Repo) *Invoice {
	return &Invoice{
		repo: repo,
	}
}

// Load loads a invoice from input
func (i *Invoice) Load(input port.CreateInputDto, business port.InvoiceClient, customer port.InvoiceClient) error {
	i.id = uuid.New().String()
	i.reference = input.GetReference()
	var err error
	i.amount, err = input.GetAmount()
	if err != nil {
		return err
	}
	i.date, err = input.GetDate()
	if err != nil {
		return err
	}
	i.due, err = input.GetDue()
	if err != nil {
		return err
	}
	i.business = business
	i.customer = customer
	i.status_id = status_map["New"]
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	return nil
}

// SetAmount sets the amount of invoice
func (i *Invoice) SetAmount(amount float64) {
	i.amount = amount
}

// Save stores the invoice on the repository
func (i *Invoice) Save() error {
	return i.repo.SaveInvoice(i)
}

// GetId returns the id of invoice
func (i *Invoice) GetId() string {
	return i.id
}

// GetReference returns the reference of invoice
func (i *Invoice) GetReference() string {
	return i.reference
}

// GetBusiness returns the business of invoice
func (i *Invoice) GetBusinessId() string {
	return i.business.GetId()
}

// GetCustomer returns the customer of invoice
func (i *Invoice) GetCustomerId() string {
	return i.customer.GetId()
}

// GetAmount returns the amount of invoice
func (i *Invoice) GetAmount() float64 {
	return i.amount
}

// GetDate returns the date of invoice
func (i *Invoice) GetDate() time.Time {
	return i.date
}

// GetDue returns the due of invoice
func (i *Invoice) GetDue() time.Time {
	return i.due
}

func (i *Invoice) GetNoteId() *string {
	return nil
}

// GetStatusId returns the status id of invoice
func (i *Invoice) GetStatusId() uint {
	return i.status_id
}

// GetCreatedAt returns the created at of invoice
func (i *Invoice) GetCreatedAt() time.Time {
	return i.CreatedAt
}

// GetUpdatedAt returns the updated at of invoice
func (i *Invoice) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}
