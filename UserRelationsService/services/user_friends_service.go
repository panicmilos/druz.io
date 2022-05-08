package services

import (
	"github.com/panicmilos/druz.io/UserRelationsService/dto"
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type UserFriendsService struct {
	repository *repository.Repository

	UserFriendReplicator *UserFriendReplicator
}

func (userFriendService *UserFriendsService) ReadByUserId(id uint) *[]models.UserFriend {
	return userFriendService.repository.UserFriends.ReadByUserId(id)
}

func (userFriendService *UserFriendsService) Create(userFriend *models.UserFriend) (*models.UserFriend, error) {

	reversedUserFriend := &models.UserFriend{
		UserId:   userFriend.FriendId,
		FriendId: userFriend.UserId,
	}

	userFriendService.repository.UserFriends.Create(userFriend)
	userFriendService.repository.UserFriends.Create(reversedUserFriend)

	userFriendService.UserFriendReplicator.Replicate(&dto.UserFriendReplication{
		ReplicationType: "Add",
		UserFriend:      userFriend,
	})

	userFriendService.UserFriendReplicator.Replicate(&dto.UserFriendReplication{
		ReplicationType: "Add",
		UserFriend:      reversedUserFriend,
	})

	return userFriend, nil
}

func (userFriendService *UserFriendsService) Delete(userFriend *models.UserFriend) (*models.UserFriend, error) {

	existingUserFriend := userFriendService.repository.UserFriends.ReadByIds(userFriend.UserId, userFriend.FriendId)
	if existingUserFriend == nil {
		return nil, errors.NewErrNotFound("You are not friend with that user.")
	}

	existingUserFriendReversed := userFriendService.repository.UserFriends.ReadByIds(userFriend.FriendId, userFriend.UserId)
	userFriendService.repository.UserFriends.Delete(existingUserFriendReversed.ID)
	userFriendService.repository.UserFriends.Delete(existingUserFriend.ID)

	userFriendService.UserFriendReplicator.Replicate(&dto.UserFriendReplication{
		ReplicationType: "Remove",
		UserFriend:      existingUserFriend,
	})

	userFriendService.UserFriendReplicator.Replicate(&dto.UserFriendReplication{
		ReplicationType: "Remove",
		UserFriend:      existingUserFriendReversed,
	})

	return existingUserFriend, nil
}
