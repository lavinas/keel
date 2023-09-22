package service

import (
	"testing"
	"strings"

	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/util"
)


func TestCreate(t *testing.T) {
	log := LogMock{}
	repo := RepoMock{}
	util := util.NewUtil()
	s := NewService(&log, &repo, util)
	input := domain.CreateInputDto{
		Name:     "Test XXXX",
		Nickname: "Test",
		Document: "947.869.840-00",
		Phone:    "11999999999",
		Email:    "teste@teste.com",
	}
	_, err := s.Create(input)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if log.mtype != "Info" {
		t.Errorf("Expected Info, got %s", log.mtype)
	}

	if !strings.Contains(log.msg, "created") {
		t.Errorf("Expected created, got %s", log.msg)
	}

}


