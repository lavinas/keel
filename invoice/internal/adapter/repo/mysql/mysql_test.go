package mysql

import (
	"os"
	"testing"
)

func TestNewRepoMysql(t *testing.T) {
	t.Run("should create new repo mysql", func(t *testing.T) {
		_, err := NewRepoMysql()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when MYSQL_USER is empty", func(t *testing.T) {
		user := os.Getenv(mysql_user)
		os.Setenv(mysql_user, "")
		_, err := NewRepoMysql()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		os.Setenv(mysql_user, user)
	})
	t.Run("should return error when MYSQL_INVOICE_DATABASE is empty", func(t *testing.T) {
		dbname := os.Getenv(mysql_dbname)
		os.Setenv(mysql_dbname, "")
		_, err := NewRepoMysql()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		os.Setenv(mysql_dbname, dbname)
	})
}

func TestBegin(t *testing.T) {
	t.Run("should commit without error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.Begin()
		if err != nil {
			t.Errorf("Excepected nil, got %s", err.Error())
		}
	})

	t.Run("should return error when already openned", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		err := repo.Begin()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction already started" {
			t.Errorf("Expected transaction already started, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no connection", func(t *testing.T) {
		repo := RepoMysql{}
		err := repo.Begin()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when wrong connection", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		err := repo.db.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		err = repo.Begin()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
}

func TestCommit(t *testing.T) {
	t.Run("should commit without error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		err := repo.Commit()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.Commit()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.Commit()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.Commit()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestRollback(t *testing.T) {
	t.Run("should rollback without error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		err := repo.Rollback()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.Rollback()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		err := repo.db.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		err = repo.Rollback()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.Rollback()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.Rollback()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestClose(t *testing.T) {
	t.Run("should close without error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		err := repo.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.db = nil
		err := repo.Close()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should close with error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.db.Close()
		err := repo.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
}

func TestSaveInvoiceClient(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		repo.TruncateInvoiceClient()
		client := InvoiceClientMock{}
		err := repo.SaveInvoiceClient(&client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.TruncateInvoiceClient()
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.SaveInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.SaveInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.SaveInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestIsDuplicatedInvoice(t *testing.T) {
	t.Run("should return false when there is no duplicated invoice", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		repo.Commit()
		duplicated, err := repo.IsDuplicatedInvoice("reference")
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if duplicated {
			t.Errorf("Expected false, got true")
		}
		repo.Begin()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.Commit()
	})
	t.Run("should return true when there is duplicated invoice", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.Begin()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		invoice := InvoiceMock{}
		repo.SaveInvoice(&invoice)
		repo.Commit()
		duplicated, err := repo.IsDuplicatedInvoice(invoice.GetReference())
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if !duplicated {
			t.Errorf("Expected true, got false")
		}
		repo.Begin()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.Commit()
	})
	t.Run("should return error when there is no db", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.db = nil
		_, err := repo.IsDuplicatedInvoice("reference")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when dabase is closed", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.db.Close()
		_, err := repo.IsDuplicatedInvoice("reference")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})

}

func TestSaveInvoice(t *testing.T) {
	t.Run("should save invoice", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		defer repo.Close()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		invoice := InvoiceMock{}
		err := repo.SaveInvoice(&invoice)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.SaveInvoice(&InvoiceMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.SaveInvoice(&InvoiceMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.SaveInvoice(&InvoiceMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestSaveInvoiceItem(t *testing.T) {
	t.Run("should save invoice item", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		repo.TruncateInvoiceItem()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		repo.SaveInvoice(&InvoiceMock{})
		invoiceItem := InvoiceItemMock{}
		err := repo.SaveInvoiceItem(&invoiceItem)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.TruncateInvoiceItem()
		repo.TruncateInvoice()
		repo.TruncateInvoiceClient()
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.SaveInvoiceItem(&InvoiceItemMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.SaveInvoiceItem(&InvoiceItemMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.SaveInvoiceItem(&InvoiceItemMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

}

func TestTruncateInvoiceClient(t *testing.T) {
	t.Run("should truncate without error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		repo.Begin()
		err := repo.TruncateInvoiceClient()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		defer repo.Close()
		err := repo.TruncateInvoiceClient()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.db = nil
		err := repo.TruncateInvoiceClient()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql()
		repo.Begin()
		repo.tx.Commit()
		err := repo.TruncateInvoiceClient()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
