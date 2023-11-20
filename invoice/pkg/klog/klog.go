package klog

import (
	"fmt"
	"os"
	"time"
)

const (
	log_path = "LOG_PATH"
)

// logFile is a struct that wraps the methods to interact with the log file
type Klog struct {
	path      string
	name      string
	component string
	file      *os.File
	date      string
	info      bool
}

// NewlogFile creates a new logFile
func NewKlog(component string, info bool) (*Klog, error) {
	path := os.Getenv(log_path)
	if path == "" {
		path = "."
	}
	l := Klog{path: path, component: component, info: info}
	if err := l.initFile(path, component); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetFile returns the log os file
func (l *Klog) GetFile() *os.File {
	return l.file
}

func (l *Klog) GetName() string {
	return l.name
}

// Info writes a info message to the log file
func (l *Klog) Info(message string) {
	if l.info {
		l.write(message)
	}
}

// Infof writes a info message to the log file attaching a input struct
func (l *Klog) Infof(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	l.Info(message)
}

// Error writes a error message to the log file
func (l *Klog) Error(err error) {
	l.write(err.Error())
}

// Errorf writes a error message to the log file attaching a input struct
func (l *Klog) Errorf(format string, a ...any) {
	l.write(fmt.Sprintf(format, a...))
}

// Close closes the log file
func (l *Klog) Close() {
	l.file.Close()
}

// write writes a message to the log file
func (l *Klog) write(message string) {
	l.shiftFile()
	d := time.Now().Format("2006/01/02 - 15:04:05")
	t := "[LOG] " + d + " | " + message + "\n"
	l.file.Write([]byte(t))
	l.file.Sync()
}

// initFile creates a new log file
func (l *Klog) initFile(path string, component string) error {
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
func (l *Klog) shiftFile() {
	d := time.Now().Format("2006-01-02")
	if d != l.date {
		l.file.Close()
		l.initFile(l.path, l.component)
	}
}
