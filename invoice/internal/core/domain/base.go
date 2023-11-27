package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Base represents the base of the model
type Base struct {
	BusinessID string    `json:"-"            gorm:"primaryKey;type:varchar(50); not null"`
	ID         string    `json:"id"           gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time `json:"created_at"   gorm:"type:timestamp; not null"`
	Updated_at time.Time `json:"updated_at"   gorm:"type:timestamp; not null"`
}

// Validate validates the base of the model
func (b *Base) Validate(repo port.Repository) error {
	valOrder := []func(repo port.Repository) error{
		b.ValidateID,
		b.ValidateBusinessID,
		b.ValidateCreated_at,
		b.ValidateUpdated_at,
	}
	return ValidateLoop(valOrder, repo)
}

// SetBusinessID sets the business id of the model
func (b *Base) SetBusinessID(business_id string) {
	b.BusinessID = business_id
}

// SetID sets the id of the model
func (b *Base) SetCreatedAt(date time.Time) {
	b.Created_at = date
}

// SetID sets the id of the model
func (b *Base) SetUpdatedAt(date time.Time) {
	b.Updated_at = date
}

// Marshal marshals the base of the model
func (b *Base) Marshal() error {
	return nil
}

// GetBusinessID gets the business id of the model
func (b *Base) GetID() string {
	return b.ID
}

// ValidateBusinessID validates the business id of the model
func (b *Base) ValidateBusinessID(repo port.Repository) error {
	if b.BusinessID == "" {
		return errors.New(ErrBaseBusinessIDIsRequired)
	}
	if len(strings.Split(b.BusinessID, " ")) > 1 {
		return errors.New(ErrBaseBusinessIDLength)
	}
	if strings.ToLower(b.BusinessID) != b.BusinessID {
		return errors.New(ErrBaseBusinessIDLower)
	}
	return nil
}

// ValidateID validates the id of the model
func (b *Base) ValidateID(repo port.Repository) error {
	if b.ID == "" {
		return errors.New(ErrBaseIDIsRequired)
	}
	if len(strings.Split(b.ID, " ")) > 1 {
		return errors.New(ErrBaseIDLength)
	}
	if strings.ToLower(b.ID) != b.ID {
		return errors.New(ErrBaseIDLower)
	}
	return nil
}

// Validate Created_at validates the created_at of the model
func (b *Base) ValidateCreated_at(repo port.Repository) error {
	if b.Created_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// Validate Updated_at validates the updated_at of the model
func (b *Base) ValidateUpdated_at(repo port.Repository) error {
	if b.Updated_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// ValidateLoopP is a function that pass a slice of validation functions and execute them in order
func ValidateLoop(orderExec []func(repo port.Repository) error, repo port.Repository) error {
	errMsg := ""
	for _, val := range orderExec {
		if err := val(repo); err != nil {
			errMsg += err.Error() + " | "
		}
	}
	if errMsg != "" {
		errMsg = strings.TrimSuffix(errMsg, " | ")
		return errors.New(errMsg)
	}
	return nil
}
