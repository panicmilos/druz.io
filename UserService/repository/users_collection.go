package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type UsersCollection struct {
	DB *gorm.DB
}

func (userCollection *UsersCollection) ReadAccountByEmail(email string) *models.Account {
	account := &models.Account{}

	result := userCollection.DB.Preload("Profile").Where("email = ?", email).First(account)
	if result.RowsAffected == 0 {
		return nil
	}

	return account
}

func (userCollection *UsersCollection) ReadAccountByProfileId(id uint) *models.Account {
	profile := &models.Profile{}

	result := userCollection.DB.First(profile, id)
	if result.RowsAffected == 0 {
		return nil
	}

	account := &models.Account{}
	result = userCollection.DB.Preload("Profile").First(account, profile.AccountID)
	if result.RowsAffected == 0 {
		return nil
	}

	return account
}

func (userCollection *UsersCollection) Create(user *models.Account) *models.Profile {
	userCollection.DB.Create(user)

	return &user.Profile
}

func (userCollection *UsersCollection) UpdateAccount(account *models.Account) *models.Account {
	userCollection.DB.Save(account)

	return account
}
