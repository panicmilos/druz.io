package setup

import (
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/services"
	"gorm.io/gorm"
)

func SetupDatabase() {
	db := services.Provider.Get(services.AppDatabaseInstance).(*gorm.DB)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserBlock{})
	db.AutoMigrate(&models.UserFriend{})
	db.AutoMigrate(&models.FriendRequest{})
}
