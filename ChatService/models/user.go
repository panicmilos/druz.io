package models

// swagger:model User
type User struct {
	ID        string
	FirstName string
	LastName  string
	Disabled  bool
	Image     string
}
