package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type UsersCollection struct {
	DB *gorm.DB
}

func (userCollection *UsersCollection) ReadByEmail(email string) *models.Profile {
	account := &models.Account{}

	result := userCollection.DB.Where("email = ?", email).First(account)
	if result.RowsAffected == 0 {
		return nil
	}

	return &account.Profile
}

func (userCollection *UsersCollection) Create(user *models.Account) *models.Profile {
	userCollection.DB.Create(user)

	return &user.Profile
}
