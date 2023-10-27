package port

type Repo interface {
	Begin() error
	Commit() error
	Rollback() error
	SaveInvoiceClient(client InvoiceClient) error
	SaveInvoice(invoice Invoice) error
	SaveInvoiceItem(item InvoiceItem) error
	Close()
}
