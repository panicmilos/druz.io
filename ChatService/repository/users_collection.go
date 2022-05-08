package repository

import (
	"fmt"
	"strings"

	"github.com/panicmilos/druz.io/ChatService/models"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type UsersCollection struct {
	Session *ravendb.DocumentSession
}

const users_prefix = "users"

func formUsersKey(id string) string {
	return fmt.Sprintf("%s/%s", users_prefix, id)
}

func (usersCollection *UsersCollection) ReadById(id string) *models.User {
	user := &models.User{}

	err := usersCollection.Session.Load(user, formUsersKey(id))
	if err != nil || user.Disabled {
		return nil
	}

	return user
}

func (usersCollection *UsersCollection) Create(user *models.User) *models.User {
	user.ID = formUsersKey(user.ID)

	usersCollection.Session.Store(user)
	usersCollection.Session.SaveChanges()

	return user
}

func (usersCollection *UsersCollection) Update(user *models.User) *models.User {
	if !strings.HasPrefix(user.ID, users_prefix) {
		user.ID = formUsersKey(user.ID)
	}

	usersCollection.Session.Store(user)
	usersCollection.Session.SaveChanges()

	return user
}

func (usersCollection *UsersCollection) Delete(id string) *models.User {
	user := usersCollection.ReadById(id)

	usersCollection.Session.Delete(user)

	return user
}

func (usersCollection *UsersCollection) Disable(id string) *models.User {
	user := usersCollection.ReadById(id)

	user.Disabled = true

	return usersCollection.Update(user)
}

func (usersCollection *UsersCollection) Reactivate(id string) *models.User {
	user := usersCollection.ReadById(id)

	user.Disabled = false

	return usersCollection.Update(user)
}
