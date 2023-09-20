package port

import (
	"os"
)

// Log is a port is a interface that wraps the methods to interact with the log
type Log interface {
	GetFile() *os.File
	Info(message string)
	Error(message string)
	Close()
}
