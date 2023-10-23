package port

import (
	"time"
)

type Domain interface {
	GetInvoice() Invoice
}

type Invoice interface {
	Create(reference, client_id, client_name, client_email string,
		client_phone uint64, value float64, date, due time.Time, notes string) error
	Get() (string, string, string, string, string, uint64, float64, float64, time.Time, time.Time, string)
	SetSent(time.Time) error
	SetPaid() error
	SetCancelled() error
	GetStatus() (uint, string)
}
