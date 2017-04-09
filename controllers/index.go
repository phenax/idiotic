
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

	ctx.send("<h1>Cool awesomeness</h1>");
}
