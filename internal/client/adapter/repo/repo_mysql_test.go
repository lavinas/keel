package repo

import (
	"testing"

	"github.com/lavinas/keel/internal/client/adapter/dto"
	"github.com/lavinas/keel/internal/client/core/domain"
)

func TestClientSave(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	defer repo.Close()
	repo.ClientTruncate()
	input := dto.NewCreateInputDto("Test Xxxx", "test", "94786984000", "5511999999999", "test@test.com") 
	client, err := domain.NewClient(repo, &input)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.ClientSave(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	repo.ClientTruncate()
}

func TestClientDocumentDuplicity(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	defer repo.Close()
	repo.ClientTruncate()
	// check not duplicated
	b, err := repo.ClientDocumentDuplicity(94786984000)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Document should not be duplicated")
	}
	// check duplicated
	input := dto.NewCreateInputDto("Test Xxxx", "test", "94786984000", "5511999999999", "test@test.com") 
	client, err := domain.NewClient(repo, &input)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.ClientSave(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	b, err = repo.ClientDocumentDuplicity(94786984000)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Document should be duplicated")
	}
	repo.ClientTruncate()
}

func TestClientEmailDuplicityQuery(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	defer repo.Close()
	repo.ClientTruncate()
	// check not duplicated
	b, err := repo.ClientEmailDuplicity("test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Email should not be duplicated")
	}
	// check duplicated
	input := dto.NewCreateInputDto("Test Xxxx", "test", "94786984000", "5511999999999", "test@test.com") 
	client, err := domain.NewClient(repo, &input)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.ClientSave(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	b, err = repo.ClientEmailDuplicity("test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Email should be duplicated")
	}
}
