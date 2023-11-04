package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// CreateMock is a mock of Log
type LogMock struct {
	Type    string
	Message string
}

func (l *LogMock) GetFile() *os.File {
	return nil
}
func (l *LogMock) Info(message string) {
	l.Type = "info"
	l.Message = message
}
func (l *LogMock) Infof(input any, message string) {
	l.Info(fmt.Sprintf("%s:%v", message, input))
}
func (l *LogMock) Error(message string) {
	l.Type = "error"
	l.Message = message
}
func (l *LogMock) Errorf(input any, err error) {
	l.Error(fmt.Sprintf("%s:%v", err.Error(), input))
}
func (l *LogMock) Close() {
}

type RestConsumerMock struct {
	Status string
}
func (r *RestConsumerMock) GetClientByNickname(nickname string, client port.GetClientByNicknameInputDto) (bool, error) {
	if r.Status == "get client error" {
		return false, errors.New("get client error")
	}
	if r.Status == "get client not found" {
		return false, nil
	}
	return true, nil
}

// DomainMock is a mock of Domain
type DomainMock struct {
}

func (d *DomainMock) GetInvoice() port.Invoice {
	return &InvoiceMock{}
}

type InvoiceClientMock struct {
	status string
}

func (ic *InvoiceClientMock) Load(nickname, clientId, name, email string, phone, document uint64) {
}
func (ic *InvoiceClientMock) LoadGetClientNicknameDto(input port.GetClientByNicknameInputDto) error {
	if ic.status == "load error" {
		return errors.New("load error")
	}
	return nil
}
func (ic *InvoiceClientMock) Save() error {
	if ic.status == "save error" {
		return errors.New("save error")
	}
	return nil
}
func (ic *InvoiceClientMock) Update() error {
	if ic.status == "update error" {
		return errors.New("update error")
	}
	return nil
}
func (ic *InvoiceClientMock) GetId() string {
	return "id"
}
func (ic *InvoiceClientMock) GetNickname() string {
	return "nickname"
}
func (ic *InvoiceClientMock) GetClientId() string {
	return "client_id"
}
func (ic *InvoiceClientMock) GetName() string {
	return "name"
}
func (ic *InvoiceClientMock) GetDocument() uint64 {
	return 1
}
func (ic *InvoiceClientMock) GetPhone() uint64 {
	return 1
}
func (ic *InvoiceClientMock) GetEmail() string {
	return "email"
}

// InoviceMock is a mock of Invoice Domain
type InvoiceMock struct {
	Status string
}

func (i *InvoiceMock) Load(input port.CreateInputDto) error {
	if i.Status == "load error" {
		return errors.New("load error")
	}
	return nil
}
func (i *InvoiceMock) SetAmount(amount float64) {
}
func (i *InvoiceMock) Save() error {
	if i.Status == "save error" {
		return errors.New("save error")
	}
	return nil
}
func (i *InvoiceMock) LoadBusiness(dto port.GetClientByNicknameInputDto) error {
	if i.Status == "load business error" {
		return errors.New("load business error")
	}
	return nil
}
func (i *InvoiceMock) LoadCustomer(dto port.GetClientByNicknameInputDto) error {
	if i.Status == "load customer error" {
		return errors.New("load customer error")
	}
	return nil
}
func (i *InvoiceMock) Update() error {
	if i.Status == "update invoice error" {
		return errors.New("update error")
	}
	return nil
}
func (i *InvoiceMock) GetId() string {
	return ""
}
func (i *InvoiceMock) GetReference() string {
	return ""
}
func (i *InvoiceMock) GetBusinessId() string {
	return ""
}
func (i *InvoiceMock) GetBusiness() port.InvoiceClient {
	return &InvoiceClientMock{}
}
func (i *InvoiceMock) GetCustomerId() string {
	return ""
}
func (i *InvoiceMock) GetCustomer() port.InvoiceClient {
	return &InvoiceClientMock{}
}
func (i *InvoiceMock) GetAmount() float64 {
	return 0
}
func (i *InvoiceMock) GetDate() time.Time {
	return time.Time{}
}
func (i *InvoiceMock) GetDue() time.Time {
	return time.Time{}
}
func (i *InvoiceMock) GetNoteId() *string {
	return nil
}
func (i *InvoiceMock) GetStatusId() uint {
	return 0
}
func (i *InvoiceMock) GetCreatedAt() time.Time {
	return time.Time{}
}
func (i *InvoiceMock) GetUpdatedAt() time.Time {
	return time.Time{}
}
func (i *InvoiceMock) IsDuplicated() (bool, error) {
	if i.Status == "duplicity" {
		return true, nil
	}
	if i.Status == "duplicity error" {
		return false, errors.New("duplicated error")
	}
	return false, nil
}

// CreateInputDtoMock is a mock of CreateInputDto
type CreateInputDtoMock struct {
	Status string
}

func (i *CreateInputDtoMock) Validate() error {
	if i.Status == "validate error" {
		return errors.New("validate error")
	}
	return nil
}
func (i *CreateInputDtoMock) GetReference() string {
	return "ref"
}
func (i *CreateInputDtoMock) GetBusinessNickname() string {
	return "business"
}
func (i *CreateInputDtoMock) GetCustomerNickname() string {
	return "customer"
}
func (i *CreateInputDtoMock) GetAmount() (float64, error) {
	return 10.33, nil
}
func (i *CreateInputDtoMock) GetDate() (time.Time, error) {
	return time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC), nil
}
func (i *CreateInputDtoMock) GetDue() (time.Time, error) {
	return time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), nil
}
func (i *CreateInputDtoMock) GetNoteReference() string {
	return ""
}
func (i *CreateInputDtoMock) GetItems() []port.CreateInputItemDto {
	return nil
}

// CreateOutputItemDtoMock is a mock of CreateOutputItemDto
type CreateOutputDtoMock struct {
	status    string
	reference string
}

func (i *CreateOutputDtoMock) Load(status string, reference string) {
	i.status = status
	i.reference = reference
}
