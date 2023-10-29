package main

import (
	"os"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		os.Setenv("KEEL_LOG_PATH", ".")
		go main()
	})
}

func TestRunPanics(t *testing.T) {
	t.Run("should panic", func(t *testing.T) {
		time.Sleep(1 * time.Second)
		path := os.Getenv("KEEL_LOG_PATH")
		os.Setenv("KEEL_LOG_PATH", "./nada")
		defer os.Setenv("KEEL_LOG_PATH", path)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		main()
	})
}
