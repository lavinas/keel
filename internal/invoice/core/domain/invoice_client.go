package domain

import (
	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/invoice/core/port"
)

// InvoiceClient is the domain entity for client data in invoice
type InvoiceClient struct {
	repo     port.Repo
	id       string
	nickname string
	clientId string
	name     string
	document uint64
	phone    uint64
	email    string
}

// NewInvoiceClient creates a new invoice client
func NewInvoiceClient(repo port.Repo) *InvoiceClient {
	return &InvoiceClient{
		repo: repo,
	}
}

// Load loads a invoice client from input
func (i *InvoiceClient) Load(nickname, clientId, name, email string, phone, document uint64) {
	i.id = uuid.New().String()
	i.nickname = nickname
	i.clientId = clientId
	i.name = name
	i.email = email
	i.phone = phone
	i.document = document
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

// Save stores the invoice client on the repository
func (i *InvoiceClient) Save() error {
	return nil
}
