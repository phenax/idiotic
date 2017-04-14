
package db;

import (
	"labix.org/v2/mgo"
	"time"
);


func GetProdConfig() (*mgo.DialInfo) {

	return &mgo.DialInfo{

		Addrs: []string{ "localhost" },

		Timeout:  60 * time.Second,

		Database: "idiotic",

		Username: "",
		Password: "",
	};
}
