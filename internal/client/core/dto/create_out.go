package dto

// CreateOutputDto is the output DTO used to create a client
type CreateOutputDto struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func  NewCreateOutputDto(id, name, nickname, document, phone, email string) CreateOutputDto {
	return CreateOutputDto{
		Id:       id,
		Name:     name,
		Nickname: nickname,
		Document: document,
		Phone:    phone,
		Email:    email,
	}
}