package libs

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

//
// TemplateHelpers - Template helper functions
//
var TemplateHelpers template.FuncMap

func init() {

	TemplateHelpers = template.FuncMap{
		"GetLink": GetLink,
	}
}

//
// ParseAllTemplates - Parses all templates
//
func ParseAllTemplates() *template.Template {

	tmpl := template.New("wrapper").Funcs(TemplateHelpers)

	root := filepath.Base(".")

	var files []string

	walker := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".gohtml") {
			return nil
		}

		files = append(files, filepath.Join(root, path))

		return err
	}

	err := filepath.Walk(filepath.Join(root, "views"), walker)

	if err != nil {
		fmt.Println("File walk Error", err)
	}

	return template.Must(tmpl.ParseFiles(files...))
}

//
// GetLink - Template helper function to get named router links
//
// params
// -- routeName {string}  Router name to craft the link for
// -- fields {...string}  A flat map of all fields and values to pass to the router
//
// returns
// -- {string}  The url
//
func GetLink(routeName string, fields ...string) string {

	router := GetRouter()

	url, err := router.Get(routeName).URL(fields...)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return url.String()
}

//
// GetTemplatePath - Get the full path to the template
//
// params
// -- templateName {string}
//
// returns
// -- {string}
//
func GetTemplatePath(templateName string) string {

	return filepath.Join(
		filepath.Base("."),
		"views",
		templateName+".gohtml",
	)
}
