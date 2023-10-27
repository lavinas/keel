package domain

import (
	"github.com/lavinas/keel/internal/invoice/core/port"
)

// Domain is the domain factory model for a invoice
type Domain struct {
	repo port.Repo
}

// New creates a new invoice
func NewDomain(repo port.Repo) *Domain {
	return &Domain{
		repo: repo,
	}
}

// GetInvoice returns a invoice
func (d *Domain) GetInvoice() port.Invoice {
	return NewInvoice(d.repo)
}
