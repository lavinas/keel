package dto

import (
	"testing"
)

func TestOutputLoad(t *testing.T) {
	t.Run("should load the fields", func(t *testing.T) {
		dto := CreateOutputDto{}
		dto.Load("status", "reference")
		if dto.Status != "status" {
			t.Errorf("Expected status 'status', got %s", dto.Status)
		}
		if dto.Reference != "reference" {
			t.Errorf("Expected reference 'reference', got %s", dto.Reference)
		}
	})
}
