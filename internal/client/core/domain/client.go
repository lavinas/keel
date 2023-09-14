package domain

import (
	"github.com/google/uuid"
)

type Client struct {
	ID       string
	Name     string
	NickName string
	Document uint64
	Phone    uint64
	Email    string
}

func NewClient(name, nickName string, document, phone uint64, email string) *Client {
	return &Client{
		ID:       uuid.New().String(),
		Name:     name,
		NickName: nickName,
		Document: document,
		Phone:    phone,
		Email:    email,
	}
}
