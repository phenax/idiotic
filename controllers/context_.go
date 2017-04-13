
package controllers;

import (
	"fmt"
	"bytes"
	"log"
	"reflect"
	"net/http"
	"path/filepath"
	"io/ioutil"
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
);


/**
 * The server context
 *
 * fields
 * -- res {http.ResponseWriter}
 * -- req {*http.Request}
 */
type Context struct {
	Response http.ResponseWriter;
	Request *http.Request;
	Router *mux.Router;
	Params map[string]string;
};



/**
 * Craft a custom response
 * TODO: Docs
 *
 * fields
 * -- 
 */
type ResponseConfig struct {

	StatusCode int;
	Body string;

	ContentType string `default:"text/plain; charset=utf-8"`;
	ContentEncoding string;
}



/**
 * Respond to the request using the config
 *
 * params
 * -- config {*ResponseConfig}
 */
func (ctx *Context) Respond(config *ResponseConfig) {

	headers := ctx.Response.Header();

	// Set the content type
	headers.Set("Content-Type", getConf(config, "ContentType", config.ContentType));

	// Apply encoding
	if(config.ContentEncoding != "") {
		headers.Set("Content-Encoding", config.ContentEncoding);
	}


	// Status code
	status := 200;
	if(config.StatusCode != 0) {
		status = config.StatusCode;
	}

	ctx.Response.WriteHeader(status);

	// Response body
	if(config.Body != "") {
		fmt.Fprint(ctx.Response, config.Body);
	}
}



/**
 * ctx.Send
 * Writes a string of html to response
 *
 * params
 * -- str  {string}  The string to send
 * -- ...config  {*ResponseConfig}  Optional configuration
 */
func (ctx *Context) Send(str string, configs ...*ResponseConfig) {

	var config *ResponseConfig;

	// Default to just status code config
	if(len(configs) > 0) {
		config = configs[0];
	} else {
		config = &ResponseConfig{ StatusCode: 200 };
	}

	config.Body = str;

	ctx.Respond(config);
}



func (ctx *Context) JSON(obj interface{}, configs ...*ResponseConfig) {

	jsonContent, err := json.Marshal(obj);

	var config *ResponseConfig;

	// Default to just status code config
	if(len(configs) > 0) {
		config = configs[0];
	} else {
		config = &ResponseConfig{
			StatusCode: 200,
		};
	}

	if(err != nil) {
		config.StatusCode = 500;
	}

	config.ContentType = "application/json";

	ctx.Send(string(jsonContent), config);
}




/**
 * ctx.Render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) Render(templateName string, data interface{}, configs ...*ResponseConfig) {

	var config *ResponseConfig;

	// Default to just status code config
	if(len(configs) > 0) {
		config = configs[0];
	} else {
		config = &ResponseConfig{
			StatusCode: 200,
		};
	}

	// IF the content type is not set, default it to html
	if(config.ContentType == "") {
		config.ContentType = "text/html; charset=utf-8";
	}


	// Open the template file
	html, err := ioutil.ReadFile(getTemplatePath(templateName));

	// Throw error if no template found
	if(err != nil) {
		fmt.Fprint(ctx.Response, "Didnt render");
		return;
	}

	// Parse template
	tpl := template.Must(template.New("homepage").Parse(string(html)));

	// The template content
	buf := new(bytes.Buffer);

	if err := tpl.ExecuteTemplate(buf, "homepage", data); err != nil {
		fmt.Fprint(ctx.Response, "Didnt render");
		return;
	}

	config.Body = string(buf.Bytes());

	// Respond with the stuff
	ctx.Respond(config);
}







/**
 * Send an error message
 */
func (ctx *Context) ErrorMessage(statusCode int, err error) {

	fmt.Println("Error " + string(statusCode))
	log.Fatal(err);

	ctx.Send(err.Error(), &ResponseConfig{
		StatusCode: statusCode,
	});
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


/**
 * Get config and apply the default value(if any)
 */
func getConf(config *ResponseConfig, key string, value string) string {


	if value == "" {

		configType := reflect.TypeOf(*config)

		field, _ := configType.FieldByName(key)
		value = field.Tag.Get("default")
	}

	return value;
}

