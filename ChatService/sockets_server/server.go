package sockets_server

import (
	"encoding/json"
	"log"

	"github.com/ambelovsky/gosf"
	"github.com/jellydator/ttlcache/v3"
	"github.com/panicmilos/druz.io/ChatService/services"
)

var _server *SocketsServer

type SocketsServer struct {
	Clients *ttlcache.Cache[string, *gosf.Client]
}

func New() *SocketsServer {
	_server = &SocketsServer{
		Clients: services.Provider.Get(services.ClientsCache).(*ttlcache.Cache[string, *gosf.Client]),
	}

	_server.addListeners()

	return _server
}

func (server *SocketsServer) addListeners() {
	gosf.Listen("init-user", initUser)

	gosf.OnDisconnect(func(client *gosf.Client, request *gosf.Request) {
		log.Println("Client disconnected.")
	})
}

func initUser(client *gosf.Client, request *gosf.Request) *gosf.Message {
	text := request.Message.Text

	initUser := &InitUser{}
	json.Unmarshal([]byte(text), initUser)

	client.Join(initUser.UserId)

	_server.Clients.Set(initUser.UserId, client, ttlcache.NoTTL)

	return gosf.NewSuccessMessage("Successfull init")
}

func (server *SocketsServer) Start() {

	go gosf.Startup(map[string]interface{}{"port": 8010})
}
