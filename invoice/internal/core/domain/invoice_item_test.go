package domain

import (
	"testing"
)

func TestInvoiceItemLoad(t *testing.T) {
	t.Run("should load invoice item without eerror", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		item := NewInvoiceItem(repo)
		dto := CreateInputItemDtoMock{}
		if err := item.Load(&dto, invoice); err != nil {
			t.Errorf("error on load invoice item")
		}
	})
	t.Run("should return quantity error", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		item := NewInvoiceItem(repo)
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
		invoice := NewInvoice(repo)
		item := NewInvoiceItem(repo)
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
		item := NewInvoiceItem(repo)
		if err := item.Save(); err != nil {
			t.Errorf("error on save")
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceItemError"
		item := NewInvoiceItem(repo)
		if err := item.Save(); err == nil {
			t.Errorf("should not be nil")
		}
	})
}

func TestInvoiceItemGetId(t *testing.T) {
	t.Run("should return id", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		item.id = "id"
		if item.GetId() != "id" {
			t.Errorf("id is empty")
		}
	})
}

func TestInvoiceItemGetInvoiceId(t *testing.T) {
	t.Run("should return invoice id", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		invoice := NewInvoice(repo)
		invoice.id = "id"
		item.invoice = invoice
		if item.GetInvoiceId() != "id" {
			t.Errorf("id is empty")
		}
	})
}

func TestInvoiceItemGetServiceReference(t *testing.T) {
	t.Run("should return service reference", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		item.serviceReference = "ref"
		if item.GetServiceReference() != "ref" {
			t.Errorf("service reference is empty")
		}
	})
}

func TestInvoiceItemGetDescription(t *testing.T) {
	t.Run("should return description", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		item.description = "description"
		if item.GetDescription() != "description" {
			t.Errorf("description is empty")
		}
	})
}

func TestInvoiceItemGetAmount(t *testing.T) {
	t.Run("should return amount", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		item.amount = 10
		if item.GetAmount() != 10 {
			t.Errorf("amount is empty")
		}
	})
}

func TestInvoiceItemGetQuantity(t *testing.T) {
	t.Run("should return quantity", func(t *testing.T) {
		repo := new(RepoMock)
		item := NewInvoiceItem(repo)
		item.quantity = 10
		if item.GetQuantity() != 10 {
			t.Errorf("quantity is empty")
		}
	})
}
