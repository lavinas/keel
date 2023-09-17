package repo

import (
	"database/sql"

	"github.com/lavinas/keel/internal/client/core/domain"
)

// Repo is a service to interact with the database
type RepoMysql struct{
	db *sql.DB
}

// NewRepo creates a new Repo service
func NewRepoMysql(db *sql.DB) *RepoMysql {
	return &RepoMysql{db: db}
}

// Create creates a new client
func (r *RepoMysql) Create(client *domain.Client) error {
	_, err := r.db.Exec("insert into client (id, name, nickname, document, phone, email) values (?, ?, ?, ?, ?, ?)",
		client.ID, client.Name, client.Nickname, client.Document, client.Phone, client.Email)
	if err != nil {
		return err
	}
	return nil
}

