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

var PasswordRecovery = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var request *api_contracts.PasswordRecoveryRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	passwordRecoveryService := di.Get(r, services.PasswordRecoveriesService).(*services.PasswordRecoveryService)
	_, err2 := passwordRecoveryService.Request(request.Email)
	if errors.Handle(err2, w) {
		return
	}

	helpers.JSONResponse(w, 200, nil)
})

var RecoverPassword = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var request *api_contracts.RecoverPasswordRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	token := r.URL.Query().Get("token")

	passwordRecoveryService := di.Get(r, services.PasswordRecoveriesService).(*services.PasswordRecoveryService)
	updatedUser, err := passwordRecoveryService.Recover(uint(id), token, request.NewPassword)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, updatedUser)
})
