package domain

import (
	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/invoice/core/port"
)

// InvoiceItem is the domain entity for invoice item
type InvoiceItem struct {
	repo             port.Repo
	id               string
	invoice          port.Invoice
	serviceReference string
	description      string
	amount           float64
	quantity         uint64
}

// NewInvoiceItem creates a new invoice item
func NewInvoiceItem(repo port.Repo) *InvoiceItem {
	return &InvoiceItem{
		repo: repo,
	}
}

// Load loads a invoice item
func (i *InvoiceItem) Load(dto port.CreateInputItemDto, invoice port.Invoice) error {
	i.id = uuid.New().String()
	i.invoice = invoice
	i.serviceReference = dto.GetReference()
	i.description = dto.GetDescription()
	var err error
	if i.quantity, err = dto.GetQuantity(); err != nil {
		return err
	}
	if i.amount, err = dto.GetPrice(); err != nil {
		return err
	}
	return nil
}

// Save stores the invoice item on the repository
func (i *InvoiceItem) Save() error {
	return i.repo.SaveInvoiceItem(i)
}

// GetId returns the invoice item id
func (i *InvoiceItem) GetId() string {
	return i.id
}

// GetInvoiceId returns the id of invoice
func (i *InvoiceItem) GetInvoiceId() string {
	return i.invoice.GetId()
}

// GetServiceReference returns the service reference of invoice item
func (i *InvoiceItem) GetServiceReference() string {
	return i.serviceReference
}

// GetDescription returns the description of invoice item
func (i *InvoiceItem) GetDescription() string {
	return i.description
}

// GetAmount returns the amount of invoice item
func (i *InvoiceItem) GetAmount() float64 {
	return i.amount
}

// GetQuantity returns the quantity of invoice item
func (i *InvoiceItem) GetQuantity() uint64 {
	return i.quantity
}
