
package main;

import (
	"fmt"
	"net/http"
	_ "github.com/phenax/idiotic/routes"
);


const (
	Port = "8080";
);


func main() {

	fmt.Println(`Server has started on port ` + Port);
	http.ListenAndServe(":" + Port, nil);
}
