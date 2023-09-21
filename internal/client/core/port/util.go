package port

type Util interface {
	ValidateName(name string) (bool, string)
	ValidateNickname (nickname string) (bool, string)
	ValidateDocument(document string) (bool, string)
	ValidateEmail(email string) (bool, string)
	ValidatePhone(phone string) (bool, string)
	ValidateAll(name, nickname, document, phone, email string) (bool, string)
	ClearName(string) (string, error)
	ClearNickname(string) (string, error)
	ClearDocument(string) (uint64, error)
	ClearPhone(string) (uint64, error)
	ClearEmail(email string) (string, error)
	ClearAll(name, nickname, document, phone, email string) (string, string, uint64, uint64, string, error)
}
