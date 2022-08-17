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

var ReadFriendsList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userFriendsService := di.Get(r, services.UserFriendService).(*services.UserFriendsService)

	helpers.JSONResponse(w, 200, userFriendsService.ReadByUserId(uint(id)))
})

var ReadByIds = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["id"])
	friendId, _ := strconv.Atoi(mux.Vars(r)["friendId"])

	userFriend := &models.UserFriend{
		UserId:   uint(userId),
		FriendId: uint(friendId),
	}

	userFriendsService := di.Get(r, services.UserFriendService).(*services.UserFriendsService)
	existingUserFriend, err := userFriendsService.ReadByIds(userFriend)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, existingUserFriend)
})

var UnfriendUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.UnfriendUserRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	userId, _ := strconv.Atoi(mux.Vars(r)["id"])

	userFriend := &models.UserFriend{
		UserId:   uint(userId),
		FriendId: uint(request.FriendId),
	}

	userFriendsService := di.Get(r, services.UserFriendService).(*services.UserFriendsService)
	deletedUserFriend, err := userFriendsService.Delete(userFriend)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, deletedUserFriend)
})
