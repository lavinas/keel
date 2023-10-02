package port

type Domain interface {
	GetClient(input CreateInputDto) (Client, error)
}

type Client interface {
	DocumentDuplicity() (bool, error)
	EmailDuplicity() (bool, error)
	Get() (string, string, string, uint64, uint64, string)
	GetFormatted() (string, string, string, string, string, string)
	Save() error
}
