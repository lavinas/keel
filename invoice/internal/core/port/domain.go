package port

import (
	"time"
)

type Domain interface {
	GetInvoice() Invoice
}

type InvoiceClient interface {
	Load(id, nickname, clientId, name, email string, phone, document uint64, created_at time.Time)
	LoadGetClientNicknameDto(input GetClientByNicknameInputDto) error
	GetLastInvoiceClient(nickname string, created_after time.Time) (bool, error)
	Save() error
	Update() error
	GetId() string
	GetNickname() string
	GetClientId() string
	GetName() string
	GetDocument() uint64
	GetPhone() uint64
	GetEmail() string
	GetCreatedAt() time.Time
	IsNew() bool
}

type Invoice interface {
	Load(input CreateInputDto) error
	SetAmount(amount float64)
	IsDuplicated() (bool, error)
	Save() error
	GetId() string
	GetReference() string
	GetBusinessId() string
	GetBusiness() InvoiceClient
	GetCustomer() InvoiceClient
	GetCustomerId() string
	GetAmount() float64
	GetDate() time.Time
	GetDue() time.Time
	GetNoteId() *string
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

type InvoiceStatusGraph interface {
	LoadRepository() error
	ChangeStatus(class string, status string, description string, author string) error
	SaveLog() error
	AddVertex(class, id, name, description string)
	AddEdge(class, vertexFrom, vertexTo, description string)
	CheckEdge(class, vertexFrom, vertexTo string) bool
	EnqueueEdge (class, vertexFrom, vertexTo, description, author string)
	DequeueEdge (class string) (bool, string, string, string, string, time.Time)
}