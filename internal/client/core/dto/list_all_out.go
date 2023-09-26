package dto

// ListAllOutputDto is the output DTO used to list all clients
type ListAllOutputDto struct {
	Clients []CreateOutputDto `json:"clients" binding:"required"`
}