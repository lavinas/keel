package domain

import (
	"strconv"
	"fmt"

	"github.com/google/uuid"
	"github.com/lavinas/keel/internal/client/core/port"
)

// Client is the domain model for a client
type Client struct {
	repo    port.Repo
	ID       string
	Name     string
	Nickname string
	Document uint64
	Phone    uint64
	Email    string
}

// NewClient creates a new client
func NewClient(repo port.Repo, input port.CreateInputDto) (*Client, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	doc, err := strconv.ParseUint(input.GetDocument(), 10, 64)
	if err != nil {
		return nil, err
	}
	ph, err := strconv.ParseUint(input.GetPhone(), 10, 64)
	if err != nil {
		return nil, err
	}
	return &Client{
		repo:    repo,
		ID:       id.String(),
		Name:     input.GetName(),
		Nickname: input.GetNickname(),
		Document: doc,
		Phone:    ph,
		Email:    input.GetEmail(),
	}, nil
}

// DocumentDuplicity checks if a document is already registered
func (c *Client) DocumentDuplicity() (bool, error) {
	return c.repo.ClientDocumentDuplicity(c.Document)
}

// EmailDuplicity checks if a email is already registered
func (c *Client) EmailDuplicity() (bool, error) {
	return c.repo.ClientEmailDuplicity(c.Email)
}

func (c *Client) Get() (string, string, string, uint64, uint64, string) {
	return c.ID, c.Name, c.Nickname, c.Document, c.Phone, c.Email
}

func (c *Client) GetFormatted() (string, string, string, string, string, string) {
	doc := fmt.Sprintf("%d", c.Document)
	if len(doc) <= 11 {
		doc = fmt.Sprintf("%011s", doc)
	} else {
		doc = fmt.Sprintf("%014s", doc)
	}
	return c.ID, c.Name, c.Nickname, doc, strconv.FormatUint(c.Phone, 10), c.Email
}

// Save saves a client on the repository
func (c *Client) Save() error {
	return c.repo.ClientSave(c)
}

