package repository

import (
	"strings"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_TYPE = "DB_TYPE"
	DB_DNS  = "DB_DNS"
)

// RepoMySql is the repository handler for the application
type MySql struct {
	Db *gorm.DB
}

// NewRepository creates a new repository handler
func NewRepository(config port.Config) (*MySql, error) {
	db, err := gorm.Open(mysql.Open(config.Get(DB_DNS)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&domain.Product{},
		&domain.Instruction{},
		&domain.Client{},
		&domain.Invoice{},
	)
	return &MySql{Db: db}, nil
}

// Close closes the database connection
func (r *MySql) Close() {
}

// Add adds a object to the database
func (r *MySql) Add(obj interface{}) error {
	return r.Db.Create(obj).Error
}

func (r *MySql) IsDuplicatedError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062")
}
