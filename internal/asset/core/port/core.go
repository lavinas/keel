package port

import (
	"github.com/lavinas/keel/pkg/kerror"
)

// CreateDtoIn is the input data for Create usecase method
type CreateDtoIn interface {
	Validate(repo Repository) *kerror.KError
	GetDomain() (Domain, *kerror.KError)
}

// CreateDtoOut is the output data for Create usecase method
type CreateDtoOut interface {
	SetDomain(d Domain) *kerror.KError
}

// Domain is the interface that represents the system generic domain
type Domain interface {
	Validate() *kerror.KError
	SetCreate() *kerror.KError
}

type UseCase interface {
	Create(dtoIn CreateDtoIn, DtoOut CreateDtoOut) *kerror.KError
}
