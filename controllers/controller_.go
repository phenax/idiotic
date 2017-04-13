
package controllers;

import (
	// "bytes"
	// "time"
	// "fmt"
	// "io/ioutil"
	// "strings"
	"net/http"
	"github.com/gorilla/mux"
)


/**
 * Static file serving configuration
 *
 * fields
 * -- Pathprefix {string}  The prefix for the url pathname for static content
 * -- Directory  {string}  The static content root directory on the server
 */
type StaticConfig struct {
	Pathprefix string;
	Directory string;
};




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




/**
 * Static routing for serving files
 *
 * params
 * -- router   {*mux.Router}    mux Router instance
 * -- options  {*StaticConfig}  Static file serving configuration
 *
 * returns
 * -- {*mux.Route}              To extend the chain
 */
func StaticRouter(router *mux.Router, options *StaticConfig) (*mux.Route) {

	// Default configuration
	defaultConfig := &StaticConfig{
		Pathprefix: "/public",
		Directory: "./public",
	};


	pathprefix := options.Pathprefix;
	directory := options.Directory;

	if(pathprefix == "") {
		pathprefix = defaultConfig.Pathprefix;
	}

	if(directory == "") {
		directory = defaultConfig.Directory;
	}

	return router.
		PathPrefix(pathprefix).
		Handler(
			http.StripPrefix(
				pathprefix,
				http.FileServer(http.Dir(directory)),
			),
		);
}


