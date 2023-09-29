package service

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/dto"
)

func TestCreateOk(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	s := NewService(&log, &repo)
	input := dto.CreateInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "teste@teste.com",
	}
	cli := domain.Client{
		ID:       "",
		Name:     "Test Xxxx",
		Nickname: "test",
		Document: 94786984000,
		Phone:    5511999999999,
		Email:    "teste@teste.com",
	}
	output := dto.CreateOutputDto{
		Id:       cli.ID,
		Name:     cli.Name,
		Nickname: cli.Nickname,
		Document: "94786984000",
		Phone:    "5511999999999",
		Email:    cli.Email,
	}

	var res dto.CreateOutputDto
	err := s.Create(&input, &res)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected Info, got %s", log.mtype)
	}
	if !strings.Contains(log.msg, "created") {
		t.Errorf("Expected 'created', got '%s'", log.msg)
	}
	if len(repo.client.ID) != 36 {
		t.Errorf("Expected '36', got '%d'", len(repo.client.ID))
	}
	cli.ID = repo.client.ID
	if !reflect.DeepEqual(cli, *repo.client) {
		t.Errorf("Expected '%v', got '%v'", cli, repo.client)
	}
	if len(res.Id) != 36 {
		t.Errorf("Expected '36', got '%d'", len(output.Id))
	}
	output.Id = repo.client.ID
	if !reflect.DeepEqual(output, res) {
		t.Errorf("Expected '%v', got '%v'", output, res)
	}
}

func TestCreateError(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	s := NewService(&log, &repo)
	input := dto.CreateInputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: "947.869.840-01",
		Phone:    "11299999999",
		Email:    "teste",
	}

	var res dto.CreateOutputDto
	err := s.Create(&input, &res)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected 'Info', got '%s'", log.mtype)
	}
	msg := "bad request: name should have at least two parts | invalid document | invalid cell phone | invalid email"
	if err.Error() != msg {
		t.Errorf("Expected '%s', Got '%s'", msg, err.Error())
	}
}
