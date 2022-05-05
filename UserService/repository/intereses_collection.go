package repository

import (
	"github.com/panicmilos/druz.io/UserService/models"

	"gorm.io/gorm"
)

type InteresesCollection struct {
	DB *gorm.DB
}

func (interesesCollection *InteresesCollection) DeleteByProfileId(profileId uint) *[]models.Interes {
	intereses := &[]models.Interes{}

	interesesCollection.DB.Where("profile_id = ?", profileId).Delete(intereses)

	return intereses
}
