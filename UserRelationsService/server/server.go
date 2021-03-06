package server

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/UserRelationsService/controllers"
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

	router.Handle("/users/{id}/block-list", controllers.ReadBlockList).Methods("GET")
	router.Handle("/users/{id}/block-list", controllers.BlockUser).Methods("POST")
	router.Handle("/users/{id}/block-list", controllers.UnblockUser).Methods("DELETE")

	router.Handle("/users/{id}/friends", controllers.ReadFriendsList).Methods("GET")
	router.Handle("/users/{id}/friends", controllers.UnfriendUser).Methods("DELETE")

	router.Handle("/users/{id}/friends/requests/sent", controllers.ReadSentFriendRequests).Methods("GET")
	router.Handle("/users/{id}/friends/requests/received", controllers.ReadReceivedFriendRequests).Methods("GET")
	router.Handle("/users/{id}/friends/requests", controllers.SendFriendRequests).Methods("POST")
	router.Handle("/users/{id}/friends/requests/accept", controllers.AcceptFriendRequest).Methods("POST")
	router.Handle("/users/{id}/friends/requests/decline", controllers.DeclineFriendRequest).Methods("DELETE")

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

	http.ListenAndServe(":8001", router)
}
