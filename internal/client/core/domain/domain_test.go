package domain

import (
	"testing"
)

func TestDomainGetClient(t *testing.T) {
	t.Run("should return a client", func(t *testing.T) {
		repo := &RepoMock{}
		domain := NewDomain(repo)
		client := domain.GetClient()
		if client == nil {
			t.Errorf("Error: client should not be nil")
		}
	})
}

func TestDomainGetClientSet(t *testing.T) {
	t.Run("should return a client set", func(t *testing.T) {
		repo := &RepoMock{}
		domain := NewDomain(repo)
		clientSet := domain.GetClientSet()
		if clientSet == nil {
			t.Errorf("Error: client set should not be nil")
		}
	})
}