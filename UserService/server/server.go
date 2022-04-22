package server

import (
	"UserService/controllers"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func New() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}

	server.addHandlers()
	server.addMiddlewares()
	server.addSwagger()

	return server
}

func (server *Server) addHandlers() {
	router := server.Router

	router.HandleFunc("/", controllers.YourGetHandler).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}/password", controllers.ChangePassword).Methods("PUT")
}

func (server *Server) addMiddlewares() {
	router := server.Router

	router.Use(DiMiddleware)
}

func (server *Server) addSwagger() {
	router := server.Router

	router.Handle("/swagger.json", http.FileServer(http.Dir("./docs/")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
}

func (server *Server) Start() {
	router := server.Router

	http.ListenAndServe(":8000", router)
}
