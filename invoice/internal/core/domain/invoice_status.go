package domain

import (
	"errors"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/ktools"
)

const (
	INVOICE_NONE              = "none"
	INVOICE_CREATING          = "creating"
	INVOICE_CONSULTING_CLIENT = "consulting"
	INVOICE_WAITING_CLIENT    = "waiting"
	INVOICE_CREATED           = "created"
	INVOICE_DELIVERED         = "delivered"
	INVOICE_SAW               = "saw"
	INVOICE_CANCELLED         = "cancelled"
	PAYMENT_NONE              = "none"
	PAYMENT_OPENED            = "opened"
	PAYMENT_UNDERPAID         = "underpaid"
	PAYMENT_PAID              = "paid"
	PAYMENT_OVERPAID          = "overpaid"
	PAYMENT_EXPIRED           = "expired"
	PAYMENT_REVERSED          = "reversed"
)

// InvoiceStatus controls the status of an invoice
type InvoiceStatusGraph struct {
	repo port.Repo
	ktools.KGraph
	lastInvoiceStatusVertex string
	lastPaymentStatusVertex string
}

// NewInvoiceGraph creates a new graph of the status of an invoice
func NewInvoiceGraph(repo port.Repo) *InvoiceStatusGraph {
	return &InvoiceStatusGraph{
		repo:                    repo,
		lastInvoiceStatusVertex: INVOICE_NONE,
		lastPaymentStatusVertex: PAYMENT_NONE,
	}
}

// LoadRepository loads the graph of the status of an invoice
func (g *InvoiceStatusGraph) LoadRepository() error {
	if err := g.repo.LoadInvoiceVertex(g); err != nil {
		return err
	}
	return g.repo.LoadInvoiceEdge(g)
}

// ChangeStatus changes the status of an invoice class graph
func (g *InvoiceStatusGraph) ChangeStatus(class string, status string, description string, author string) error {
	if !g.CheckEdge(class, g.lastInvoiceStatusVertex, status) {
		return errors.New("invalid status")
	}
	g.EnqueueEdge(class, g.lastInvoiceStatusVertex, status, description, author)
	g.lastInvoiceStatusVertex = status
	return nil
}

// SaveLog saves the log of the invoice class graph
func (g *InvoiceStatusGraph) SaveLog() error {
	g.repo.Begin()
	defer g.repo.Rollback()
	if err := g.repo.LogInvoiceEdge("invoice", g); err != nil {
		return err
	}
	if err := g.repo.LogInvoiceEdge("payment", g); err != nil {
		return err
	}
	g.repo.Commit()
	return nil
}
