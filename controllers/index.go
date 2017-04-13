
package controllers;

import (
	"github.com/phenax/idiotic/libs"
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

	options :=
		struct{
			Cool string;
		}{
			Cool: "Foobar",
		};


	ctx.Render("index", options);
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
	}

	ctx.Send(string(gzipppedContent), &ResponseConfig{
		ContentType: "text/plain",
		ContentEncoding: "gzip",
	});
}

