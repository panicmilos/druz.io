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

var ReadReceivedFriendRequests = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	friendRequestService := di.Get(r, services.FriendRequestService).(*services.FriendRequestsService)

	helpers.JSONResponse(w, 200, friendRequestService.ReadByFriendId(uint(id)))
})

var ReadSentFriendRequests = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	friendRequestService := di.Get(r, services.FriendRequestService).(*services.FriendRequestsService)

	helpers.JSONResponse(w, 200, friendRequestService.ReadByUserId(uint(id)))
})

var SendFriendRequests = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.SendFriendRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	userId, _ := strconv.Atoi(mux.Vars(r)["id"])

	friendRequest := &models.FriendRequest{
		UserId:   uint(userId),
		FriendId: uint(request.FriendId),
	}

	friendRequestService := di.Get(r, services.FriendRequestService).(*services.FriendRequestsService)
	createdFriendRequest, err := friendRequestService.Create(friendRequest)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdFriendRequest)
})
