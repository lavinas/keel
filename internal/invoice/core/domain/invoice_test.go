package domain

import (
	"testing"
	"time"
)

func TestInvoiceLoad(t *testing.T) {
	t.Run("should load a invoice", func(t *testing.T) {
		repo := new(RepoMock)
		business := NewInvoiceClient(repo)
		business.Load("business", "clientId", "name", "email", 123456789, 123456789)
		customer := NewInvoiceClient(repo)
		customer.Load("customer", "clientId", "name", "email", 123456789, 123456789)
		dto := CreateInputDtoMock{}
		invoice := NewInvoice(repo)
		if err := invoice.Load(&dto, business, customer); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
	})
	t.Run("should return amount error", func(t *testing.T) {
		repo := new(RepoMock)
		business := NewInvoiceClient(repo)
		business.Load("business", "clientId", "name", "email", 123456789, 123456789)
		customer := NewInvoiceClient(repo)
		customer.Load("customer", "clientId", "name", "email", 123456789, 123456789)
		dto := CreateInputDtoMock{}
		dto.Status = "amountError"
		invoice := NewInvoice(repo)
		err := invoice.Load(&dto, business, customer)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "amount error" {
			t.Errorf("expected amount error, got %v", err.Error())
		}
	})
	t.Run("should return date error", func(t *testing.T) {
		repo := new(RepoMock)
		business := NewInvoiceClient(repo)
		business.Load("business", "clientId", "name", "email", 123456789, 123456789)
		customer := NewInvoiceClient(repo)
		customer.Load("customer", "clientId", "name", "email", 123456789, 123456789)
		dto := CreateInputDtoMock{}
		dto.Status = "dateError"
		invoice := NewInvoice(repo)
		err := invoice.Load(&dto, business, customer)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "date error" {
			t.Errorf("expected date error, got %v", err.Error())
		}
	})
	t.Run("should return due error", func(t *testing.T) {
		repo := new(RepoMock)
		business := NewInvoiceClient(repo)
		business.Load("business", "clientId", "name", "email", 123456789, 123456789)
		customer := NewInvoiceClient(repo)
		customer.Load("customer", "clientId", "name", "email", 123456789, 123456789)
		dto := CreateInputDtoMock{}
		dto.Status = "dueError"
		invoice := NewInvoice(repo)
		err := invoice.Load(&dto, business, customer)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err != nil && err.Error() != "due error" {
			t.Errorf("expected due error, got %v", err.Error())
		}
	})
}

func TestInvoiceSetAmount(t *testing.T) {
	t.Run("should set amount", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		invoice.SetAmount(123.23)
		if invoice.amount != 123.23 {
			t.Errorf("expected amount, got empty")
		}
	})
}

func TestInvoiceSave(t *testing.T) {
	t.Run("should save invoice", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		if err := invoice.Save(); err != nil {
			t.Errorf("expected nil, got %v", err.Error())
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceError"
		invoice := NewInvoice(repo)
		if err := invoice.Save(); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestInvoiceGetId(t *testing.T) {
	t.Run("should return id", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		invoice.id = "id"
		if invoice.GetId() != "id" {
			t.Errorf("expected id, got empty")
		}
	})
}

func TestInvoiceGetReference(t *testing.T) {
	t.Run("should return reference", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		invoice.reference = "reference"
		if invoice.GetReference() != "reference" {
			t.Errorf("expected reference, got empty")
		}
	})
}

func TestInvoiceGetBusinessId(t *testing.T) {
	t.Run("should return business id", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		business := NewInvoiceClient(repo)
		business.id = "id"
		invoice.business = business
		if invoice.GetBusinessId() != "id" {
			t.Errorf("expected id, got empty")
		}
	})
}

func TestInvoiceGetCustomerId(t *testing.T) {
	t.Run("should return customer id", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		customer := NewInvoiceClient(repo)
		customer.id = "id"
		invoice.customer = customer
		if invoice.GetCustomerId() != "id" {
			t.Errorf("expected id, got empty")
		}
	})
}

func TestInvoiceGetAmount(t *testing.T) {
	t.Run("should return amount", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		invoice.amount = 123
		if invoice.GetAmount() != 123 {
			t.Errorf("expected amount, got empty")
		}
	})
}

func TestInvoiceGetDate(t *testing.T) {
	t.Run("should return date", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		date, _ := time.Parse("2006-01-02", "2020-01-01")
		invoice.date = date
		if invoice.GetDate() != date {
			t.Errorf("expected date, got empty")
		}
	})
}

func TestInvoiceGetDue(t *testing.T) {
	t.Run("should return due", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		due, _ := time.Parse("2006-01-02", "2020-01-01")
		invoice.due = due
		if invoice.GetDue() != due {
			t.Errorf("expected due, got empty")
		}
	})
}

func TestInvoiceGetNoteId(t *testing.T) {
	t.Run("should return note id", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		noteId := invoice.GetNoteId()
		if noteId != nil {
			t.Errorf("expected nil, got %s", *noteId)
		}
	})
}

func TestInvoiceGetStatusId(t *testing.T) {
	t.Run("should return status id", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
		invoice.status_id = 1
		if invoice.GetStatusId() != 1 {
			t.Errorf("expected status id, got empty")
		}
	})
}

func TestInvoiceGetCreatedAt(t *testing.T) {
	t.Run("should return created at", func(t *testing.T) {
		repo := new(RepoMock)
		invoice := NewInvoice(repo)
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
		invoice := NewInvoice(repo)
		updatedAt, _ := time.Parse("2006-01-02", "2020-01-01")
		invoice.UpdatedAt = updatedAt
		if invoice.GetUpdatedAt() != updatedAt {
			t.Errorf("expected updated at, got empty")
		}
	})
}

