package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	ClientDocumentDuplicity(document uint64) (bool, error)
	ClientEmailDuplicity(email string) (bool, error)
	ClientSave(client Client) error
}
