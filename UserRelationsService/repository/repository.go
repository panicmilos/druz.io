package repository

import "gorm.io/gorm"

type Repository struct {
	DB             *gorm.DB
	Users          *UsersCollection
	UserBlocks     *UserBlocksCollection
	UserFriends    *UserFriendsCollection
	FriendRequests *FriendRequestsCollection
}
