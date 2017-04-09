
package controllers;

import (
	// "net/http"
)

/**
 * Homepage controller
 *
 * params
 * -- ctx {*Context}  Server request context
 */
func Homepage(ctx *Context) {

	options :=
		struct{
			Title string;
			Cool string;
		}{
			Title: "This is a cool title",
			Cool: "Foobar",
		};

	ctx.render("index", options);
}
