package domain

// Sender is the struct that contains the business information
type Sender struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
