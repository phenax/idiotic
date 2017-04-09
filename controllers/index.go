
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

	ctx.render("wrapper");
}
