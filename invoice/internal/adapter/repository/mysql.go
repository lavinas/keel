package repository

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
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
	db, err := gorm.Open(config.Get(DB_TYPE), config.Get(DB_DNS))
	if err != nil {
		return nil, err
	}
	db.LogMode(false)
	db.AutoMigrate(&domain.Client{}, &domain.Instruction{}, &domain.Product{})
	return &MySql{Db: db}, nil
}

// Close closes the database connection
func (r *MySql) Close() {
	r.Db.Close()
}

// Add adds a object to the database
func (r *MySql) Add(obj interface{}) error {
	return r.Db.Create(obj).Error
}

func (r *MySql) IsDuplicatedError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062")
}
