package domain

import (
	"strconv"
	"errors"
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
func NewClient(name, nickName, document, phone, email string) (*Client, error) {
	doc, err := strconv.ParseUint(document, 10, 64)
	if err != nil {
		return nil, errors.New("invalid document")
	}
	ph, err := strconv.ParseUint(phone, 10, 64)
	if err != nil {
		return nil, errors.New("invalid cell phone")
	}
	return &Client{
		ID:       uuid.New().String(),
		Name:     name,
		Nickname: nickName,
		Document: doc,
		Phone:    ph,
		Email:    email,
	}, nil
}
