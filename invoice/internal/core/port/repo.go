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
	GetLastInvoiceClientId(nickname string, created_after time.Time) (string, error)
	UpdateInvoiceClient(client InvoiceClient) error
	SaveInvoice(invoice Invoice) error
	SaveInvoiceItem(item InvoiceItem) error
	Close() error
}
