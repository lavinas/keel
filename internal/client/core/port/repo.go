package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	ClientSave(client Client) error
	ClientDocumentDuplicity(document uint64, id string) (bool, error)
	ClientEmailDuplicity(email, id string) (bool, error)
	ClientNickDuplicity(nick, id string) (bool, error)
	ClientLoadSet(page, perPage uint64, name, nick, doc, email string, set ClientSet) error
	ClientGetById(id string, client Client) error
	ClientGetByNick(nick string, client Client) error
	ClientGetByEmail(email string, client Client) error
	ClientGetByDoc(doc uint64, client Client) error
	ClientGetByPhone(phone uint64, client Client) error
	ClientUpdate(client Client) error
	Close() error
}
