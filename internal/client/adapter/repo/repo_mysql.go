package repo

import (
	"context"
	"database/sql"
	"time"

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

// Update updates a client on the repository
func (r *RepoMysql) ClientUpdate(client port.Client) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	id, name, nick, doc, phone, email := client.Get()
	_, err = tx.Exec(clientUpdateQuery, name, nick, doc, phone, email, id)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// ClientDocumentDuplicity checks if a document is already registered
func (r *RepoMysql) ClientDocumentDuplicity(document uint64, id string) (bool, error) {
	var count int
	row := r.db.QueryRow(clientDocumentDuplicityQuery, document, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientEmailDuplicity checks if an email is already registered
func (r *RepoMysql) ClientEmailDuplicity(email, id string) (bool, error) {
	var count int
	row := r.db.QueryRow(clientEmailDuplicityQuery, email, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientNickDuplicity checks if a nick is already registered
func (r *RepoMysql) ClientNickDuplicity(nick, id string) (bool, error) {
	var count int
	row := r.db.QueryRow(clientNickDuplicityQuery, nick, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientGetAll gets all clients
func (r *RepoMysql) ClientLoadSet(page, perPage uint64, name, nick, doc, email string, set port.ClientSet) error {
	query, args := clientLoadSetQuery(page, perPage, name, nick, doc, email)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return err
	}
	defer row.Close()
	if err := clientLoadSetInterate(row, set); err != nil {
		return err
	}
	return nil
}

// ClientGetById gets a client by id
func (r *RepoMysql) ClientGetById(id string, client port.Client) (bool, error) {
	row := r.db.QueryRow(clientGetById, id)
	var rid, name, nick, email string
	var doc, phone uint64
	if err := row.Scan(&rid, &name, &nick, &doc, &phone, &email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(rid, name, nick, doc, phone, email)
	return true, nil
}

// ClientGetByNick gets a client by nick
func (r *RepoMysql) ClientGetByNick(nick string, client port.Client) (bool, error) {
	row := r.db.QueryRow(clientGetByNick, nick)
	var id, rnick, name, email string
	var doc, phone uint64
	if err := row.Scan(&id, &rnick, &name, &doc, &phone, &email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(id, name, rnick, doc, phone, email)
	return true, nil
}

// ClientGetByEmail gets a client by email
func (r *RepoMysql) ClientGetByEmail(email string, client port.Client) (bool, error) {
	row := r.db.QueryRow(clientGetByEmail, email)
	var id, name, nick, remail string
	var doc, phone uint64
	if err := row.Scan(&id, &name, &nick, &doc, &phone, &remail); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(id, name, nick, doc, phone, remail)
	return true, nil
}

// ClientGetByDoc gets a client by doc
func (r *RepoMysql) ClientGetByDoc(doc uint64, client port.Client) (bool, error) {
	row := r.db.QueryRow(clientGetByDoc, doc)
	var id, name, nick, email string
	var rdoc, phone uint64
	if err := row.Scan(&id, &name, &nick, &rdoc, &phone, &email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(id, name, nick, rdoc, phone, email)
	return true, nil
}

// ClientGetByPhone gets a client by phone
func (r *RepoMysql) ClientGetByPhone(phone uint64, client port.Client) (bool, error) {
	row := r.db.QueryRow(clientGetByPhone, phone)
	var id, name, nick, email string
	var doc, rphone uint64
	if err := row.Scan(&id, &name, &nick, &doc, &rphone, &email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(id, name, nick, doc, rphone, email)
	return true, nil
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
func (r *RepoMysql) Close() error {
	if err := r.db.Close(); err != nil {
		return err
	}
	return nil
}

// getField gets a field from a group in config file
func getField(c port.Config, field string) string {
	r, err := c.GetField(groupMysql, field)
	if err != nil {
		panic(err)
	}
	return r
}

// clientLoadSetQuery prepate the query for Load Set
func clientLoadSetQuery(page, perPage uint64, name, nick, doc, email string) (string, []interface{}) {
	query := clientListBase
	q, args := clientLoadSetFilters(name, nick, doc, email)
	query += q
	q, a := clientLoadSetPagination(page, perPage)
	query += q
	args = append(args, a...)
	return query, args

}

// clientLoadSetInterate iterates over the rows and append to the set
func clientLoadSetInterate(row *sql.Rows, set port.ClientSet) error {
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

// clientLoadSetFilters prepate the filters query Load
func clientLoadSetFilters(name, nick, doc, email string) (string, []interface{}) {
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

// clientLoadSetPagination prepate the pagination query Load
func clientLoadSetPagination(page, perPage uint64) (string, []interface{}) {
	query := ""
	args := make([]interface{}, 0)
	// Pagination
	query += clientListPagination
	args = append(args, perPage)
	args = append(args, (page-1)*perPage)
	return query, args
}
