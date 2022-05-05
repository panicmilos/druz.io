package repository

import (
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"gorm.io/gorm"
)

type UserFriendsCollection struct {
	DB *gorm.DB
}

func (userFriendsCollection *UserFriendsCollection) ReadById(id uint) *models.UserFriend {
	userFriend := &models.UserFriend{}

	userFriendsCollection.DB.Preload("Friend").First(userFriend, id)

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) ReadByUserId(userId uint) *[]models.UserFriend {
	userFriends := &[]models.UserFriend{}

	userFriendsCollection.DB.Preload("Friend").Where("user_id = ?", userId).Find(userFriends)

	return userFriends
}

func (userFriendsCollection *UserFriendsCollection) ReadByIds(userId uint, friendId uint) *models.UserFriend {
	userFriend := &models.UserFriend{}

	result := userFriendsCollection.DB.Preload("Friend").Where("user_id = ? AND friend_id = ?", userId, friendId).First(userFriend)
	if result.RowsAffected == 0 {
		return nil
	}

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) Create(userFriend *models.UserFriend) *models.UserFriend {
	userFriendsCollection.DB.Create(userFriend)

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) Delete(id uint) *models.UserFriend {
	userFriend := userFriendsCollection.ReadById(id)

	userFriendsCollection.DB.Delete(userFriend)

	return userFriend
}
