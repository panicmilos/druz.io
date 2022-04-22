package dto

import "UserService/models"

type AuthenticatedUser struct {
	Jwt     string
	Profile *models.Profile
}
