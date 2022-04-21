package config

import (
	"fmt"
	"os"

	"github.com/sarulabs/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DatabaseSeeder     = "DatabaseSeeder"
	DatabaseConnection = "DatabaseConnection"
)

var ServiceContainer = []di.Def{
	{
		Name:  DatabaseSeeder,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_NAME"))
			return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		},
		Close: func(obj interface{}) error {
			db, err := obj.(*gorm.DB).DB()
			db.Close()

			return err
		},
	},
	{
		Name:  DatabaseConnection,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_NAME"))
			return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		},
		Close: func(obj interface{}) error {
			db, err := obj.(*gorm.DB).DB()
			db.Close()

			return err
		},
	},
}
