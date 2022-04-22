package controllers

import (
	"UserService/api_contracts"
	"UserService/errors"
	"UserService/helpers"
	"UserService/models"
	"UserService/services"
	"net/http"
	"strconv"

	"github.com/devfeel/mapper"
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
func YourGetHandler(w http.ResponseWriter, r *http.Request) {
	// service := di.Get(r, services.UsersService).(*services.UserService)

	// helpers.JSONResponse(w, 200, service.ReadUsers())
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	mapper.Mapper(request, user)

	userService := di.Get(r, services.UsersService).(*services.UserService)
	createdUser, err := userService.Create(user)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdUser)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
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
}
