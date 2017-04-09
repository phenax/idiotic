
package controllers;

import (
	"fmt"
	"net/http"
)

type Context struct {
	res http.ResponseWriter;
	req *http.Request;
};


func (ctx *Context) send(str string) {
	fmt.Fprint(ctx.res, str);
}

func (ctx *Context) render(templateName string) {

}

