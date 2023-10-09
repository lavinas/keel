package dto

// UpdateOutputDto is the output DTO used to update a client
type UpdateOutputDto struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// NewUpdateOutputDto creates a new output DTO for update service
func NewUpdateOutputDto() *UpdateOutputDto {
	return &UpdateOutputDto{
		Id:       "",
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
}

// Fill fills the output DTO with the given values
func (o *UpdateOutputDto) Fill(id, name, nick, doc, phone, email string) {
	o.Id = id
	o.Name = name
	o.Nickname = nick
	o.Document = doc
	o.Phone = phone
	o.Email = email
}
