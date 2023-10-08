package port

// Domain is the interface that wraps the methods to interact with the database for the client domain
type Domain interface {
	GetClient() Client
	GetClientSet() ClientSet
}

// Client is the interface that wraps the methods to interact with the database for the client domain
type Client interface {
	Create(name, nick string, doc, phone uint64, email string) error
	Load(id, name, nick string, doc, phone uint64, email string)
	LoadById(id string) (bool, error)
	LoadByNick(nick string) (bool, error)
	LoadByEmail(email string) (bool, error)
	LoadByDoc(doc uint64) (bool, error)
	LoadByPhone(phone uint64) (bool, error)
	DocumentDuplicity() (bool, error)
	EmailDuplicity() (bool, error)
	NickDuplicity() (bool, error)
	Get() (string, string, string, uint64, uint64, string)
	GetFormatted() (string, string, string, string, string, string)
	Save() error
	Update() error
}

// ClientSet is the interface that wraps the methods to interact with the database for the client domain
type ClientSet interface {
	Load(page, perPage uint64, name, nick, doc, email string) error
	Append(id, name, nick string, doc, phone uint64, email string)
	SetOutput(output ClientListOutputDto)
}
