package repository

import (
	"fmt"

	"github.com/panicmilos/druz.io/UserService/models"
	"gorm.io/gorm"
)

type UserBlocksCollection struct {
	DB *gorm.DB
}

func (userBlocksCollection *UserBlocksCollection) ReadById(id uint) *models.UserBlock {
	userBlock := &models.UserBlock{}

	result := userBlocksCollection.DB.First(userBlock, id)
	if result.RowsAffected == 0 {
		return nil
	}

	return userBlock
}

func (userBlocksCollection *UserBlocksCollection) Create(userBlock *models.UserBlock) *models.UserBlock {
	userBlocksCollection.DB.Create(userBlock)

	return userBlock
}

func (userBlocksCollection *UserBlocksCollection) Delete(id uint) *models.UserBlock {
	fmt.Println(id)
	userBlock := userBlocksCollection.ReadById(id)

	userBlocksCollection.DB.Delete(userBlock)

	return userBlock
}
