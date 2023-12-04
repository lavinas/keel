package domain

// Receiver is the struct that contains the client information
type Receiver struct {
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
