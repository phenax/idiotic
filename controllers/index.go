
package controllers;

import (
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

	// ctx.Send("Cool");

	options :=
		struct{
			Cool string;
		}{
			Cool: "Foobar",
		};


	ctx.Render("index", options);
}
