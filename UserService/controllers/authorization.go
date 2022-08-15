package controllers

import (
	"net/http"

	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/helpers"
	"github.com/panicmilos/druz.io/UserService/services"
	"github.com/sarulabs/di"
)

var Authorize = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	sessionStorage := di.Get(r, services.SessionStorage).(*helpers.SessionStorage)

	helpers.JSONResponse(w, 200, &dto.AuthorizedUser{
		Id:   sessionStorage.AuthenticatedUserId,
		Role: sessionStorage.Role,
	})
})
