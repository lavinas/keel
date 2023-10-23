package gin_wrapper

import (
	"encoding/json"
	"os"
)

// Log Mock
type LogMock struct {
	msgs []string
	mtype []string
}

func (l *LogMock) GetFile() *os.File {
	return nil
}
func (l *LogMock) Info(msg string) {
	l.msgs = append(l.msgs, msg)
	l.mtype = append(l.mtype, "Info")
}
func (l *LogMock) Infof(input any, message string) {
	b, _ := json.Marshal(input)
	l.Info(message + " | " + string(b))
}
func (l *LogMock) Error(msg string) {
	l.mtype = append(l.mtype, "Error")
	l.msgs = append(l.msgs, msg)
}
func (l *LogMock) Errorf(input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}
func (l *LogMock) Close() {
}
