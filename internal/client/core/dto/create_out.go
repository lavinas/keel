package dto

// CreateOutputDto is the output DTO used to create a client
type CreateOutputDto struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document uint64 `json:"document" binding:"required"`
	Phone    uint64 `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
