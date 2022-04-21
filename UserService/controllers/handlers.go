package controllers

import (
	"UserService/api_contracts"
	"UserService/helpers"
	"UserService/models"
	"UserService/services"
	"fmt"
	"net/http"

	"github.com/devfeel/mapper"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

// swagger:route DELETE /admin/company/{id} admin deleteCompany
// Delete company
//
// security:
// - Bearer: []
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
	db := di.Get(r, services.DatabaseConnection).(*gorm.DB)
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
