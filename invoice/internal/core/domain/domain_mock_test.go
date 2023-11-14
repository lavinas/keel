package domain

import (
	"errors"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// RepoMock
type RepoMock struct {
	Status string
}

func (r *RepoMock) Begin() error {
	if r.Status == "beginError" {
		return errors.New("begin error")
	}
	return nil
}
func (r *RepoMock) Commit() error {
	if r.Status == "commitError" {
		return errors.New("commit error")
	}
	return nil
}
func (r *RepoMock) Rollback() error {
	return nil
}
func (r *RepoMock) IsDuplicatedInvoice(reference string) (bool, error) {
	return false, nil
}
func (r *RepoMock) SaveInvoiceClient(client port.InvoiceClient) error {
	if r.Status == "saveInvoiceClientError" {
		return errors.New("save error")
	}
	if r.Status == "saveBusinessError" && client.GetNickname() == "businessNickname" {
		return errors.New("save business error")
	}
	if r.Status == "saveCustomerError" && client.GetNickname() == "customerNickname" {
		return errors.New("save customer error")
	}
	return nil
}
func (r *RepoMock) GetLastInvoiceClient(nickname string, created_after time.Time, client port.InvoiceClient) (bool, error) {
	if r.Status == "getLastInvoiceBusinessError" && nickname == "businessNickname" {
		return false, errors.New("get last invoice client error")
	}
	if r.Status == "getLastInvoiceCustomerError" && nickname == "customerNickname" {
		return false, errors.New("get last invoice client error")
	}
	return false, nil
}
func (r *RepoMock) UpdateInvoiceClient(client port.InvoiceClient) error {
	if r.Status == "updateInvoiceClientError" {
		return errors.New("update error")
	}
	if r.Status == "updateBusinessError" && client.GetNickname() == "businessNickname" {
		return errors.New("update business error")
	}
	if r.Status == "updateCustomerError" && client.GetNickname() == "customerNickname" {
		return errors.New("update customer error")
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
func (r *RepoMock) LoadInvoiceVertex(graph port.InvoiceStatusGraph) error {
	if r.Status == "loadInvoiceVertexError" {
		return errors.New("load error")
	}
	return nil
}
func (r *RepoMock) LoadInvoiceEdge(graph port.InvoiceStatusGraph) error {
	if r.Status == "loadInvoiceEdgeError" {
		return errors.New("load error")
	}
	return nil
}
func (r *RepoMock) LogInvoiceEdge(class string, graph port.InvoiceStatusGraph) error {
	if r.Status == "logInvoiceEdgeError" {
		return errors.New("log error")
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
func (d *CreateInputDtoMock) Format() error {
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
	if d.Status == "itemsError" {
		return []port.CreateInputItemDto{
			&CreateInputItemDtoMock{Status: "quantityError"},
			&CreateInputItemDtoMock{Status: "priceError"},
		}
	}
	items := []port.CreateInputItemDto{
		&CreateInputItemDtoMock{},
		&CreateInputItemDtoMock{},
	}
	return items
}

type GetClientByNicknameInputDtoMock struct {
	Status string
}

func (d *GetClientByNicknameInputDtoMock) GetId() string {
	return "0000"
}
func (d *GetClientByNicknameInputDtoMock) GetName() string {
	return "name"
}
func (d *GetClientByNicknameInputDtoMock) GetNickname() string {
	return "nickname"
}
func (d *GetClientByNicknameInputDtoMock) GetDocument() (uint64, error) {
	if d.Status == "documentError" {
		return 0, errors.New("document error")
	}
	return 1, nil
}
func (d *GetClientByNicknameInputDtoMock) GetPhone() (uint64, error) {
	if d.Status == "phoneError" {
		return 0, errors.New("phone error")
	}
	return 1, nil
}
func (d *GetClientByNicknameInputDtoMock) GetEmail() string {
	return "email"
}
