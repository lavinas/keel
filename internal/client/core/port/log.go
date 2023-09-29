package port

import (
	"os"
)

// Log is a port is a interface that wraps the methods to interact with the log
type Log interface {
	GetFile() *os.File
	Info(message string)
	Infof(input any, message string)
	Error(message string)
	Errorf(input any, err error)
	Close()
}
