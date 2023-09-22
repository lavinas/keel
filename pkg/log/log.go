package log

import (
	"os"
	"time"
	"encoding/json"
)

type logFile struct {
	path      string
	component string
	file      *os.File
	date      string
	info      bool
}

func NewlogFile(path string, component string, info bool) *logFile {
	f, d := initFile(path, component)
	return &logFile{path: path, file: f, component: component, date: d, info: info}
}

func (l *logFile) GetFile() *os.File {
	return l.file
}

func (l *logFile) Info(message string) {
	if l.info {
		write(l, message)
	}
}

func (l *logFile) Infof(input any, message string) {
	b, _ := json.Marshal(input)
	l.Info(message + " | " + string(b))
}

func (l *logFile) Error(message string) {
	write(l, message)
}

func (l *logFile) Errorf (input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}

func (l *logFile) Close() {
	l.file.Close()
}

func write(l *logFile, message string) {
	shiftFile(l)
	d := time.Now().Format("2006/01/02 - 15:04:05")
	t := "[LOG] " + d + " | " + message + "\n"
	l.file.Write([]byte(t))
	l.file.Sync()
}

func initFile(path string, component string) (*os.File, string) {
	d := time.Now().Format("2006-01-02")
	fp := path + "/" + d + "-" + component + ".log"
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic("creating log file error - " + err.Error())
	}
	return f, d
}

func shiftFile(l *logFile) {
	d := time.Now().Format("2006-01-02")
	if d != l.date {
		l.file.Close()
		l.file, l.date = initFile(l.path, l.component)
	}
}
