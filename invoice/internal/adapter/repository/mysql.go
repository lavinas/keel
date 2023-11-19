package repository

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

// RepoMySql is the repository handler for the application
type MySql struct {
	Db *gorm.DB
}

// NewRepository creates a new repository handler
func NewRepository(config port.Config) (*MySql, error) {
	db, err := connectDB(config)
	if err != nil {
		return nil, err
	}
	return &MySql{Db: db}, nil
}

// AddClient adds a new client to the database
func (r *MySql) AddClient(client *domain.Client) error {
	return r.Db.Create(client).Error
}

// Close closes the database connection
func (r *MySql) Close() {
	r.Db.Close()
}

// ConnectDB connects to the database
func connectDB(config port.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("dbname=%s sslmode=%s user=%s password=%s host=%s", config.Get(MYSQL_DATABASE),
		config.Get(MYSQL_SSLMODE), config.Get(MYSQL_USER), config.Get(MYSQL_PASSWORD), config.Get(MYSQL_HOST))
	db, err := gorm.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}
