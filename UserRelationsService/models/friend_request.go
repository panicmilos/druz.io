package models

import "gorm.io/gorm"

type FriendRequest struct {
	*gorm.Model

	UserId uint
	User   *User

	FriendId uint
	Friend   *User
}
