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
func (b *Base) Validate(p interface{}) error {
	valOrder := []func(interface{}) error{
		b.ValidateID,
		b.ValidateBusinessID,
		b.ValidateCreated_at,
		b.ValidateUpdated_at,
	}
	repo := p.(port.Repository)
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

// ValidateBusinessID validates the business id of the model
func (b *Base) ValidateBusinessID(p interface{}) error {
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
func (b *Base) ValidateID(p interface{}) error {
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
func (b *Base) ValidateCreated_at(p interface{}) error {
	if b.Created_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// Validate Updated_at validates the updated_at of the model
func (b *Base) ValidateUpdated_at(p interface{}) error {
	if b.Updated_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the client
func (b *Base) ValidateDuplicity(p interface{}) error {
	if b.ID == "" {
		return nil
	}
	repo := p.(port.Repository)
	if repo.FindByID(b, b.ID) {
		return errors.New(ErrDuplicatedID)
	}
	return nil
}

// ValidateLoopP is a function that pass a slice of validation functions and execute them in order
func ValidateLoop(orderExec []func(interface{}) error, p interface{}) error {
	errMsg := ""
	for _, val := range orderExec {
		if err := val(p); err != nil {
			errMsg += err.Error() + " | "
		}
	}
	if errMsg != "" {
		errMsg = strings.TrimSuffix(errMsg, " | ")
		return errors.New(errMsg)
	}
	return nil
}