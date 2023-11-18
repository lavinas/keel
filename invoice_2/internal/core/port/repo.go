package port

import (
	"time"
)

type Repo interface {
	Begin() error
	Commit() error
	Rollback() error
	IsDuplicatedInvoice(reference string) (bool, error)
	SaveInvoiceClient(client InvoiceClient) error
	GetLastInvoiceClient(nickname string, created_after time.Time, client InvoiceClient) (bool, error)
	UpdateInvoiceClient(client InvoiceClient) error
	SaveInvoice(invoice Invoice) error
	SaveInvoiceItem(item InvoiceItem) error
	GetInvoiceVertex(graph InvoiceStatus) error
	GetInvoiceEdge(graph InvoiceStatus) error
	StoreInvoiceStatus(class string, graph InvoiceStatus) error
	CreateInvoiceStatusLog(class string, graph InvoiceStatus) error
	Close() error
}
