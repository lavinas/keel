package domain

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// ClientSet is the domain model for a client set
type ClientSet struct {
	repo port.Repo
	set  []Client
}

// NewClientSet creates a new client set
func NewClientSet(repo port.Repo) *ClientSet {
	return &ClientSet{
		repo: repo,
		set:  []Client{},
	}
}

// GetJson returns the json representation of the client set
func (c *ClientSet) Load() error {
	return c.repo.ClientLoad(c)
}

func (c *ClientSet) Append(id, name, nick string, doc, phone uint64, email string) {
	client := NewClient(c.repo)
	client.Create(name, nick, doc, phone, email)
	c.set = append(c.set, *client)
}

func (c *ClientSet) SetOutput(output port.ClientListOutputDto) {
	for _, client := range c.set {
		id, name, nick, doc, phone, email := client.GetFormatted()
		output.Append(id, name, nick, doc, phone, email)
	}
}
