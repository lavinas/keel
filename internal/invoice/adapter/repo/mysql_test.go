package repo

import (
	"testing"
)

func TestSaveInvoiceClient(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo := NewRepoMysql()
		repo.Truncate()
		client := InvoiceClientMock{}
		err := repo.SaveInvoiceClient(&client)
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		repo.Truncate()
	})
}
