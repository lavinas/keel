package domain

import (
	"strings"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrVariableNameIsRequired  = "variable name is required"
	ErrVariableNameLength      = "variable name must have less than 50 characters"
	ErrVariableNameCompose     = "variable name must be composed by one word"
	ErrVariableValueIsRequired = "variable value is required"
)

// Variable is the struct that contains the variables for email template
type Variable struct {
	Base
	EmailID string `json:"-" gorm:"type:varchar(50); not null"`
	Name    string `json:"name"  gorm:"type:varchar(50); not null"`
	Value   string `json:"value" gorm:"type:varchar(500); not null"`
}

// SetCreate set information for create a new variable item
func (v *Variable) SetCreate() {
	v.Base.SetCreate(true)
}

// Validate validate the variable information
func (v *Variable) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		v.Base.Validate,
		v.ValidateName,
		v.ValidateValue,
	})
}

// ValidateName validate the variable name
func (v *Variable) ValidateName() *kerror.KError {
	if v.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrVariableNameIsRequired)
	}
	if len(v.Name) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrVariableNameLength)
	}
	if len(strings.Split(v.Name, " ")[0]) > 1 {
		return kerror.NewKError(kerror.BadRequest, ErrVariableNameCompose)
	}
	return nil
}

// ValidateValue validate the variable value
func (v *Variable) ValidateValue() *kerror.KError {
	if v.Value == "" {
		return kerror.NewKError(kerror.BadRequest, ErrVariableValueIsRequired)
	}
	if len(v.Value) > 500 {
		return kerror.NewKError(kerror.BadRequest, ErrVariableNameLength)
	}
	return nil
}

// GetResult returns the variable value
func (b *Variable) GetResult() string {
	return b.Name + ": " + b.Value
}

// TableName returns the table name for gorm
func (b *Variable) TableName() string {
	return "variable"
}
