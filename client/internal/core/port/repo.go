package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	Save(client Client) error
	DocumentDuplicity(document uint64, id string) (bool, error)
	EmailDuplicity(email, id string) (bool, error)
	NickDuplicity(nick, id string) (bool, error)
	LoadSet(page, perPage uint64, name, nick, doc, phone, email string, set ClientSet) error
	GetById(id string, client Client) (bool, error)
	GetByNick(nick string, client Client) (bool, error)
	GetByEmail(email string, client Client) (bool, error)
	GetByDoc(doc uint64, client Client) (bool, error)
	GetByPhone(phone uint64, client Client) (bool, error)
	Update(client Client) error
	Close() error
}
