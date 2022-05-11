package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/panicmilos/druz.io/ChatService/helpers"
	"github.com/panicmilos/druz.io/ChatService/models"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type MessagesCollection struct {
	Session        *ravendb.DocumentSession
	SessionStorage *helpers.SessionStorage
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

func formCollection(fromId string, toId string) string {
	return fmt.Sprintf("%s-%s-%s", messages_prefix, toId, fromId)
}

func formMessageKeyFromChat(chat string, messageId string) string {
	return fmt.Sprintf("%s/%s/%s", messages_prefix, chat, messageId)
}

func formChatCollection(chat string) string {
	return fmt.Sprintf("%s-%s", messages_prefix, chat)
}

func (messagesCollection *MessagesCollection) ReadMessage(chat string, messageId string) *models.Message {
	message := &models.Message{}

	err := messagesCollection.Session.Load(&message, formMessageKeyFromChat(chat, messageId))
	authenticatedUser := messagesCollection.SessionStorage.AuthenticatedUserId
	if err != nil || message.DeletedBy1 == authenticatedUser || message.DeletedBy2 == authenticatedUser {
		return nil
	}

	return message
}

func (messagesCollection *MessagesCollection) ReadChat(chat string) []*models.Message {
	q := messagesCollection.Session.QueryCollection(strings.Replace(formChatCollection(chat), "/", "//", -1))

	q.WaitForNonStaleResults(0)
	q.WhereNotEquals("DeletedBy1", messagesCollection.SessionStorage.AuthenticatedUserId)
	q.AndAlso()
	q.WhereNotEquals("DeletedBy2", messagesCollection.SessionStorage.AuthenticatedUserId)

	var messages []*models.Message
	err := q.GetResults(&messages)
	if err != nil || len(messages) == 0 {
		return nil
	}

	return messages
}

func (messagesCollection *MessagesCollection) Create(message *models.Message) *models.Message {
	message.ID = formMessagesKey(message.FromId, message.ToId)

	messagesCollection.Session.Store(message)
	metadata, _ := messagesCollection.Session.Advanced().GetMetadataFor(message)
	metadata.Put("@collection", formCollection(message.FromId, message.ToId))
	messagesCollection.Session.SaveChanges()

	return message
}

func (messagesCollection *MessagesCollection) DeleteMessage(chat string, messageId string, deleteFor []string) *models.Message {
	message := messagesCollection.ReadMessage(chat, messageId)

	for _, val := range deleteFor {
		if message.DeletedBy1 == "" {
			message.DeletedBy1 = val
		} else {
			message.DeletedBy2 = val
		}
	}

	messagesCollection.Session.SaveChanges()

	return message
}

func (messagesCollection *MessagesCollection) DeleteChat(chat string, deleteFor []string) []*models.Message {
	messages := messagesCollection.ReadChat(chat)

	for _, message := range messages {
		for _, val := range deleteFor {
			if message.DeletedBy1 == "" {
				message.DeletedBy1 = val
			} else {
				message.DeletedBy2 = val
			}
		}
	}

	messagesCollection.Session.SaveChanges()

	return messages
}
