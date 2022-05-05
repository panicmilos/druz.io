package repository

import (
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"gorm.io/gorm"
)

type FriendRequestsCollection struct {
	DB *gorm.DB
}

func (friendRequestsCollection *FriendRequestsCollection) ReadById(id uint) *models.FriendRequest {
	friendRequest := &models.FriendRequest{}

	friendRequestsCollection.DB.Preload("User").Preload("Friend").First(friendRequest, id)

	return friendRequest
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByUserId(userId uint) *[]models.FriendRequest {
	friendRequests := &[]models.FriendRequest{}

	friendRequestsCollection.DB.Preload("Friend").Where("user_id = ?", userId).Find(friendRequests)

	return friendRequests
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByFriendId(friendId uint) *[]models.FriendRequest {
	friendRequests := &[]models.FriendRequest{}

	friendRequestsCollection.DB.Preload("User").Where("friend_id = ?", friendId).Find(friendRequests)

	return friendRequests
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByIds(userId uint, friendId uint) *models.FriendRequest {
	friendRequest := &models.FriendRequest{}

	result := friendRequestsCollection.DB.Preload("User").Preload("Friend").Where("user_id = ? AND friend_id = ?", userId, friendId).First(friendRequest)
	if result.RowsAffected == 0 {
		return nil
	}

	return friendRequest
}

func (friendRequestsCollection *FriendRequestsCollection) Create(friendRequest *models.FriendRequest) *models.FriendRequest {
	friendRequestsCollection.DB.Create(friendRequest)

	return friendRequest
}

func (friendRequestsCollection *FriendRequestsCollection) Delete(id uint) *models.FriendRequest {
	friendRequest := friendRequestsCollection.ReadById(id)

	friendRequestsCollection.DB.Delete(friendRequest)

	return friendRequest
}
