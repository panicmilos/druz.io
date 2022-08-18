package repository

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/panicmilos/druz.io/ChatService/dto"
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
	fromIdParsed, _ := strconv.ParseUint(fromId, 10, 32)
	toIdParsed, _ := strconv.ParseUint(toId, 10, 32)
	if fromIdParsed < toIdParsed {
		return fmt.Sprintf("%s-%s-%s", messages_prefix, fromId, toId)
	}
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
	if err != nil || (message != nil && message.ID == "") || message.DeletedBy1 == authenticatedUser || message.DeletedBy2 == authenticatedUser {
		return nil
	}

	return message
}

func (messagesCollection *MessagesCollection) ReadChat(chat string) []*models.Message {
	q := messagesCollection.Session.QueryCollection(formChatCollection(chat))

	q.WaitForNonStaleResults(0)
	q.WhereNotEquals("DeletedBy1", messagesCollection.SessionStorage.AuthenticatedUserId)
	q.WhereNotEquals("DeletedBy2", messagesCollection.SessionStorage.AuthenticatedUserId)
	q.OrderBy("CreatedAt")

	var messages []*models.Message
	err := q.GetResults(&messages)
	if err != nil || len(messages) == 0 {
		return nil
	}

	return messages
}

func (messagesCollection *MessagesCollection) ChatsWith(forId string) *[]dto.Chat {
	q := messagesCollection.Session.QueryCollection("@all_docs")

	q.WhereExists("ToId")
	q.WhereExists("FromId")

	q.Include("ToId")
	q.Include("FromId")
	q.SelectFields(reflect.TypeOf(""), "ToId", "FromId")
	q.Distinct()

	var messages []*models.Message
	q.GetResults(&messages)

	chats := []dto.Chat{}

	for _, message := range messages {
		if message.ToId == messagesCollection.SessionStorage.AuthenticatedUserId {
			chat := messagesCollection.makeChat(message.FromId)
			if !helpers.ContainsChat(chats, *chat) {
				chats = append(chats, *chat)
			}
		}

		if message.FromId == messagesCollection.SessionStorage.AuthenticatedUserId {
			chat := messagesCollection.makeChat(message.ToId)
			if !helpers.ContainsChat(chats, *chat) {
				chats = append(chats, *chat)
			}
		}
	}

	return &chats
}

func (messagesCollection *MessagesCollection) makeChat(fromId string) *dto.Chat {
	from := &models.User{}
	messagesCollection.Session.Load(&from, formUsersKey(fromId))

	chat := ""
	fromIdParsed, _ := strconv.ParseUint(fromId, 10, 32)
	toIdParsed, _ := strconv.ParseUint(messagesCollection.SessionStorage.AuthenticatedUserId, 10, 32)
	if fromIdParsed < toIdParsed {
		chat = fmt.Sprintf("%s-%s", fromId, messagesCollection.SessionStorage.AuthenticatedUserId)
	} else {
		chat = fmt.Sprintf("%s-%s", messagesCollection.SessionStorage.AuthenticatedUserId, fromId)
	}

	return &dto.Chat{
		Chat: chat,
		User: from,
	}
}

func (messagesCollection *MessagesCollection) SearchChat(chat string, searchParams *dto.ChatSearchParams) []*models.Message {
	q := messagesCollection.Session.QueryCollection(formChatCollection(chat))

	q.WaitForNonStaleResults(0)
	q.WhereNotEquals("DeletedBy1", messagesCollection.SessionStorage.AuthenticatedUserId)
	q.WhereNotEquals("DeletedBy2", messagesCollection.SessionStorage.AuthenticatedUserId)
	if len(strings.TrimSpace(searchParams.Keywoard)) != 0 {
		q.Search("Message", searchParams.Keywoard)
	}
	if searchParams.Take > 0 {
		q.Take(searchParams.Take)
	}
	q.OrderBy("CreatedAt")

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

	setDeleteFor(message, deleteFor)

	messagesCollection.Session.SaveChanges()

	return message
}

func (messagesCollection *MessagesCollection) DeleteChat(chat string, deleteFor []string) []*models.Message {
	messages := messagesCollection.readChatForDelete(chat)

	for _, message := range messages {
		setDeleteFor(message, deleteFor)
	}

	messagesCollection.Session.SaveChanges()

	return messages
}

func (messagesCollection *MessagesCollection) readChatForDelete(chat string) []*models.Message {
	q := messagesCollection.Session.QueryCollection(formChatCollection(chat))

	q.WaitForNonStaleResults(0)

	var messages []*models.Message
	err := q.GetResults(&messages)
	if err != nil || len(messages) == 0 {
		return nil
	}

	return messages
}

func setDeleteFor(message *models.Message, deleteFor []string) {
	for _, val := range deleteFor {
		if message.DeletedBy1 == "" {
			message.DeletedBy1 = val
		} else {
			if message.DeletedBy1 == val {
				continue
			}
			message.DeletedBy2 = val
		}
	}
}
