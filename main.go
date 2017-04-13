
package main;

import (
	"fmt"
	"net/http"
	_ "github.com/phenax/idiotic/routes"
);


const (
	// The host for the server
	host = "0.0.0.0";
	// The server port to listen to
	port = "8080";
);


func main() {

	fmt.Println("Server has started on " + host + ":" + port);

	// Start the server
	http.ListenAndServe(host + ":" + port, nil);
}
