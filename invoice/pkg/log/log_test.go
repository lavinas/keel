package log

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewlogFileOk(t *testing.T) {
	t.Run("should create a new log file", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		if l == nil {
			t.Errorf("Log file should not be nil")
		}
		if _, err := os.Stat(l.GetName()); os.IsNotExist(err) {
			t.Errorf("Log file should exist in %s", l.GetName())
		}
		if l.GetFile() == nil {
			t.Errorf("Log file should not be nil")
		}
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
	t.Run("should create a new log file", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, "")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		if l == nil {
			t.Errorf("Log file should not be nil")
		}
		if _, err := os.Stat(l.GetName()); os.IsNotExist(err) {
			t.Errorf("Log file should exist in %s", l.GetName())
		}
		if l.GetFile() == nil {
			t.Errorf("Log file should not be nil")
		}
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
	// should return error
	t.Run("should return error", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, "./noexists")
		_, err := NewlogFile("component", true)
		if err == nil {
			t.Errorf("Log file should return error")
		}
		os.Setenv(log_path, lp)
	})
}

func TestInfo(t *testing.T) {
	t.Run("should write a info message to the log file", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		l.Info("message")
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
}

func TestError(t *testing.T) {
	t.Run("should write a error message to the log file", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		l.Error("message")
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
}

func TestInfof(t *testing.T) {
	t.Run("should write a info message to the log file attaching a input struct", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		l.Infof("input", "message")
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
}

func TestErrorf(t *testing.T) {
	t.Run("should write a error message to the log file attaching a input struct", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		err = errors.New("error")
		l.Errorf("input", err)
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
}

func TestShiftFile(t *testing.T) {
	t.Run("should shift log file", func(t *testing.T) {
		lp := os.Getenv(log_path)
		os.Setenv(log_path, ".")
		l, err := NewlogFile("component", true)
		if err != nil {
			t.Errorf("Log file should be created")
		}
		l.date = "2010-10-10"
		l.shiftFile()
		name := l.GetName()
		today := time.Now().Format("2006-01-02")
		if !strings.Contains(name, today) {
			t.Errorf("Log file should be shifted")
		}
		l.Close()
		os.Remove(l.GetName())
		os.Setenv(log_path, lp)
	})
}
