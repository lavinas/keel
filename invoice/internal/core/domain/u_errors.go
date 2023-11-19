package domain

// Common errors for all entities

// Base errors
const (
	ErrBaseIDLength = "id must have only one word. Use underscore to separate words"
	ErrBaseIDLower  = "id must be lower case"
)

// Client errors
const (
	ErrClientIDLength          = "client id must have only one word"
	ErrClientNameIsRequired    = "client name is required"
	ErrClientNameLength        = "client name must have at least a name and surname"
	ErrClientEmailIsRequired   = "client email is required"
	ErrClientEmailIsInvalid    = "client email is invalid"
	ErrClientDocumentIsInvalid = "client document is invalid"
	ErrClientPhoneIsInvalid    = "client cell phone number is invalid"
	ErrClientIDNotLower        = "client id must be lower case"
)

// Instruction errors
const (
	ErrInstructionBusinessIsRequired = "instruction business is required"
)

// Invoice errors
const (
	ErrInvoiceBusinessIsRequired = "invoice business is required"
	ErrInvoiceCustomerIsRequired = "invoice customer is required"
	ErrInvoiceDateIsRequired     = "invoice date is required"
	ErrInvoiceDueIsRequired      = "invoice due is required"
	ErrInvoiceAmountNotMatch     = "invoice amount not match with items values"
	ErrInvoiceAmountIsInvalid    = "invoice amount should be greater than 0"
	ErrInvoiceItemIDLength       = "invoice item id must have only one word"
	ErrInvoiceItemIDLower        = "invoice item id must be lower case"
	ErrInvoiceItemQuantity       = "invoice item quantity must be greater than 0"
	ErrInvoiceItemPrice          = "invoice item price must be not equal to 0"
)

// Product errors
const (
	ErrProductIDLength           = "product id must have only one word"
	ErrProductIDLower            = "product id must be lower case"
	ErrProductBusinessIsRequired = "product business is required"
)

// Invoice Status errors
const (
	ErrInvoiceStatusIDIsRequired   = "invoice status id is required"
	ErrInvoiceStatusNameIsRequired = "invoice status name is required"
	ErrInvoiceStatusIDIsInvalid    = "invoice status id is invalid"
)

// Payment Status errors
const (
	ErrPaymentStatusIDIsRequired   = "payment status id is required"
	ErrPaymentStatusNameIsRequired = "payment status name is required"
	ErrPaymentStatusIDIsInvalid    = "payment status id is invalid"
)
