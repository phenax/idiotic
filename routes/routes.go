
package routes;

import (
	"net/http"
	ctrlr "github.com/phenax/idiotic/controllers"
)


// Initialize the routes
func init() {

	// Index route
	http.HandleFunc("/", ctrlr.Call(ctrlr.Homepage));
}



