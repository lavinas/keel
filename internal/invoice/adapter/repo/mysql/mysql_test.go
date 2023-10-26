package mysql

import (
	"testing"
)

func TestSaveInvoiceClient(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo := NewRepoMysql()
		defer repo.Close()
		if err := repo.Truncate(); err != nil {
			t.Errorf("truncate error: %s", err.Error())
		}
		client := InvoiceClientMock{}
		err := repo.SaveInvoiceClient(&client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Truncate()
	})
}
