package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Base represents the base of the model
type Base struct {
	BusinnessID string    `json:"-"            gorm:"primaryKey;type:varchar(50); not null"`
	ID          string    `json:"id"           gorm:"primaryKey;type:varchar(50); not null"`
	Created_at  time.Time `json:"created_at"   gorm:"type:timestamp; not null"`
	Updated_at  time.Time `json:"updated_at"   gorm:"type:timestamp; not null"`
}

func NewBase(businness_id, id string, created_at time.Time, updated_at time.Time) Base {
	if id == "" {
		id = uuid.New().String()
	}
	if created_at.IsZero() {
		created_at = time.Now()
	}
	if updated_at.IsZero() {
		updated_at = time.Now()
	}
	return Base{
		BusinnessID: businness_id,
		ID:          id,
		Created_at:  created_at,
		Updated_at:  updated_at,
	}
}

// Validate validates the base of the model
func (b *Base) Validate() error {
	valOrder := []func() error{
		b.ValidateID,
	}
	errMsg := ""
	for _, val := range valOrder {
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

// SetBusinessID sets the business id of the model
func (b *Base) SetBusinessID(businness_id string) {
	b.BusinnessID = businness_id
}

// Marshal marshals the base of the model
func (b *Base) Marshal() error {
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
