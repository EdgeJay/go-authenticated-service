package middlewares

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Logger(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc) {
	fmt.Println("Logging before")
	next(nil)
	fmt.Println("Logging after")
}
