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

	query := userFriendsCollection.DB.Table("user_friends")
	addUserFriendsFilters(query)
	query.Preload("Friend").First(userFriend, id)

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) ReadByUserId(userId uint) *[]models.UserFriend {
	userFriends := &[]models.UserFriend{}

	query := userFriendsCollection.DB.Table("user_friends")
	addUserFriendsFilters(query)
	query.Preload("Friend").Where("user_id = ?", userId).Find(userFriends)

	return userFriends
}

func (userFriendsCollection *UserFriendsCollection) ReadByIds(userId uint, friendId uint) *models.UserFriend {
	userFriend := &models.UserFriend{}

	query := userFriendsCollection.DB.Table("user_friends")
	addUserFriendsFilters(query)
	result := query.Preload("Friend").Where("user_id = ? AND friend_id = ?", userId, friendId).First(userFriend)
	if result.RowsAffected == 0 {
		return nil
	}

	return userFriend
}

func addUserFriendsFilters(query *gorm.DB) {
	query.Joins("JOIN users u ON user_friends.user_id = u.id").Where("(u.disabled is NULL OR u.disabled = 0) AND u.deleted_at is NULL")
	query.Joins("JOIN users u2 ON user_friends.friend_id = u2.id").Where("(u2.disabled is NULL OR u2.disabled = 0) AND u2.deleted_at is NULL")
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
