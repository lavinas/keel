package repository

import (
	"errors"

	"github.com/lavinas/keel/internal/invoice/core/domain"
	"github.com/lavinas/keel/internal/invoice/core/port"
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
	// Open
	db, err := Open(config)
	if err != nil {
		return nil, err
	}
	// Migrate
	m := &MySql{Db: db}
	if err := m.Migrate(); err != nil {
		return nil, err
	}
	// Return
	return m, nil
}

// Open opens the database connection
func Open(config port.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Get(DB_DNS)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

// Migrate migrates the database
func (r *MySql) Migrate() error {
	for _, domain := range domain.GetDomain() {
		if err := r.Db.AutoMigrate(domain); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the database connection
func (r *MySql) Close() {
}

// Add adds a object to the database
func (r *MySql) Add(obj interface{}) error {
	return r.Db.Create(obj).Error
}

// FindByID finds a object by id
func (r *MySql) Exists(obj interface{}, business_id string, id string) (bool, error) {
	tx := r.Db.First(obj, "business_id = ? and ID = ?", business_id, id)
	if tx.Error == nil {
		return true, nil
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, tx.Error
}
