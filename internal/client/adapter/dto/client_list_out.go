package dto

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// ListAllOutputDto is the output DTO used to list all clients
type ClientListOutputDto struct {
	Page    uint64                      `json:"page" binding:"required"`
	PerPage uint64                      `json:"per_page" binding:"required"`
	Clients []port.ClientInserOutputDto `json:"clients" binding:"required"`
}

// Append appends a new client to the set
func (dto *ClientListOutputDto) Append(id, name, nick, doc, phone, email string) {
	client := NewClientInserOutputDto()
	client.Fill(id, name, nick, doc, phone, email)
	dto.Clients = append(dto.Clients, client)
}

// Count returns the number of clients in the set
func (dto *ClientListOutputDto) Count() int {
	return len(dto.Clients)
}

// SetPage sets the page and perPage for the client set
func (dto *ClientListOutputDto) SetPage(page, perPage uint64) {
	dto.Page = page
	dto.PerPage = perPage
}
