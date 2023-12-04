package domain

// Email is the struct that contains the email information
type Email struct {
	Base
	Sender     Sender            `json:"sender"`
	Receiver   Receiver          `json:"receiver"`
	Template   Template          `json:"template"`
	SMTPServer SMTPServer        `json:"smtp_server"`
	Variables  map[string]string `json:"variables"`
}
