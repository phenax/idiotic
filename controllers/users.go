package controllers

import (
	"errors"
	// "fmt"

	"github.com/phenax/idiotic/models"
	"labix.org/v2/mgo/bson"
)

//
// UserWrapper - Wrapper for users data to send to the client
//
type UserWrapper struct {
	User  models.User
	Users []models.User
	Ctx   *Context
}

//
// GetLink - Get the profile link of the user
//
// params
// -- users {interface{}} Optional user parameter
//
// returns
// -- {string} The link to the profile
//
func (resp UserWrapper) GetLink(users ...interface{}) string {

	var username string

	if len(users) > 0 {
		username = users[0].(string)
	} else {
		username = ""
	}

	return resp.User.ProfileLink(resp.Ctx.Router)(username)
}

//
// ProfilePage - Users profile page controller
//
// params
// -- ctx {*Context}
//
func ProfilePage(ctx *Context) {

	var user models.User

	// Get the route url parameter
	username := ctx.Params["name"]

	// Get the user with that username
	models.
		Users.Find(bson.M{
		"username": username,
	}).One(&user)

	// If nothing came back from the fetch
	if user.Username == "" {
		ctx.ErrorMessage(404, errors.New("No user with that username"))
		return
	}

	config := &ResponseConfig{
		ContentType: "text/html",
	}

	renderContent := &UserWrapper{
		User: user,
		Ctx:  ctx,
	}

	ctx.Render("user", renderContent, config)
}
