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
func (r *RepoMysql) ClientSave(client port.Client) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	id, name, nick, doc, phone, email := client.Get()
	_, err = tx.Exec(clientSaveQuery, id, name, nick, doc, phone, email)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// ClientDocumentDuplicity checks if a document is already registered
func (r *RepoMysql) ClientDocumentDuplicity(document uint64) (bool, error) {
	var count int
	row := r.db.QueryRow(clientDocumentDuplicityQuery, document)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientEmailDuplicity checks if an email is already registered
func (r *RepoMysql) ClientEmailDuplicity(email string) (bool, error) {
	var count int
	row := r.db.QueryRow(clientEmailDuplicityQuery, email)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientGetAll gets all clients
func (r *RepoMysql) ClientLoad(page, perPage uint64, set port.ClientSet) error {
	limit := perPage
	offset := (page - 1) * perPage
	row, err := r.db.Query(clientGetAll, limit, offset)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		var id, name, nick, email string
		var doc, phone uint64
		if err := row.Scan(&id, &name, &nick, &doc, &phone, &email); err != nil {
			return err
		}
		set.Append(id, name, nick, doc, phone, email)
	}
	return nil
}

// ClientTruncate truncates the client table
func (r *RepoMysql) ClientTruncate() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(clientTruncateQuery)
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
