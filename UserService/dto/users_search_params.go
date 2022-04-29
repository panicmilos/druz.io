package dto

import "UserService/models"

type UsersSearchParams struct {
	Name   string
	Gender *models.Gender

	LivePlace string
	WorkPlace string
	Education string
	Interes   string
}
