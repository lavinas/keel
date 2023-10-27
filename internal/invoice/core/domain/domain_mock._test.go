package domain

import (
	"errors"
	"time"

	"github.com/lavinas/keel/internal/invoice/core/port"
)

// RepoMock
type RepoMock struct {
	Status string
}
func (r *RepoMock) Begin() error {
	return nil
}
func (r *RepoMock) Commit() error {
	return nil
}
func (r *RepoMock) Rollback() error {
	return nil
}
func (r *RepoMock) SaveInvoiceClient(client port.InvoiceClient) error {
	if r.Status == "saveInvoiceClientError" {
		return errors.New("save error")
	}
	return nil
}
func (r *RepoMock) SaveInvoice(invoice port.Invoice) error {
	if r.Status == "saveInvoiceError" {
		return errors.New("save error")
	}
	return nil
}
func (r *RepoMock) SaveInvoiceItem(item port.InvoiceItem) error {
	if r.Status == "saveInvoiceItemError" {
		return errors.New("save error")
	}
	return nil	
}
func (r *RepoMock) Close() error {
	return nil
}

// CreateInputItemDtoMock
type CreateInputItemDtoMock struct {
	Status string
}
func (d *CreateInputItemDtoMock) Validate() error {
	return nil
}
func (d *CreateInputItemDtoMock) GetReference() string {
	return "ref"
}
func (d *CreateInputItemDtoMock) GetDescription() string {
	return "desc"
}
func (d *CreateInputItemDtoMock) GetQuantity() (uint64, error) {
	if d.Status == "quantityError" {
		return 0, errors.New("quantity error")
	}
	return 1, nil
}
func (d *CreateInputItemDtoMock) GetPrice() (float64, error) {
	if d.Status == "priceError" {
		return 0, errors.New("price error")
	}
	return 1.33, nil
}

// CreateInputDtoMock
type CreateInputDtoMock struct {
	Status string
}
func (d *CreateInputDtoMock) Validate() error {
	return nil
}
func (d *CreateInputDtoMock) GetReference() string {
	return "ref"
}
func (d *CreateInputDtoMock) GetBusinessNickname() string {
	return "businessNickname"
}
func (d *CreateInputDtoMock) GetCustomerNickname() string {
	return "customerNickname"
}

func (d *CreateInputDtoMock) GetAmount() (float64, error) {
	if d.Status == "amountError" {
		return 0, errors.New("amount error")
	}
	return 1.33, nil
}
func (d *CreateInputDtoMock) GetDate() (time.Time, error) {
	if d.Status == "dateError" {
		return time.Now(), errors.New("date error")
	}
	return time.Parse("2006-01-02", "2023-10-10")
}
func (d *CreateInputDtoMock) GetDue() (time.Time, error) {
	if d.Status == "dueError" {
		return time.Now(), errors.New("due error")
	}
	return time.Parse("2006-01-02", "2023-10-20")
}
func (d *CreateInputDtoMock) GetNoteReference() string {
	return "noteReference"
}
func (d *CreateInputDtoMock) GetItems() []port.CreateInputItemDto {
	return []port.CreateInputItemDto{&CreateInputItemDtoMock{}}
}


