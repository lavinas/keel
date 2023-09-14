package port

import (
	"github.com/lavinas/keel/internal/client/core/domain"
)

type Repo interface {
	Create(client *domain.Client) error
}