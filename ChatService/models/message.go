package models

import "time"

type Message struct {
	ID        string
	CreatedAt time.Time

	FromId     string
	ToId       string
	Message    string
	Type       string
	DeletedBy1 string
	DeletedBy2 string
}
