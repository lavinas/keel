package domain

// Email is the struct that contains the email information
type Email struct {
	Base
	SenderID   string            `json:"sender_id"   gorm:"type:varchar(50); not null"`
	Sender     Sender            `json:"sender"      gorm:"foreignKey:SenderID"`
	ReceiverID string            `json:"receiver_id" gorm:"type:varchar(50); not null"`
	Receiver   Receiver          `json:"receiver"    gorm:"foreignKey:ReceiverID"`
	TemplateID string            `json:"template_id" gorm:"type:varchar(50); not null"`
	Template   Template          `json:"template"    gorm:"foreignKey:TemplateID"`
	SMTPServerID string          `json:"smtp_server_id" gorm:"type:varchar(50); not null"`
	SMTPServer SMTPServer        `json:"smtp_server" gorm:"foreignKey:SMTPServerID"`
	Variables  map[string]string `json:"variables"   gorm:"type:varchar(50); not null"`
	StatusID   string            `json:"-"           gorm:"type:varchar(50); not null"`
	Status     Status            `json:"status"      gorm:"foreignKey:StatusID"`
}
