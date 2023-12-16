package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
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
	ID         string    `json:"id" gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time `json:"-"  gorm:"type:timestamp; not null"`
	Updated_at time.Time `json:"-"  gorm:"type:timestamp; not null"`
}

// SetCreate set information for create a new client
func (b *Base) SetCreate() {
	if b.ID == "" {
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

// ValidateLoop validate the base information
func validateLoop(val []func() *kerror.KError) *kerror.KError {
	err := kerror.NewKError(kerror.None, "")
	for _, f := range val {
		err.JoinKError(f())
	}
	if err.IsEmpty() {
		return nil
	}
	return err
}

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
