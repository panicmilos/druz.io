package services

import (
	"log"
	"os"

	"github.com/panicmilos/druz.io/ChatService/repository"
	ravendb "github.com/ravendb/ravendb-go-client"
	"github.com/sarulabs/di"
)

var Provider = buildServiceContainer()

const (
	DocumentStore        = "DocumentStore"
	AppDatabaseInstance  = "AppDatabaseInstance"
	DatabaseConnection   = "DatabaseConnection"
	Repository           = "Repository"
	UserReplicator       = "UserReplicator"
	UserFriendReplicator = "UserFriendReplicator"
)

var serviceContainer = []di.Def{
	{
		Name:  DocumentStore,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			serverNodes := []string{os.Getenv("DB_URL")}
			store := ravendb.NewDocumentStore(serverNodes, os.Getenv("DB_NAME"))
			if err := store.Initialize(); err != nil {
				return nil, err
			}
			return store, nil
		},
		Close: func(obj interface{}) error {
			store := obj.(*ravendb.DocumentStore)
			store.Close()

			return nil
		},
	},
	{
		Name:  AppDatabaseInstance,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			store := ctn.Get(DocumentStore).(*ravendb.DocumentStore)

			return store.OpenSession("")
		},
		Close: func(obj interface{}) error {
			session := obj.(*ravendb.DocumentSession)
			session.Close()

			return nil
		},
	},
	{
		Name:  DatabaseConnection,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			store := ctn.Get(DocumentStore).(*ravendb.DocumentStore)

			return store.OpenSession("")
		},
		Close: func(obj interface{}) error {
			session := obj.(*ravendb.DocumentSession)
			session.Close()

			return nil
		},
	},
	{
		Name:  Repository,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			session := ctn.Get(DatabaseConnection).(*ravendb.DocumentSession)

			return &repository.Repository{
				Session: session,
				Users: &repository.UsersCollection{
					Session: session,
				},
			}, nil
		},
	},
	{
		Name:  UserReplicator,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			session := ctn.Get(AppDatabaseInstance).(*ravendb.DocumentSession)
			usersReplicator := &UsersReplicator{
				Users: &repository.UsersCollection{
					Session: session,
				},
			}
			usersReplicator.Initialize()

			return usersReplicator, nil
		},
		Close: func(obj interface{}) error {
			usersReplicator := obj.(*UsersReplicator)
			usersReplicator.Deinitialize()

			return nil
		},
	},
	{
		Name:  UserFriendReplicator,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			session := ctn.Get(AppDatabaseInstance).(*ravendb.DocumentSession)
			userFriendsReplicator := &UserFriendsReplicator{
				UserFriends: &repository.UserFriendsCollection{
					Session: session,
				},
			}
			userFriendsReplicator.Initialize()

			return userFriendsReplicator, nil
		},
		Close: func(obj interface{}) error {
			userFriendsReplicator := obj.(*UserFriendsReplicator)
			userFriendsReplicator.Deinitialize()

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
