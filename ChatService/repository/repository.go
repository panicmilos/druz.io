package repository

import ravendb "github.com/ravendb/ravendb-go-client"

type Repository struct {
	Session *ravendb.DocumentSession
	Users   *UsersCollection
}
