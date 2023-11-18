package domain

import (
	"encoding/json"
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
	repo      port.Repo      `json:"-"`
	Status    *InvoiceStatus `json:"status"`
	Id        string         `json:"id"`
	Reference string         `json:"reference"`
	Business  *InvoiceClient `json:"business"`
	Customer  *InvoiceClient `json:"customer"`
	Amount    float64        `json:"amount"`
	Date      time.Time      `json:"date"`
	Due       time.Time      `json:"due"`
	Items     []*InvoiceItem `json:"items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time
}

// NewInvoice creates a new invoice
func NewInvoice(repo port.Repo) *Invoice {
	invoice := &Invoice{repo: repo}
	status := NewInvoiceGraph(repo, invoice)
	status.LoadRepository()
	invoice.Status = status
	return invoice
}

// Load loads a invoice from input
func (i *Invoice) Load(input port.CreateInputDto) error {
	i.Id = uuid.New().String()
	i.Reference = input.GetReference()
	if err := i.loadClients(input); err != nil {
		return err
	}
	if err := ktools.MergeError(i.loadAmount(input), i.loadDate(input), i.loadDue(input),
		i.loadItems(input.GetItems())); err != nil {
		return err
	}
	i.Status.Change(INVOICE_CLASS, INVOICE_GETTING_CLIENT, "", "")
	i.Status.Change(PAYMENT_CLASS, PAYMENT_OPENED, "", "")
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Invoice) GetJson() ([]byte, error) {
	return json.Marshal(i)
}

// IsDuplicated returns true if the invoice is duplicated
func (i *Invoice) IsDuplicated() (bool, error) {
	return i.repo.IsDuplicatedInvoice(i.Reference)
}

// SetAmount sets the amount of invoice
func (i *Invoice) SetAmount(amount float64) {
	i.Amount = amount
}

// Save stores the invoice on the repository
func (i *Invoice) Save() error {
	execOrder := []func() error{
		i.repo.Begin,
		i.saveClients,
		i.saveInvoice,
		i.Status.Save,
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
	return i.Id
}

// GetReference returns the reference of invoice
func (i *Invoice) GetReference() string {
	return i.Reference
}

// GetBusiness returns the business of invoice
func (i *Invoice) GetBusinessId() string {
	return i.Business.GetId()
}

// GetBusiness returns the business client object of invoice
func (i *Invoice) GetBusiness() port.InvoiceClient {
	return i.Business
}

// Getcustomerreturns the customerof invoice
func (i *Invoice) GetCustomerId() string {
	return i.Customer.GetId()
}

// Getcustomerreturns the customerclient object of invoice
func (i *Invoice) GetCustomer() port.InvoiceClient {
	return i.Customer
}

// GetAmount returns the amount of invoice
func (i *Invoice) GetAmount() float64 {
	return i.Amount
}

// GetDate returns the date of invoice
func (i *Invoice) GetDate() time.Time {
	return i.Date
}

// GetDue returns the due of invoice
func (i *Invoice) GetDue() time.Time {
	return i.Due
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
	i.Business = NewInvoiceClient(i.repo)
	createdAfter := time.Now().Add(time.Duration(-client_expiration_seconds) * time.Second)
	if b, err := i.Business.GetLastInvoiceClient(input.GetBusinessNickname(), createdAfter); err != nil {
		return err
	} else if !b {
		i.Business.Load("", input.GetBusinessNickname(), "", "", "", 0, 0, time.Time{})
	}
	i.Customer = NewInvoiceClient(i.repo)
	if b, err := i.Customer.GetLastInvoiceClient(input.GetCustomerNickname(), createdAfter); err != nil {
		return err
	} else if !b {
		i.Customer.Load("", input.GetCustomerNickname(), "", "", "", 0, 0, time.Time{})
	}
	return nil
}

// loadAmount loads the amount from input
func (i *Invoice) loadAmount(input port.CreateInputDto) error {
	var err error
	i.Amount, err = input.GetAmount()
	if err != nil {
		return err
	}
	return nil
}

// loadDate loads the date from input
func (i *Invoice) loadDate(input port.CreateInputDto) error {
	var err error
	i.Date, err = input.GetDate()
	if err != nil {
		return err
	}
	return nil
}

// loadDue loads the due from input
func (i *Invoice) loadDue(input port.CreateInputDto) error {
	var err error
	i.Due, err = input.GetDue()
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
		i.Items = append(i.Items, item)
	}
	return nil
}

// saveInvoice saves the invoice on the repository
func (i *Invoice) saveInvoice() error {
	return i.repo.SaveInvoice(i)
}

// saveClients saves the clients (business and customer) on the repository
func (i *Invoice) saveClients() error {
	if i.Business.IsNew() {
		if err := i.Business.Save(); err != nil {
			return err
		}
	}
	if i.Customer.IsNew() {
		if err := i.Customer.Save(); err != nil {
			return err
		}
	}
	return nil
}

// saveItems saves the items on the repository
func (i *Invoice) saveItems() error {
	for _, item := range i.Items {
		if err := item.Save(); err != nil {
			return err
		}
	}
	return nil
}
