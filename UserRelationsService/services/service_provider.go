package services

import (
	"fmt"
	"log"
	"os"

	"github.com/panicmilos/druz.io/UserRelationsService/repository"
	"github.com/sarulabs/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Provider = buildServiceContainer()

const (
	AppDatabaseInstance  = "AppDatabaseInstance"
	DatabaseConnection   = "DatabaseConnection"
	Repository           = "Repository"
	UsersReplicator      = "UsersReplicator"
	UserBlockService     = "UserBlockService"
	UserFriendService    = "UserFriendService"
	UserService          = "UserService"
	FriendRequestService = "FriendRequestService"
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
				UserBlocks: &repository.UserBlocksCollection{
					DB: db,
				},
				UserFriends: &repository.UserFriendsCollection{
					DB: db,
				},
				FriendRequests: &repository.FriendRequestsCollection{
					DB: db,
				},
			}, nil
		},
	},
	{
		Name:  UsersReplicator,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			db := ctn.Get(AppDatabaseInstance).(*gorm.DB)
			userReplicator := &UserReplicator{
				Users: &repository.UsersCollection{
					DB: db,
				},
			}
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
		Name:  UserBlockService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			userService := ctn.Get(UserService).(*UsersService)

			return &UserBlocksService{
				repository:   repository,
				usersService: userService,
			}, nil
		},
	},
	{
		Name:  UserFriendService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)

			return &UserFriendsService{
				repository: repository,
			}, nil
		},
	},
	{
		Name:  FriendRequestService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			usersService := ctn.Get(UserService).(*UsersService)
			userFriendsService := ctn.Get(UserFriendService).(*UserFriendsService)

			return &FriendRequestsService{
				repository:         repository,
				usersService:       usersService,
				userFriendsService: userFriendsService,
			}, nil
		},
	},
	{
		Name:  UserService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			return &UsersService{
				repository: repository,
			}, nil
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
