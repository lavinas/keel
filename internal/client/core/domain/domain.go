package domain

import (
	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/client/core/port"
	"strconv"
)

// clientData is the data structure for a client
type clientData struct {
	ID       string
	Name     string
	Nickname string
	Document uint64
	Phone    uint64
	Email    string
}

// Client is the domain model for a client
type Domain struct {
	client clientData
	repo   port.Repo
}

// NewClient creates a new client
func NewDomain(repo port.Repo) *Domain {
	return &Domain{
		client: clientData{},
		repo:   repo,
	}
}

// Create fills the client with the given data
func (c *Domain) ClientInit(name, nick, document, phone, email string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	c.client.ID = id.String()
	c.client.Name = name
	c.client.Nickname = nick
	doc, err := strconv.ParseUint(document, 10, 64)
	if err != nil {
		return "", err
	}
	c.client.Document = doc
	ph, err := strconv.ParseUint(phone, 10, 64)
	if err != nil {
		return "", err
	}
	c.client.Phone = ph
	c.client.Email = email
	return c.client.ID, nil
}

// ClientDocumentDuplicity checks if a document is already registered
func (c *Domain)  ClientDocumentDuplicity() (bool, error) {
	return c.repo.ClientDocumentDuplicity(c.client.Document)
}

// ClientEmailDuplicity checks if an email is already registered
func (c *Domain) ClientEmailDuplicity() (bool, error) {
	return c.repo.ClientEmailDuplicity(c.client.Email)
}

// SaveClient stores the client data
func (c *Domain) ClientSave() error {
	return c.repo.ClientSave(c)
}

// GetClient returns the client data
func (c *Domain) ClientGet() (string, string, string, uint64, uint64, string) {
	return c.client.ID, c.client.Name, c.client.Nickname, c.client.Document, c.client.Phone, c.client.Email
}

// GetClientFormatted returns the client data formatted
func (c *Domain) ClientGetFormatted() (string, string, string, string, string, string) {
	return c.client.ID, c.client.Name, c.client.Nickname, strconv.FormatUint(c.client.Document, 10), 
			strconv.FormatUint(c.client.Phone, 10), c.client.Email
}
