package port

type Domain interface {
	CreateClient(name, nickName, document, phone, email string) error
	GetClient() (string, string, string, uint64, uint64, string)
}
