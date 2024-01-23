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

// Repository is an interface that represents the system generic repository
type Repository interface {
	Exists(obj interface{}, id string) (bool, error)
	GetByID(obj interface{}, id string) (bool, error)
	Add(obj interface{}) error
	Close()
}
