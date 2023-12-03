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
	ErrBaseIDAlreadyExists      = "id already exists"
)

// Client errors
const (
	ErrClientNameIsRequired    = "name is required"
	ErrClientNameLength        = "name must have at least a name and surname"
	ErrClientEmailIsRequired   = "email is required"
	ErrClientEmailIsInvalid    = "email is invalid"
	ErrClientDocumentIsInvalid = "document is invalid"
	ErrClientPhoneIsInvalid    = "cell phone number is invalid"
)

// Instruction Status errors
const (
	ErrInstructionDescriptionIsRequired = "description is required"
)

// Product errors
const (
	ErrProductDescriptionIsRequired = "description is required"
)

// Invoice errors
const (
	ErrInvoiceClientIsRequired         = "client id is required"
	ErrInvoiceClientInformedTwice      = "client id and client are informed. Choose only one"
	ErrInvoiceClientNotFound           = "client id not found"
	ErrInvoiceDateIsRequired           = "date is required"
	ErrInvoiceDateIsInvalid            = "date is invalid. It must be in the format YYYY-MM-DD"
	ErrInvoiceDueIsRequired            = "due is required"
	ErrInvoiceDueIsInvalid             = "due is invalid. It must be in the format YYYY-MM-DD"
	ErrInvoiceDueBeforeDate            = "due date should not be before invoice date"
	ErrInvoiceAmountIsRequired         = "amount is required"
	ErrInvoiceAmountIsInvalid          = "amount is invalid. It must be a number (ex: 100.00)"
	ErrInvoiceInstructionIDIsRequired  = "instruction id is required"
	ErrInvoiceInstructionInformedTwice = "instruction id and instruction are informed. Choose only one"
	ErrInvoiceInstructionNotFound      = "instruction id not found"
	ErrInvoiceAmountUnmatch            = "invoice amount does not match with the sum of the items"
)

// Invoice Item errors
const (
	ErrItemProductRequired         = "product is required"
	ErrItemProductConflict         = "product and product id are informed. Choose only one"
	ErrItemProductInvalid          = "product is invalid"
	ErrItemProductNotFound         = "product not found"
	ErrItemQuantityRequired        = "quantity is required"
	ErrItemQuantityInvalid         = "quantity is invalid"
	ErrItemQuantityLessOrEqualZero = "quantity must be greater than 0"
	ErrItemPriceRequired           = "price is required"
	ErrItemPriceInvalid            = "price is invalid"
	ErrItemPriceLessOrEqualZero    = "price must be greater than 0"
)

// Invoice Status errors
const (
	ErrInvoiceStatusIDIsRequired   = "id is required"
	ErrInvoiceStatusNameIsRequired = "name is required"
	ErrInvoiceStatusIDIsInvalid    = "id is invalid"
)

// Payment Status errors
const (
	ErrPaymentStatusIDIsRequired   = "id is required"
	ErrPaymentStatusNameIsRequired = "name is required"
	ErrPaymentStatusIDIsInvalid    = "id is invalid"
)
