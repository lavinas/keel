package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

const (
	mysql_user   = "MYSQL_INVOICE_USER"
	mysql_pass   = "MYSQL_INVOICE_PASSWORD"
	mysql_host   = "MYSQL_INVOICE_HOST"
	mysql_port   = "MYSQL_INVOICE_PORT"
	mysql_dbname = "MYSQL_INVOICE_DATABASE"
)

// Repo is a service to interact with the database Mysql
type RepoMysql struct {
	db *sql.DB
	tx *sql.Tx
}

// NewRepo creates a new Repo service
func NewRepoMysql() (*RepoMysql, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv(mysql_user), os.Getenv(mysql_pass), os.Getenv(mysql_host), os.Getenv(mysql_port))
	db, _ := sql.Open("mysql", conn)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	dbname := os.Getenv(mysql_dbname)
	if dbname == "" {
		return nil, errors.New("MYSQL_DATABASE is empty")
	}
	for i, q := range querieMap {
		querieMap[i] = strings.Replace(q, "{DB}", dbname, -1)
	}
	return &RepoMysql{db: db}, nil
}

// Begin starts a transaction
func (r *RepoMysql) Begin() error {
	if r.tx != nil {
		return errors.New("transaction already started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	tx, err := r.db.Begin()
	if err != nil {
		r.tx = nil
		return err
	}
	r.tx = tx
	return nil
}

// Commit commits a transaction
func (r *RepoMysql) Commit() error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	defer func() { r.tx = nil }()
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	err := r.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Rollback rollbacks a transaction
func (r *RepoMysql) Rollback() error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	defer func() { r.tx = nil }()
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	err := r.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

// IsDuplicatedInvoice checks if the invoice is duplicated
func (r *RepoMysql) IsDuplicatedInvoice(reference string) (bool, error) {
	if r.db == nil {
		return false, errors.New("sql: database is closed")
	}
	q := querieMap["IsDuplicatedInvoice"]
	var count int
	err := r.db.QueryRow(q, reference).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// SaveInvoiceClient stores the invoice client on the repository
func (r *RepoMysql) SaveInvoiceClient(client port.InvoiceClient) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["SaveInvoiceClient"]
	c := client
	_, err := r.tx.Exec(q, c.GetId(), c.GetNickname(), c.GetClientId(),
		c.GetName(), c.GetDocument(), c.GetPhone(), c.GetEmail())
	if err != nil {
		return err
	}
	return nil
}

// UpdateInvoiceClient updates the invoice client on the repository
func (r *RepoMysql) UpdateInvoiceClient(client port.InvoiceClient) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["UpadateInvoiceClient"]
	c := client
	_, err := r.tx.Exec(q, c.GetNickname(), c.GetClientId(), c.GetName(), 
							c.GetDocument(), c.GetPhone(), c.GetEmail(), c.GetId())
	if err != nil {
		return err
	}
	return nil
}

// SaveInvoice stores the invoice on the repository
func (r *RepoMysql) SaveInvoice(invoice port.Invoice) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["SaveInvoice"]
	i := invoice
	_, err := r.tx.Exec(q, i.GetId(), i.GetReference(), i.GetBusinessId(), i.GetCustomerId(),
		i.GetAmount(), i.GetDate(), i.GetDue(), i.GetNoteId(), i.GetStatusId(),
		i.GetCreatedAt(), i.GetUpdatedAt())
	if err != nil {
		return err
	}
	return nil
}

// SaveInvoiceItem stores the invoice item on the repository
func (r *RepoMysql) SaveInvoiceItem(item port.InvoiceItem) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["SaveInvoiceItem"]
	i := item
	_, err := r.tx.Exec(q, i.GetId(), i.GetInvoiceId(), i.GetServiceReference(),
		i.GetDescription(), i.GetAmount(), i.GetQuantity())
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database connection
func (r *RepoMysql) Close() error {
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	err := r.db.Close()
	r.db = nil
	return err
}

// TruncateInvoiceClient cleans the invoice client table
func (r *RepoMysql) TruncateInvoiceItem() error {
	return r.truncate("TruncateInvoiceItem")
}

// TruncateInvoiceClient cleans the invoice client table
func (r *RepoMysql) TruncateInvoice() error {
	return r.truncate("TruncateInvoice")
}

// TruncateInvoiceClient cleans the invoice client table
func (r *RepoMysql) TruncateInvoiceClient() error {
	return r.truncate("TruncateInvoiceClient")
}

// truncate cleans a table
func (r *RepoMysql) truncate(querieName string) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	_, err := r.tx.Exec(querieMap[querieName])
	if err != nil {
		return err
	}
	return nil
}
