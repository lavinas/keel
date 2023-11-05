package restconsumer

import (
	"os"

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






