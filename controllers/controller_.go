
package controllers;

import (
	"net/http"
	"github.com/gorilla/mux"
)


/**
 * Wrapper function for passing server context to controllers
 * 
 * params
 * -- ctrlrFn {func(*Context)}  The route controller
 *
 * returns
 * -- {func(http.ResponseWriter, *http.Request)}
 */
func Call(ctrlrFn func(*Context)) func(http.ResponseWriter, *http.Request) {


	/**
	 * The real route action handler func
	 *
	 * params
	 * -- res {http.ResponseWriter}
	 * -- req {*http.Request}
	 */
	return func(res http.ResponseWriter, req *http.Request) {

		// Create a context
		ctx := &Context{
			res: res,
			req: req,
			params: mux.Vars(req),
		};

		// Call the controller
		ctrlrFn(ctx);
	};
}



