package repo

import (
	"testing"

	"github.com/lavinas/keel/internal/client/core/domain"

)

func TestCreateOk(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	defer repo.Close()
	cli, err := domain.NewClient("Test Xxxx", "test", "94786984000", "5511999999999", "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Create(cli)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}