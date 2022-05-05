package repository

import (
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"gorm.io/gorm"
)

type UserBlocksCollection struct {
	DB *gorm.DB
}

func (userBlocksCollection *UserBlocksCollection) ReadById(id uint) *models.UserBlock {
	userBlock := &models.UserBlock{}

	query := userBlocksCollection.DB.Table("user_blocks")
	addUserBlocksFilters(query)
	query.Preload("Blocked").First(userBlock, id)

	return userBlock
}

func (userBlocksCollection *UserBlocksCollection) ReadByBlockedById(blockedById uint) *[]models.UserBlock {
	userBlocks := &[]models.UserBlock{}

	query := userBlocksCollection.DB.Table("user_blocks")
	addUserBlocksFilters(query)
	query.Preload("Blocked").Where("blocked_by_id = ?", blockedById).Find(userBlocks)

	return userBlocks
}

func (userBlocksCollection *UserBlocksCollection) ReadByIds(blockedById uint, blockedId uint) *models.UserBlock {
	userBlock := &models.UserBlock{}

	query := userBlocksCollection.DB.Table("user_blocks")
	addUserBlocksFilters(query)
	result := query.Preload("Blocked").Where("blocked_by_id = ? AND blocked_id = ?", blockedById, blockedId).First(userBlock)
	if result.RowsAffected == 0 {
		return nil
	}

	return userBlock
}

func addUserBlocksFilters(query *gorm.DB) {
	query.Joins("JOIN users u ON user_blocks.blocked_by_id = u.id").Where("(u.disabled is NULL OR u.disabled = 0) AND u.deleted_at is NULL")
	query.Joins("JOIN users u2 ON user_blocks.blocked_id = u2.id").Where("(u2.disabled is NULL OR u2.disabled = 0) AND u2.deleted_at is NULL")
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
