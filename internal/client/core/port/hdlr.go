package port

import (
	"github.com/lavinas/keel/internal/client/core/domain"
)

// Handler is the interface that wraps the methods to interact with the handler for the client domain
type Handler interface {
	Create(client *domain.Client) error
}