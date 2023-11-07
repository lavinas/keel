package port

type Config interface {
	Get(param string) string
}