package port

type Domain interface {
	ClientInit(name, nickName, document, phone, email string) (string, error)
	ClientGet() (string, string, string, uint64, uint64, string)
	ClientSave() error
}
