package domain

import (
	"strings"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/kerror"
)

// Base represents the base of the model
type Base struct {
	BusinessID string    `json:"-"   gorm:"primaryKey;type:varchar(50); not null"`
	ID         string    `json:"id"  gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time `json:"-"   gorm:"type:timestamp; not null"`
	Updated_at time.Time `json:"-"   gorm:"type:timestamp; not null"`
}

// SetCreate sets the created_at and updated_at of the model
func (b *Base) SetCreate(business_id string) {
	b.BusinessID = business_id
	b.Created_at = time.Now()
	b.Updated_at = time.Now()
	b.Fit()
}

// Validate validates the base of the model
func (b *Base) Validate(repo port.Repository) *kerror.KError {
	valOrder := []func(repo port.Repository) *kerror.KError{
		b.ValidateID,
		b.ValidateBusinessID,
		b.ValidateCreated_at,
		b.ValidateUpdated_at,
	}
	return ValidateLoop(valOrder, repo)
}

// Fit fits the base of the model
func (b *Base) Fit() {
	b.ID = strings.TrimSpace(b.ID)
	b.BusinessID = strings.TrimSpace(b.BusinessID)
}

// ValidateBusinessID validates the business id of the model
func (b *Base) ValidateBusinessID(repo port.Repository) *kerror.KError {
	if b.BusinessID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrBaseBusinessIDIsRequired)
	}
	if len(strings.Split(b.BusinessID, " ")) > 1 {
		return kerror.NewKError(kerror.BadRequest, ErrBaseBusinessIDLength)
	}
	if strings.ToLower(b.BusinessID) != b.BusinessID {
		return kerror.NewKError(kerror.BadRequest, ErrBaseBusinessIDLower)
	}
	return nil
}

// ValidateID validates the id of the model
func (b *Base) ValidateID(repo port.Repository) *kerror.KError {
	if b.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrBaseIDIsRequired)
	}
	if len(strings.Split(b.ID, " ")) > 1 {
		return kerror.NewKError(kerror.BadRequest, ErrBaseIDLength)
	}
	if strings.ToLower(b.ID) != b.ID {
		return kerror.NewKError(kerror.BadRequest, ErrBaseIDLower)
	}
	return nil
}

// Validate Created_at validates the created_at of the model
func (b *Base) ValidateCreated_at(repo port.Repository) *kerror.KError {
	if b.Created_at.IsZero() {
		return kerror.NewKError(kerror.BadRequest, ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// Validate Updated_at validates the updated_at of the model
func (b *Base) ValidateUpdated_at(repo port.Repository) *kerror.KError {
	if b.Updated_at.IsZero() {
		return kerror.NewKError(kerror.BadRequest, ErrBaseCreatedAtIsRequired)
	}
	return nil
}

// ValidateLoopP is a function that pass a slice of validation functions and execute them in order
func ValidateLoop(orderExec []func(repo port.Repository) *kerror.KError, repo port.Repository) *kerror.KError {
	jerr := kerror.NewKError(kerror.None, "")
	for _, val := range orderExec {
		if err := val(repo); err != nil {
			jerr.Join(err.GetType(), err.Error())
		}
	}
	if jerr.IsEmpty() {
		return nil
	}
	return jerr
}

// GetBase returns a new base object
func GetDomain() []interface{} {
	return []interface{}{
		&Client{},
		&Instruction{},
		&Product{},
		&Invoice{},
		&Item{},
	}
}

// ValidateDuplicity validates the duplicity of the model
func (b *Base) ValidateDuplicity(base interface{}, repo port.Repository) *kerror.KError {
	exists, err := repo.Exists(base, b.BusinessID, b.ID)
	if err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	}
	if exists {
		return kerror.NewKError(kerror.Conflict, ErrBaseIDAlreadyExists)
	}
	return nil
}
