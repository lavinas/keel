package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/lavinas/keel/internal/client/core/port"
)

const (
	// GroupMysql is the group name for mysql
	groupMysql = "mysql"
)

// Repo is a service to interact with the database
type RepoMysql struct {
	db *sql.DB
}

// NewRepo creates a new Repo service
func NewRepoMysql(c port.Config) *RepoMysql {
	user := getField(c, "user")
	pass := getField(c, "pass")
	host := getField(c, "host")
	port := getField(c, "port")
	dbname := getField(c, "dbname")
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	return &RepoMysql{db: db}
}

// Create creates a new client
func (r *RepoMysql) CreateClient(domain port.Domain) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	id, name, nick, doc, phone, email := domain.GetClient()
	_, err = tx.Exec("insert into client (id, name, nickname, document, phone, email) values (?, ?, ?, ?, ?, ?)",
		id, name, nick, doc, phone, email)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// Close closes the database connection
func (r *RepoMysql) Close() {
	r.db.Close()
}

// getField gets a field from a group in config file
func getField(c port.Config, field string) string {
	r, err := c.GetField(groupMysql, field)
	if err != nil {
		panic(err)
	}
	return r
}
