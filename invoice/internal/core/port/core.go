package port

type RegisterInvoiceClient interface {
	Validate() error
	Get() (string, string, string, string, string)
}

type UseCase interface {
	RegisterClient(dto RegisterInvoiceClient) error
}