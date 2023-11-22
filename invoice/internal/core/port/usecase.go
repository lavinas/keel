package port

type UseCase interface {
	Register(dto Register, result DefaultResult)
}
