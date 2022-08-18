package http_server

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/ChatService/controllers"
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

	router.Handle("/users/{id}/message", AuthenticateMiddlewate(controllers.SendMessage)).Methods("POST")
	router.Handle("/users/chats", AuthenticateMiddlewate(controllers.ChatsWith)).Methods("GET")
	router.Handle("/users/chats/statuses", AuthenticateMiddlewate(controllers.ReadStatuses)).Methods("GET")
	router.Handle("/users/chats/{chat}", AuthenticateMiddlewate(controllers.ReadChat)).Methods("GET")
	router.Handle("/users/chats/{chat}/{messageId}", AuthenticateMiddlewate(controllers.DeleteMessage)).Methods("DELETE")
	router.Handle("/users/chats/{chat}", AuthenticateMiddlewate(controllers.DeleteChat)).Methods("DELETE")
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
