package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrTemplateNameIsRequired    = "name is required"
	ErrTemplateSubjectIsRequired = "subject is required"
	ErrTemplateSubjectLength     = "subject must have less than 50 characters"
	ErrTemplateBodyIsRequired    = "body is required"
	ErrTemplateBodyLength        = "body must have less than 50 characters"
)

// Template is the struct that contains the email template information
type Template struct {
	Base
	Subject string `json:"subject" gorm:"type:varchar(50); not null"`
	Body    string `json:"body"    gorm:"type:varchar(50); not null"`
}

// SetCreate set information for create a new email template
func (t *Template) SetCreate() {
	t.Base.SetShortenID(t, t.Subject)
	t.Base.SetCreate()
}

// Validate validate the email template information
func (t *Template) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		t.Base.Validate,
		t.ValidateSubject,
		t.ValidateBody,
		t.ValidateDuplicity,
	})
}

// ValidateSubject validate the email template information
func (t *Template) ValidateSubject() *kerror.KError {
	if t.Subject == "" {
		return kerror.NewKError(kerror.BadRequest, ErrTemplateSubjectIsRequired)
	}
	if len(t.Subject) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrTemplateSubjectLength)
	}
	return nil
}

// ValidateBody validate the email template information
func (t *Template) ValidateBody() *kerror.KError {
	if t.Body == "" {
		return kerror.NewKError(kerror.BadRequest, ErrTemplateBodyIsRequired)
	}
	if len(t.Body) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrTemplateBodyLength)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (r *Template) ValidateDuplicity() *kerror.KError {
	return r.Base.ValidateDuplicity(r)
}

// GetByID returns the model by its ID
func (r *Template) GetByID() *kerror.KError {
	return r.Base.GetByID(r)
}

// GetResult returns the result that is the template itself
func (t *Template) GetResult() interface{} {
	return t
}

// TableName returns the table name for gorm
func (b *Template) TableName() string {
	return "template"
}
