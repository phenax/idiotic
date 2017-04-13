
package routes;

import (
	"net/http"
	"github.com/gorilla/mux"
	ctrlr "github.com/phenax/idiotic/controllers"
)


// Initialize the routes
func init() {

	router := mux.NewRouter();
	indexRoutes(router);


	config := &ctrlr.StaticConfig{
		Pathprefix: "/public",
		Directory: "./public",
	};

	ctrlr.StaticRouter(router, config);

	// start with the base
	http.Handle("/", router);

}


func indexRoutes(router *mux.Router) {

	// Homepage
	router.HandleFunc("/{name}", ctrlr.Call(ctrlr.Homepage)).Methods("GET");
	router.HandleFunc("/people/{name}", ctrlr.Call(ctrlr.Homepage));

}



