package setup

import (
	"UserService/models"
	"UserService/services"

	"gorm.io/gorm"
)

func SetupDatabase() {
	db := services.Provider.Get(services.DatabaseSeeder).(*gorm.DB)

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.LivePlace{})
	db.AutoMigrate(&models.WorkPlace{})
	db.AutoMigrate(&models.Education{})
	db.AutoMigrate(&models.Interes{})

}
