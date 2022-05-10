package models

type Message struct {
	ID string

	FromId string
	ToId   string

	Message string
	Type    string

	DeletedFor []string
}
