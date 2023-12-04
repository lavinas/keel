package domain

// Receiver is the struct that contains the client information
type Receiver struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
