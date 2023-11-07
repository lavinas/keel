package mysql

import (
	"errors"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Invoice Client Mock
type InvoiceClientMock struct {
	status string
}

func (i *InvoiceClientMock) Load(id, nickname, clientId, name, email string, phone, document uint64, created_id time.Time) {
}
func (i *InvoiceClientMock) LoadGetClientNicknameDto(input port.GetClientByNicknameInputDto) error {
	if i.status == "get client error" {
		return errors.New("get client error")
	}
	return nil
}
func (i *InvoiceClientMock) GetLastInvoiceClientId(nickname string, created_after time.Time) (string, error) {
	return "1", nil
}
func (i *InvoiceClientMock) Save() error {
	return nil
}
func (i *InvoiceClientMock) Update() error {
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
func (i *InvoiceClientMock) GetCreatedAt() time.Time {
	return time.Time{}
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
func (*InvoiceMock) LoadBusiness(dto port.GetClientByNicknameInputDto) error {
	return nil
}
func (*InvoiceMock) LoadCustomer(dto port.GetClientByNicknameInputDto) error {
	return nil
}
func (*InvoiceMock) Update() error {
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
func (*InvoiceMock) GetBusiness() port.InvoiceClient {
	return &InvoiceClientMock{}
}
func (*InvoiceMock) GetCustomerId() string {
	return "1"
}
func (*InvoiceMock) GetCustomer() port.InvoiceClient {
	return &InvoiceClientMock{}
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
