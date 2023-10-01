package port

// Repo is the interface that wraps the methods to interact with the database for the client domain
type Repo interface {
	CreateClient(domain Domain) error
}
