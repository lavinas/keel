package domain

// CreateInputOutputDto is the input DTO used to create a client
type CreateInputDto struct {
	Name string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// CreateInputOutputDto is the output DTO used to create a client
type CreateOutputDto struct {
	Name string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document uint64 `json:"document" binding:"required"`
	Phone uint64 `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// ListAllOutputDto is the output DTO used to list all clients
type ListAllOutputDto struct {
	Clients []CreateOutputDto `json:"clients" binding:"required"`
}