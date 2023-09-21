package port

type Util interface {
	ValidateDocument(document string) bool
	ValidateEmail(email string) bool
	ClearNumber(document string) string
	ClearEmail(email string) string
}
