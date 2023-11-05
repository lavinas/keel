package restconsumer

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/keel/invoice/pkg/gin_wrapper"
)


// LogMock is a mock of Log interface
type LogMock struct {
	lType string
	message string
}
func (l *LogMock) GetFile() *os.File {
	return nil
}
func (l *LogMock) Info(message string) {
	l.lType = "info"
	l.message = message
}
func (l *LogMock) Infof(input any, message string) {
	l.lType = "infof"
	l.message = message
}
func (l *LogMock) Error(message string) {
	l.lType = "error"
	l.message = message
}
func (l *LogMock) Errorf(input any, err error) {
	l.lType = "errorf"
	l.message = err.Error()
}
func (l *LogMock) Close() {
}

// ServerMock is a mock of Server interface
type ServerMock struct {
	handlerName string
	outCode int
	outBody interface{}
}
func (s *ServerMock) Run() {
	log := LogMock{}
	h := gin_wrapper.NewGinEngineWrapper(&log)
	h.POST(s.handlerName, s.GetOutput)
	h.Run()
}
func (s *ServerMock) GetOutput(c *gin.Context) {
	c.JSON(s.outCode, s.outBody)
}






