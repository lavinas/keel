package port

import (
	"os"
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
	Add(obj interface{}) error
	IsDuplicatedError(err error) bool
	FindByID(obj interface{}, id string) bool
}
