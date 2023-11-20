package port

type RegisterClient interface {
	Validate() error
	Get() (string, string, string, string, string)
}

type UseCase interface {
	RegisterClient(dto RegisterClient) error
}
