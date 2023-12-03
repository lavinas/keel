package port

type UseCase interface {
	Create(domain Domain, result DefaultResult)
}
