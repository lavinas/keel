package register

import (
	"errors"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/domain"
	"github.com/lavinas/keel/invoice/internal/core/port"
)

type RegisterProduct struct {
	RegisterBase
	Description string `json:"description"`
}

// Validate validates the product
func (r *RegisterProduct) Validate() error {
	return ValidateLoop([]func() error{
		r.ValidateBase,
		r.ValidateDescription,
	})
}

func (r *RegisterProduct) GetDomain(businnes_id string) port.Domain {
	return domain.NewProduct(businnes_id, r.ID, r.Description, time.Time{}, time.Time{})
}

// Get returns the ID and Description
func (r *RegisterProduct) Get() (string, string) {
	return r.ID, r.Description
}

// ValidateDescription validates the description of the product
func (r *RegisterProduct) ValidateDescription() error {
	if r.Description == "" {
		return errors.New(ErrRegisterProductDescriptionIsRequired)
	}
	return nil
}
