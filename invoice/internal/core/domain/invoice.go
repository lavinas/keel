package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/ktools"
)

var (
	client_expiration_seconds = 60 * 60 * 24 * 30
)

// Invoice is the domain model for a invoice
type Invoice struct {
	repo      port.Repo
	status    *InvoiceStatus
	id        string
	reference string
	business  *InvoiceClient
	customer  *InvoiceClient
	amount    float64
	date      time.Time
	due       time.Time
	items     []*InvoiceItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewInvoice creates a new invoice
func NewInvoice(repo port.Repo) *Invoice {
	invoice := &Invoice{repo: repo}
	status := NewInvoiceGraph(repo, invoice)
	status.LoadRepository()
	invoice.status = status
	return invoice
}

// Load loads a invoice from input
func (i *Invoice) Load(input port.CreateInputDto) error {
	i.id = uuid.New().String()
	i.reference = input.GetReference()
	if err := i.loadClients(input); err != nil {
		return err
	}
	if err := ktools.MergeError(i.loadAmount(input), i.loadDate(input), i.loadDue(input),
		i.loadItems(input.GetItems())); err != nil {
		return err
	}
	i.status.Change(INVOICE_CLASS, INVOICE_GETTING_CLIENT, "", "")
	i.status.Change(PAYMENT_CLASS, PAYMENT_OPENED, "", "")
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	return nil
}

// IsDuplicated returns true if the invoice is duplicated
func (i *Invoice) IsDuplicated() (bool, error) {
	return i.repo.IsDuplicatedInvoice(i.reference)
}

// SetAmount sets the amount of invoice
func (i *Invoice) SetAmount(amount float64) {
	i.amount = amount
}

// Save stores the invoice on the repository
func (i *Invoice) Save() error {
	execOrder := []func() error{
		i.repo.Begin,
		i.saveClients,
		i.saveInvoice,
		i.status.Save,
		i.saveItems,
		i.repo.Commit,
	}
	defer i.repo.Rollback()
	for _, exec := range execOrder {
		if err := exec(); err != nil {
			return err
		}
	}
	return nil
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

// GetBusiness returns the business client object of invoice
func (i *Invoice) GetBusiness() port.InvoiceClient {
	return i.business
}

// Getcustomerreturns the customerof invoice
func (i *Invoice) GetCustomerId() string {
	return i.customer.GetId()
}

// Getcustomerreturns the customerclient object of invoice
func (i *Invoice) GetCustomer() port.InvoiceClient {
	return i.customer
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
	return 1
}

// GetCreatedAt returns the created at of invoice
func (i *Invoice) GetCreatedAt() time.Time {
	return i.CreatedAt
}

// GetUpdatedAt returns the updated at of invoice
func (i *Invoice) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}

// loadClients loads the clients (business and customer) from input
func (i *Invoice) loadClients(input port.CreateInputDto) error {
	i.business = NewInvoiceClient(i.repo)
	createdAfter := time.Now().Add(time.Duration(-client_expiration_seconds) * time.Second)
	if b, err := i.business.GetLastInvoiceClient(input.GetBusinessNickname(), createdAfter); err != nil {
		return err
	} else if !b {
		i.business.Load("", input.GetBusinessNickname(), "", "", "", 0, 0, time.Time{})
	}
	i.customer = NewInvoiceClient(i.repo)
	if b, err := i.customer.GetLastInvoiceClient(input.GetCustomerNickname(), createdAfter); err != nil {
		return err
	} else if !b {
		i.customer.Load("", input.GetCustomerNickname(), "", "", "", 0, 0, time.Time{})
	}
	return nil
}

// loadAmount loads the amount from input
func (i *Invoice) loadAmount(input port.CreateInputDto) error {
	var err error
	i.amount, err = input.GetAmount()
	if err != nil {
		return err
	}
	return nil
}

// loadDate loads the date from input
func (i *Invoice) loadDate(input port.CreateInputDto) error {
	var err error
	i.date, err = input.GetDate()
	if err != nil {
		return err
	}
	return nil
}

// loadDue loads the due from input
func (i *Invoice) loadDue(input port.CreateInputDto) error {
	var err error
	i.due, err = input.GetDue()
	if err != nil {
		return err
	}
	return nil
}

// loadItems loads the items from input
func (i *Invoice) loadItems(inputItems []port.CreateInputItemDto) error {
	for _, inputItem := range inputItems {
		item := NewInvoiceItem(i.repo)
		err := item.Load(inputItem, i)
		if err != nil {
			return err
		}
		i.items = append(i.items, item)
	}
	return nil
}

// saveInvoice saves the invoice on the repository
func (i *Invoice) saveInvoice() error {
	return i.repo.SaveInvoice(i)
}

// saveClients saves the clients (business and customer) on the repository
func (i *Invoice) saveClients() error {
	if i.business.IsNew() {
		if err := i.business.Save(); err != nil {
			return err
		}
	}
	if i.customer.IsNew() {
		if err := i.customer.Save(); err != nil {
			return err
		}
	}
	return nil
}

// saveItems saves the items on the repository
func (i *Invoice) saveItems() error {
	for _, item := range i.items {
		if err := item.Save(); err != nil {
			return err
		}
	}
	return nil
}
