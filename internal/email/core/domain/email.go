package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrEmailSenderIsRequired     = "sender is required"
	ErrEmailSenderIsTwice        = "sender id and sender is informed. Only one is allowed"
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
	SenderID     string            `json:"sender_id"   gorm:"type:varchar(50); not null"`
	Sender       *Sender           `json:"sender"      gorm:"foreignKey:SenderID"`
	ReceiverID   string            `json:"receiver_id" gorm:"type:varchar(50); not null"`
	Receiver     *Receiver         `json:"receiver"    gorm:"foreignKey:ReceiverID"`
	TemplateID   string            `json:"template_id" gorm:"type:varchar(50); not null"`
	Template     *Template         `json:"template"    gorm:"foreignKey:TemplateID"`
	SMTPServerID string            `json:"smtp_server_id" gorm:"type:varchar(50); not null"`
	SMTPServer   *SMTPServer       `json:"smtp_server" gorm:"foreignKey:SMTPServerID"`
	Variables    map[string]string `json:"variables"   gorm:"type:varchar(50); not null"`
	StatusID     string            `json:"-"           gorm:"type:varchar(50); not null"`
	Status       *Status           `json:"status"      gorm:"foreignKey:ID,StatusID"`
}

// SetCreate set information for create a new email
func (e *Email) SetCreate() {
	e.Base.SetCreate()
	e.StatusID = Created
	e.Status = NewStatus(e.Base.ID, e.StatusID)
	if e.Variables == nil {
		e.Variables = make(map[string]string)
	}
	if e.Sender != nil {
		e.Sender.SetCreate()
	}
	if e.Receiver != nil {
		e.Receiver.SetCreate()
	}
	if e.Template != nil {
		e.Template.SetCreate()
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
	})
}

// ValidateSender validate the email information
func (e *Email) ValidateSender() *kerror.KError {
	if e.SenderID == "" && e.Sender == nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailSenderIsRequired)
	}
	if e.SenderID != "" && e.Sender != nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailSenderIsTwice)
	}
	if e.Sender != nil {
		if err := e.Sender.Validate(); err != nil {
			err.SetPrefix("sender")
			return err
		}
	}
	return nil
}

// ValidateReceiver validate the email information
func (e *Email) ValidateReceiver() *kerror.KError {
	if e.ReceiverID == "" && e.Receiver == nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailReceiverIsRequired)
	}
	if e.ReceiverID != "" && e.Receiver != nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailReceiverIsTwice)
	}
	if e.Receiver != nil {
		if err := e.Receiver.Validate(); err != nil {
			err.SetPrefix("receiver")
			return err
		}
	}
	return nil
}

// ValidateTemplate validate the email information
func (e *Email) ValidateTemplate() *kerror.KError {
	if e.TemplateID == "" && e.Template == nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailTemplateIsRequired)
	}
	if e.TemplateID != "" && e.Template != nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailTemplateTwice)
	}
	if e.Template != nil {
		if err := e.Template.Validate(); err != nil {
			err.SetPrefix("template")
			return err
		}
	}
	return nil
}

// ValidateSMTPServer validate the email information
func (e *Email) ValidateSMTPServer() *kerror.KError {
	if e.SMTPServerID == "" && e.SMTPServer == nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailSMTPServerIsRequired)
	}
	if e.SMTPServerID != "" && e.SMTPServer != nil {
		return kerror.NewKError(kerror.BadRequest, ErrEmailSMTPServerTwice)
	}
	if e.SMTPServer != nil {
		if err := e.SMTPServer.Validate(); err != nil {
			err.SetPrefix("smtp_server")
			return err
		}
	}
	return nil
}
