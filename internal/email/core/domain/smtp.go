package domain

// SMTPServer is the struct that contains the SMTP server information
type SMTPServer struct {
	Base
	Host string `json:"host" gorm:"type:varchar(50); not null"`
	Port int    `json:"port" gorm:"type:int; not null"`
	User string `json:"user" gorm:"type:varchar(50); not null"`
	Pass string `json:"pass" gorm:"type:varchar(50); not null"`
}
