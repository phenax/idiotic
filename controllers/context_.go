
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



type Page struct {
	Title string;
	Body []byte;
}



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


func getPage(title string, templateName string) (*Page, error) {

	body, err := ioutil.ReadFile(getTemplatePath(templateName));

	if(err == nil) {
		body = body[:];
	}

	return &Page{ Body: body, Title: title }, err;
}

func getTemplatePath(templateName string) string {

	return filepath.Join(
		filepath.Base("."),
		"views",
		templateName + ".html",
	);
}


/**
 * ctx.render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) render(templateName string, options interface{}) {

	// The path of the template
	wrapperPath := getTemplatePath("wrapper");

	page, err := getPage("Coolness", templateName);

	if(err != nil) {
		log.Fatal(err);
		return;
	}

	// Parse the template
	tmpl, err := template.ParseFiles(wrapperPath);

	// If couldnt read
	if(err != nil) {
		log.Fatal(err);
		return;
	}

	// Render the template
	err = tmpl.ExecuteTemplate(ctx.res, "Wrapper", page);

	if(err != nil) {
		log.Fatal(err);
	}
}



