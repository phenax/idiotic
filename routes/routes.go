
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
	userRoutes(router);


	// Add static router
	ctrlr.StaticRouter(router, &ctrlr.StaticConfig{
		Pathprefix: "/public",
		Directory: "./public",
	});

	// start with the base
	http.Handle("/", router);

}


func indexRoutes(router *mux.Router) {

	// Homepage
	router.HandleFunc("/", ctrlr.Call(ctrlr.HomePage));

	// gzip test
	router.HandleFunc("/gzip", ctrlr.Call(ctrlr.GzipTest));

}

func userRoutes(router *mux.Router) {

	// Profile page
	router.HandleFunc("/user/{name}", ctrlr.Call(ctrlr.ProfilePage));
}


