package controllers

import (
	"UserService/api_contracts"
	"UserService/errors"
	"UserService/helpers"
	"UserService/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

var UserReactivation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var request *api_contracts.UserReactivationRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	userReactivationsService := di.Get(r, services.UserReactivationService).(*services.UserReactivationsService)
	_, err2 := userReactivationsService.Request(request.Email)
	if errors.Handle(err2, w) {
		return
	}

	helpers.JSONResponse(w, 200, nil)
})

var ReactivateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	token := r.URL.Query().Get("token")

	userReactivationsService := di.Get(r, services.UserReactivationService).(*services.UserReactivationsService)
	updatedUser, err := userReactivationsService.Reactivate(uint(id), token)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, updatedUser)
})
