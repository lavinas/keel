package dto

import (
	"strconv"
)

// GetClientByNicknameInputDto is the input DTO used to get a client by nickname
type GetClientByNicknameInputDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// NewGetClientByNicknameInputDto is the constructor of GetClientByNicknameInputDto
func NewGetClientByNicknameInputDto() *GetClientByNicknameInputDto {
	return &GetClientByNicknameInputDto{
		Id:       "",
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
}

// GetId returns the Id field
func (dto *GetClientByNicknameInputDto) GetId() string {
	return dto.Id
}

// GetName returns the Name field
func (dto *GetClientByNicknameInputDto) GetName() string {
	return dto.Name
}

// GetNickname returns the Nickname field
func (dto *GetClientByNicknameInputDto) GetNickname() string {
	return dto.Nickname
}

// GetDocument returns the Document field
func (dto *GetClientByNicknameInputDto) GetDocument() (uint64, error) {
	idoc, error := strconv.ParseUint(dto.Document, 10, 64)
	if error != nil {
		return 0, error
	}
	return idoc, nil
}

// GetPhone returns the Phone field
func (dto *GetClientByNicknameInputDto) GetPhone() (uint64, error) {
	iphone, err := strconv.ParseUint(dto.Phone, 10, 64)
	if err != nil {
		return 0, err
	}
	return iphone, nil
}

// GetEmail returns the Email field
func (dto *GetClientByNicknameInputDto) GetEmail() string {
	return dto.Email
}
