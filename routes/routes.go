
package routes;

import (
	"net/http"
	ctrlr "github.com/phenax/idiotic/controllers"
)

func init() {

	http.HandleFunc("/", ctrlr.Call(ctrlr.Homepage));
}



