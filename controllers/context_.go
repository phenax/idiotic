
package controllers;

import (
	"fmt"
	// "log"
	// "reflect"
	"net/http"
	"path/filepath"
	// "io/ioutil"
	// "github.com/gorilla/mux"
	// "html/template"
);


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
	params map[string]string;
};


/**
 * ctx.Send
 * Writes a string of html to response
 *
 * params
 * -- str {string}  The string to send
 */
func (ctx *Context) Send(str string) {

	fmt.Fprint(ctx.res, str);
}


/**
 * Write some data to the response
 * 
 * params
 * -- data {[]byte}  Data to write to the response
 */
func (ctx *Context) Write(data []byte) {

	header := ctx.res.Header();

	if header.Get("Content-Type") == "" {
		header.Set("Content-Type", http.DetectContentType(data));
	}

	header.Del("Content-Length");

	fmt.Fprint(ctx.res, data);
}


/**
 * ctx.Render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) Render(templateName string) {

	// buf := new(bytes.Buffer)
	// if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
	// 	rest.ServerError(w, r, err)
	// 	return
	// }
	// w.Write(buf.Bytes())
	
}







/**
 * getTemplatePath
 * Get the full path to the template
 *
 * params
 * -- templateName {string}
 *
 * returns
 * -- {string}
 */
func getTemplatePath(templateName string) string {

	return filepath.Join(
		filepath.Base("."),
		"views",
		templateName + ".html",
	);
}

