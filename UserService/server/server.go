package server

import (
	"net/http"
	"os"

	"github.com/panicmilos/druz.io/UserService/controllers"
	"github.com/rs/cors"

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

	router.Handle("/users/search", AuthenticateMiddlewate(controllers.SearchUsers)).Methods("GET")
	router.Handle("/users/{id}", AuthenticateMiddlewate(controllers.ReadUserById)).Methods("GET")
	router.Handle("/users", controllers.CreateUser).Methods("POST")
	router.Handle("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.Handle("/users/{id}/image", controllers.ChangeImage).Methods("PUT")
	router.Handle("/users/{id}/password", controllers.ChangePassword).Methods("PUT")
	router.Handle("/users/{id}/block", controllers.BlockUser).Methods("DELETE")

	router.Handle("/users/{id}/disable", controllers.DisableUser).Methods("PUT")
	router.Handle("/users/reactivation/request", controllers.UserReactivation).Methods("POST")
	router.Handle("/users/{id}/reactivation", controllers.ReactivateUser).Methods("PUT")

	router.Handle("/users/password/recover/request", controllers.PasswordRecovery).Methods("POST")
	router.Handle("/users/{id}/password/recover", controllers.RecoverPassword).Methods("PUT")
	router.Handle("/users/{id}/report", AuthenticateMiddlewate(controllers.ReportUser)).Methods("POST")

	router.Handle("/reports/search", controllers.SearchReports).Methods("GET")
	router.Handle("/reports/{id}/ignore", controllers.IgnoreReport).Methods("DELETE")

	router.Handle("/auth", controllers.Authenticate).Methods("POST")

	router.Handle("/authorize", AuthenticateMiddlewate(controllers.Authorize)).Methods("POST")
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
