package dto

import "github.com/panicmilos/druz.io/UserService/models"

type AuthenticatedUser struct {
	Jwt     string
	Profile *models.Profile
}
