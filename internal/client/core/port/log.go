package port

import (
	"os"
)

type Log interface {
	GetFile() *os.File
	Info(message string)
	Error(message string)
	Close()
}