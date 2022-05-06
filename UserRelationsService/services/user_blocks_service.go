package services

import (
	"github.com/panicmilos/druz.io/UserRelationsService/dto"
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type UserBlocksService struct {
	repository *repository.Repository

	usersService        *UsersService
	userBlockReplicator *UserBlockReplicator
}

func (userBlockService *UserBlocksService) ReadByBlockedById(id uint) *[]models.UserBlock {
	return userBlockService.repository.UserBlocks.ReadByBlockedById(id)
}

func (userBlockService *UserBlocksService) Create(userBlock *models.UserBlock) (*models.UserBlock, error) {

	_, err := userBlockService.usersService.ReadById(uint(userBlock.BlockedById))
	if err != nil {
		return nil, err
	}

	blocked, err := userBlockService.usersService.ReadById(uint(userBlock.BlockedId))
	if err != nil {
		return nil, err
	}

	if userBlock.BlockedById == userBlock.BlockedId {
		return nil, errors.NewErrBadRequest("You can not block yourself.")
	}

	if userBlockService.repository.UserBlocks.ReadByIds(userBlock.BlockedById, userBlock.BlockedId) != nil {
		return nil, errors.NewErrBadRequest("User is already blocked.")
	}

	if userBlockService.repository.UserBlocks.ReadByIds(userBlock.BlockedId, userBlock.BlockedById) != nil {
		return nil, errors.NewErrBadRequest("User has blocked you.")
	}

	userBlock.Blocked = blocked

	userBlockService.DeleteFriendsOrRequest(userBlock)

	userBlockService.userBlockReplicator.Replicate(&dto.UserBlockReplication{
		ReplicationType: "Block",
		UserBlock:       userBlock,
	})

	return userBlockService.repository.UserBlocks.Create(userBlock), nil
}

func (userBlockService *UserBlocksService) DeleteFriendsOrRequest(userBlock *models.UserBlock) {
	friendRequest := userBlockService.repository.FriendRequests.ReadByIds(userBlock.BlockedId, userBlock.BlockedById)
	if friendRequest != nil {
		userBlockService.repository.FriendRequests.Delete(friendRequest.ID)
	}

	friendRequest = userBlockService.repository.FriendRequests.ReadByIds(userBlock.BlockedById, userBlock.BlockedId)
	if friendRequest != nil {
		userBlockService.repository.FriendRequests.Delete(friendRequest.ID)
	}

	userFriend := userBlockService.repository.UserFriends.ReadByIds(userBlock.BlockedId, userBlock.BlockedById)
	if userFriend != nil {
		userBlockService.repository.UserFriends.Delete(userFriend.ID)
	}

	userFriend = userBlockService.repository.UserFriends.ReadByIds(userBlock.BlockedById, userBlock.BlockedId)
	if userFriend != nil {
		userBlockService.repository.UserFriends.Delete(userFriend.ID)
	}
}

func (userBlockService *UserBlocksService) Delete(userBlock *models.UserBlock) (*models.UserBlock, error) {

	existingUserBlock := userBlockService.repository.UserBlocks.ReadByIds(userBlock.BlockedById, userBlock.BlockedId)
	if existingUserBlock == nil {
		return nil, errors.NewErrNotFound("User is not blocked.")
	}

	userBlockService.userBlockReplicator.Replicate(&dto.UserBlockReplication{
		ReplicationType: "Unblock",
		UserBlock:       existingUserBlock,
	})

	return userBlockService.repository.UserBlocks.Delete(existingUserBlock.ID), nil
}
