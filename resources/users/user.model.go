package users

import (
	"github.com/IkeyBenz/CodeReviewCLI/database"
	"gopkg.in/mgo.v2/bson"
)

// User struct serves as the model for all user objects in our database
type User struct {
	Email    string
	Password string
	Inbox    []bson.ObjectId
	Outbox   []bson.ObjectId
	History  []bson.ObjectId
}

// NewUser creates an new User object with empty inbox, outbox, and history slices
func NewUser(email, password string) error {
	d := database.MgoSession.DB("code-review")
	newUser := User{
		Email:    email,
		Password: password,
		Inbox:    make([]bson.ObjectId, 0),
		Outbox:   make([]bson.ObjectId, 0),
		History:  make([]bson.ObjectId, 0),
	}
	return d.C("users").Insert(newUser)
}
