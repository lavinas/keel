package port

import (
	"github.com/lavinas/keel/pkg/kerror"
)

type Domain interface {
	SetCreate()
	Validate() *kerror.KError
}

type DefaultResult interface {
	Set(code int, message string)
	Get() (int, string)
}

type UseCase interface {
	Create(domain Domain) *kerror.KError
}
