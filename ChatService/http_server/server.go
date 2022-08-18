package http_server

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/ChatService/controllers"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/rs/cors"
)

type HttpServer struct {
	Router *mux.Router
}

func New() *HttpServer {
	server := &HttpServer{
		Router: mux.NewRouter(),
	}

	server.addHandlers()
	server.addMiddlewares()
	server.addSwagger()

	return server
}

func (server *HttpServer) addHandlers() {
	router := server.Router

	router.Handle("/users/{id}/message", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.SendMessage, models.NormalUser))).Methods("POST")
	router.Handle("/users/chats", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ChatsWith, models.NormalUser))).Methods("GET")
	router.Handle("/users/chats/statuses", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadStatuses, models.NormalUser))).Methods("GET")
	router.Handle("/users/chats/{chat}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadChat, models.NormalUser))).Methods("GET")
	router.Handle("/users/chats/{chat}/{messageId}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.DeleteMessage, models.NormalUser))).Methods("DELETE")
	router.Handle("/users/chats/{chat}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.DeleteChat, models.NormalUser))).Methods("DELETE")
}

func (server *HttpServer) addMiddlewares() {
	router := server.Router

	router.Use(DiMiddleware)
	router.Use(AccessControlMiddleware)
}

func (server *HttpServer) addSwagger() {
	router := server.Router

	router.Handle("/swagger.json", http.FileServer(http.Dir("./docs/")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
}

func (server *HttpServer) Start() {
	router := server.Router

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})

	http.ListenAndServe(os.Getenv("PORT"), corsOpts.Handler(router))
}
