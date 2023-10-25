package port

import (
	"time"
)

type Domain interface {
	GetInvoice() Invoice
}

type InvoiceClient interface {
	Load(nickname, clientId, name, email string, phone, document uint64)
	Save() error
	GetId() string
	GetNickname() string
	GetClientId() string
	GetName() string
	GetDocument() uint64
	GetPhone() uint64
	GetEmail() string
}

type Invoice interface {
	Load(input CreateInputDto, business InvoiceClient, customer InvoiceClient) error
	SetAmount(amount float64)
	Save() error
	GetId() string
	GetReference() string
	GetBusinessId() string
	GetCustomerId() string
	GetAmount() float64
	GetDate() time.Time
	GetDue() time.Time
	GetNoteId() string
	GetStatusId() uint
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type InvoiceItem interface {
	Load(dto CreateInputItemDto, invoice Invoice) error
	Save() error
	GetId() string
	GetInvoiceId() string
	GetServiceReference() string
	GetDescription() string
	GetAmount() float64
	GetQuantity() uint64
}
