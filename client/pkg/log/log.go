package log

import (
	"encoding/json"
	"os"
	"time"
)

const (
	log_path = "KEEL_CLIENT_LOG_PATH"
)

// logFile is a struct that wraps the methods to interact with the log file
type LogFile struct {
	path      string
	name      string
	component string
	file      *os.File
	date      string
	info      bool
}

// NewlogFile creates a new logFile
func NewlogFile(component string, info bool) (*LogFile, error) {
	path := os.Getenv(log_path)
	if path == "" {
		path = "."
	}
	l := LogFile{path: path, component: component, info: info}
	if err := l.initFile(path, component); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetFile returns the log os file
func (l *LogFile) GetFile() *os.File {
	return l.file
}

func (l *LogFile) GetName() string {
	return l.name
}

// Info writes a info message to the log file
func (l *LogFile) Info(message string) {
	if l.info {
		l.write(message)
	}
}

// Infof writes a info message to the log file attaching a input struct
func (l *LogFile) Infof(input any, message string) {
	b, _ := json.Marshal(input)
	l.Info(message + " | " + string(b))
}

// Error writes a error message to the log file
func (l *LogFile) Error(message string) {
	l.write(message)
}

// Errorf writes a error message to the log file attaching a input struct
func (l *LogFile) Errorf(input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}

// Close closes the log file
func (l *LogFile) Close() {
	l.file.Close()
}

// write writes a message to the log file
func (l *LogFile) write(message string) {
	l.shiftFile()
	d := time.Now().Format("2006/01/02 - 15:04:05")
	t := "[LOG] " + d + " | " + message + "\n"
	l.file.Write([]byte(t))
	l.file.Sync()
}

// initFile creates a new log file
func (l *LogFile) initFile(path string, component string) error {
	l.date = time.Now().Format("2006-01-02")
	l.name = path + "/" + l.date + "-" + component + ".log"
	var err error
	l.file, err = os.OpenFile(l.name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	return nil
}

// shiftFile checks if the log file needs to be shifted and shifts it if necessary
func (l *LogFile) shiftFile() {
	d := time.Now().Format("2006-01-02")
	if d != l.date {
		l.file.Close()
		l.initFile(l.path, l.component)
	}
}
