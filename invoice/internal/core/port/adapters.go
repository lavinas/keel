package port

import (
	"os"

	"github.com/lavinas/keel/invoice/internal/core/domain"
)

type Config interface {
	Get(param string) string
}

type Logger interface {
	GetFile() *os.File
	GetName() string
	Info(message string)
	Infof(format string, a ...any)
	Error(err error)
	Fatal(err error)
	Errorf(format string, a ...any)
	Close()
}

type Repository interface {
	AddClient(client *domain.Client) error
	IsDuplicatedError(err error) bool
}
