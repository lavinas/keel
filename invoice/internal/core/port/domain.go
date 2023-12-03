package port

import (
	"github.com/lavinas/keel/invoice/pkg/kerror"
)

type Domain interface {
	SetCreate(businessID string)
	Validate(repo Repository) *kerror.KError
	Fit()
}
