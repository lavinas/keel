package port

type UseCase interface {
	RegisterClient(dto RegisterClient, result DefaultResult)
}
