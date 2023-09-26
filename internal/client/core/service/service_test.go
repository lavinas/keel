package service

import (
	"testing"
	"strings"
	"reflect"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/dto"
	"github.com/lavinas/keel/internal/client/util"
)


func TestCreateOk(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	util := util.NewUtil()
	s := NewService(&log, &repo, util)
	input := dto.CreateInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "teste@teste.com",
	}
	cli := domain.Client{
		ID: "",
		Name:     "Test Xxxx",
		Nickname: "test",
		Document: 94786984000,
		Phone:    5511999999999,
		Email: "teste@teste.com",
	}
	output := dto.CreateOutputDto{
		Id: cli.ID,
		Name:     cli.Name,
		Nickname: cli.Nickname,
		Document: cli.Document,
		Phone:    cli.Phone,
		Email:    cli.Email,
	}
		
	res, err := s.Create(input)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected Info, got %s", log.mtype)
	}
	if !strings.Contains(log.msg, "created") {
		t.Errorf("Expected created, got %s", log.msg)
	}
	if len(repo.client.ID) != 36 {
		t.Errorf("Expected 36, got %d", len(repo.client.ID))
	}
	cli.ID = repo.client.ID
	if !reflect.DeepEqual(cli, *repo.client) {
		t.Errorf("Expected %v, got %v", cli, repo.client)
	}
	if len(res.Id) != 36 {
		t.Errorf("Expected 36, got %d", len(output.Id))
	}
	output.Id = repo.client.ID
	if !reflect.DeepEqual(output, *res) {
		t.Errorf("Expected %v, got %v", output, res)
	}
}

func TestCreate(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	util := util.NewUtil()
	s := NewService(&log, &repo, util)
	input := dto.CreateInputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: "947.869.840-01",
		Phone:    "11299999999",
		Email:    "teste",
	}
	_, err := s.Create(input)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected Info, got %s", log.mtype)
	}
	if err.Error() != "bad request: name should have at least two parts || invalid document || invalid cell phone || invalid email" {
		t.Errorf("Expected bad request: email is invalid, got %s", err.Error())
	}
}



