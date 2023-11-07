package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

// InvoiceClient is the domain entity for client data in invoice
type InvoiceClient struct {
	repo       port.Repo
	id         string
	nickname   string
	clientId   string
	name       string
	document   uint64
	phone      uint64
	email      string
	created_at time.Time
	isNew      bool
}

// NewInvoiceClient creates a new invoice client
func NewInvoiceClient(repo port.Repo) *InvoiceClient {
	return &InvoiceClient{
		repo: repo,
	}
}

// Load loads a invoice client from input
func (i *InvoiceClient) Load(id, nickname, clientId, name, email string, phone, 
								document uint64, created_at time.Time) {
	if id == "" {
		i.id = uuid.New().String()
		i.isNew = true
	} else {
		i.id = id
		i.isNew = false
	}
	i.id = uuid.New().String()
	i.nickname, i.clientId, i.name, i.email, i.phone, i.document = nickname, clientId, 
				name, email, phone, document
	if created_at.IsZero() {
		i.created_at = time.Now()
	} else {
		i.created_at = created_at
	}
}

// GetLastInvoiceClient returns the last invoice client from repository
func (i *InvoiceClient) GetLastInvoiceClient(nickname string, 
												created_after time.Time) (bool, error) {
	return i.repo.GetLastInvoiceClient(nickname, created_after, i)
}

func (i *InvoiceClient) LoadGetClientNicknameDto(input port.GetClientByNicknameInputDto) error {
	i.nickname, i.clientId, i.name, i.email = input.GetNickname(), input.GetId(), 
			input.GetName(), input.GetEmail()
	phone, err := input.GetPhone()
	if err != nil {
		return err
	}
	doc, err := input.GetDocument()
	if err != nil {
		return err
	}
	i.phone = phone
	i.document = doc
	i.created_at = time.Now()
	return nil
}

// Save stores the invoice client on the repository
func (i *InvoiceClient) Save() error {
	return i.repo.SaveInvoiceClient(i)
}

// Update updates the invoice client on the repository
func (i *InvoiceClient) Update() error {
	return i.repo.UpdateInvoiceClient(i)
}

// GetId returns the id of invoice client
func (i *InvoiceClient) GetId() string {
	return i.id
}

// GetNickname returns the nickname of invoice client
func (i *InvoiceClient) GetNickname() string {
	return i.nickname
}

// GetClientId returns the client id of invoice client
func (i *InvoiceClient) GetClientId() string {
	return i.clientId
}

// GetName returns the name of invoice client
func (i *InvoiceClient) GetName() string {
	return i.name
}

// GetDocument returns the document of invoice client
func (i *InvoiceClient) GetDocument() uint64 {
	return i.document
}

// GetPhone returns the phone of invoice client
func (i *InvoiceClient) GetPhone() uint64 {
	return i.phone
}

// GetEmail returns the email of invoice client
func (i *InvoiceClient) GetEmail() string {
	return i.email
}

// GetCreatedAt returns the created_at of invoice client
func (i *InvoiceClient) GetCreatedAt() time.Time {
	return i.created_at
}

func (i *InvoiceClient) IsNew() bool {
	return i.isNew
}