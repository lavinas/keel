package domain

import (
	"fmt"
	"reflect"

	"github.com/lavinas/keel/internal/email/core/port"
	"github.com/lavinas/keel/pkg/kerror"

)

const (
	ErrEmailSenderIsRequired     = "sender is required"
	ErrEmailSenderIsTwice        = "sender id and sender is informed. Only one is allowed"
	ErrEmailSenderNotFound 	     = "sender id not found"
	ErrEmailReceiverIsRequired   = "receiver is required"
	ErrEmailReceiverIsTwice      = "receiver id and receiver is informed. Only one is allowed"
	ErrEmailTemplateIsRequired   = "template is required"
	ErrEmailTemplateTwice        = "template id and template is informed. Only one is allowed"
	ErrEmailSMTPServerIsRequired = "smtp server is required"
	ErrEmailSMTPServerTwice      = "smtp server id and smtp server is informed. Only one is allowed"
)

// Email is the struct that contains the email information
type Email struct {
	Base
	SenderID     string      `json:"sender_id"      gorm:"type:varchar(50); not null"`
	Sender       *Sender     `json:"sender"         gorm:"foreignKey:SenderID"`
	ReceiverID   string      `json:"receiver_id"    gorm:"type:varchar(50); not null"`
	Receiver     *Receiver   `json:"receiver"       gorm:"foreignKey:ReceiverID"`
	TemplateID   string      `json:"template_id"    gorm:"type:varchar(50); not null"`
	Template     *Template   `json:"template"       gorm:"foreignKey:TemplateID"`
	SMTPServerID string      `json:"smtp_server_id" gorm:"type:varchar(50); not null"`
	SMTPServer   *SMTPServer `json:"smtp_server"    gorm:"foreignKey:SMTPServerID"`
	Variables    []*Variable `json:"variables"      gorm:"foreignKey:ID;associationForeignKey:EmailID"`
	StatusID     string      `json:"-"              gorm:"type:varchar(50); not null"`
	Status       *Status     `json:"status"         gorm:"foreignKey:ID,StatusID"`
}

// EmailResult is the struct that contains the email information for presentation
type EmailResult struct {
	ID           string `json:"id"`
	SenderID     string `json:"sender_id"`
	ReceiverID   string `json:"receiver_id"`
	TemplateID   string `json:"template_id"`
	SMTPServerID string `json:"smtp_server_id"`
	Variables    string `json:"variables"`
	Status       string `json:"status"`
}

// SetCreate set information for create a new email
func (e *Email) SetCreate() {
	e.Base.SetCreate(true)
	e.StatusID = Created
	e.Status = NewStatus(e.Base.ID)
	e.Status.SetCreated("")
	setList := []port.Domain{e.Sender, e.Receiver, e.Template, e.SMTPServer}
	for _, v := range setList {
		if !reflect.ValueOf(v).IsNil() {
			v.SetRepository(e.Repo)
			v.SetCreate()
		}
	}
	if e.Variables != nil {
		for _, v := range e.Variables {
			v.SetCreate()
		}
	}
}

// Validate validate the email information
func (e *Email) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		e.Base.Validate,
		e.ValidateSender,
		e.ValidateReceiver,
		e.ValidateTemplate,
		e.ValidateSMTPServer,
		e.ValidateVariables,
		e.ValidateDuplicity,
	})
}

// ValidateSender validate sender information
func (e *Email) ValidateSender() *kerror.KError {
	return e.validateItem(e.SenderID, e.Sender, "sender ", ErrEmailSenderIsRequired)
}

// ValidateReceiver validate receiver information
func (e *Email) ValidateReceiver() *kerror.KError {
	return e.validateItem(e.ReceiverID, e.Receiver, "receiver ", ErrEmailReceiverIsRequired)
}

// ValidateTemplate validate template information
func (e *Email) ValidateTemplate() *kerror.KError {
	return e.validateItem(e.TemplateID, e.Template, "template ", ErrEmailTemplateIsRequired)
}

// ValidateSMTPServer validate smtp server information
func (e *Email) ValidateSMTPServer() *kerror.KError {
	return e.validateItem(e.SMTPServerID, e.SMTPServer, "smtp server ", ErrEmailSMTPServerIsRequired)
}

// ValidateVariables validate the email information
func (e *Email) ValidateVariables() *kerror.KError {
	if e.Variables == nil {
		return nil
	}
	count := 0
	for _, v := range e.Variables {
		if err := v.Validate(); err != nil {
			msg := fmt.Sprintf("variable %v ", count)
			err.SetPrefix(msg)
			return err
		}
		count++
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (e *Email) ValidateDuplicity() *kerror.KError {
	return e.Base.ValidateDuplicity(e)
}

// GetByID returns the model by its ID
func (r *Email) GetByID() *kerror.KError {
	return r.Base.GetByID(r)
}

// TableName returns the table name for gorm
func (b *Email) TableName() string {
	return "email"
}

// GetResult returns the email information for presentation without extra information
func (b *Email) GetResult() any {
	vresult := ""
	for _, v := range b.Variables {
		vresult += v.GetResult() + " | "
	}
	if vresult != "" {
		vresult = vresult[:len(vresult)-3]
	}
	return &EmailResult{
		ID:           b.ID,
		SenderID:     b.getItemID(b.Sender, b.SenderID),
		ReceiverID:   b.getItemID(b.Receiver, b.ReceiverID),
		TemplateID:   b.getItemID(b.Template, b.TemplateID),
		SMTPServerID: b.getItemID(b.SMTPServer, b.SMTPServerID),
		Variables:    vresult,
		Status:       b.StatusID,
	}
}

// validateItem validate the email sub information by id or object
func (e *Email) validateItem(id string, sub port.Domain, prefix string, errMessage string) *kerror.KError {
	if !reflect.ValueOf(sub).IsNil() {
		if err := sub.Validate(); err != nil {
			err.SetPrefix(prefix)
			return err
		}
	} else if id != "" {
		sub2 := reflect.New(reflect.TypeOf(sub).Elem()).Interface().(port.Domain)
		sub2.SetRepository(e.Repo)
		sub2.SetID(id)
		if err := sub2.GetByID(); err != nil {
			err.SetPrefix(prefix)
			return err
		}
	} else {
		return kerror.NewKError(kerror.BadRequest, errMessage)
	}
	return nil
}

// getItemID returns the email sub information id
func (e *Email) getItemID(sub port.Domain, id string) string {
	if !reflect.ValueOf(sub).IsNil() {
		return sub.GetID()
	}
	return id
}
