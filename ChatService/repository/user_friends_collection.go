package repository

import (
	"fmt"

	"github.com/panicmilos/druz.io/ChatService/models"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type UserFriendsCollection struct {
	Session *ravendb.DocumentSession
}

const user_friends_prefix = "friends"

func formUserFriendsKey(userId string, friendId string) string {
	return fmt.Sprintf("%s/%s/%s", user_friends_prefix, userId, friendId)
}

func (userFriendsCollection *UserFriendsCollection) ReadByIds(userID string, friendID string) *models.UserFriend {
	userFriend := &models.UserFriend{}

	err := userFriendsCollection.Session.Include(formUsersKey(friendID)).Load(&userFriend, formUserFriendsKey(userID, friendID))
	if err != nil || (userFriend != nil && userFriend.ID == "") {
		return nil
	}

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) ReadByUserId(userId string) []*models.UserFriend {
	q := userFriendsCollection.Session.QueryCollection("UserFriends")
	q.WhereEquals("UserId", userId)

	var userFriends []*models.UserFriend
	q.GetResults(&userFriends)

	return userFriends
}

func (userFriendsCollection *UserFriendsCollection) Create(userFriend *models.UserFriend) *models.UserFriend {
	userFriend.ID = formUserFriendsKey(userFriend.UserId, userFriend.FriendId)

	userFriendsCollection.Session.Store(userFriend)
	userFriendsCollection.Session.SaveChanges()

	return userFriend
}

func (userFriendsCollection *UserFriendsCollection) Delete(userID string, friendID string) *models.UserFriend {
	userFriend := userFriendsCollection.ReadByIds(userID, friendID)

	userFriendsCollection.Session.Delete(userFriend)

	return userFriend
}
