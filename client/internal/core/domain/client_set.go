package domain

import (
	"github.com/lavinas/keel/client/internal/core/port"
)

// ClientSet is the domain model for a client set
type ClientSet struct {
	repo    port.Repo
	set     []Client
	page    uint64
	perPage uint64
	name    string
	nick    string
	doc     string
	phone   string
	email   string
}

// NewClientSet creates a new client set
func NewClientSet(repo port.Repo) *ClientSet {
	return &ClientSet{
		repo: repo,
		set:  []Client{},
	}
}

// Load loads the client set from the repository
func (c *ClientSet) Load(page, perPage uint64, name, nick, doc, phone, email string) error {
	c.page, c.perPage, c.name, c.nick, c.doc, c.phone, c.email = page, perPage, name, nick, doc, phone, email
	return c.repo.LoadSet(page, perPage, name, nick, doc, phone, email, c)
}

// Append appends a new client to the set
func (c *ClientSet) Append(id, name, nick string, doc, phone uint64, email string) {
	client := NewClient(c.repo)
	client.Load(id, name, nick, doc, phone, email)
	c.set = append(c.set, *client)
}

// SetOutput sets the output for the client set
func (c *ClientSet) SetOutput(output port.FindOutputDto) {
	output.SetPage(c.page, c.perPage)
	for _, client := range c.set {
		id, name, nick, doc, phone, email := client.GetFormatted()
		output.Append(id, name, nick, doc, phone, email)
	}
}

// Count returns the number of clients in the set
func (c *ClientSet) Count() int {
	return len(c.set)
}
