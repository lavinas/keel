package dto

// CreateOutputDto is the DTO for the create the crate a new invoice
type CreateOutputDto struct {
	Status    string `json:"status"`
	Reference string `json:"reference"`
}


func (dto *CreateOutputDto) Load(status string, reference string) {
	dto.Status = status
	dto.Reference = reference
}