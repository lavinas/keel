package port

type Domain interface {
	ClientInit(name, nickName, document, phone, email string) (string, error)
	ClientDocumentDuplicity() (bool, error)
	ClientEmailDuplicity() (bool, error)
	ClientSave() error
	ClientGet() (string, string, string, uint64, uint64, string)
	ClientGetFormatted() (string, string, string, string, string, string)
}
