package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/lavinas/keel/internal/email/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrBaseIDIsRequired        = "id is required"
	ErrBaseIDLength            = "id must have only one word. Use underscore to separate words"
	ErrBaseIDLower             = "id must be lower case"
	ErrBaseCreatedAtIsRequired = "created_at is required"
	ErrBaseUpdatedAtIsRequired = "updated_at is required"
	ErrBaseIDAlreadyExists     = "id already exists"
)

// Base is the struct that contains the base information
type Base struct {
	Repo       port.Repository `json:"-"  gorm:"-"`
	ID         string          `json:"id" gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time       `json:"-"  gorm:"type:timestamp; not null"`
	Updated_at time.Time       `json:"-"  gorm:"type:timestamp; not null"`
}

// SetRepository set the repository for the base
func (b *Base) SetRepository(repo port.Repository) {
	b.Repo = repo
}

func (b *Base) GetRepository() port.Repository {
	return b.Repo
}

// SetCreate set information for create a new client
func (b *Base) SetCreate(genID bool) {
	if b.ID == "" && genID {
		b.ID = uuid.New().String()
	}
	b.Created_at = time.Now()
	b.Updated_at = time.Now()
}

// Validate validate the base information
func (b *Base) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		b.ValidateID,
		b.ValidateCreatedAt,
		b.ValidateUpdatedAt,
	})
}

// ValidateID validate the base information
func (b *Base) ValidateID() *kerror.KError {
	if b.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrBaseIDIsRequired)
	}
	if len(strings.Split(b.ID, " ")) > 1 {
		return kerror.NewKError(kerror.BadRequest, ErrBaseIDLength)
	}
	return nil
}

// ValidateCreatedAt validate the base information
func (b *Base) ValidateCreatedAt() *kerror.KError {
	if b.Created_at.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// ValidateUpdatedAt validate the base information
func (b *Base) ValidateUpdatedAt() *kerror.KError {
	if b.Updated_at.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrBaseUpdatedAtIsRequired)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (b *Base) ValidateDuplicity(base interface{}) *kerror.KError {
	exists, err := b.Repo.Exists(base, b.ID)
	if err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	}
	if exists {
		return kerror.NewKError(kerror.Conflict, ErrBaseIDAlreadyExists)
	}
	return nil
}

// SetID set the base information
func (b *Base) SetID(id string) {
	b.ID = id
}

// getID returns the base ID information
func (b *Base) GetID() string {
	return b.ID
}

// Get returns the base information from repository
func (b *Base) GetByID(obj interface{}) *kerror.KError {
	if b.Repo == nil {
		panic("cagou")
	}
	found, err := b.Repo.GetByID(obj)
	if err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	}
	if !found {
		return kerror.NewKError(kerror.BadRequest, "id not found")
	}
	return nil
}

// Domain functions

// GetBase returns a new base object
func GetDomain() []interface{} {
	return []interface{}{
		&Template{},
		&Sender{},
		&Receiver{},
		&SMTPServer{},
		&Status{},
		&Variable{},
		&Email{},
	}
}

// ValidateLoop validate the base information
func validateLoop(val []func() *kerror.KError) *kerror.KError {
	jerr := kerror.NewKError(kerror.None, "")
	for _, f := range val {
		jerr.JoinKError(f())
	}
	if jerr.IsEmpty() {
		return nil
	}
	return jerr
}
