package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/ktools"
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
	items     []port.InvoiceItem
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
func (i *Invoice) Load(input port.CreateInputDto) error {
	i.id = uuid.New().String()
	i.reference = input.GetReference()
	i.loadClients(input)
	if err := ktools.MergeError(i.loadAmount(input), i.loadDate(input), i.loadDue(input),
		i.loadItems(input.GetItems())); err != nil {
		return err
	}
	i.status_id = status_map["New"]
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
	if err := i.repo.Begin(); err != nil {
		return err
	}
	defer i.repo.Rollback()
	if err := i.saveClients(); err != nil {
		return err
	}
	if err := i.repo.SaveInvoice(i); err != nil {
		return err
	}
	if err := i.saveItems(); err != nil {
		return err
	}
	if err := i.repo.Commit(); err != nil {
		return err
	}
	return nil
}

// UpdateBusiness updates the business invoice from dto
func (i *Invoice) LoadBusiness(dto port.GetClientByNicknameInputDto) error {
	if i.business == nil {
		i.business = NewInvoiceClient(i.repo)
	}
	return i.business.LoadGetClientNicknameDto(dto)
}

// Updatecustomerupdates the customerinvoice from dto
func (i *Invoice) LoadCustomer(dto port.GetClientByNicknameInputDto) error {
	if i.customer == nil {
		i.customer = NewInvoiceClient(i.repo)
	}
	return i.customer.LoadGetClientNicknameDto(dto)
}

func (i *Invoice) Update() error {
	if err := i.repo.Begin(); err != nil {
		return err
	}
	defer i.repo.Rollback()
	if i.customer == nil {
		i.customer = NewInvoiceClient(i.repo)
	}
	if err := i.customer.Update(); err != nil {
		return err
	}
	if i.business == nil {
		i.business = NewInvoiceClient(i.repo)
	}
	if err := i.business.Update(); err != nil {
		return err
	}
	if err := i.repo.Commit(); err != nil {
		return err
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

// loadClients loads the clients (business and customer) from input
func (i *Invoice) loadClients(input port.CreateInputDto) {
	i.business = NewInvoiceClient(i.repo)
	i.business.Load("",input.GetBusinessNickname(), "", "", "", 0, 0, time.Time{})
	i.customer = NewInvoiceClient(i.repo)
	i.customer.Load("",input.GetCustomerNickname(), "", "", "", 0, 0, time.Time{})
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

// saveClients saves the clients (business and customer) on the repository
func (i *Invoice) saveClients() error {
	if err := i.business.Save(); err != nil {
		return err
	}
	if err := i.customer.Save(); err != nil {
		return err
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
