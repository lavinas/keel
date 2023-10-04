package dto

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// ListAllOutputDto is the output DTO used to list all clients
type ClientListOutputDto struct {
	Clients []port.ClientCreateOutputDto `json:"clients" binding:"required"`
}

func NewClientListOutDto() *ClientListOutputDto {
	return &ClientListOutputDto{
		Clients: []port.ClientCreateOutputDto{},
	}
}

func (dto *ClientListOutputDto) Append(id, name, nick, doc, phone, email string) {
	client := NewClientCreateOutputDto()
	client.Fill(id, name, nick, doc, phone, email)
	dto.Clients = append(dto.Clients, client)
}
