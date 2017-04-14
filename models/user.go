
package models;

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/phenax/idiotic/config"
);


type User struct {

	ID        bson.ObjectId `bson:"_id,omitempty"`;

	Name      string        `bson:"name"`;

	Username  string        `bson:"username"`;

	Email     string        `bson:"email"`;

	Password  string        `bson:"password"`;
};


func (user *User) SetPassword(password string) {
	user.Password = password;
}




var Users *mgo.Collection;

const (
	UserCollectionName = "users";
);

func init() {

	db, _ := config.GetDB();

	Users = db.C(UserCollectionName);

}


func NewUser(user *User) (*User) {

	newUser := *user;

	newUser.SetPassword(user.Password);

	return &newUser;
}


func SaveUser(user *User) {

}


