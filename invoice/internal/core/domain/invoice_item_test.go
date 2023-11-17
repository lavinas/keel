package domain_test

import (
	"testing"

	"github.com/lavinas/keel/invoice/internal/core/domain"
)

func TestInvoiceItemLoad(t *testing.T) {
	t.Run("should load invoice item without eerror", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		item := domain.NewInvoiceItem(repo)
		dto := CreateInputItemDtoMock{}
		if err := item.Load(&dto, invoice); err != nil {
			t.Errorf("error on load invoice item")
		}
	})
	t.Run("should return quantity error", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		item := domain.NewInvoiceItem(repo)
		dto := CreateInputItemDtoMock{}
		dto.Status = "quantityError"
		err := item.Load(&dto, invoice)
		if err == nil {
			t.Errorf("should not be nil")
		}
		if err != nil && err.Error() != "quantity error" {
			t.Errorf("should be quantity error")
		}
	})
	t.Run("should return price error", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		item := domain.NewInvoiceItem(repo)
		dto := CreateInputItemDtoMock{}
		dto.Status = "priceError"
		err := item.Load(&dto, invoice)
		if err == nil {
			t.Errorf("should not be nil")
		}
		if err != nil && err.Error() != "price error" {
			t.Errorf("should be price error")
		}
	})
}

func TestInvoiceItemSave(t *testing.T) {
	t.Run("should save invoice item", func(t *testing.T) {
		repo := new(RepoMock)
		item := domain.NewInvoiceItem(repo)
		if err := item.Save(); err != nil {
			t.Errorf("error on save")
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceItemError"
		item := domain.NewInvoiceItem(repo)
		if err := item.Save(); err == nil {
			t.Errorf("should not be nil")
		}
	})
}
