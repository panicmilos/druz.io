package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"gorm.io/gorm"

	"UserService/api_contracts"
	"UserService/config"
	"UserService/helpers"
	"UserService/models"

	"github.com/devfeel/mapper"
)

// swagger:route DELETE /admin/company/{id} admin deleteCompany
// Delete company
//
// security:
// - apiKey: []
// responses:
//  401: Account
//  200: Account
// Create handles Delete get company
func YourGetHandler(w http.ResponseWriter, r *http.Request) {
	user2 := &models.Account{}

	user := &models.Account{Username: "Milos"}

	mapper.Mapper(user, user2)

	fmt.Printf(user2.Username)

	w.Write([]byte("Gorilla!\n"))
}

func YourPostHandler(w http.ResponseWriter, r *http.Request) {
	db := di.Get(r, config.DatabaseConnection).(*gorm.DB)
	result := db.Find(&[]models.Account{})
	fmt.Println(result.RowsAffected)

	var input *api_contracts.CreateUserRequest
	err := helpers.ReadJSONBody(r, &input)
	if err != nil {
		return
	}
	fmt.Println(input.Username)
	err2 := input.Validate()
	if err2 != nil {
		fmt.Println("NIJE NIL")
		helpers.JSONResponse(w, 400, err2.Error())
	}

	user2 := &models.Account{}
	mapper.Mapper(input, user2)
	fmt.Println(user2.Username)

	db.Create(&models.Account{Username: "CAO", Password: "CAO", Salt: "NEsto"})

	return
}

func main() {
	mapper.Register(&models.Account{})
	mapper.Register(&models.Profile{})

	godotenv.Load(".env")

	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
	}
	err = builder.Add(config.ServiceContainer...)
	app := builder.Build()
	defer app.Delete()

	db := app.Get(config.DatabaseSeeder).(*gorm.DB)

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Profile{})

	r := mux.NewRouter()

	diMiddleware := func(next http.Handler) http.Handler {
		return di.HTTPMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)
		}), app, func(msg string) {
			fmt.Println(msg)
		})

	}
	r.Use(diMiddleware)

	// Routes consist of a path and a handler function.
	r.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	r.HandleFunc("/", YourGetHandler).Methods("GET")
	r.HandleFunc("/", YourPostHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
