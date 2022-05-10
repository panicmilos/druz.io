package controllers

import (
	"net/http"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/ChatService/api_contracts"
	"github.com/panicmilos/druz.io/ChatService/errors"
	"github.com/panicmilos/druz.io/ChatService/helpers"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/panicmilos/druz.io/ChatService/services"
	"github.com/sarulabs/di"
)

var SendMessage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.SendMessageRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	toId := mux.Vars(r)["id"]
	fromId := r.Header.Get("name")

	message := &models.Message{
		ToId:   toId,
		FromId: fromId,
	}
	dtoMapper.Map(message, request)

	messagesService := di.Get(r, services.MessageService).(*services.MessagesService)
	createdMessage, err := messagesService.Create(message)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdMessage)
})
