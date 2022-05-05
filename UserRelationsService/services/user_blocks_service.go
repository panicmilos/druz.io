package services

import (
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type UserBlocksService struct {
	repository *repository.Repository

	usersService *UsersService
}

func (userBlockService *UserBlocksService) ReadByBlockedById(id uint) *[]models.UserBlock {
	return userBlockService.repository.UserBlocks.ReadByBleckedById(id)
}

func (userBlockService *UserBlocksService) Create(userBlock *models.UserBlock) (*models.UserBlock, error) {

	blockedBy, err := userBlockService.usersService.ReadById(uint(userBlock.BlockedById))
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

	userBlock.BlockedBy = blockedBy
	userBlock.Blocked = blocked

	return userBlockService.repository.UserBlocks.Create(userBlock), nil
}

func (userBlockService *UserBlocksService) Delete(userBlock *models.UserBlock) (*models.UserBlock, error) {

	existingUserBlock := userBlockService.repository.UserBlocks.ReadByIds(userBlock.BlockedById, userBlock.BlockedId)
	if existingUserBlock == nil {
		return nil, errors.NewErrNotFound("User is not blocked.")
	}

	return userBlockService.repository.UserBlocks.Delete(existingUserBlock.ID), nil
}
