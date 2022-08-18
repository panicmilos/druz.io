package server

import (
	"net/http"
	"os"

	"github.com/panicmilos/druz.io/UserService/controllers"
	"github.com/panicmilos/druz.io/UserService/models"
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

	router.Handle("/users/search", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.SearchUsers, models.User, models.Administrator))).Methods("GET")
	router.Handle("/users/{id}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReadUserById, models.User, models.Administrator))).Methods("GET")
	router.Handle("/users", controllers.CreateUser).Methods("POST")
	router.Handle("/users/{id}", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.UpdateUser, models.User))).Methods("PUT")
	router.Handle("/users/{id}/image", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ChangeImage, models.User))).Methods("PUT")
	router.Handle("/users/{id}/password", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ChangePassword, models.User))).Methods("PUT")
	router.Handle("/users/{id}/block", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.BlockUser, models.Administrator))).Methods("DELETE")

	router.Handle("/users/{id}/disable", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.DisableUser, models.User))).Methods("PUT")
	router.Handle("/users/reactivation/request", controllers.UserReactivation).Methods("POST")
	router.Handle("/users/{id}/reactivation", controllers.ReactivateUser).Methods("PUT")

	router.Handle("/users/password/recover/request", controllers.PasswordRecovery).Methods("POST")
	router.Handle("/users/{id}/password/recover", controllers.RecoverPassword).Methods("PUT")
	router.Handle("/users/{id}/report", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.ReportUser, models.User))).Methods("POST")

	router.Handle("/reports/search", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.SearchReports, models.Administrator))).Methods("GET")
	router.Handle("/reports/{id}/ignore", AuthenticateMiddlewate(AuthorizeMiddlewate(controllers.IgnoreReport, models.Administrator))).Methods("DELETE")

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
