package service

import (
	"testing"
)

func TestServiceCreate(t *testing.T) {
	s := NewService(&RepoMock{}, &LogMock{}, &RestConsumerMock{}, &DomainMock{})
	t.Run("should create without errors", func(t *testing.T) {
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		err := s.Create(&i, &o)
		if err != nil {
			t.Errorf("expected no errors, got %v", err)
		}
		if o.status != "created" {
			t.Errorf("expected status to be created, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
}
