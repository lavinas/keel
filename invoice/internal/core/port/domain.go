package port

import (
	"time"

	"github.com/lavinas/keel/invoice/pkg/kerror"
)

type Domain interface {
	Validate(repo Repository) *kerror.KError
	SetBusinessID(string)
	SetCreatedAt(date time.Time)
	SetUpdatedAt(date time.Time)
	Fit()
	GetID() string
}
