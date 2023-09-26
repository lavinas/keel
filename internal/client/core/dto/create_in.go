package dto

// CreateInputDto is the input DTO used to create a client
type CreateInputDto struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// NewCreateInputDto creates a new CreateInputDto
func NewCreateInputDto(name, nickname, document, phone, email string) CreateInputDto {
	return CreateInputDto{
		Name:     name,
		Nickname: nickname,
		Document: document,
		Phone:    phone,
		Email:    email,
	}
}

func (c CreateInputDto) Validate() (bool, string) {
	if c.Name == "" || c.Nickname == "" || c.Document == "" || c.Phone == "" || c.Email == "" {
		return false, "name, nickname, document, phone and email are required"
	}
	return true, ""
}

