package port

// Repository is an interface that represents the system generic repository
type Repository interface {
	Exists(obj interface{}, id string) (bool, error)
}
