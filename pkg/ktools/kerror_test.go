package ktools

import (
	"errors"
	"testing"
)

func TestMergeError(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		if err := MergeError(); err != nil {
			t.Errorf("should be nil")
		}
	})
	t.Run("should return nil", func(t *testing.T) {
		err := MergeError(nil, nil, nil, nil)
		if err != nil {
			t.Errorf("should be nil")
		}
	})
	t.Run("should return error", func(t *testing.T) {
		err := MergeError(errors.New("error"))
		if err == nil {
			t.Errorf("should not be nil")
		}
		if err != nil && err.Error() != "error" {
			t.Errorf("should be error")
		}
	})
	t.Run("should return error", func(t *testing.T) {
		err := MergeError(errors.New("error 1"), errors.New("error 2"))
		if err == nil {
			t.Errorf("should not be nil")
		}
		if err != nil && err.Error() != "error 1 | error 2" {
			t.Errorf("should be error 1 | error 2")
		}
	})
}