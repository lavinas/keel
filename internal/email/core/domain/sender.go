package domain

// Sender is the struct that contains the business information
type Sender struct {
	Base
	Name  string `json:"name"  gorm:"type:varchar(50); not null"`
	Email string `json:"email" gorm:"type:varchar(50); not null"`
}
