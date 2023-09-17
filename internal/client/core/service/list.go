package service

import (
	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/port"
)

// List is a service to list clients
type List struct {
	Repo port.Repo
}

func NewList(repo port.Repo) *List {
	return &List{
		Repo: repo,
	}
}

func (l *List) ListAll() (*domain.ListAllOutputDto, error) {
	c := domain.CreateOutputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: 12321232222,
		Phone:    11999999999,
		Email:    "test@test.com.br",		
	}
	r := domain.ListAllOutputDto {
		Clients: []domain.CreateOutputDto{c},
	}
	return &r, nil
}


