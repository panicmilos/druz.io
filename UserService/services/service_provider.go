package services

import (
	"fmt"
	"log"
	"os"

	"github.com/panicmilos/druz.io/UserService/repository"

	"github.com/sarulabs/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Provider = buildServiceContainer()

const (
	AppDatabaseInstance       = "AppDatabaseInstance"
	DatabaseConnection        = "DatabaseConnection"
	Repository                = "Repository"
	UsersService              = "UsersService"
	AuthService               = "AuthenticationService"
	EmailDispatcher           = "EmailDispatcher"
	PasswordRecoveriesService = "PasswordRecoveryService"
	UserReportService         = "UserReportsService"
	UserReactivationService   = "UserReactivationService"
	UsersReplicator           = "UsersReplicator"
	UserBlocksReplicator      = "UsersBlocksReplicator"
)

var serviceContainer = []di.Def{
	{
		Name:  AppDatabaseInstance,
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
				PasswordRecoveriesCollection: &repository.PasswordRecoveriesCollection{
					DB: db,
				},
				UserReportsCollection: &repository.UserReportsCollection{
					DB: db,
				},
				UserReactivationsCollection: &repository.UserReactivationsCollection{
					DB: db,
				},
				UserBlocksCollection: &repository.UserBlocksCollection{
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
			userReplicator := ctn.Get(UsersReplicator).(*UserReplicator)
			return &UserService{
				repository:     repository,
				userReplicator: userReplicator,
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
		Name:  PasswordRecoveriesService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			emailDispatcher := ctn.Get(EmailDispatcher).(*EmailService)
			return &PasswordRecoveryService{
				repository:      repository,
				emailDispatcher: emailDispatcher,
			}, nil
		},
	},
	{
		Name:  UserReportService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			return &UserReportsService{
				repository: repository,
			}, nil
		},
	},
	{
		Name:  UserReactivationService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			emailDispatcher := ctn.Get(EmailDispatcher).(*EmailService)
			userReplicator := ctn.Get(UsersReplicator).(*UserReplicator)
			return &UserReactivationsService{
				repository:      repository,
				emailDispatcher: emailDispatcher,
				userReplicator:  userReplicator,
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
	{
		Name:  UsersReplicator,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			userReplicator := &UserReplicator{}
			userReplicator.Initialize()

			return userReplicator, nil
		},
		Close: func(obj interface{}) error {
			userReplicator := obj.(*UserReplicator)
			userReplicator.Deinitialize()

			return nil
		},
	},
	{
		Name:  UserBlocksReplicator,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get(AppDatabaseInstance).(*gorm.DB)
			userBlockReplicator := &UserBlockReplicator{
				UserBlocks: &repository.UserBlocksCollection{
					DB: db,
				},
			}
			userBlockReplicator.Initialize()

			return userBlockReplicator, nil
		},
		Close: func(obj interface{}) error {
			userBlockReplicator := obj.(*UserBlockReplicator)
			userBlockReplicator.Deinitialize()

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
