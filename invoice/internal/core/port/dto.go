package port

type RegisterClient interface {
	Validate() error
	Get() (string, string, string, string, string)
}

type DefaultResult interface {
	Set(code int, message string)
	Get() (int, string)
}
