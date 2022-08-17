package controllers

import (
	"net/http"
	"strconv"

	"github.com/panicmilos/druz.io/UserService/api_contracts"
	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"
	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/services"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

var SearchUsers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	gender, err := models.GenderFromString(r.URL.Query().Get("Gender"))

	params := &dto.UsersSearchParams{
		Name:      r.URL.Query().Get("Name"),
		Gender:    map[bool]*models.Gender{true: &gender, false: nil}[err == nil],
		LivePlace: r.URL.Query().Get("LivePlace"),
		WorkPlace: r.URL.Query().Get("WorkPlace"),
		Education: r.URL.Query().Get("Education"),
		Interes:   r.URL.Query().Get("Interes"),
	}

	userService := di.Get(r, services.UsersService).(*services.UserService)

	helpers.JSONResponse(w, 200, userService.Search(params))
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
	dtoMapper.Map(user, request)

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
	dtoMapper.Map(profile, request)
	profile.ID = uint(id)

	userService := di.Get(r, services.UsersService).(*services.UserService)

	updatedUser, err := userService.Update(profile)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, updatedUser)
})

var ChangeImage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.ChangeImageRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := di.Get(r, services.UsersService).(*services.UserService)
	updatedUser, err := userService.ChangeImage(uint(id), request.Image)
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
	updatedUser, err := userService.ChangePassword(uint(id), request.CurrentPassword, request.NewPassword)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, updatedUser)
})

var BlockUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := di.Get(r, services.UsersService).(*services.UserService)

	blockedUser, err := userService.Delete(uint(id))
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, blockedUser)
})

var DisableUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := di.Get(r, services.UsersService).(*services.UserService)

	disabledUser, err := userService.Disable(uint(id))
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, disabledUser)
})
