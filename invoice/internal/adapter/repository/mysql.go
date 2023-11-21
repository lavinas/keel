package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"

)

const (
	DB_TYPE        = "DB_HOST"
	DB_DNS         = "DB_DNS"
)

// RepoMySql is the repository handler for the application
type MySql struct {
	Db *gorm.DB
}

// NewRepository creates a new repository handler
func NewRepository(config port.Config) (*MySql, error) {
	db, err := gorm.Open(config.Get(DB_TYPE), config.Get(DB_DNS))
	if err != nil {
		return nil, err
	}
	return &MySql{Db: db}, nil
}

// Close closes the database connection
func (r *MySql) Close() {
	r.Db.Close()
}

// AddClient adds a new client to the database
func (r *MySql) AddClient(client *domain.Client) error {
	return r.Db.Create(client).Error
}

