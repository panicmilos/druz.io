package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/UserRelationsService/api_contracts"
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/helpers"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/services"
	"github.com/sarulabs/di"
)

var ReadBlockList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userBlocksService := di.Get(r, services.UserBlockService).(*services.UserBlocksService)

	helpers.JSONResponse(w, 200, userBlocksService.ReadByBlockedById(uint(id)))
})

var BlockUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.BlockUserRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	blockedById, _ := strconv.Atoi(mux.Vars(r)["id"])

	userBlock := &models.UserBlock{
		BlockedId:   uint(request.BlockedId),
		BlockedById: uint(blockedById),
	}

	userBlocksService := di.Get(r, services.UserBlockService).(*services.UserBlocksService)
	createdUserBlock, err := userBlocksService.Create(userBlock)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdUserBlock)
})

var UnblockUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.UnblockUserRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	blockedById, _ := strconv.Atoi(mux.Vars(r)["id"])

	userBlock := &models.UserBlock{
		BlockedId:   uint(request.BlockedId),
		BlockedById: uint(blockedById),
	}

	userBlocksService := di.Get(r, services.UserBlockService).(*services.UserBlocksService)
	deletedUserBlock, err := userBlocksService.Delete(userBlock)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, deletedUserBlock)
})
