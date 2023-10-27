package domain

import (
	"testing"
)

func TestGetInvoice(t *testing.T) {
	repo := new(RepoMock)
	domain := NewDomain(repo)
	invoice := domain.GetInvoice()
	if invoice == nil {
		t.Errorf("invoice is nil")
	}
}