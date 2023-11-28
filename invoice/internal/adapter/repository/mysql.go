package repository

import (
	"errors"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(mysql.Open(config.Get(DB_DNS)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&domain.Product{},
		&domain.Instruction{},
		&domain.Client{},
		&domain.Invoice{},
		&domain.InvoiceItem{},
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

// FindByID finds a object by id
func (r *MySql) Exists(obj interface{}, id string) bool {
	tx := r.Db.First(obj, "ID = ?", id)
	return !errors.Is(tx.Error, gorm.ErrRecordNotFound)
}
