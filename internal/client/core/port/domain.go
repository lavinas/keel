package port

type Domain interface {
	GetClient(input ClientCreateInputDto) Client
}

type Client interface {
	LoadInput(input ClientCreateInputDto) error
	DocumentDuplicity() (bool, error)
	EmailDuplicity() (bool, error)
	Get() (string, string, string, uint64, uint64, string)
	GetFormatted() (string, string, string, string, string, string)
	Save() error
}
