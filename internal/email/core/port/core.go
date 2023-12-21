package port

import (
	"github.com/lavinas/keel/pkg/kerror"
)

type Domain interface {
	SetID(id string)
	GetID() string
	SetRepository(repo Repository)
	GetRepository() Repository
	SetCreate()
	Validate() *kerror.KError
	GetByID() *kerror.KError
	GetResult() any
}

type DefaultResult interface {
	Set(code int, message string)
	Get() (int, string)
}

type UseCase interface {
	Create(domain Domain) *kerror.KError
}
