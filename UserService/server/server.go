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

	router.Handle("/", controllers.YourGetHandler).Methods("GET")
	router.Handle("/users/{id}", controllers.ReadUserById).Methods("GET")
	router.Handle("/users", controllers.CreateUser).Methods("POST")
	router.Handle("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.Handle("/users/{id}/password", controllers.ChangePassword).Methods("PUT")
	router.Handle("/users/{id}/block", controllers.BlockUser).Methods("DELETE")

	router.Handle("/users/password/recover/request", controllers.PasswordRecovery).Methods("GET")
	router.Handle("/users/{id}/password/recover", controllers.RecoverPassword).Methods("POST")
	router.Handle("/users/{id}/report", AuthenticateMiddlewate(controllers.ReportUser)).Methods("POST")

	router.Handle("/reports/search", controllers.SearchReports).Methods("GET")
	router.Handle("/reports/{id}/ignore", controllers.IgnoreReport).Methods("DELETE")

	router.Handle("/auth", controllers.Authenticate).Methods("POST")
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
