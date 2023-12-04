package domain

import (
	"time"
)

// Base is the struct that contains the base information
type Base struct {
	ID         string    `json:"id" gorm:"primaryKey;type:varchar(50); not null"`
	Created_at time.Time `json:"-"  gorm:"type:timestamp; not null"`
	Updated_at time.Time `json:"-"  gorm:"type:timestamp; not null"`
}
