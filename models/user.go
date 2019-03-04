package models

// User struct serves as the model for all user objects in our database
type User struct {
	_id      string
	email    string
	password string
	inbox    []Request
	outbox   []Request
	history  []Request
}
