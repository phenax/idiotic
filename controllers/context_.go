
package controllers;

import (
	"fmt"
	"bytes"
	// "log"
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

	ContentType string `default:"text/plain; charset=utf-8"`;

	ContentEncoding string;
}




func (ctx *Context) ApplyConfig(config *ResponseConfig) {

	headers := ctx.res.Header();

	// Set the content type
	headers.Set("Content-Type", getConf(config, "ContentType", config.ContentType));

	// Apply encoding
	if(config.ContentEncoding != "") {
		headers.Set("Content-Encoding", config.ContentEncoding);
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

	if(len(configs) > 0) {
		ctx.ApplyConfig(configs[0]);
	}

	fmt.Fprint(ctx.res, str);
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
func (ctx *Context) ErrorMessage(number int, err error) {
	ctx.res.WriteHeader(number);
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

