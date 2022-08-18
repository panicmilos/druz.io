package server

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/panicmilos/druz.io/UserRelationsService/controllers"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/rs/cors"
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

	router.Handle("/users/{id}/block-list", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadBlockList, models.NormalUser))).Methods("GET")
	router.Handle("/users/{id}/block-list", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.BlockUser, models.NormalUser))).Methods("POST")
	router.Handle("/users/{id}/block-list", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.UnblockUser, models.NormalUser))).Methods("DELETE")

	router.Handle("/users/{id}/friends", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadFriendsList, models.NormalUser))).Methods("GET")
	router.Handle("/users/{id}/friends/{friendId}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadByIds, models.NormalUser))).Methods("GET")
	router.Handle("/users/{id}/friends", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.UnfriendUser, models.NormalUser))).Methods("DELETE")

	router.Handle("/users/{id}/friends/requests/sent", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadSentFriendRequests, models.NormalUser))).Methods("GET")
	router.Handle("/users/{id}/friends/requests/sent", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.DeleteSentFriendRequests, models.NormalUser))).Methods("DELETE")
	router.Handle("/users/{id}/friends/requests/received", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadReceivedFriendRequests, models.NormalUser))).Methods("GET")
	router.Handle("/users/{id}/friends/requests", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.SendFriendRequests, models.NormalUser))).Methods("POST")
	router.Handle("/users/{id}/friends/requests/accept", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.AcceptFriendRequest, models.NormalUser))).Methods("POST")
	router.Handle("/users/{id}/friends/requests/decline", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.DeclineFriendRequest, models.NormalUser))).Methods("DELETE")

}

func (server *Server) addMiddlewares() {
	router := server.Router

	router.Use(DiMiddleware)
	router.Use(AccessControlMiddleware)
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
