
package controllers;

import (
	"net/http"
)


func Call(ctrlrFn func(*Context)) func(http.ResponseWriter, *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {

		ctx := &Context{
			res: res,
			req: req,
		};

		ctrlrFn(ctx);
	};
}



