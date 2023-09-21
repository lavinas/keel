package domain

import (
	"github.com/google/uuid"
)

// Client is the domain model for a client
type Client struct {
	ID       string
	Name     string
	Nickname string
	Document uint64
	Phone    uint64
	Email    string
}

// NewClient creates a new client
func NewClient(name, nickName string, document, phone uint64, email string) *Client {
	return &Client{
		ID:       uuid.New().String(),
		Name:     name,
		Nickname: nickName,
		Document: document,
		Phone:    phone,
		Email:    email,
	}
}