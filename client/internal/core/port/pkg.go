package port

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

type GinEngineWrapper interface {
	Run() *http.Server
	ShutDown()
	MapError(message string)
	POST(relativePath string, handlers ...gin.HandlerFunc)
	PUT(relativePath string, handlers ...gin.HandlerFunc)
	GET(relativePath string, handlers ...gin.HandlerFunc)
	DELETE(relativePath string, handlers ...gin.HandlerFunc)
}
