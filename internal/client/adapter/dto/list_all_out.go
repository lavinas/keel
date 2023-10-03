package dto

// ListAllOutputDto is the output DTO used to list all clients
type ListAllOutputDto struct {
	Clients []ClientCreateOutputDto `json:"clients" binding:"required"`
}
