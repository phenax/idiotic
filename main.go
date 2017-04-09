
package main;

import (
	"fmt"
	"net/http"
	_ "github.com/phenax/idiotic/routes"
);


const (
	Host = "0.0.0.0";
	Port = "8080";
);


func main() {

	fmt.Println(`Server has started on ` + Host + ":" + Port);

	// Start the server
	http.ListenAndServe(Host + ":" + Port, nil);
}
