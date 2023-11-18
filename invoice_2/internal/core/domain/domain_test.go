package domain_test

import (
	"testing"

	"github.com/lavinas/keel/invoice/internal/core/domain"
)

func TestGetInvoice(t *testing.T) {
	t.Run("should return invoice", func(t *testing.T) {
		repo := new(RepoMock)
		domain := domain.NewDomain(repo)
		invoice := domain.GetInvoice()
		if invoice == nil {
			t.Errorf("invoice is nil")
		}
	})
}
