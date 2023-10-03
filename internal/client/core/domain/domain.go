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

func (c *Domain) GetClient(input port.ClientCreateInputDto) port.Client {
	return NewClient(c.repo)
}
