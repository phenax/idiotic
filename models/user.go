
package models;

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/phenax/idiotic/db"
);


/**
 * User model
 */
type User struct {

	ID        bson.ObjectId `bson:"_id,omitempty"`;

	Name      string        `bson:"name"`;

	Username  string        `bson:"username"`;

	Email     string        `bson:"email"`;

	Password  string        `bson:"password"`;
};


/**
 * Setter for the password field
 * TODO: Encrypt the password
 * 
 * params
 * -- password {string}
 */
func (user *User) SetPassword(password string) {
	user.Password = password;
}







// User Collection
var Users *mgo.Collection;

const (
	// Collection name
	UserCollectionName = "users";
);

func init() {

	dbObj, _ := db.GetDB();

	// Cache it on start
	Users = dbObj.C(UserCollectionName);
}



/**
 * Create a new user(make a copy of this user and modify for saving)
 * (mostly for triggering setters)
 *
 * params
 * -- user {*User}  User
 *
 * returns
 * -- {*User} The better and save ready user
 */
func NewUser(user *User) (*User) {

	newUser := *user;

	newUser.SetPassword(user.Password);

	return &newUser;
}

