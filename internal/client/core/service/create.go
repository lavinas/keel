package service

import (
	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Create is the service to create a client
type Create struct {
	Repo port.Repo
}

//NewCreate creates a new Create service
func NewCreate(repo port.Repo) *Create {
	return &Create{
		Repo: repo,
	}
}

// Execute creates a new client
func (u *Create) Execute(input domain.CreateInputDto) (*domain.CreateOutputDto, error) {
	d := clearNumber(input.Document)
	p := clearNumber(input.Phone)

	client := domain.NewClient(input.Name, input.Nickname, p, d, input.Email)
	// if err := u.Repo.Create(client); err != nil {
	// 	return nil, err
	// }
	return &domain.CreateOutputDto{
		Name:     client.Name,
		Nickname: client.Nickname,
		Document: client.Document,
		Phone:    client.Phone,
		Email:    client.Email,
	}, nil
}

//clearNumber removes all non-numeric characters from a string
func clearNumber(number string) uint64 {
	var result uint64
	for _, n := range number {
		if n >= '0' && n <= '9' {
			result = result * 10 + uint64(n-'0')
		}
	}
	return result
}
