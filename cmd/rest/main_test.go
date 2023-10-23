package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		go main()
	})
}
