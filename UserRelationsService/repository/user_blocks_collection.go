package repository

import (
	"fmt"

	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"gorm.io/gorm"
)

type UserBlocksCollection struct {
	DB *gorm.DB
}

func (userBlocksCollection *UserBlocksCollection) ReadById(id uint) *models.UserBlock {
	userBlock := &models.UserBlock{}

	userBlocksCollection.DB.Preload("Blocked").Preload("BlockedBy").First(userBlock, id)

	return userBlock
}

func (userBlocksCollection *UserBlocksCollection) ReadByBleckedById(blockedById uint) *[]models.UserBlock {
	userBlock := &[]models.UserBlock{}

	userBlocksCollection.DB.Preload("Blocked").Preload("BlockedBy").Where("blocked_by_id = ?", blockedById).Find(userBlock)

	return userBlock
}

func (userBlocksCollection *UserBlocksCollection) ReadByIds(blockedById uint, blockedId uint) *models.UserBlock {
	userBlock := &models.UserBlock{}

	result := userBlocksCollection.DB.Preload("Blocked").Preload("BlockedBy").Where("blocked_by_id = ? AND blocked_id = ?", blockedById, blockedId).First(userBlock)
	fmt.Println(result.RowsAffected)
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
	userBlock := userBlocksCollection.ReadById(id)

	userBlocksCollection.DB.Delete(userBlock)

	return userBlock
}
