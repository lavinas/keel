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
func (r *RepoMysql) ClientLoad(page, perPage uint64, name, nick, doc, email string, set port.ClientSet) error {
	query := clientListBase
	q, args := loadFilters(name, nick, doc, email)
	query += q
	q, a := loadPagination(page, perPage)
	query += q
	args = append(args, a...)
	row, err := r.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer row.Close()
	if err := loadInterate(row, set); err != nil {
		return err
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

// loadFilters prepate the filters query Load
func loadFilters(name, nick, doc, email string) (string, []interface{}) {
	query := ""
	args := make([]interface{}, 0)
	if name != "" {
		query += clientListFilterName
		args = append(args, "%"+name+"%")
	}
	if nick != "" {
		query += clientListFilterNick
		args = append(args, "%"+nick+"%")
	}
	if doc != "" {
		query += clientListFilterDoc
		args = append(args, "%"+doc+"%")
	}
	if email != "" {
		query += clientListFilterEmail
		args = append(args, "%"+email+"%")
	}
	return query, args
}

// loadPagination prepate the pagination query Load
func loadPagination(page, perPage uint64) (string, []interface{}) {
	query := ""
	args := make([]interface{}, 0)
	// Pagination
	query += clientListPagination
	args = append(args, perPage)
	args = append(args, (page-1)*perPage)
	return query, args
}

// loadInterate iterates over the rows and append to the set
func loadInterate(row *sql.Rows, set port.ClientSet) error {
	var id, name, nick, email string
	var doc, phone uint64
	for row.Next() {
		if err := row.Scan(&id, &name, &nick, &doc, &phone, &email); err != nil {
			return err
		}
		set.Append(id, name, nick, doc, phone, email)
	}
	return nil
}
