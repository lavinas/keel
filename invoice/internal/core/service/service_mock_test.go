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

// DomainMock is a mock of Domain
type DomainMock struct {
}

func (d *DomainMock) GetInvoice() port.Invoice {
	return &InvoiceMock{}
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
func (i *InvoiceMock) GetId() string {
	return ""
}
func (i *InvoiceMock) GetReference() string {
	return ""
}
func (i *InvoiceMock) GetBusinessId() string {
	return ""
}
func (i *InvoiceMock) GetCustomerId() string {
	return ""
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