package service

import (
	_ "reflect"
	"strings"
	"testing"

	"github.com/lavinas/keel/internal/client/adapter/dto"
	"github.com/lavinas/keel/internal/client/adapter/repo"
	"github.com/lavinas/keel/internal/client/core/domain"
)

func TestInsertOk(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	config := ConfigMock{}
	domain := domain.NewDomain(&repo)
	s := NewService(domain, &config, &log, &repo)
	input := dto.InsertInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "teste@teste.com",
	}
	var res dto.InsertOutputDto
	err := s.Insert(&input, &res)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected Info, got %s", log.mtype)
	}
	if !strings.Contains(log.msg, "created") {
		t.Errorf("Expected 'created', got '%s'", log.msg)
	}
}

func TestInsertError(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	config := ConfigMock{}
	domain := domain.NewDomain(&repo)
	s := NewService(domain, &config, &log, &repo)
	input := dto.InsertInputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: "947.869.840-01",
		Phone:    "11299999999",
		Email:    "teste",
	}
	var res dto.InsertOutputDto
	err := s.Insert(&input, &res)
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

func TestInsertDuplicity(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	config := ConfigMock{}
	repo.ClientDocumentDuplicityReturn = true
	repo.ClientEmailDuplicityReturn = true
	domain := domain.NewDomain(&repo)
	s := NewService(domain, &config, &log, &repo)
	input := dto.InsertInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "test@test.com",
	}
	var res dto.InsertOutputDto
	err := s.Insert(&input, &res)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected 'Info', got '%s'", log.mtype)
	}
	msg := "conflict: document already registered | email already registered"
	if err.Error() != msg {
		t.Errorf("Expected '%s', Got '%s'", msg, err.Error())
	}
}

func TestWithDB(t *testing.T) {
	c := ConfigMock{}
	l := LogMock{}
	r := repo.NewRepoMysql(&c)
	defer r.Close()
	r.Truncate()
	d := domain.NewDomain(r)

	s := NewService(d, &c, &l, r)

	input := dto.InsertInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "teste@teste.com",
	}
	var res dto.InsertOutputDto
	err := s.Insert(&input, &res)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l.mtype != "Info" {
		t.Errorf("Expected Info, got %s", l.mtype)
	}
	if !strings.Contains(l.msg, "created") {
		t.Errorf("Expected 'created', got '%s'", l.msg)
	}
	r.Truncate()
}