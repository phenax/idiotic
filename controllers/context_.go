
package controllers;

import (
	"fmt"
	"bytes"
	// "log"
	// "reflect"
	"net/http"
	"path/filepath"
	"io/ioutil"
	// "github.com/gorilla/mux"
	"html/template"
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
func (ctx *Context) Render(templateName string, data interface{}) {

	html, err := ioutil.ReadFile(getTemplatePath(templateName));

	if(err != nil) {
		fmt.Fprint(ctx.res, "Didnt render");
		return;
	}

	tpl := template.Must(template.New("homepage").Parse(string(html)));

	buf := new(bytes.Buffer);

	if err := tpl.ExecuteTemplate(buf, "homepage", data); err != nil {
		fmt.Fprint(ctx.res, "Didnt render");
		return;
	}

	ctx.res.Header().Set("Content-Type", "text/html; charset=utf-8");

	ctx.res.Write(buf.Bytes());
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

