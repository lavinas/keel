package dto

// Register DTO errors

// Invoice Create DTO errors
const (
	ErrRegisterInvoiceDtoBusinessRequired      = "invoice business is required. You must provide either a business id or a business object"
	ErrRegisterInvoiceDtoBusinessDuplicity     = "invoice business is duplicated. You must provide just one business id or a business object"
	ErrRegisterInvoiceDtoCustomerRequired      = "invoice customer id is required. You must provide either a customer id or a customer object"
	ErrRegisterInvoiceDtoCustomerIDDuplicity   = "invoice customer is duplicated. You must provide just one customer id or a customer object"
	ErrRegisterInvoiceDtoNumberRequired        = "invoice number is required"
	ErrRegisterInvoiceDtoNumberInvalid         = "invoice number is invalid"
	ErrRegisterInvoiceDtoDateRequired          = "invoice date is required"
	ErrRegisterInvoiceDtoDateInvalid           = "invoice date is invalid"
	ErrRegisterInvoiceDtoDueRequired           = "invoice due is required"
	ErrRegisterInvoiceDtoDueInvalid            = "invoice due date is invalid"
	ErrRegisterInvoiceDtoAmountRequired        = "invoice amount is required"
	ErrRegisterInvoiceDtoAmountInvalid         = "invoice amount is invalid"
	ErrRegisterInvoiceDtoItemsRequired         = "invoice items is required"
	ErrRegisterInvoiceDtoInstructionIDRequired = "invoice instruction id is required"
)

// Invoice Item Create DTO errors
const (
	ErrRegisterInvoiceItemDtoProductIDRequired          = "invoice item product id is required"
	ErrRegisterInvoiceItemDtoProductDescriptionRequired = "invoice item product description is required"
	ErrRegisterInvoiceItemDtoDescriptionRequired        = "invoice item description is required"
	ErrRegisterInvoiceItemDtoQuantityRequired           = "invoice item quantity is required"
	ErrRegisterInvoiceItemDtoQuantityInvalid            = "invoice item quantity is invalid"
	ErrRegisterInvoiceItemDtoUnitPriceRequired          = "invoice item unit price is required"
	ErrRegisterInvoiceItemDtoUnitPriceInvalid           = "invoice item unit price is invalid"
)

// Invoice Client Create DTO errors
const (
	ErrRegisterInvoiceClientIDLength          = "invoice client id must have just one world. Please use underscore to separate words"
	ErrRegisterInvoiceClientIDLower           = "invoice client id must be lower case"
	ErrRegisterInvoiceClientNameIsRequired    = "invoice client name is required"
	ErrRegisterInvoiceClientNameLength        = "client name must have at least a name and surname"
	ErrRegisterInvoiceClientEmailIsRequired   = "invoice client email is required"
	ErrRegisterInvoiceClientEmailIsInvalid    = "invoice client email is invalid"
	ErrRegisterInvoiceClientDocumentIsInvalid = "invoice client document is invalid"
	ErrRegisterInvoiceClientPhoneIsInvalid    = "invoice client phone is invalid"
)
