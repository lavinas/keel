package adapter

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

const (
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_DATABASE = "MYSQL_DATABASE"
	MYSQL_SSLMODE  = "MYSQL_SSLMODE"
)

// Repository is the repository handler for the application
type Repository struct {
	Db *gorm.DB
}

// NewRepository creates a new repository handler
func NewRepository(config port.Config) (*Repository, error) {
	db, err := ConnectDB(config)
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}

// AddClient adds a new client to the database
func (r *Repository) AddClient(client *domain.Client) error {
	return r.Db.Create(client).Error
}

// Close closes the database connection
func (r *Repository) Close() {
	r.Db.Close()
}

// ConnectDB connects to the database
func ConnectDB(config port.Config) (*gorm.DB, error) {
	host := config.Get(MYSQL_HOST)
	if host == "" {
		host = "localhost"
	}
	user := config.Get(MYSQL_USER)
	if user == "" {
		user = "root"
	}
	pass := config.Get(MYSQL_PASSWORD)
	if pass == "" {
		pass = "root"
	}
	dab := config.Get(MYSQL_DATABASE)
	if dab == "" {
		dab = "keel_invoice"
	}
	mode := config.Get(MYSQL_SSLMODE)
	if mode == "" {
		mode = "disable"
	}
	dns := fmt.Sprintf("dbname=%s sslmode=%s user=%s password=%s host=%s", dab, mode, user, pass, host)
	db, err := gorm.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}
