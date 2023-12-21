package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrReceiverNameLength      = "name must have less than 50 characters"
	ErrReceiverEmailIsRequired = "email is required"
	ErrReceiverEmailLength     = "email must have less than 50 characters"
)

// Receiver is the struct that contains the client information
type Receiver struct {
	Base
	Name  string `json:"name"  gorm:"type:varchar(50); not null"`
	Email string `json:"email" gorm:"type:varchar(50); not null"`
}

// SetCreate set information for create a new client
func (r *Receiver) SetCreate() {
	r.Base.SetCreate(false)
}

// Validate validate the client information
func (r *Receiver) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		r.Base.Validate,
		r.ValidateName,
		r.ValidateEmail,
		r.Base.Validate,
		r.ValidateDuplicity,
	})
}

// ValidateName validate the client information
func (r *Receiver) ValidateName() *kerror.KError {
	if r.Name == "" {
		return nil
	}
	if len(r.Name) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrReceiverNameLength)
	}
	return nil
}

// ValidateEmail validate the client information
func (r *Receiver) ValidateEmail() *kerror.KError {
	if r.Email == "" {
		return kerror.NewKError(kerror.BadRequest, ErrReceiverEmailIsRequired)
	}
	if len(r.Email) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrReceiverEmailLength)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (r *Receiver) ValidateDuplicity() *kerror.KError {
	return r.Base.ValidateDuplicity(r)
}

// GetByID returns the model by its ID
func (r *Receiver) GetByID() *kerror.KError {
	return r.Base.GetByID(r)
}

// GetResult returns the result that is the receiver itself
func (r *Receiver) GetResult() interface{} {
	return r
}

// TableName returns the table name for gorm
func (b *Receiver) TableName() string {
	return "receiver"
}
