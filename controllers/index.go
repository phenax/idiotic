
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
func Homepage(ctx *Context) {

	var title string;

	if(ctx.params["name"] != "") {
		title = "<h1>Hey, " + ctx.params["name"] + "</h1>";
	} else {
		title = "<h1>This is cool</h1>";
	}

	ctx.Send(title);

	// options :=
	// 	struct{
	// 		Title string;
	// 		Cool string;
	// 	}{
	// 		Title: title,
	// 		Cool: "Foobar",
	// 	};

	// ctx.Render("index", options);
}
