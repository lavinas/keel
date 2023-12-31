package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrSenderNameLength      = "name must have less than 50 characters"
	ErrSenderEmailIsRequired = "email is required"
	ErrSenderEmailLength     = "email must have less than 50 characters"
)

// Sender is the struct that contains the business information
type Sender struct {
	Base
	Name  string `json:"name"  gorm:"type:varchar(50); not null"`
	Email string `json:"email" gorm:"type:varchar(50); not null"`
}

// SetCreate set information for create a new business
func (s *Sender) SetCreate() {
	s.Base.SetShortenID(s, s.Name)
	s.Base.SetCreate()
}

// Validate validate the business information
func (s *Sender) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		s.Base.Validate,
		s.ValidateName,
		s.ValidateEmail,
		s.Base.Validate,
		s.ValidateDuplicity,
	})
}

// ValidateName validate the business information
func (s *Sender) ValidateName() *kerror.KError {
	if s.Name == "" {
		return nil
	}
	if len(s.Name) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrSenderNameLength)
	}
	return nil
}

// ValidateEmail validate the business information
func (s *Sender) ValidateEmail() *kerror.KError {
	if s.Email == "" {
		return kerror.NewKError(kerror.BadRequest, ErrSenderEmailIsRequired)
	}
	if len(s.Email) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrSenderEmailLength)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (r *Sender) ValidateDuplicity() *kerror.KError {
	return r.Base.ValidateDuplicity(r)
}

// GetByID returns the model by its ID
func (r *Sender) GetByID() *kerror.KError {
	return r.Base.GetByID(r)
}

// GetResult returns the result that is the business itself
func (s *Sender) GetResult() interface{} {
	return s
}

// TableName returns the table name for gorm
func (b *Sender) TableName() string {
	return "sender"
}
