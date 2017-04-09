
package routes;

import (
	"net/http"
	"github.com/gorilla/mux"
	ctrlr "github.com/phenax/idiotic/controllers"
)


// Initialize the routes
func init() {

	router := mux.NewRouter();

	// Index route
	router.HandleFunc("/", ctrlr.Call(ctrlr.Homepage));

	// start with the base
	http.Handle("/", router);
}



