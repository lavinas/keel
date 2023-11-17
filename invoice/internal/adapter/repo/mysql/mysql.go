package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

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
func NewRepoMysql(config port.Config) (*RepoMysql, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Get(mysql_user), config.Get(mysql_pass),
		config.Get(mysql_host), config.Get(mysql_port))
	db, _ := sql.Open("mysql", conn)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	dbname := config.Get(mysql_dbname)
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
		c.GetName(), c.GetDocument(), c.GetPhone(), c.GetEmail(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

// GetLastInvoiceClient gets the last invoice client on the repository
func (r *RepoMysql) GetLastInvoiceClient(nickname string, created_after time.Time, client port.InvoiceClient) (bool, error) {
	if r.db == nil {
		return false, errors.New("sql: database is closed")
	}
	q := querieMap["GetInvoiceClient"]
	var rId, rNickname, rClientId, rName, rEmail string
	var rDocument, rPhone uint64
	var rCreatedAt []uint8
	err := r.db.QueryRow(q, nickname, created_after).Scan(&rId, &rNickname, &rClientId,
		&rName, &rDocument, &rPhone, &rEmail, &rCreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	rCreatedAtTime, err := time.Parse("2006-01-02 15:04:05", string(rCreatedAt))
	if err != nil {
		return false, err
	}
	client.Load(rId, rNickname, rClientId, rName, rEmail, rDocument, rPhone, rCreatedAtTime)
	return true, nil
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
		i.GetAmount(), i.GetDate(), i.GetDue(), i.GetNoteId(), i.GetCreatedAt(), i.GetUpdatedAt())
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

// LoadInvoiceVertex loads the invoice status graph vertex
func (r *RepoMysql) GetInvoiceVertex(graph port.InvoiceStatus) error {
	q := querieMap["GetInvoiceVertex"]
	rows, err := r.db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var class, id, name, description string
		err = rows.Scan(&class, &id, &name, &description)
		if err != nil {
			return err
		}
		graph.AddVertex(class, id, name, description)
	}
	return nil
}

// LoadInvoiceEdge loads the invoice status graph edge
func (r *RepoMysql) GetInvoiceEdge(graph port.InvoiceStatus) error {
	q := querieMap["GetInvoiceEdge"]
	rows, err := r.db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var class, vertexFrom, vertexTo, description string
		err = rows.Scan(&class, &vertexFrom, &vertexTo, &description)
		if err != nil {
			return err
		}
		graph.AddEdge(class, vertexFrom, vertexTo, description)
	}
	return nil
}

// LogInvoiceEdge logs the invoice status graph edge
func (r *RepoMysql) CreateInvoiceStatusLog(class string, graph port.InvoiceStatus) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["CreateInvoiceStatusLog"]
	invoice_id := graph.GetInvoiceId()
	for {
		next, id, from, to, description, author, createdAt := graph.DequeueEdge(class)
		if !next {
			return nil
		}
		_, err := r.tx.Exec(q, id, invoice_id, class, from, to, createdAt, description, author)
		if err != nil {
			return err
		}
	}
}

// StoreInvoiceStatus stores the invoice status graph. If exists, updates. If not, creates.
func (r *RepoMysql) StoreInvoiceStatus(class string, graph port.InvoiceStatus) error {
	if r.tx == nil {
		return errors.New("transaction not started")
	}
	if r.db == nil {
		return errors.New("sql: database is closed")
	}
	q := querieMap["UpdateInvoiceStatus"]
	count, err := r.tx.Exec(q, graph.GetLastStatusId(class), graph.GetInvoiceId(), class)
	if err != nil {
		return err
	}
	if c, _ := count.RowsAffected(); c == 0 {
		q := querieMap["CreateInvoiceStatus"]
		_, err := r.tx.Exec(q, graph.GetInvoiceId(), class, graph.GetLastStatusId(class))
		if err != nil {
			return err
		}
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

// Truncate deletes the database operational data
func (r *RepoMysql) Truncate() error {
	tables := []string{
		"invoice_item",
		"invoice_status_log",
		"invoice_status",
		"invoice_payment",
		"invoice_delivery",
		"invoice",
		"invoice_client",
		"invoice_note",
	}
	if err := r.Begin(); err != nil {
		return err
	}
	defer r.Rollback()
	for _, table := range tables {
		q := strings.Replace(querieMap["Truncate"], "{TABLE}", table, -1)
		_, err := r.tx.Exec(q)
		if err != nil {
			return err
		}
	}
	if err := r.Commit(); err != nil {
		return err
	}
	return nil
}
