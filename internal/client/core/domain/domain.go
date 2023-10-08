package domain

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// Domain is the domain model for a client services
type Domain struct {
	repo port.Repo
}

// NewClient creates a new client
func NewDomain(repo port.Repo) *Domain {
	return &Domain{
		repo: repo,
	}
}

// GetClient returns a new client domain model when a input dto is provided
func (c *Domain) GetClient() port.Client {
	return NewClient(c.repo)
}

// GetClientSet returns a new client set domain model stored in the repo
func (c *Domain) GetClientSet() port.ClientSet {
	return NewClientSet(c.repo)
}
