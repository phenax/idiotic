
package controllers;

import (
	"fmt"
	"bytes"
	"log"
	"reflect"
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




type ResponseConfig struct {

	StatusCode int;
	Body string;

	ContentType string `default:"text/plain; charset=utf-8"`;

	ContentEncoding string;
}




func (ctx *Context) Respond(config *ResponseConfig) {

	headers := ctx.res.Header();

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

	ctx.res.WriteHeader(status);

	if(config.Body != "") {
		fmt.Fprint(ctx.res, config.Body);
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

	if(len(configs) > 0) {
		config = configs[0];
	} else {
		config = &ResponseConfig{ StatusCode: 200 };
	}

	config.Body = str;

	ctx.Respond(config);
}



/**
 * ctx.Render
 * Render a template and write to response
 *
 * params
 * -- templateName {string}   Name of the template to render
 */
func (ctx *Context) Render(templateName string, data interface{}) {

	// Open the template file
	html, err := ioutil.ReadFile(getTemplatePath(templateName));

	// Throw error if no template found
	if(err != nil) {
		fmt.Fprint(ctx.res, "Didnt render");
		return;
	}

	// Parse template
	tpl := template.Must(template.New("homepage").Parse(string(html)));

	// The template content
	buf := new(bytes.Buffer);

	if err := tpl.ExecuteTemplate(buf, "homepage", data); err != nil {
		fmt.Fprint(ctx.res, "Didnt render");
		return;
	}

	// Write it to the response
	ctx.res.Header().Set("Content-Type", "text/html; charset=utf-8");
	ctx.res.Write(buf.Bytes());
}







/**
 * Send an error message
 */
func (ctx *Context) ErrorMessage(statusCode int, err error) {

	fmt.Println("Error " + string(statusCode))
	log.Fatal(err);

	ctx.res.WriteHeader(statusCode);
	ctx.Send(err.Error());
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


func getConf(config *ResponseConfig, key string, value string) string {


	if value == "" {

		configType := reflect.TypeOf(*config)

		field, _ := configType.FieldByName(key)
		value = field.Tag.Get("default")
	}

	return value;
}

