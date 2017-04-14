
package controllers;

import (
	"github.com/phenax/idiotic/models"
	"labix.org/v2/mgo/bson"
	"errors"
);


/**
 * Users profile page
 *
 * params
 * -- ctx {*Context}
 */
func ProfilePage(ctx *Context) {

	var user models.User;

	// Get the route url parameter
	username := ctx.Params["name"];

	// Get the user with that username
	models.
		Users.Find(bson.M{
			"username": username,
		}).One(&user);

	// If nothing came back from the fetch
	if(user.Username == "") {
		ctx.ErrorMessage(404, errors.New("No user with that username"));
		return;
	}

	config := &ResponseConfig{
		ContentType: "text/html",
	};

	ctx.Render("user", user, config);
}
