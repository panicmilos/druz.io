package services

import (
	"log"
	"os"
	"time"

	"github.com/ambelovsky/gosf"
	"github.com/jellydator/ttlcache/v3"
	"github.com/panicmilos/druz.io/ChatService/helpers"
	"github.com/panicmilos/druz.io/ChatService/repository"
	ravendb "github.com/ravendb/ravendb-go-client"
	"github.com/sarulabs/di"
)

var Provider = buildServiceContainer()

const (
	ClientsCache         = "ClientsCache"
	DocumentStore        = "DocumentStore"
	AppDatabaseInstance  = "AppDatabaseInstance"
	DatabaseConnection   = "DatabaseConnection"
	Repository           = "Repository"
	UserReplicator       = "UserReplicator"
	UserFriendReplicator = "UserFriendReplicator"
	UserService          = "UserService"
	MessageService       = "MessageService"
	SessionStorage       = "SessionStorage"
	StatusService        = "StatusService"
	AppStatusService     = "AppStatusService"
)

var serviceContainer = []di.Def{
	{
		Name:  ClientsCache,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cache := ttlcache.New(
				ttlcache.WithTTL[string, *gosf.Client](30 * time.Second),
			)

			go cache.Start()

			return cache, nil
		},
		Close: func(obj interface{}) error {
			cashe := obj.(*ttlcache.Cache[string, *gosf.Client])
			cashe.DeleteAll()

			return nil
		},
	},
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
			sessionStorage := ctn.Get(SessionStorage).(*helpers.SessionStorage)

			return &repository.Repository{
				Session: session,
				Users: &repository.UsersCollection{
					Session: session,
				},
				UserFriends: &repository.UserFriendsCollection{
					Session: session,
				},
				Messages: &repository.MessagesCollection{
					Session:        session,
					SessionStorage: sessionStorage,
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
	{
		Name:  MessageService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			usersService := ctn.Get(UserService).(*UsersService)
			clientsCache := ctn.Get(ClientsCache).(*ttlcache.Cache[string, *gosf.Client])
			sessionStorage := ctn.Get(SessionStorage).(*helpers.SessionStorage)

			return &MessagesService{
				repository:     repository,
				UsersService:   usersService,
				Clients:        clientsCache,
				SessionStorage: sessionStorage,
			}, nil
		},
	},
	{
		Name:  SessionStorage,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &helpers.SessionStorage{}, nil
		},
	},
	{
		Name:  StatusService,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			repository := ctn.Get(Repository).(*repository.Repository)
			usersService := ctn.Get(UserService).(*UsersService)
			clientsCache := ctn.Get(ClientsCache).(*ttlcache.Cache[string, *gosf.Client])

			return &StatusesService{
				repository:   repository,
				UsersService: usersService,
				Clients:      clientsCache,
			}, nil
		},
	},
	{
		Name:  AppStatusService,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			appDatabaseInstance := ctn.Get(AppDatabaseInstance).(*ravendb.DocumentSession)
			repository := &repository.Repository{
				Session: appDatabaseInstance,
				UserFriends: &repository.UserFriendsCollection{
					Session: appDatabaseInstance,
				},
				Users: &repository.UsersCollection{
					Session: appDatabaseInstance,
				},
			}
			usersService := &UsersService{
				repository: repository,
			}
			clientsCache := ctn.Get(ClientsCache).(*ttlcache.Cache[string, *gosf.Client])

			return &StatusesService{
				repository:   repository,
				UsersService: usersService,
				Clients:      clientsCache,
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
