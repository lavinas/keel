package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrSMTPServerHostIsRequired = "host is required"
	ErrSMTPServerHostLength     = "host must have less than 50 characters"
	ErrSMTPPortIsRequired       = "port is required"
	ErrSMTPUserIsRequired       = "user is required"
	ErrSMTPUserLength           = "user must have less than 50 characters"
	ErrSMTPPassIsRequired       = "pass is required"
	ErrSMTPPassLength           = "pass must have less than 50 characters"
)

// SMTPServer is the struct that contains the SMTP server information
type SMTPServer struct {
	Base
	Host string `json:"host" gorm:"type:varchar(50); not null"`
	Port int    `json:"port" gorm:"type:int; not null"`
	User string `json:"user" gorm:"type:varchar(50); not null"`
	Pass string `json:"pass" gorm:"type:varchar(50); not null"`
}

// SetCreate set information for create a new SMTP server
func (s *SMTPServer) SetCreate() {
	s.Base.SetShortenID(s, s.User)
	s.Base.SetCreate()
}

// Validate validate the SMTP server information
func (s *SMTPServer) Validate() *kerror.KError {
	return validateLoop([]func() *kerror.KError{
		s.Base.Validate,
		s.ValidateHost,
		s.ValidatePort,
		s.ValidateUser,
		s.ValidatePass,
		s.ValidateDuplicity,
	})
}

// ValidateHost validate the SMTP server information
func (s *SMTPServer) ValidateHost() *kerror.KError {
	if s.Host == "" {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPServerHostIsRequired)
	}
	if len(s.Host) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPServerHostLength)
	}
	return nil
}

// ValidatePort validate the SMTP port information
func (s *SMTPServer) ValidatePort() *kerror.KError {
	if s.Port == 0 {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPPortIsRequired)
	}
	return nil
}

// ValidateUser validate the SMTP user information
func (s *SMTPServer) ValidateUser() *kerror.KError {
	if s.User == "" {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPUserIsRequired)
	}
	if len(s.User) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPUserLength)
	}
	return nil
}

// ValidatePass validate the SMTP pass information
func (s *SMTPServer) ValidatePass() *kerror.KError {
	if s.Pass == "" {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPPassIsRequired)
	}
	if len(s.Pass) > 50 {
		return kerror.NewKError(kerror.BadRequest, ErrSMTPPassLength)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (r *SMTPServer) ValidateDuplicity() *kerror.KError {
	return r.Base.ValidateDuplicity(r)
}

// GetByID returns the model by its ID
func (r *SMTPServer) GetByID() *kerror.KError {
	return r.Base.GetByID(r)
}

// GetResult returns the result that is the SMTP server itself
func (s *SMTPServer) GetResult() interface{} {
	return s
}

// TableName returns the table name for gorm
func (b *SMTPServer) TableName() string {
	return "smtp_server"
}
