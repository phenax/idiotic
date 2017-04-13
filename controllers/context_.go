
package controllers;

import (
	"fmt"
	"log"
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
 * The page related data
 *
 * fields
 * -- Title {string}  The title of the page
 * -- Body  {[]byte}  The content to go inside the body
 */
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
func (ctx *Context) Send(str string) {

	fmt.Fprint(ctx.res, str);
}

/**
 * ctx.render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) Render(templateName string, options interface{}) {


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






/**
 * getPage
 * Create a new Page instance
 *
 * params
 * -- title {string}  The title of the page
 * -- templateName {string}  The name of the template to render
 *
 * returns
 * -- {*Page}  The page data
 * -- {error}  Error is nil if the body was loaded successfully
 */
func getPage(title string, templateName string) (*Page, error) {

	body, err := ioutil.ReadFile(getTemplatePath(templateName));

	return &Page{ Body: body, Title: title }, err;
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

