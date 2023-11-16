package mysql

import (
	"testing"
	"time"
)

func TestNewRepoMysql(t *testing.T) {
	t.Run("should create new repo mysql", func(t *testing.T) {
		_, err := NewRepoMysql(&ConfigMock{})
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when MYSQL_INVOICE_USER is empty", func(t *testing.T) {
		c := ConfigMock{}
		user := c.Get(mysql_user)
		c.Set(mysql_user, "")
		_, err := NewRepoMysql(&c)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		c.Set(mysql_user, user)
	})
	t.Run("should return error when MYSQL_DATABASE is empty", func(t *testing.T) {
		c := ConfigMock{}
		dbname := c.Get(mysql_dbname)
		c.Set(mysql_dbname, "")
		_, err := NewRepoMysql(&c)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		c.Set(mysql_dbname, dbname)
	})
}

func TestBegin(t *testing.T) {
	t.Run("should commit without error", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		err := repo.Begin()
		if err != nil {
			t.Errorf("Excepected nil, got %s", err.Error())
		}
	})

	t.Run("should return error when already openned", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Begin()
		err := repo.Commit()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Begin()
		err := repo.Rollback()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		err := repo.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.db.Close()
		err := repo.Close()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})
}

func TestSaveInvoiceClient(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		repo.Begin()
		client := InvoiceClientMock{}
		err := repo.SaveInvoiceClient(&client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
		repo.Truncate()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.Begin()
		repo.tx.Commit()
		err := repo.SaveInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdateInvoiceClient(t *testing.T) {
	t.Run("should update invoice client", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		client := InvoiceClientMock{}
		repo.SaveInvoiceClient(&client)
		err := repo.UpdateInvoiceClient(&client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		err := repo.UpdateInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "transaction not started" {
			t.Errorf("Expected transaction not started, got %s", err.Error())
		}
	})
	t.Run("should return error when connection is nil", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.Begin()
		repo.db = nil
		err := repo.UpdateInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when connection error", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.Begin()
		repo.tx.Commit()
		err := repo.UpdateInvoiceClient(&InvoiceClientMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetLastInvoiceClientId(t *testing.T) {
	t.Run("should return last invoice client id", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		if err := repo.Truncate(); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		defer repo.Truncate()
		repo.Begin()
		if err := repo.SaveInvoiceClient(&InvoiceClientMock{}); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if err := repo.Commit(); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		created_after := time.Now().Add(-time.Hour * 24)
		client := InvoiceClientMock{}
		ok, err := repo.GetLastInvoiceClient("nickname", created_after, &client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if !ok {
			t.Errorf("Expected ok, got !ok")
		}
		repo.Commit()
	})
	t.Run("should return error blank where no rows found", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		created_after := time.Now().Add(-time.Hour * 24)
		client := InvoiceClientMock{}
		ok, err := repo.GetLastInvoiceClient("nickname", created_after, &client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if ok {
			t.Errorf("Expected !ok, got ok")
		}
		repo.Commit()
	})
	t.Run("should return error when there is no db", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.db = nil
		client := InvoiceClientMock{}
		_, err := repo.GetLastInvoiceClient("nickname", time.Now(), &client)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
	t.Run("should return error when dabase is closed", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.db.Close()
		client := InvoiceClientMock{}
		_, err := repo.GetLastInvoiceClient("nickname", time.Now(), &client)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != "sql: database is closed" {
			t.Errorf("Expected sql: database is closed, got %s", err.Error())
		}
	})
}

func TestIsDuplicatedInvoice(t *testing.T) {
	t.Run("should return false when there is no duplicated invoice", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		if err := repo.Begin(); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if err := repo.SaveInvoiceClient(&InvoiceClientMock{}); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if err := repo.SaveInvoice(&InvoiceMock{}); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if err := repo.Commit(); err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		duplicated, err := repo.IsDuplicatedInvoice("reference")
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if duplicated {
			t.Errorf("Expected false, got true")
		}
	})
	t.Run("should return true when there is duplicated invoice", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
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
	})
	t.Run("should return error when there is no db", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		invoice := InvoiceMock{}
		err := repo.SaveInvoice(&invoice)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		repo.SaveInvoice(&InvoiceMock{})
		invoiceItem := InvoiceItemMock{}
		err := repo.SaveInvoiceItem(&invoiceItem)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
	})
	t.Run("should return error when there is no transaction", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
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
		repo, _ := NewRepoMysql(&ConfigMock{})
		repo.Begin()
		repo.tx.Commit()
		err := repo.SaveInvoiceItem(&InvoiceItemMock{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

}

func TestGetInvoiceVertex(t *testing.T) {
	t.Run("should return invoice vertex", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		status := NewInvoiceStatusMock()
		err := repo.GetInvoiceVertex(status)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if len(status.vertex) != 13 {
			t.Errorf("Expected 13, got %v", len(status.vertex))
		}
	})
}

func TestGetInvoiceEdge(t *testing.T) {
	t.Run("should return invoice edge", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		status := NewInvoiceStatusMock()
		err := repo.GetInvoiceEdge(status)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if len(status.edge) != 20 {
			t.Errorf("Expected 20, got %v", len(status.edge))
		}
	})
}

func TestStoreInvoiceStatus(t *testing.T) {
	t.Run("should store invoice status", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		repo.SaveInvoice(&InvoiceMock{})
		status := NewInvoiceStatusMock()
		err := repo.StoreInvoiceStatus("invoice", status)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
	})
}

func TestCreateInvoiceStatusLog(t *testing.T) {
	t.Run("should create invoice status log", func(t *testing.T) {
		repo, _ := NewRepoMysql(&ConfigMock{})
		defer repo.Close()
		repo.Truncate()
		defer repo.Truncate()
		repo.Begin()
		repo.SaveInvoiceClient(&InvoiceClientMock{})
		repo.SaveInvoice(&InvoiceMock{})
		repo.Commit()
		repo.Begin()
		status := NewInvoiceStatusMock()
		err := repo.CreateInvoiceStatusLog("invoice", status)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Commit()
	})
}
