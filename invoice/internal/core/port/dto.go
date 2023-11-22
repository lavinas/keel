package port

type Register interface {
	Validate() error
	GetDomain(businnes_id string) Domain
}

type RegisterClient interface {
	Validate() error
	Get() (string, string, string, string, string)
}

type RegisterInstruction interface {
	Validate() error
	Get() (string, string)
}

type RegisterProduct interface {
	Validate() error
	Get() (string, string)
}

type DefaultResult interface {
	Set(code int, message string)
	Get() (int, string)
}
