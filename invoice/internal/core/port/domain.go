package port

type Domain interface {
	Validate() error
	SetBusinessID(string)
}
