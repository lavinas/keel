package domain

// Template is the struct that contains the email template information
type Template struct {
	Base
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
