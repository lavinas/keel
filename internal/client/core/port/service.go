package port

import (
	"github.com/lavinas/keel/internal/client/core/dto"
)

type Service interface {
	Create(input dto.CreateInputDto) (*dto.CreateOutputDto, error)
	ListAll() (*dto.ListAllOutputDto, error)
}
