package port

import (
	"github.com/lavinas/keel/internal/client/core/domain"
)

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	Create(client *domain.Client) error
}

