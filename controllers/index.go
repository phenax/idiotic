
package controllers;

import (
	"github.com/phenax/idiotic/libs"
	"github.com/phenax/idiotic/models"
	// "labix.org/v2/mgo/bson"
	// "errors"
	// "net/http"
	// "fmt"
	// "reflect"
	// "github.com/gorilla/mux"
)

/**
 * Homepage controller
 *
 * params
 * -- ctx {*Context}  Server request context
 */
func HomePage(ctx *Context) {

	var users []models.User;

	models.Users.Find(nil).All(&users);

	options :=
		struct{
			Cool string;
			Users []models.User;
		}{
			Cool: "Foobar",
			Users: users,
		};

	ctx.Render("index", options);
}



/**
 * Send json test data to the client
 *
 * params
 * -- ctx {*Context}
 */
func JSONTest(ctx *Context) {

	type Product struct {
		Id int;
		Name string;
	};

	type Response struct {
		Access bool;
		Products []*Product;
	};

	obj := &Response{
		Access: true,
		Products: []*Product{
			&Product{
				Id: 1,
				Name: "Soap",
			},
			&Product{
				Id: 2,
				Name: "Shampoo",
			},
		},
	};

	ctx.JSON(obj);
}



/**
 * Testing gzip compression on a string of content
 *
 * params
 * -- ctx {*Context}
 */
func GzipTest(ctx *Context) {

	content := `
		Disposable city market rain pistol saturation point hacker grenade engine range-rover.
		Neural gang dome nano-faded beef noodles bicycle footage kanji advert courier garage singularity.
		Rain concrete weathered industrial grade knife tank-traps sign RAF nodal point alcohol tower.
		Ablative spook neural military-grade engine cyber-carbon media shoes Kowloon knife.
		Numinous fluidity into market silent physical crypto-sprawl tanto euro-pop.
	`;

	// Gzip the string
	gzipppedContent, err := libs.GzipString(content);

	if(err != nil) {
		ctx.ErrorMessage(500, err);
		return;
	}

	ctx.Send(string(gzipppedContent), &ResponseConfig{
		ContentType: "text/plain",
		ContentEncoding: "gzip",
	});
}

