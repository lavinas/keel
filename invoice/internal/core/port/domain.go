package port

import (
	"time"
)

type Domain interface {
	Validate(p interface{}) error
	SetBusinessID(string)
	SetCreatedAt(date time.Time)
	SetUpdatedAt(date time.Time)
	Marshal() error
}
