package port

import (
	"time"

)

type Domain interface {
	Validate(repo Repository) error
	SetBusinessID(string)
	SetCreatedAt(date time.Time)
	SetUpdatedAt(date time.Time)
	Marshal() error
	GetID() string
}
