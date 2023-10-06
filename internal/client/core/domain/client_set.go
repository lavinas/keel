package domain

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientSet is the domain model for a client set
type ClientSet struct {
	repo    port.Repo
	set     []Client
	page    uint64
	perPage uint64
}

// NewClientSet creates a new client set
func NewClientSet(repo port.Repo) *ClientSet {
	return &ClientSet{
		repo: repo,
		set:  []Client{},
	}
}

// Load loads the client set from the repository
func (c *ClientSet) Load(page, perPage uint64) error {
	c.page, c.perPage = page, perPage
	return c.repo.ClientLoad(page, perPage, c)
}

// Append appends a new client to the set
func (c *ClientSet) Append(id, name, nick string, doc, phone uint64, email string) {
	client := NewClient(c.repo)
	client.Load(id, name, nick, doc, phone, email)
	c.set = append(c.set, *client)
}

// SetOutput sets the output for the client set
func (c *ClientSet) SetOutput(output port.ClientListOutputDto) {
	output.SetPage(c.page, c.perPage)
	for _, client := range c.set {
		id, name, nick, doc, phone, email := client.GetFormatted()
		output.Append(id, name, nick, doc, phone, email)
	}
}
