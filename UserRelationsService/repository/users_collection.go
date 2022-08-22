package repository

import (
	"github.com/panicmilos/druz.io/UserRelationsService/models"

	"gorm.io/gorm"
)

type UsersCollection struct {
	DB *gorm.DB
}

func (userCollection *UsersCollection) ReadById(id uint) *models.User {
	user := &models.User{}

	result := userCollection.DB.First(user, id)
	if result.RowsAffected == 0 || user.Disabled {
		return nil
	}

	return user
}

func (userCollection *UsersCollection) ReadByIdEvenDisabled(id uint) *models.User {
	user := &models.User{}

	result := userCollection.DB.First(user, id)
	if result.RowsAffected == 0 {
		return nil
	}

	return user
}

func (userCollection *UsersCollection) Create(user *models.User) *models.User {
	userCollection.DB.Create(user)

	return user
}

func (userCollection *UsersCollection) Update(user *models.User) *models.User {
	userCollection.DB.Save(user)

	return user
}

func (userCollection *UsersCollection) Delete(id uint) *models.User {
	user := userCollection.ReadByIdEvenDisabled(id)

	userCollection.DB.Delete(user)

	return user
}

func (userCollection *UsersCollection) Disable(id uint) *models.User {
	user := userCollection.ReadByIdEvenDisabled(id)

	user.Disabled = true

	userCollection.DB.Save(user)

	return user
}

func (userCollection *UsersCollection) Reactivate(id uint) *models.User {
	user := userCollection.ReadByIdEvenDisabled(id)

	user.Disabled = false

	userCollection.DB.Save(user)

	return user
}
