package mysql

import (
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Invoice Client Mock
type InvoiceClientMock struct {
}

func (i *InvoiceClientMock) Load(nickname, clientId, name, email string, phone, document uint64) {
}
func (i *InvoiceClientMock) Save() error {
	return nil
}
func (i *InvoiceClientMock) GetId() string {
	return "1"
}
func (i *InvoiceClientMock) GetNickname() string {
	return "nickname"
}
func (i *InvoiceClientMock) GetClientId() string {
	return "client_id"
}
func (i *InvoiceClientMock) GetName() string {
	return "name"
}
func (i *InvoiceClientMock) GetDocument() uint64 {
	return 1
}
func (i *InvoiceClientMock) GetPhone() uint64 {
	return 1
}
func (i *InvoiceClientMock) GetEmail() string {
	return "email"
}

// Invoice mock
type InvoiceMock struct {
}

func (*InvoiceMock) Load(input port.CreateInputDto) error {
	return nil
}
func (*InvoiceMock) SetAmount(amount float64) {
}
func (*InvoiceMock) Save() error {
	return nil
}
func (*InvoiceMock) GetId() string {
	return "id"
}
func (*InvoiceMock) GetReference() string {
	return "ref"
}
func (*InvoiceMock) GetBusinessId() string {
	return "1"
}
func (*InvoiceMock) GetCustomerId() string {
	return "1"
}
func (*InvoiceMock) GetAmount() float64 {
	return 1.66
}
func (*InvoiceMock) GetDate() time.Time {
	t, _ := time.Parse("2006-01-02", "2023-10-10")
	return t
}
func (*InvoiceMock) GetDue() time.Time {
	t, _ := time.Parse("2006-01-02", "2023-10-20")
	return t
}
func (*InvoiceMock) GetNoteId() *string {
	return nil
}

func (*InvoiceMock) GetStatusId() uint {
	return 0
}
func (*InvoiceMock) GetCreatedAt() time.Time {
	t, _ := time.Parse("02 Jan 06 15:04 -0700", "10 Oct 23 10:10 -0300")
	return t

}
func (*InvoiceMock) GetUpdatedAt() time.Time {
	t, _ := time.Parse("02 Jan 06 15:04 -0700", "10 Oct 23 10:10 -0300")
	return t
}
func (*InvoiceMock) IsDuplicated() (bool, error) {
	return false, nil
}

// Invoice Item Mock
type InvoiceItemMock struct {
}

func (i *InvoiceItemMock) Load(dto port.CreateInputItemDto, invoice port.Invoice) error {
	return nil
}
func (i *InvoiceItemMock) Save() error {
	return nil
}
func (i *InvoiceItemMock) GetId() string {
	return "id"
}
func (i *InvoiceItemMock) GetInvoiceId() string {
	return "id"
}
func (i *InvoiceItemMock) GetServiceReference() string {
	return "ref"
}
func (i *InvoiceItemMock) GetDescription() string {
	return "description"
}
func (i *InvoiceItemMock) GetAmount() float64 {
	return 1.66
}
func (i *InvoiceItemMock) GetQuantity() uint64 {
	return 1
}