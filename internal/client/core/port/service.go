package port

import (
	"github.com/lavinas/keel/internal/client/core/domain"
)

type Service interface {
	Create(input domain.CreateInputDto) (*domain.CreateOutputDto, error)
	ListAll() (*domain.ListAllOutputDto, error)
}
