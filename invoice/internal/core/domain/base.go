package domain

import (
	"errors"
	"strings"
	"time"
)

// Base represents the base of the model
type Base struct {
	BusinessID string    `json:"-"            gorm:"primaryKey;type:varchar(50); not null"`
	ID         string    `json:"id"           gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time `json:"created_at"   gorm:"type:timestamp; not null"`
	Updated_at time.Time `json:"updated_at"   gorm:"type:timestamp; not null"`
}

// Validate validates the base of the model
func (b *Base) Validate() error {
	valOrder := []func() error{
		b.ValidateID,
		b.ValidateBusinessID,
		b.ValidateCreated_at,
		b.ValidateUpdated_at,
	}
	return ValidateLoop(valOrder)
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
func (b *Base) ValidateBusinessID() error {
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
func (p *Base) ValidateID() error {
	if p.ID == "" {
		return errors.New(ErrBaseIDIsRequired)
	}
	if len(strings.Split(p.ID, " ")) > 1 {
		return errors.New(ErrBaseIDLength)
	}
	if strings.ToLower(p.ID) != p.ID {
		return errors.New(ErrBaseIDLower)
	}
	return nil
}

// Validate Created_at validates the created_at of the model
func (p *Base) ValidateCreated_at() error {
	if p.Created_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// Validate Updated_at validates the updated_at of the model
func (p *Base) ValidateUpdated_at() error {
	if p.Updated_at.IsZero() {
		return errors.New(ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// ValidateLoop is a function that pass a slice of validation functions and execute them in order
func ValidateLoop(orderExec []func() error) error {
	errMsg := ""
	for _, val := range orderExec {
		if err := val(); err != nil {
			errMsg += err.Error() + " | "
		}
	}
	if errMsg != "" {
		errMsg = strings.TrimSuffix(errMsg, " | ")
		return errors.New(errMsg)
	}
	return nil
}
