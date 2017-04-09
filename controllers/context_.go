
package controllers;

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"io/ioutil"
	// "github.com/gorilla/mux"
	"html/template"
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

	templatePath := filepath.Join(filepath.Base("."), "views", templateName + ".html");

	templateContent, err := ioutil.ReadFile(templatePath);

	if(err != nil) {
		log.Fatal(err);
		return;
	}

	tmpl, err :=
		template.
			New("poo").
			Parse(string(templateContent[:]));

	options := struct{
		Cool string;
	}{
		Cool: "ness",
	};

	err = tmpl.ExecuteTemplate(ctx.res, "Wrapper", options);

	if(err != nil) {
		log.Fatal(err);
	}
}



