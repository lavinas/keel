package port

type Repo interface {
	Begin() error
	Commit() error
	Rollback() error
	IsDuplicatedInvoice(reference string) (bool, error)
	SaveInvoiceClient(client InvoiceClient) error
	SaveInvoice(invoice Invoice) error
	SaveInvoiceItem(item InvoiceItem) error
	Close() error
}
