package controllers

import (
	"UserService/api_contracts"
	"UserService/errors"
	"UserService/helpers"
	"UserService/models"
	"UserService/services"
	"net/http"
	"strconv"

	"github.com/dranikpg/dto-mapper"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// swagger:route DELETE /admin/company/{id} admin deleteCompany
// Delete company
//
// security:
// - Bearer: []
// responses:
//  401: Account
//  200: Account
// Create handles Delete get company
var YourGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// service := di.Get(r, services.UsersService).(*services.UserService)

	// helpers.JSONResponse(w, 200, service.ReadUsers())
})

var ReadUserById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := di.Get(r, services.UsersService).(*services.UserService)
	user, err := userService.ReadById(uint(id))
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, user)
})

var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.CreateAccountRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	user := &models.Account{}
	dto.Map(user, request)

	userService := di.Get(r, services.UsersService).(*services.UserService)
	createdUser, err := userService.Create(user)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdUser)
})

var UpdateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.UpdateProfileRequest
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

	profile := &models.Profile{}
	dto.Map(profile, request)
	profile.ID = uint(id)

	userService := di.Get(r, services.UsersService).(*services.UserService)

	updatedUser, err := userService.Update(profile)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, updatedUser)
})

var ChangePassword = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.ChangePasswordRequest
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

	userService := di.Get(r, services.UsersService).(*services.UserService)
	createdUser, err := userService.ChangePassword(uint(id), request.CurrentPassword, request.NewPassword)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdUser)
})
