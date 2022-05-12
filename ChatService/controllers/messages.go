package controllers

import (
	"net/http"
	"strconv"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/ChatService/api_contracts"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/errors"
	"github.com/panicmilos/druz.io/ChatService/helpers"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/panicmilos/druz.io/ChatService/services"
	"github.com/sarulabs/di"
)

var ChatsWith = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	forId := r.Header.Get("name")

	messagesService := di.Get(r, services.MessageService).(*services.MessagesService)

	helpers.JSONResponse(w, 200, messagesService.ChatsWith(forId))
})

var ReadStatuses = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("name")

	statusesService := di.Get(r, services.StatusService).(*services.StatusesService)

	helpers.JSONResponse(w, 200, statusesService.ReadStatuses(userId))
})

var ReadChat = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	chat := mux.Vars(r)["chat"]
	take, _ := strconv.Atoi(r.URL.Query().Get("take"))
	params := &dto.ChatSearchParams{
		Take:     int(take),
		Keywoard: r.URL.Query().Get("keywoard"),
	}

	messagesService := di.Get(r, services.MessageService).(*services.MessagesService)

	helpers.JSONResponse(w, 200, messagesService.SearchChat(chat, params))
})

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

var DeleteMessage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	chat := params["chat"]
	messageId := params["messageId"]
	mode := r.URL.Query().Get("mode")

	messagesService := di.Get(r, services.MessageService).(*services.MessagesService)
	deletedMessage, err := messagesService.DeleteMessage(chat, messageId, mode)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, deletedMessage)
})

var DeleteChat = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	chat := params["chat"]
	mode := r.URL.Query().Get("mode")

	messagesService := di.Get(r, services.MessageService).(*services.MessagesService)
	deletedChat, err := messagesService.DeleteChat(chat, mode)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, deletedChat)
})
