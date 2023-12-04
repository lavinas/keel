package domain

// Sender is the struct that contains the business information
type Sender struct {
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
