package port

type Repo interface {
	SaveInvoiceClient(client InvoiceClient) error
	SaveInvoice(invoice Invoice) error
	SaveInvoiceItem(item InvoiceItem) error
}
