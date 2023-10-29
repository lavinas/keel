package port

import (
	"time"
)

type CreateInputItemDto interface {
	Validate() error
	GetReference() string
	GetDescription() string
	GetQuantity() (uint64, error)
	GetPrice() (float64, error)
}

type CreateInputDto interface {
	Validate() error
	GetReference() string
	GetBusinessNickname() string
	GetCustomerNickname() string
	GetAmount() (float64, error)
	GetDate() (time.Time, error)
	GetDue() (time.Time, error)
	GetNoteReference() string
	GetItems() []CreateInputItemDto
}

type CreateOutputDto interface {
	Load(status string, reference string)
}
