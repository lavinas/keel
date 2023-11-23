package port

type UseCase interface {
	Register(domain Domain, result DefaultResult)
}
