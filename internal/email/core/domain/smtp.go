package domain

// SMTPServer is the struct that contains the SMTP server information
type SMTPServer struct {
	ID   string `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}
