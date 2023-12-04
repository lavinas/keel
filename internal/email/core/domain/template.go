package domain

// Template is the struct that contains the email template information
type Template struct {
	Base
	Name    string `json:"name"    gorm:"type:varchar(50); not null"`
	Subject string `json:"subject" gorm:"type:varchar(50); not null"`
	Body    string `json:"body"    gorm:"type:varchar(50); not null"`
}
