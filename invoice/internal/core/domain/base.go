package domain

import (
	"errors"
	"strings"
	"time"
)

// Base represents the base of the model
type Base struct {
	ID         string    `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// Validate validates the base of the model
func (b *Base) Validate() error {
	valOrder := []func() error{
		b.ValidateID,
		b.ValidateCreatedAt,
		b.ValidateUpdatedAt,
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

// ValidateID validates the id of the model
func (p *Base) ValidateID() error {
	if p.ID == "" {
		return nil
	}
	if len(strings.Split(p.ID, " ")) > 1 {
		return errors.New(ErrBaseIDLength)
	}
	if strings.ToLower(p.ID) != p.ID {
		return errors.New(ErrBaseIDLower)
	}
	return nil
}

// ValidateCreatedAt validates the created at of the model
func (p *Base) ValidateCreatedAt() error {
	if p.Created_at.IsZero() {
		return nil
	}
	return nil
}

// ValidateUpdatedAt validates the updated at of the model
func (p *Base) ValidateUpdatedAt() error {
	if p.Updated_at.IsZero() {
		return nil
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
