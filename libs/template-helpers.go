package libs

import (
	"fmt"
	"html/template"
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
