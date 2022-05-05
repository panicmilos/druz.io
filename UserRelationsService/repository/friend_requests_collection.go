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

	query := friendRequestsCollection.DB.Table("friend_requests")
	addFriendRequestsFilters(query)
	query.Preload("User").Preload("Friend").First(friendRequest, id)

	return friendRequest
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByUserId(userId uint) *[]models.FriendRequest {
	friendRequests := &[]models.FriendRequest{}

	query := friendRequestsCollection.DB.Table("friend_requests")
	addFriendRequestsFilters(query)
	query.Preload("Friend").Where("user_id = ?", userId).Find(friendRequests)

	return friendRequests
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByFriendId(friendId uint) *[]models.FriendRequest {
	friendRequests := &[]models.FriendRequest{}

	query := friendRequestsCollection.DB.Table("friend_requests")
	addFriendRequestsFilters(query)
	query.Preload("User").Where("friend_id = ?", friendId).Find(friendRequests)

	return friendRequests
}

func (friendRequestsCollection *FriendRequestsCollection) ReadByIds(userId uint, friendId uint) *models.FriendRequest {
	friendRequest := &models.FriendRequest{}

	query := friendRequestsCollection.DB.Table("friend_requests")
	addFriendRequestsFilters(query)
	result := query.Preload("User").Preload("Friend").Where("user_id = ? AND friend_id = ?", userId, friendId).First(friendRequest)
	if result.RowsAffected == 0 {
		return nil
	}

	return friendRequest
}

func addFriendRequestsFilters(query *gorm.DB) {
	query.Joins("JOIN users u ON friend_requests.user_id = u.id").Where("(u.disabled is NULL OR u.disabled = 0) AND u.deleted_at is NULL")
	query.Joins("JOIN users u2 ON friend_requests.friend_id = u2.id").Where("(u2.disabled is NULL OR u2.disabled = 0) AND u2.deleted_at is NULL")
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
