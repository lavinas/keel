package service

import (
	"os"
	"fmt"
	"time"

	"github.com/lavinas/keel/internal/invoice/core/port"
)

// CreateMock is a mock of Log
type LogMock struct {
	Type string
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

// InoviceMock is a mock of Invoice Domain
type InvoiceMock struct {
}
func (i *InvoiceMock) Load(input port.CreateInputDto) error {
	return nil
}
func (i *InvoiceMock) SetAmount(amount float64) {
}
func (i *InvoiceMock) Save() error {
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

// CreateInputDtoMock is a mock of CreateInputDto
type CreateInputDtoMock struct {
}
func (i *CreateInputDtoMock) Validate() error {
	return nil
}
func (i *CreateInputDtoMock) GetReference() string {
	return ""
}
func (i *CreateInputDtoMock) GetBusinessNickname() string {
	return ""
}
func (i *CreateInputDtoMock) GetCustomerNickname() string {
	return ""
}
func (i *CreateInputDtoMock) GetAmount() (float64, error) {
	return 0, nil
}
func (i *CreateInputDtoMock) GetDate() (time.Time, error) {
	return time.Time{}, nil
}
func (i *CreateInputDtoMock) GetDue() (time.Time, error) {
	return time.Time{}, nil
}
func (i *CreateInputDtoMock) GetNoteReference() string {
	return ""
}
func (i *CreateInputDtoMock) GetItems() []port.CreateInputItemDto {
	return nil
}

// CreateOutputItemDtoMock is a mock of CreateOutputItemDto
type CreateOutputDtoMock struct {
}
func (i *CreateOutputDtoMock) Load(status string, reference string) {
}