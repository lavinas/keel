package domain

// Common errors for all entities

// Base errors
const (
	ErrBaseBusinessIDIsRequired = "business id is required"
	ErrBaseBusinessIDLength     = "business id must have only one word. Use underscore to separate words"
	ErrBaseBusinessIDLower      = "business id must be lower case"
	ErrBaseIDIsRequired         = "id is required"
	ErrBaseIDLength             = "id must have only one word. Use underscore to separate words"
	ErrBaseIDLower              = "id must be lower case"
	ErrBaseCreatedAtIsRequired  = "created_at is required"
)

// Client errors
const (
	ErrClientNameIsRequired    = "client name is required"
	ErrClientNameLength        = "client name must have at least a name and surname"
	ErrClientEmailIsRequired   = "client email is required"
	ErrClientEmailIsInvalid    = "client email is invalid"
	ErrClientDocumentIsInvalid = "client document is invalid"
	ErrClientPhoneIsInvalid    = "client cell phone number is invalid"
	ErrClientIDNotLower        = "client id must be lower case"
	ErrDuplicatedID            = "id already exists"
)

// Instruction Status errors
const (
	ErrInstructionDescriptionIsRequired = "instruction description is required"
)

// Product errors
const (
	ErrProductDescriptionIsRequired = "product description is required"
)

// Invoice errors
const (
	ErrInvoiceClientIsRequired        = "client id is required"
	ErrInvoiceClientInformedTwice     = "client id and client are informed. Choose only one"
	ErrInvoiceClientNotFound          = "client id not found"
	ErrInvoiceBusinessIsRequired      = "invoice business is required"
	ErrInvoiceCustomerIsRequired      = "invoice customer is required"
	ErrInvoiceDateIsRequired          = "invoice date is required"
	ErrInvoiceDateIsInvalid           = "invoice date is invalid. It must be in the format YYYY-MM-DD"
	ErrInvoiceDueIsRequired           = "invoice due is required"
	ErrInvoiceDueIsInvalid            = "invoice due is invalid. It must be in the format YYYY-MM-DD"
	ErrInvoiceDueBeforeDate           = "invoice due date should not be before invoice date"
	ErrInvoiceAmountIsRequired        = "invoice amount is required"
	ErrInvoiceAmountIsInvalid         = "invoice amount is invalid. It must be a number (ex: 100.00)"
	ErrInvoiceInstructionIDIsRequired = "instruction id is required"
	ErrInvoiceInstructionNotFound     = "instruction id not found"
	ErrInvoiceItemIDLength            = "invoice item id must have only one word"
	ErrInvoiceItemIDLower             = "invoice item id must be lower case"
	ErrInvoiceItemQuantity            = "invoice item quantity must be greater than 0"
	ErrInvoiceItemPrice               = "invoice item price must be not equal to 0"
	ErrInvoiceAmountUnmatch           = "invoice amount does not match with the sum of the items"
)

// Invoice Item errors
const (
	ErrInvoiceItemQuantityRequired        = "invoice item quantity is required"
	ErrInvoiceItemQuantityInvalid         = "invoice item quantity is invalid"
	ErrInvoiceItemQuantityLessOrEqualZero = "invoice item quantity must be greater than 0"
	ErrInvoiceItemPriceRequired           = "invoice item price is required"
	ErrInvoiceItemPriceInvalid            = "invoice item price is invalid"
	ErrInvoiceItemPriceLessOrEqualZero    = "invoice item price must be greater than 0"
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
