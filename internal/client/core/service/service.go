package service

import (
	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	Repo port.Repo
}

// NewCreate creates a new Create service
func NewService(repo port.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}

// Create creates a new client
func (u *Service) Create(input domain.CreateInputDto) (*domain.CreateOutputDto, error) {
	d := clearNumber(input.Document)
	p := clearNumber(input.Phone)

	client := domain.NewClient(input.Name, input.Nickname, p, d, input.Email)
	if err := u.Repo.Create(client); err != nil {
		return nil, err
	}
	return &domain.CreateOutputDto{
		Id:       client.ID,
		Name:     client.Name,
		Nickname: client.Nickname,
		Document: client.Document,
		Phone:    client.Phone,
		Email:    client.Email,
	}, nil
}

// ListAll list all clients
func (l *Service) ListAll() (*domain.ListAllOutputDto, error) {
	c := domain.CreateOutputDto{
		Name:     "Test",
		Nickname: "Test",
		Document: 12321232222,
		Phone:    11999999999,
		Email:    "test@test.com.br",
	}
	r := domain.ListAllOutputDto{
		Clients: []domain.CreateOutputDto{c},
	}
	return &r, nil
}

// clearNumber removes all non-numeric characters from a string
func clearNumber(number string) uint64 {
	var result uint64
	for _, n := range number {
		if n >= '0' && n <= '9' {
			result = result*10 + uint64(n-'0')
		}
	}
	return result
}
