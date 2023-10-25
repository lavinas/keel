package repo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lavinas/keel/internal/invoice/core/port"
)

const (
	mysql_user   = "MYSQL_USER"
	mysql_pass   = "MYSQL_PASSWORD"
	mysql_host   = "MYSQL_HOST"
	mysql_port   = "MYSQL_PORT"
	mysql_dbname = "MYSQL_DATABASE"
)

// Repo is a service to interact with the database Mysql
type RepoMysql struct {
	db *sql.DB
}

// NewRepo creates a new Repo service
func NewRepoMysql() *RepoMysql {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv(mysql_user), os.Getenv(mysql_pass), os.Getenv(mysql_pass),
		os.Getenv(mysql_port), os.Getenv(mysql_dbname))
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	return &RepoMysql{db: db}
}

// SaveInvoiceClient stores the invoice client on the repository
func (r *RepoMysql) SaveInvoiceClient(client port.InvoiceClient) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	c := client
	_, err = tx.Exec(SaveInvoiceClient, c.GetId(), c.GetNickname(), c.GetClientId(),
		c.GetName(), c.GetDocument(), c.GetPhone(), c.GetEmail())
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// SaveInvoice stores the invoice on the repository
func (r *RepoMysql) SaveInvoice(invoice port.Invoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	i := invoice
	_, err = tx.Exec(SaveInvoice, i.GetId(), i.GetReference(), i.GetBusinessId(), i.GetCustomerId(),
		i.GetAmount(), i.GetDate(), i.GetDue(), i.GetNoteId(), i.GetStatusId(),
		i.GetCreatedAt(), i.GetUpdatedAt())
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// SaveInvoiceItem stores the invoice item on the repository
func (r *RepoMysql) SaveInvoiceItem(item port.InvoiceItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	i := item
	_, err = tx.Exec(SaveInvoiceItem, i.GetId(), i.GetInvoiceId(), i.GetServiceReference(),
		i.GetDescription(), i.GetAmount(), i.GetQuantity())
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// Truncate cleans the database
func (r *RepoMysql) Truncate() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(TruncateInvoiceItem)
	if err != nil {
		return err
	}
	_, err = tx.Exec(TruncateInvoice)
	if err != nil {
		return err
	}
	_, err = tx.Exec(TruncateInvoiceClient)
	if err != nil {
		return err
	}
	return nil
}