package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	ClientSave(client Client) error
	ClientDocumentDuplicity(document uint64, id string) (bool, error)
	ClientEmailDuplicity(email, id string) (bool, error)
	ClientNickDuplicity(nick, id string) (bool, error)
	ClientLoadSet(page, perPage uint64, name, nick, doc, email string, set ClientSet) error
	GetById(id string, client Client) (bool, error)
	GetByNick(nick string, client Client) (bool, error)
	GetByEmail(email string, client Client) (bool, error)
	GetByDoc(doc uint64, client Client) (bool, error)
	GetByPhone(phone uint64, client Client) (bool, error)
	Update(client Client) error
	Close() error
}
