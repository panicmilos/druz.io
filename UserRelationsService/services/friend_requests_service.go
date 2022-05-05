package services

import (
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type FriendRequestsService struct {
	repository *repository.Repository

	usersService       *UsersService
	userFriendsService *UserFriendsService
}

func (friendRequestService *FriendRequestsService) ReadByUserId(id uint) *[]models.FriendRequest {
	return friendRequestService.repository.FriendRequests.ReadByUserId(id)
}

func (friendRequestService *FriendRequestsService) ReadByFriendId(id uint) *[]models.FriendRequest {
	return friendRequestService.repository.FriendRequests.ReadByFriendId(id)
}

func (friendRequestService *FriendRequestsService) Create(friendRequest *models.FriendRequest) (*models.FriendRequest, error) {
	_, err := friendRequestService.usersService.ReadById(uint(friendRequest.UserId))
	if err != nil {
		return nil, err
	}

	friend, err := friendRequestService.usersService.ReadById(uint(friendRequest.FriendId))
	if err != nil {
		return nil, err
	}

	if friendRequest.UserId == friendRequest.FriendId {
		return nil, errors.NewErrBadRequest("You can not send friend requests to yourself.")
	}

	if friendRequestService.repository.UserBlocks.ReadByIds(friendRequest.UserId, friendRequest.FriendId) != nil {
		return nil, errors.NewErrBadRequest("You have blocked that user.")
	}

	if friendRequestService.repository.UserBlocks.ReadByIds(friendRequest.FriendId, friendRequest.UserId) != nil {
		return nil, errors.NewErrBadRequest("You are blocked by that user.")
	}

	if friendRequestService.repository.UserFriends.ReadByIds(friendRequest.UserId, friendRequest.FriendId) != nil {
		return nil, errors.NewErrBadRequest("You are already friend with that user.")
	}

	if friendRequestService.repository.FriendRequests.ReadByIds(friendRequest.UserId, friendRequest.FriendId) != nil {
		return nil, errors.NewErrBadRequest("You have already sent friend request to that user.")
	}

	if friendRequestService.repository.FriendRequests.ReadByIds(friendRequest.FriendId, friendRequest.UserId) != nil {
		return nil, errors.NewErrBadRequest("You have friend request from that user.")
	}

	friendRequest.Friend = friend

	return friendRequestService.repository.FriendRequests.Create(friendRequest), nil
}
