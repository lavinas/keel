package dto

// Invoice Create DTO errors
const (
	ErrInvoiceCreateDtoBusinessIDRequired    = "invoice business id is required"
	ErrInvoiceCreateDtoNumberRequired        = "invoice number is required"
	ErrInvoiceCreateDtoNumberInvalid         = "invoice number is invalid"
	ErrInvoiceCreateDtoCustomerIDRequired    = "invoice customer id is required"
	ErrInvoiceCreateDtoDateRequired          = "invoice date is required"
	ErrInvoiceCreateDtoDateInvalid           = "invoice date is invalid"
	ErrInvoiceCreateDtoDueRequired           = "invoice due is required"
	ErrInvoiceCreateDtoDueInvalid            = "invoice due date is invalid"
	ErrInvoiceCreateDtoAmountRequired        = "invoice amount is required"
	ErrInvoiceCreateDtoAmountInvalid         = "invoice amount is invalid"
	ErrInvoiceCreateDtoItemsRequired         = "invoice items is required"
	ErrInvoiceCreateDtoInstructionIDRequired = "invoice instruction id is required"
)

// Invoice Item Create DTO errors
const (
	ErrInvoiceItemCreateDtoProductIDRequired   = "invoice item product id is required"
	ErrInvoiceItemCreateDtoDescriptionRequired = "invoice item description is required"
	ErrInvoiceItemCreateDtoQuantityRequired    = "invoice item quantity is required"
	ErrInvoiceItemCreateDtoQuantityInvalid     = "invoice item quantity is invalid"
	ErrInvoiceItemCreateDtoUnitPriceRequired   = "invoice item unit price is required"
	ErrInvoiceItemCreateDtoUnitPriceInvalid    = "invoice item unit price is invalid"
)
