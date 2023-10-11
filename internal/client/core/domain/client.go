package domain

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Client is the domain model for a client
type Client struct {
	repo     port.Repo
	ID       string
	Name     string
	Nickname string
	Document uint64
	Phone    uint64
	Email    string
}

// NewClient creates a new client
func NewClient(repo port.Repo) *Client {
	return &Client{
		repo:     repo,
		ID:       "",
		Name:     "",
		Nickname: "",
		Document: 0,
		Phone:    0,
		Email:    "",
	}
}

// Insert loads a client
func (c *Client) Insert(name, nick string, doc, phone uint64, email string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	c.ID, c.Name, c.Nickname, c.Document, c.Phone, c.Email = id.String(), name, nick, doc, phone, email
	return nil
}

// Load loads a client values
func (c *Client) Load(id, name, nick string, doc, phone uint64, email string) {
	c.ID, c.Name, c.Nickname, c.Document, c.Phone, c.Email = id, name, nick, doc, phone, email
}

// LoadById loads a client by id from the repository
func (c *Client) LoadById(id string) (bool, error) {
	return c.repo.GetById(id, c)
}

// LoadByNick loads a client by nick from the repository
func (c *Client) LoadByNick(nick string) (bool, error) {
	return c.repo.GetByNick(nick, c)
}

// LoadByEmail loads a client by email from the repository
func (c *Client) LoadByEmail(email string) (bool, error) {
	return c.repo.GetByEmail(email, c)
}

// LoadByDoc loads a client by doc from the repository
func (c *Client) LoadByDoc(doc uint64) (bool, error) {
	return c.repo.GetByDoc(doc, c)
}

// LoadByDoc loads a client by doc from the repository
func (c *Client) LoadByPhone(phone uint64) (bool, error) {
	return c.repo.GetByDoc(phone, c)
}

// DocumentDuplicity checks if a document is already registered
func (c *Client) DocumentDuplicity() (bool, error) {
	return c.repo.DocumentDuplicity(c.Document, c.ID)
}

// EmailDuplicity checks if a email is already registered
func (c *Client) EmailDuplicity() (bool, error) {
	return c.repo.EmailDuplicity(c.Email, c.ID)
}

// NickDuplicity checks if a nick is already registered
func (c *Client) NickDuplicity() (bool, error) {
	return c.repo.NickDuplicity(c.Nickname, c.ID)
}

// Get returns the client values
func (c *Client) Get() (string, string, string, uint64, uint64, string) {
	return c.ID, c.Name, c.Nickname, c.Document, c.Phone, c.Email
}

// GetFormatted returns the client values formatted to string
func (c *Client) GetFormatted() (string, string, string, string, string, string) {
	doc := fmt.Sprintf("%d", c.Document)
	if len(doc) <= 11 {
		doc = fmt.Sprintf("%011s", doc)
	} else {
		doc = fmt.Sprintf("%014s", doc)
	}
	return c.ID, c.Name, c.Nickname, doc, strconv.FormatUint(c.Phone, 10), c.Email
}

// Save saves a client on the repository
func (c *Client) Save() error {
	return c.repo.Save(c)
}

// Update updates a client on the repository
func (c *Client) Update() error {
	return c.repo.Update(c)
}
