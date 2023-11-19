package port

import (
	"github.com/lavinas/keel/invoice/internal/core/domain"
)

type Repository interface {
	CheckInvoiceExists(invoice *domain.Invoice) (bool, error)
	GetClientByID(id string) (*domain.Client, error)
	GetProductByID(id string) (*domain.Product, error)
	GetInstructionByID(id string) (*domain.Instruction, error)	
	RegisterInvoice(invoice *domain.Invoice) error
}