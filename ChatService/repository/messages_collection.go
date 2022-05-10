package repository

import (
	"fmt"
	"strconv"

	"github.com/panicmilos/druz.io/ChatService/models"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type MessagesCollection struct {
	Session *ravendb.DocumentSession
}

const messages_prefix = "messages"

func formMessagesKey(fromId string, toId string) string {
	fromIdParsed, _ := strconv.ParseUint(fromId, 10, 32)
	toIdParsed, _ := strconv.ParseUint(toId, 10, 32)
	if fromIdParsed < toIdParsed {
		return fmt.Sprintf("%s/%s-%s/", messages_prefix, fromId, toId)
	}
	return fmt.Sprintf("%s/%s-%s/", messages_prefix, toId, fromId)
}

func (messagesCollection *MessagesCollection) Create(message *models.Message) *models.Message {
	message.ID = formMessagesKey(message.FromId, message.ToId)

	messagesCollection.Session.Store(message)
	messagesCollection.Session.SaveChanges()

	return message
}
