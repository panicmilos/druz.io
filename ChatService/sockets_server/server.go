package sockets_server

import (
	"context"
	"encoding/json"

	"github.com/ambelovsky/gosf"
	"github.com/jellydator/ttlcache/v3"
	"github.com/panicmilos/druz.io/ChatService/services"
)

var _server *SocketsServer

type SocketsServer struct {
	Clients         *ttlcache.Cache[string, *gosf.Client]
	StatusesService *services.StatusesService
}

func New() *SocketsServer {
	_server = &SocketsServer{
		Clients:         services.Provider.Get(services.ClientsCache).(*ttlcache.Cache[string, *gosf.Client]),
		StatusesService: services.Provider.Get(services.StatusService).(*services.StatusesService),
	}

	_server.addListeners()

	return _server
}

func (server *SocketsServer) addListeners() {
	gosf.Listen("heartbit", heartbit)

	server.Clients.OnEviction(func(context context.Context, reason ttlcache.EvictionReason, item *ttlcache.Item[string, *gosf.Client]) {
		if reason == ttlcache.EvictionReasonExpired {
			userId := item.Key()
			_server.StatusesService.NotifyWentOffline(userId)
		}
	})

}

func heartbit(client *gosf.Client, request *gosf.Request) *gosf.Message {
	text := request.Message.Text
	heartbit := &Heartbit{}
	json.Unmarshal([]byte(text), heartbit)

	if _server.Clients.Get(heartbit.UserId) == nil {
		client.Join(heartbit.UserId)
		_server.StatusesService.NotifyCameOnline(heartbit.UserId)
	}

	_server.Clients.Set(heartbit.UserId, client, ttlcache.DefaultTTL)

	return gosf.NewSuccessMessage()
}

func (server *SocketsServer) Start() {

	go gosf.Startup(map[string]interface{}{"port": 8010})
}
