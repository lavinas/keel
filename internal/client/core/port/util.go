package port

type Util interface {
	ValidateDocument(document string) bool
	ClearNumber(document string) string
}