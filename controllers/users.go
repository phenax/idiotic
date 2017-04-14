
package controllers;

import (
	"github.com/phenax/idiotic/models"
	"labix.org/v2/mgo/bson"
);


func ProfilePage(ctx *Context) {

	var user models.User;

	username := ctx.Params["name"];

	models.
		Users.Find(bson.M{
			"username": username,
		}).One(&user);

	if(user == nil) {
		ctx.Send("No user with that username");
		return;
	}

	config := &ResponseConfig{
		ContentType: "text/html",
	};

	ctx.Render("user", user, config);
}
