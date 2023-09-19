package port


import (
	"github.com/lavinas/keel/internal/client/core/domain"
)

type ClientService interface {
	Create(input domain.CreateInputDto) (*domain.CreateOutputDto, error)
}