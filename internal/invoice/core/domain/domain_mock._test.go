package domain

import (
	"github.com/lavinas/keel/internal/invoice/core/port"
)

type RepoMock struct {
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
	return nil
}
func (r *RepoMock) SaveInvoice(invoice port.Invoice) error {
	return nil
}
func (r *RepoMock) SaveInvoiceItem(item port.InvoiceItem) error {
	return nil	
}
func (r *RepoMock) Close() error {
	return nil
}
