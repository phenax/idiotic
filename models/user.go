package models

import (
	"github.com/gorilla/mux"
	"github.com/phenax/idiotic/db"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	// "fmt"
)

//
// User model
//
type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Name string `bson:"name"`

	Username string `bson:"username"`

	Email string `bson:"email"`

	Password string `bson:"password"`
}

//
// SetPassword - Setter for the password field
// TODO: Encrypt the password
//
// params
// -- password {string}
//
func (user *User) SetPassword(password string) {
	user.Password = password
}

//
// ProfileLink - Get the profile link to the user
//
func (user *User) ProfileLink(router *mux.Router) func(users ...interface{}) string {

	return func(users ...interface{}) string {

		var username string

		if len(users) > 0 {
			username = users[0].(string)
		} else {
			username = user.Username
		}

		url, _ := router.Get("profile").URL("name", username)

		return url.String()
	}
}

// Users - User Collection
var Users *mgo.Collection

const (
	// UserCollectionName - Collection name
	UserCollectionName = "users"
)

func init() {

	dbObj, _ := db.GetDB()

	// Cache it on start
	Users = dbObj.C(UserCollectionName)
}

//
// NewUser - Create a new user(make a copy of this user and modify for saving)
// (mostly for triggering setters)
//
// params
// -- user {*User}  User
//
// returns
// -- {*User} The better and save ready user
//
func NewUser(user *User) *User {

	newUser := *user

	newUser.SetPassword(user.Password)

	return &newUser
}
