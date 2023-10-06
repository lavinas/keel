package port

// Domain is the interface that wraps the methods to interact with the database for the client domain
type Domain interface {
	GetClient(input ClientCreateInputDto) Client
	GetClientSet() ClientSet
}

// Client is the interface that wraps the methods to interact with the database for the client domain
type Client interface {
	Create(name, nick string, doc, phone uint64, email string) error
	Load(id, name, nick string, doc, phone uint64, email string)
	DocumentDuplicity() (bool, error)
	EmailDuplicity() (bool, error)
	Get() (string, string, string, uint64, uint64, string)
	GetFormatted() (string, string, string, string, string, string)
	Save() error
}

// ClientSet is the interface that wraps the methods to interact with the database for the client domain
type ClientSet interface {
	Load(page, perPage uint64) error
	Append(id, name, nick string, doc, phone uint64, email string)
	SetOutput(output ClientListOutputDto)
}
