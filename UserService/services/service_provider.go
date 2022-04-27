package services

import (
	"UserService/repository"
	"fmt"
	"log"
	"os"

	"github.com/sarulabs/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Provider = buildServiceContainer()

const (
	DatabaseSeeder     = "DatabaseSeeder"
	DatabaseConnection = "DatabaseConnection"
	Repository         = "Repository"
	UsersService       = "UsersService"
	AuthService        = "AuthenticationService"
	EmailDispatcher    = "EmailDispatcher"
)

var serviceContainer = []di.Def{
	{
		Name:  DatabaseSeeder,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_NAME"))
			return gorm.Open(mysql.Open(connectionString), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
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
			return gorm.Open(mysql.Open(connectionString), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		},
		Close: func(obj interface{}) error {
			db, err := obj.(*gorm.DB).DB()
			db.Close()

			return err
		},
	},
	{
		Name:  Repository,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get(DatabaseConnection).(*gorm.DB)
			return &repository.Repository{
				DB: db,
				Users: &repository.UsersCollection{
					DB: db,
				},
				LivePlaces: &repository.LivePlacesCollection{
					DB: db,
				},
				WorkPlaces: &repository.WorkPlacesCollection{
					DB: db,
				},
				Educations: &repository.EducationsCollection{
					DB: db,
				},
				Intereses: &repository.InteresesCollection{
					DB: db,
				},
			}, nil
		},
	},
	{
		Name:  UsersService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			return &UserService{
				repository: repository,
			}, nil
		},
	},
	{
		Name:  AuthService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			return &AuthenticationService{
				repository: repository,
			}, nil
		},
	},
	{
		Name:  EmailDispatcher,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			emailService := &EmailService{}
			emailService.Initialize()

			return emailService, nil
		},
		Close: func(obj interface{}) error {
			emailService := obj.(*EmailService)
			emailService.Deinitialize()

			return nil
		},
	},
}

func buildServiceContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = builder.Add(serviceContainer...)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return builder.Build()
}
