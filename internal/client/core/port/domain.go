package port

type Domain interface {
}

type Client interface {
	DocumentDuplicity() (bool, error)
	EmailDuplicity() (bool, error)
	Get() (string, string, string, uint64, uint64, string)
	GetFormatted() (string, string, string, string, string, string)
	Save() error
}