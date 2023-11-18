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
	Format() error
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

type GetClientByNicknameInputDto interface {
	GetId() string
	GetName() string
	GetNickname() string
	GetDocument() (uint64, error)
	GetPhone() (uint64, error)
	GetEmail() string
}