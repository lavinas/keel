package mysql

import (
	"context"
	"database/sql"
	"os"
	"time"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lavinas/keel/internal/client/core/port"
)

const (
	mysql_user   = "MYSQL_USER"
	mysql_pass   = "MYSQL_PASSWORD"
	mysql_host   = "MYSQL_HOST"
	mysql_port   = "MYSQL_PORT"
	mysql_dbname = "MYSQL_CLIENT_DATABASE"
)

// Repo is a service to interact with the database
type RepoMysql struct {
	db *sql.DB
}

// NewRepo creates a new Repo service
func NewRepoMysql() *RepoMysql {
	user := os.Getenv(mysql_user)
	pass := os.Getenv(mysql_pass)
	host := os.Getenv(mysql_host)
	port := os.Getenv(mysql_port)
	dbname := os.Getenv(mysql_dbname)
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/")
	if err != nil {
		panic(err)
	}
	if dbname == "" {
		panic("MYSQL_CLIENT_BASE environment variable is empty")
	}
	for i, q := range dbQueries {
		dbQueries[i] = strings.Replace(q, "{DB}", dbname, -1)
	}
	return &RepoMysql{db: db}
}

// Insert creates a new client
func (r *RepoMysql) Save(client port.Client) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	id, name, nick, doc, phone, email := client.Get()
	q := dbQueries["clientSaveQuery"]
	_, err = tx.Exec(q, id, name, nick, doc, phone, email)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// Update updates a client on the repository
func (r *RepoMysql) Update(client port.Client) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	id, name, nick, doc, phone, email := client.Get()
	q := dbQueries["clientUpdateQuery"]
	_, err = tx.Exec(q, name, nick, doc, phone, email, id)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// ClientDocumentDuplicity checks if a document is already registered
func (r *RepoMysql) DocumentDuplicity(document uint64, id string) (bool, error) {
	var count int
	q := dbQueries["clientDocumentDuplicityQuery"]
	row := r.db.QueryRow(q, document, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientEmailDuplicity checks if an email is already registered
func (r *RepoMysql) EmailDuplicity(email, id string) (bool, error) {
	var count int
	q := dbQueries["clientEmailDuplicityQuery"]
	row := r.db.QueryRow(q, email, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// ClientNickDuplicity checks if a nick is already registered
func (r *RepoMysql) NickDuplicity(nick, id string) (bool, error) {
	var count int
	q := dbQueries["clientNickDuplicityQuery"]
	row := r.db.QueryRow(q, nick, id)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetAll gets all clients
func (r *RepoMysql) LoadSet(page, perPage uint64, name, nick, doc, phone, email string, set port.ClientSet) error {
	query, args := r.clientLoadSetQuery(page, perPage, name, nick, doc, phone, email)
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
	if err := r.clientLoadSetInterate(row, set); err != nil {
		return err
	}
	return nil
}

// GetById gets a client by id
func (r *RepoMysql) GetById(id string, client port.Client) (bool, error) {
	q := dbQueries["clientGetById"]
	row := r.db.QueryRow(q, id)
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

// GetByNick gets a client by nick
func (r *RepoMysql) GetByNick(nick string, client port.Client) (bool, error) {
	q := dbQueries["clientGetByNick"]
	row := r.db.QueryRow(q, nick)
	var id, rnick, name, email string
	var doc, phone uint64
	if err := row.Scan(&id, &name, &rnick, &doc, &phone, &email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	client.Load(id, name, rnick, doc, phone, email)
	return true, nil
}

// GetByEmail gets a client by email
func (r *RepoMysql) GetByEmail(email string, client port.Client) (bool, error) {
	q := dbQueries["clientGetByEmail"]
	row := r.db.QueryRow(q, email)
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

// GetByDoc gets a client by doc
func (r *RepoMysql) GetByDoc(doc uint64, client port.Client) (bool, error) {
	q := dbQueries["clientGetByDoc"]
	row := r.db.QueryRow(q, doc)
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

// GetByPhone gets a client by phone
func (r *RepoMysql) GetByPhone(phone uint64, client port.Client) (bool, error) {
	q := dbQueries["clientGetByPhone"]
	row := r.db.QueryRow(q, phone)
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
func (r *RepoMysql) Truncate() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := dbQueries["clientTruncateQuery"]
	_, err = tx.Exec(q)
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

// clientLoadSetQuery prepate the query for Load Set
func (r *RepoMysql) clientLoadSetQuery(page, perPage uint64, name, nick, doc, phone, email string) (string, []interface{}) {
	query := dbQueries["clientSetBase"]
	q, args := r.clientLoadSetFilters(name, nick, doc, phone, email)
	query += q
	q, a := r.clientLoadSetPagination(page, perPage)
	query += q
	args = append(args, a...)
	return query, args

}

// clientLoadSetInterate iterates over the rows and append to the set
func (r *RepoMysql) clientLoadSetInterate(row *sql.Rows, set port.ClientSet) error {
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
func (r *RepoMysql) clientLoadSetFilters(name, nick, doc, phone, email string) (string, []interface{}) {
	query := ""
	args := make([]interface{}, 0)
	if name != "" {
		query += dbQueries["clientSetFilterName"]
		args = append(args, "%"+name+"%")
	}
	if nick != "" {
		query += dbQueries["clientSetFilterNick"]
		args = append(args, "%"+nick+"%")
	}
	if doc != "" {
		query += dbQueries["clientSetFilterDoc"]
		args = append(args, "%"+doc+"%")
	}
	if phone != "" {
		query += dbQueries["clientSetFilterPhone"]
		args = append(args, "%"+phone+"%")
	}
	if email != "" {
		query += dbQueries["clientSetFilterEmail"]
		args = append(args, "%"+email+"%")
	}
	return query, args
}

// clientLoadSetPagination prepate the pagination query Load
func (r *RepoMysql) clientLoadSetPagination(page, perPage uint64) (string, []interface{}) {
	query := ""
	args := make([]interface{}, 0)
	// Pagination
	query += dbQueries["clientSetPagination"]
	args = append(args, perPage)
	args = append(args, (page-1)*perPage)
	return query, args
}
