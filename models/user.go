
package models;

import (
	"labix.org/v2/mgo/bson"
);


type User struct {

	ID        bson.ObjectId `bson:"_id,omitempty"`;

	Name      string        `bson:"name"`;

	Username  string        `bson:"username"`;

	Email     string        `bson:"email"`;

	Password  string        `bson:"password"`;
};

