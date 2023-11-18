package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/ktools"
)

const (
	INVOICE_CLASS          = "invoice"
	PAYMENT_CLASS          = "payment"
	INVOICE_NONE           = "none"
	INVOICE_GETTING_CLIENT = "getting"
	INVOICE_WAITING_CLIENT = "waiting"
	INVOICE_CREATED        = "created"
	INVOICE_DELIVERED      = "delivered"
	INVOICE_SAW            = "saw"
	INVOICE_CANCELLED      = "cancelled"
	PAYMENT_NONE           = "none"
	PAYMENT_OPENED         = "opened"
	PAYMENT_UNDERPAID      = "underpaid"
	PAYMENT_PAID           = "paid"
	PAYMENT_OVERPAID       = "overpaid"
	PAYMENT_EXPIRED        = "expired"
	PAYMENT_REVERSED       = "reversed"
)

// InvoiceStatus controls the status of an invoice
type InvoiceStatus struct {
	repo    port.Repo
	invoice *Invoice
	ktools.KGraph
	lastStatus map[string]string
}

// NewInvoiceGraph creates a new graph of the status of an invoice
func NewInvoiceGraph(repo port.Repo, invoice *Invoice) *InvoiceStatus {
	status := &InvoiceStatus{
		repo:       repo,
		invoice:    invoice,
		KGraph:     *ktools.NewKGraph(),
		lastStatus: make(map[string]string),
	}
	status.lastStatus[INVOICE_CLASS] = INVOICE_NONE
	status.lastStatus[PAYMENT_CLASS] = PAYMENT_NONE
	return status
}

// LoadRepository loads the graph of the status of an invoice
func (g *InvoiceStatus) LoadRepository() error {
	if err := g.repo.GetInvoiceVertex(g); err != nil {
		return err
	}
	return g.repo.GetInvoiceEdge(g)
}

// GetInvoiceId returns the id of the invoice
func (g *InvoiceStatus) GetInvoiceId() string {
	return g.invoice.GetId()
}

// ChangeStatus changes the status of an invoice class graph
func (g *InvoiceStatus) Change(class string, status string, description string, author string) error {
	last := g.lastStatus[class]

	if !g.CheckEdge(class, last, status) {
		return errors.New("invalid status")
	}
	id := uuid.New().String()
	g.EnqueueEdge(id, class, last, status, description, author)
	g.lastStatus[class] = status
	return nil
}

// SaveLog saves the log of the invoice class graph
func (g *InvoiceStatus) Save() error {
	for class := range g.lastStatus {
		if err := g.repo.CreateInvoiceStatusLog(class, g); err != nil {
			return err
		}
		if err := g.repo.StoreInvoiceStatus(class, g); err != nil {
			return err
		}
	}
	return nil
}

// GetLastStatusId returns the last status of the invoice class graph
func (g *InvoiceStatus) GetLastStatusId(class string) string {

	return g.lastStatus[class]
}
