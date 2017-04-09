
package controllers;

import (
	"fmt"
	// "log"
	"net/http"
	// "html/template"
)


/**
 * The server context
 *
 * fields
 * -- res {http.ResponseWriter}
 * -- req {*http.Request}
 */
type Context struct {
	res http.ResponseWriter;
	req *http.Request;
};



/**
 * ctx.send
 * Writes a string of html to response
 *
 * params
 * -- str {string}  The string to send
 */
func (ctx *Context) send(str string) {

	fmt.Fprint(ctx.res, str);
}


/**
 * ctx.render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) render(templateName string) {

	// tmpl, err :=
	// 	template.
	// 		New("poo").
	// 		Parse(`{{define "T"}}Hello, {{.}}!{{end}}`);

	// err = tmpl.ExecuteTemplate(ctx.res, "T", template.HTML("<em>Heloo</em>"));

	// if(err != nil) {
	// 	log.Fatal("Fuck");
	// }
}

