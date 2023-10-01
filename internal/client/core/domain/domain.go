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
func (c *Domain) CreateClient(name, nickName, document, phone, email string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	c.client.ID = id.String()
	c.client.Name = name
	c.client.Nickname = nickName
	doc, err := strconv.ParseUint(document, 10, 64)
	if err != nil {
		return err
	}
	c.client.Document = doc
	ph, err := strconv.ParseUint(phone, 10, 64)
	if err != nil {
		return err
	}
	c.client.Phone = ph
	c.client.Email = email
	return nil
}

// GetClient returns the client data
func (c *Domain) GetClient() (string, string, string, uint64, uint64, string) {
	return c.client.ID, c.client.Name, c.client.Nickname, c.client.Document, c.client.Phone, c.client.Email
}
