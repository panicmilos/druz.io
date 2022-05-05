package models

import "gorm.io/gorm"

type UserFriend struct {
	*gorm.Model

	UserId uint

	FriendId uint
	Friend   *User
}
