package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	ClientSave(client Client) error
	ClientDocumentDuplicity(document uint64) (bool, error)
	ClientEmailDuplicity(email string) (bool, error)
	ClientLoad(page, perPage uint64, set ClientSet) error
}
