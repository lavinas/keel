package port

type UseCase interface {
	RegisterClient(dto RegisterClient) error
}
