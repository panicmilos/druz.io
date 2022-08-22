package dto

import (
	"strconv"

	"github.com/panicmilos/druz.io/ChatService/models"
)

type UserReplication struct {
	ReplicationType string
	User            *UserDTO
}

type UserDTO struct {
	ID        uint
	FirstName string
	LastName  string
	Image     string
	Disabled  bool
}

func (userDTO *UserDTO) ToModel() *models.User {
	return &models.User{
		ID:        strconv.Itoa(int(userDTO.ID)),
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Image:     userDTO.Image,
		Disabled:  userDTO.Disabled,
	}
}
