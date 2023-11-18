package domain_test

import (
	"testing"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/domain"
)

func TestInvoiceLoad(t *testing.T) {
	t.Run("should load a invoice", func(t *testing.T) {
		repo := new(RepoMock)
		business := domain.NewInvoiceClient(repo)
		business.Load("", "business", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		customer := domain.NewInvoiceClient(repo)
		customer.Load("", "customer", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		dto := CreateInputDtoMock{}
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err == nil {
			if invoice.CreatedAt.IsZero() {
				t.Errorf("expected created at, got empty")
			}
			if invoice.UpdatedAt.IsZero() {
				t.Errorf("expected updated at, got empty")
			}
		}

	})
	t.Run("should return amount error", func(t *testing.T) {
		repo := new(RepoMock)
		dto := CreateInputDtoMock{}
		dto.Status = "amountError"
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "amount error" {
			t.Errorf("expected amount error, got %v", err.Error())
		}
	})
	t.Run("should return date error", func(t *testing.T) {
		repo := new(RepoMock)
		dto := CreateInputDtoMock{}
		dto.Status = "dateError"
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "date error" {
			t.Errorf("expected date error, got %v", err.Error())
		}
	})
	t.Run("should return due error", func(t *testing.T) {
		repo := new(RepoMock)
		dto := CreateInputDtoMock{}
		dto.Status = "dueError"
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "due error" {
			t.Errorf("expected due error, got %v", err.Error())
		}
	})
	t.Run("should return error loading items", func(t *testing.T) {
		repo := new(RepoMock)
		dto := CreateInputDtoMock{}
		dto.Status = "itemsError"
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "quantity error" {
			t.Errorf("expected quantity error, got %v", err.Error())
		}
	})
	t.Run("should return error loading business", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "getLastInvoiceBusinessError"
		dto := CreateInputDtoMock{}
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "get last invoice client error" {
			t.Errorf("expected get last invoice client error, got %v", err.Error())
		}
	})
	t.Run("should return error loading customer", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "getLastInvoiceCustomerError"
		dto := CreateInputDtoMock{}
		invoice := domain.NewInvoice(repo)
		err := invoice.Load(&dto)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "get last invoice client error" {
			t.Errorf("expected get last invoice client error, got %v", err.Error())
		}
	})
}

func TestInvoiceSave(t *testing.T) {
	t.Run("should save invoice", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return error on save businness client", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveBusinessError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return error on save customer client", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveCustomerError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return error on save invoice item", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceItemError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return error on begin transaction", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "beginError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
	t.Run("should return error on commit transaction", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "commitError"
		invoice := domain.NewInvoice(repo)
		if err := invoice.Load(&CreateInputDtoMock{}); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestInvoiceGetCreatedAt(t *testing.T) {
	t.Run("should return created at", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		createdAt, _ := time.Parse("2006-01-02", "2020-01-01")
		invoice.CreatedAt = createdAt
		if invoice.GetCreatedAt() != createdAt {
			t.Errorf("expected created at, got empty")
		}
	})
}

func TestInvoiceGetUpdatedAt(t *testing.T) {
	t.Run("should return updated at", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := domain.NewInvoice(repo)
		updatedAt, _ := time.Parse("2006-01-02", "2020-01-01")
		invoice.UpdatedAt = updatedAt
		if invoice.GetUpdatedAt() != updatedAt {
			t.Errorf("expected updated at, got empty")
		}
	})
}
