package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	ClientSave(client Client) error
	ClientDocumentDuplicity(document uint64) (bool, error)
	ClientEmailDuplicity(email string) (bool, error)
	ClientLoadSet(page, perPage uint64, name, nick, doc, email string, set ClientSet) error
	ClientLoadById(id string, client Client) error
	ClientUpdate(client Client) error
}
